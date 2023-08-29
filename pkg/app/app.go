package app

import (
	"context"

	"github.com/wregulski/log-poc/pkg/handler"
	"github.com/wregulski/log-poc/pkg/server"
	"github.com/wregulski/log-poc/pkg/service"
)

func Run() {

	s := service.NewServices()
	h := handler.NewHandler(s)

	srv := server.NewServer(h.Init())

	if err := srv.Run(); err != nil {
		panic(err)
	}

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
}
