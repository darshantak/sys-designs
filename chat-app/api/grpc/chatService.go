package grpc

import (
	"chat-app/proto"
	"log"
	"sync"

	"google.golang.org/grpc/peer"
)

type ChatServer struct {
	proto.UnimplementedChatServiceServer
	rooms map[string]*Rooms
	mu    sync.Mutex
}

type Rooms struct {
	clients map[proto.ChatService_ChatStreamServer]bool
	mu      sync.Mutex
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		rooms: make(map[string]*Rooms),
	}
}

func (s *ChatServer) joinRoom(roomId string, stream proto.ChatService_ChatStreamServer) {
	s.mu.Lock()
	if _, exists := s.rooms[roomId]; !exists {
		s.rooms[roomId] = &Rooms{clients: make(map[proto.ChatService_ChatStreamServer]bool)}
	}
	room := s.rooms[roomId]
	s.mu.Unlock()

	room.mu.Lock()
	room.clients[stream] = true
	room.mu.Unlock()
}

func (s *ChatServer) leaveRoom(roomId string, stream proto.ChatService_ChatStreamServer) {
	s.mu.Lock()
	if room, exists := s.rooms[roomId]; exists {
		room.mu.Lock()
		delete(room.clients, stream)
		room.mu.Unlock()
	}
	s.mu.Unlock()
}

func (s *ChatServer) ChatStream(stream proto.ChatService_ChatStreamServer) error {
	peerInfo, _ := peer.FromContext(stream.Context())
	log.Printf("New client connected: %v", peerInfo.Addr)
	defer func() {
		log.Printf("Client Disconnected")
	}()

	for {
		in, err := stream.Recv()
		if err != nil {
			s.leaveRoom(in.RoomId, stream)
			return err
		}
		s.joinRoom(in.RoomId, stream)
		s.broadcast(in.RoomId, in)
	}

}

func (s *ChatServer) broadcast(roomId string, msg *proto.ChatMessage) {
	s.mu.Lock()
	room, exists := s.rooms[roomId]
	s.mu.Unlock()

	if exists {
		room.mu.Lock()
		for client := range room.clients {
			if err := client.Send(msg); err != nil {
				log.Printf("error sending messages %v", err)
				s.leaveRoom(roomId, client)
			}
		}
		room.mu.Unlock()
	}
}
