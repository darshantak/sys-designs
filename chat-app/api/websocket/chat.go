package websocket

import (
	"chat-app/chat"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var chatRooms = make(map[string]*chat.ChatRoom)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func JoinChatRoom(w http.ResponseWriter, r *http.Request) {
	roomId := r.URL.Query().Get("room_id")
	if roomId == "" {
		http.Error(w, "missing room_id", http.StatusBadGateway)
		return
	}

	_, exists := chatRooms[roomId]
	if !exists {
		chatRooms[roomId] = chat.NewChatRoom(roomId, "Chat Room"+roomId)
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("websocket upgrade failed", err)
		return
	}
	room := chatRooms[roomId]
	room.AddUsers(conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error : ", err)
			break
		}
		room.Broadcast(message)
	}
}
