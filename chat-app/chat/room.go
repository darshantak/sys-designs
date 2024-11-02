package chat

import (
	"sync"

	"github.com/gorilla/websocket"
)

type ChatRoom struct {
	Id      string
	Name    string
	Users   map[*websocket.Conn]bool
	Mu      sync.Mutex
	Message chan []byte
}

func NewChatRoom(id, name string) *ChatRoom {
	return &ChatRoom{
		Id:      id,
		Name:    name,
		Users:   make(map[*websocket.Conn]bool),
		Message: make(chan []byte),
	}
}

func (room *ChatRoom) AddUsers(conn *websocket.Conn) {
	room.Mu.Lock()
	defer room.Mu.Unlock()
	room.Users[conn] = true
}

func (room *ChatRoom) RemoveUsers(conn *websocket.Conn) {
	room.Mu.Lock()
	defer room.Mu.Unlock()
	delete(room.Users, conn)
}

func (room *ChatRoom) Broadcast(message []byte) {
	for conn := range room.Users {
		conn.WriteMessage(websocket.TextMessage, message)
	}
}
