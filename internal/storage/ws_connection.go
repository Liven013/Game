package storage

import (
	"encoding/json"
	"game/internal/models"
	"sync"

	"github.com/gorilla/websocket"
)

var Users = NewConnStorage()

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

func (cs *ConnStorage) Delete(i interface{}) error {

	switch el := i.(type) {
	case string:
		err := cs.UsersStorage.Delete(el)
		if err != nil {
			return err
		}
		cs.mu.Lock()
		delete(cs.CS, el)
		cs.mu.Unlock()
	case *websocket.Conn:
		for k, val := range cs.CS {
			if val == el {
				err := cs.UsersStorage.Delete(k)
				if err != nil {
					return err
				}
				cs.UsersStorage.Delete(k)
				cs.mu.Lock()
				delete(cs.CS, k)
				cs.mu.Unlock()
				break
			}
		}
	}

	return nil
}

func (cs *ConnStorage) Broadcast(message string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	for id, client := range cs.CS {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			client.Close()
			delete(cs.CS, id)
		}
	}
}

func (cs *ConnStorage) SendAllUsers() {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	// Получаем всех пользователей из хранилища
	users := cs.GetAll()

	// Преобразуем список пользователей в JSON
	jsonData, err := json.Marshal(users)
	if err != nil {
		return
	}

	// Отправляем JSON-данные всем подключенным клиентам
	cs.Broadcast(string(jsonData))
}

func (cs *ConnStorage) GetAll() []models.User {
	return cs.UsersStorage.GetAll()
}

func (cs *ConnStorage) GetOne(id string) (models.User, error) {
	return cs.UsersStorage.GetOne(id)
}

func NewConnStorage() *ConnStorage {
	ls := NewLocalStorage()

	return &ConnStorage{UsersStorage: ls}
}
