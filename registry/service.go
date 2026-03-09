package registry

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Add(registration Registration) error {
	return reg.add(registration)
}

func (s *Service) Remove(url string) error {
	return reg.remove(url)
}
