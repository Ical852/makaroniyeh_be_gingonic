package product

type Service interface {
	GetAll() ([]Product, error)
	GetByCategory(input GetByCatInput) ([]Product, error)
	GetBestSeller() ([]Product, error)
	GetSuperSpesial() ([]Product, error)
	GetPopular() ([]Product, error)
	GetByName(input GetByNameInput) ([]Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) GetByCategory(input GetByCatInput) ([]Product, error) {
	products, err := s.repository.GetByCategory(input.ID)
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) GetBestSeller() ([]Product, error) {
	products, err := s.repository.GetBestSeller()
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) GetSuperSpesial() ([]Product, error) {
	products, err := s.repository.GetSuperSpesial()
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) GetPopular() ([]Product, error) {
	products, err := s.repository.GetPopular()
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) GetByName(input GetByNameInput) ([]Product, error) {
	products, err := s.repository.GetByName(input.Name)
	if err != nil {
		return products, err
	}

	return products, nil
}
