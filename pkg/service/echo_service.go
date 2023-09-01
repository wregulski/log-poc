package service

import (
	"errors"

	"github.com/wregulski/log-poc/pkg/log"
)

type EchoService struct {
	logger log.Logger
}

func (s *EchoService) Echo() string {
	s.logger.Info("Info message from logger!")
	return "Echo, it's working!"
}

func (s *EchoService) ErrorEcho() string {
	err := errors.New("something crashed i guess :/")
	s.logger.Errorf("Error message from logger: %s", err)
	return "Error, it's working!"
}

func NewEchoService(logger log.Logger) *EchoService {
	return &EchoService{
		logger: logger,
	}
}
