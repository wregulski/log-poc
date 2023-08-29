package service

type Echo interface {
	Echo() string
}

type Services struct {
	Echo Echo
}

func NewServices() *Services {
	return &Services{
		Echo: NewEchoService(),
	}
}
