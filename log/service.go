package log

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Write(msg string) {
	write(msg)
}
