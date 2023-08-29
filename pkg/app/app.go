package app

import (
	"context"

	"github.com/wregulski/log-poc/pkg/config"
	"github.com/wregulski/log-poc/pkg/handler"
	"github.com/wregulski/log-poc/pkg/log"
	"github.com/wregulski/log-poc/pkg/server"
	"github.com/wregulski/log-poc/pkg/service"
)

func Run() {
	config.NewViperConfig()
	lf := log.DefaultLoggerFactory()
	log := lf.NewLogger("main")

	s := service.NewServices(lf)
	h := handler.NewHandler(s, lf)

	srv := server.NewServer(h.Init())

	if err := srv.Run(); err != nil {
		log.Criticalf("error running server: %s", err)
		panic(err)
	}

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Criticalf("error shutting down server: %s", err)
		panic(err)
	}
}
