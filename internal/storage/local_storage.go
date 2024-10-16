package storage

import (
	"errors"
	"fmt"
	"game/internal/models"
	"strconv"
	"sync"
)

type LocalStorage struct {
	S       map[string]models.User
	counter int
	mu      sync.Mutex
}

func (ls *LocalStorage) FindFreePositin() string {
	for i := 1; i < len(ls.S); i++ {
		id := fmt.Sprint(i)
		if _, ok := ls.S[id]; !ok {
			return id
		}
	}
	return fmt.Sprint(len(ls.S) + 1)
}

func (ls *LocalStorage) Create(user models.User) {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	user.ID = ls.FindFreePositin()
	ls.S[user.ID] = user
}

func (ls *LocalStorage) GetOne(id string) (models.User, error) {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	i, err := strconv.Atoi(id)
	if err != nil {
		return models.User{}, errors.New("ошибка преобразования индекса")
	}
	if i >= 0 && i < len(ls.S) {
		return ls.S[id], nil
	}
	return models.User{}, errors.New("ошибка значения индекса")
}

func (ls *LocalStorage) GetAll() []models.User {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	users := make([]models.User, len(ls.S))
	for k, val := range ls.S {
		i, _ := strconv.Atoi(k)
		users[i] = val
	}
	return users
}

func (ls *LocalStorage) Delete(id string) error {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	i, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("ошибка преобразования индекса")
	}
	if i <= 0 && i > len(ls.S) {
		return errors.New("ошибка значения индекса")
	}
	delete(ls.S, id)
	return nil
}

func NewLocalStorage() *LocalStorage {
	var host models.User = models.User{ID: "0", Name: "Server", Role: "host"}
	counter := 0
	players := map[string]models.User{
		host.ID: host,
	}
	return &LocalStorage{S: players, counter: counter}
}
