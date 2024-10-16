package storage

import "game/internal/models"

var Players Storage = NewLocalStorage()

type Storage interface {
	Create(user models.User)
	GetOne(id string) (models.User, error)
	GetAll() []models.User
	Delete(id string) error
}
