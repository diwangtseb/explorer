// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"explorer/internal/biz"
	"explorer/internal/conf"
	"explorer/internal/server"
	"explorer/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, data *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	explorerUseCase := biz.NewFooUseCase(logger)
	explorerService := service.NewFooService(explorerUseCase)
	httpServer := server.NewHTTPServer(confServer, explorerService, logger)
	grpcServer := server.NewGRPCServer(confServer, explorerService)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
	}, nil
}
