package product

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Product, error)
	GetByCategory(categoryID int) ([]Product, error)
	GetBestSeller() ([]Product, error)
	GetSuperSpesial() ([]Product, error)
	GetPopular() ([]Product, error)
	GetByName(name string) ([]Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Product, error) {
	var products []Product
	err := r.db.Preload("Category").Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) GetByCategory(categoryID int) ([]Product, error) {
	var products []Product
	err := r.db.Where("category_id = ?", categoryID).Preload("Category").Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) GetBestSeller() ([]Product, error) {
	var products []Product
	err := r.db.Preload("Category").Limit(3).Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) GetSuperSpesial() ([]Product, error) {
	var products []Product
	err := r.db.Where("id > ? AND id < ?", "3", "8").Preload("Category").Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) GetPopular() ([]Product, error) {
	var products []Product
	err := r.db.Where("id > ? AND id < ?", "7", "11").Preload("Category").Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) GetByName(name string) ([]Product, error) {
	var products []Product
	err := r.db.Where("name LIKE ?", "%"+ name +"%").Preload("Category").Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}
