package product

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return []Product{
		{Title: "One"},
		{Title: "Two"},
		{Title: "Three"},
	}
}
