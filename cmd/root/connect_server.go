package root

import (
	"google.golang.org/grpc/credentials"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectChatServer() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect chat server: %v", err)
	}
	return conn
}

func ConnectAuthServer() *grpc.ClientConn {

	creds, err := credentials.NewClientTLSFromFile("certificates/service.pem", "")
	if err != nil {
		log.Fatalf("failed to connect login server: %v", err)
	}

	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalf("failed to connect login server: %v", err)
	}
	return conn
}
