package main

import (
	gr "chat-app/api/grpc"
	"chat-app/api/rest"
	"chat-app/api/websocket"
	"chat-app/proto"
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	//******************************//
	//**********REST API************//
	//******************************//
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the chat service")
	})

	http.HandleFunc("/api/register", rest.RegisterUser)
	http.HandleFunc("/api/login", rest.LoginUser)

	http.HandleFunc("/ws/join", websocket.JoinChatRoom)

	port := ":7007"
	fmt.Printf("Starting server at %s", port)
	log.Fatal(http.ListenAndServe(port, nil))

	//**********GRPC************//
	grpcServer := grpc.NewServer()
	chatServer := gr.NewChatServer()
	proto.RegisterChatServiceServer(grpcServer, chatServer)
	reflection.Register(grpcServer)

	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen to the port %v", err)
		}
		log.Println("gRPC server listening on port 50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()
	select {}
}
