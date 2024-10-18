package storage

import "game/internal/models"

type Storage interface {
	Create(user models.User) models.User
	GetOne(id string) (models.User, error)
	GetAll() []models.User
	Delete(id string) error
}
