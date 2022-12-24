package biz

import (
	pb "chat/api"
	"chat/internal/conf"
	"chat/internal/mq"
	"context"
	v1 "github.com/gomicroim/go-timeline/api/v1"
	"github.com/gomicroim/gomicroim/pkg/log"
	"github.com/gomicroim/gomicroim/protos/chat"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"time"
)

// MessageSyncUseCase timeline消息存储(收件箱)
type MessageSyncUseCase interface {
	// WriteMsg 使用mq异步扩散写存储到mongo db
	WriteMsg(ctx context.Context, msg *pb.SendMsgRequest)
}

type messageSyncUseCase struct {
	timeline v1.TimelineClient
	logger   *log.Logger

	producer mq.MsgProducer
	consumer mq.ChatConsumerGroup
}

var _ MessageSyncUseCase = (*messageSyncUseCase)(nil)

func NewMessageSyncUseCase(config *conf.Bootstrap, l *log.Logger, producer mq.MsgProducer, groupConsumer mq.ChatConsumerGroup) MessageSyncUseCase {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	conn, err := grpc.DialContext(ctx, config.Data.Timeline.GrpcAddr)
	if err != nil {
		panic(err)
	}
	m := &messageSyncUseCase{
		timeline: v1.NewTimelineClient(conn),
		producer: producer,
		consumer: groupConsumer,
		logger:   l,
	}
	groupConsumer.ObserverOutboxWriteMsg(m.onHandleOutboxWriteMsg)
	return m
}

func (m messageSyncUseCase) WriteMsg(ctx context.Context, msg *pb.SendMsgRequest) {
	// persistent to mongo use kafka (outbox, write diffusion)
	_, _, err := m.producer.SendOutboxMsg(ctx, msg.To, msg)
	if err != nil {
		log.L.Error("SendOutboxMsg err", zap.Error(err))
	}
}

func (m messageSyncUseCase) onHandleOutboxWriteMsg(ctx context.Context, req *pb.SendMsgRequest) {
	m.logger.Info("onHandleOutboxWriteMsg", zap.String("req", req.String()))

	var (
		buffer         []byte
		err            error
		sendReply      *v1.SendMessageReply
		sendGroupReply *v1.SendGroupReply
	)

	switch req.SessionType {
	case chat.IMSessionType_SessionTypeSingle:
		sendReply, err = m.timeline.Send(context.Background(), &v1.SendMessageRequest{
			From:    "",
			To:      "",
			Message: "",
		})
		if sendReply != nil {
			m.logger.Info("timeline.Send reply", zap.Any("reply", sendReply.String()))
		}
	case chat.IMSessionType_SessionTypeNormalGroup:
		sendGroupReply, err = m.timeline.SendGroup(context.Background(), &v1.SendGroupRequest{})
		if sendGroupReply != nil {
			m.logger.Info("timeline.SendGroup reply", zap.Any("reply", sendReply.String()))
		}
	default:
		m.logger.Warn("error write to mongodb: unknown session type")
		return
	}

	// 写入到死信队列，手动处理
	if err != nil {
		buffer, err = protojson.Marshal(req)
		if err != nil {
			m.logger.Error("SendDeadMsg Marshal error", zap.Error(err))
			return
		}
		_, _, err = m.producer.SendDeadMsg(ctx, req.To, mq.DeadChatMsg{
			BizType: mq.BizTypeChatMsg,
			Content: string(buffer),
		})
		if err != nil {
			m.logger.Error("SendDeadMsg error", zap.Error(err))
		}
	}
}
