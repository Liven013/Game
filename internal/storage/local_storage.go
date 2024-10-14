package storage

import (
	"errors"
	"game/internal/models"
	"strconv"
)

type LocalStorage struct {
	S []models.User
}

func (ls *LocalStorage) Create(user models.User) {
	ls.S = append(ls.S, user)
}

func (ls *LocalStorage) GetOne(id string) (models.User, error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return models.User{}, errors.New("ошибка преобразования индекса")
	}
	i--
	if i > 0 && i < len(ls.S) {
		return ls.S[i], nil
	}
	return models.User{}, errors.New("ошибка значения индекса")
}

func (ls *LocalStorage) GetAll() []models.User {
	return ls.S
}

func (ls *LocalStorage) Delete(id string) error {
	i, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("ошибка преобразования индекса")
	}
	i--
	if i < 0 && i > len(ls.S) {
		return errors.New("ошибка значения индекса")

	}

	ls.S = append(ls.S[:i], ls.S[i+1:]...)
	return nil
}

func NewLocalStorage() *LocalStorage {
	players := []models.User{
		{ID: "1", Name: "Liven", Role: "leader"},
		{ID: "2", Name: "NoName", Role: "player"},
	}
	return &LocalStorage{S: players}
}
