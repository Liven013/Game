package storage

import (
	"game/internal/models"
	"sync"

	"github.com/gorilla/websocket"
)

type ConnStorage struct {
	UsersStorage Storage
	CS           map[string]*websocket.Conn
	mu           sync.Mutex
}

func (cs *ConnStorage) Create(user models.User, ws *websocket.Conn) {
	user = cs.UsersStorage.Create(user)
	cs.mu.Lock()
	cs.CS[user.ID] = ws
	cs.mu.Unlock()
}

func (cs *ConnStorage) Delete(id string) error {
	err := cs.UsersStorage.Delete(id)
	if err != nil {
		return err
	}
	cs.mu.Lock()
	delete(cs.CS, id)
	cs.mu.Unlock()
	return nil
}

func NewConnStorage() *ConnStorage {
	ls := NewLocalStorage()

	return &ConnStorage{UsersStorage: ls}
}
