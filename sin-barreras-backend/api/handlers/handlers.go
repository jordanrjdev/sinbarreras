package handlers

import "github.com/JGurus/sin-barreras-backend/models"

type UserHandler interface {
	Migrate() error
	Create(u *models.User) error
	Update(id int, u *models.User) error
	Delete(id int) error
	GetAll() ([]models.User, error)
	GetByID(id int) (models.User, error)
	GetByUsername(username string) (models.User, error)
}