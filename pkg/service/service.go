package service

import "github.com/wregulski/log-poc/pkg/log"

type Echo interface {
	Echo() string
	ErrorEcho() string
}

type Services struct {
	logger log.Logger
	Echo   Echo
}

func NewServices(lf log.LoggerFactory) *Services {
	logger := lf.NewLogger("services")
	return &Services{
		logger: logger,
		Echo:   NewEchoService(logger),
	}
}
