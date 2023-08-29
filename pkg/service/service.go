package service

import "github.com/wregulski/log-poc/pkg/log"

type Echo interface {
	Echo() string
}

type Services struct {
	logger log.Logger
	Echo   Echo
}

func NewServices(lf log.LoggerFactory) *Services {
	return &Services{
		logger: lf.NewLogger("services"),
		Echo:   NewEchoService(),
	}
}
