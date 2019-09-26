import 'package:cc_flutter_app/gui/model/model.dart';
import 'package:cc_flutter_app/imsdk/im_client.dart';
import 'package:cc_flutter_app/imsdk/proto/CIM.Def.pb.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

/// 代表一条消息的组件
class MsgItem extends StatefulWidget {
  final MessageModel msg;
  final UserModel fromUser;

  MsgItem(this.msg, this.fromUser);

  @override
  State<StatefulWidget> createState() {
    if (ImClient.singleton.isSelf(fromUser.userId)) {
      return MsgItemMeState(msg, fromUser);
    }
    return MsgItemOtherState(msg, fromUser);
  }
}

class _MsgItemState extends State<MsgItem> {
  final MessageModel msg;
  final UserModel fromUser;

  _MsgItemState(this.msg, this.fromUser);

  @override
  Widget build(BuildContext context) {
    return null;
  }

  // 生成聊天内容
  Widget _buildMsgContent() {
    return null;
  }

  // 头像
  Widget _buildAvatar(edge) {
    return Container(
      width: 36,
      margin: edge,
      child: ClipRRect(
        borderRadius: BorderRadius.circular(3.0),
        child: FadeInImage(
          image: NetworkImage(fromUser.avatarURL),
          placeholder: AssetImage("assets/default_avatar.png"),
        ),
      ),
    );
  }

  void _reSendMsg() {}
}

/// 别人的消息
class MsgItemOtherState extends _MsgItemState {
  MsgItemOtherState(MessageModel msg, UserModel fromUser)
      : super(msg, fromUser);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: <Widget>[
          _buildAvatar(EdgeInsets.only(right: 8.0, top: 8.0)),
          Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: <Widget>[
              Text(fromUser.nickName,
                  style: Theme.of(context).textTheme.subhead),
              _buildMsgContent(),
            ],
          )
        ],
      ),
    );
  }
}

/// 自己的消息
class MsgItemMeState extends _MsgItemState {
  MsgItemMeState(MessageModel msg, UserModel fromUser) : super(msg, fromUser);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.end,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: <Widget>[
          Column(
            crossAxisAlignment: CrossAxisAlignment.end,
            children: <Widget>[
              Text(fromUser.nickName,
                  style: Theme.of(context).textTheme.subhead),
              Row(
                children: <Widget>[
                  _buildSendMsgContent(),
                  _buildMsgContent(),
                ],
              )
            ],
          ),
          _buildAvatar(EdgeInsets.only(left: 8.0, top: 8.0))
        ],
      ),
    );
  }

  // 我发发送消息内容
  Widget _buildSendMsgContent() {
    return msg.msgStatus == CIMMsgStatus.kCIM_MSG_STATUS_SENDING
        ? CircularProgressIndicator(
            strokeWidth: 1.0,
            valueColor: AlwaysStoppedAnimation<Color>(Colors.black12),
          )
        : (msg.msgStatus == CIMMsgStatus.kCIM_MSG_STATUS_FAILED
            ? IconButton(
                icon: Icon(
                  Icons.error,
                  color: Colors.red,
                ),
                onPressed: () {
                  _reSendMsg();
                },
              )
            : Center());
  }
}