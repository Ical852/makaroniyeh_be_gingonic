package category

type Service interface {
	GetAll() ([]Category, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]Category, error) {
	categories, err := s.repository.GetAll()
	if err != nil {
		return categories, err
	}

	return categories, nil
}
