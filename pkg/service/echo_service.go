package service

type EchoService struct{}

func (s *EchoService) Echo() string {
	return "Echo, it's working!"
}

func NewEchoService() *EchoService {
	return &EchoService{}
}
