package repository

import "wobble_tracker/internal/models"

type CatRepository interface {
	Create(cat *models.Cat) (*models.Cat, error)
	GetByID(id int) (*models.Cat, error)
	Update(id int, updates *models.UpdateCatRequest) (*models.Cat, error)
	Delete(id int) error
	List(limit, offset int) ([]*models.Cat, error)
}

// ! additional repos ... 

type Repositories struct {
	Cat CatRepository
	// ! additional repos ...
}