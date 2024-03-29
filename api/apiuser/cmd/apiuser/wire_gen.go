// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"apiuser/internal/conf"
	"apiuser/internal/server"
	"apiuser/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	log2 "github.com/gomicroim/gomicroim/pkg/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, discover *conf.Discover, logger log.Logger, logLogger *log2.Logger, discovery registry.Discovery) (*kratos.App, func(), error) {
	authClient := service.NewAuthClient(discover, discovery)
	apiUserService := service.NewApiUserService(authClient)
	httpServer := server.NewHTTPServer(confServer, apiUserService, logger)
	app := newApp(logLogger, httpServer)
	return app, func() {
	}, nil
}
