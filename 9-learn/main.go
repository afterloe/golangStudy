package main

import (
	"GoGrpcDemo/pb"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"net"
)

type userServerImpl struct {
	pb.UnimplementedUserServerServer
}

func (*userServerImpl) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Received request: %v", request)
	res := &pb.LoginResponse{
		SessionID:    uuid.NewString(),
		AccessToken:  uuid.NewString(),
		RefreshToken: uuid.NewString(),
	}
	return res, nil
}

func main() {
	listen, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServerServer(s, &userServerImpl{})
	defer func() {
		s.Stop()
		err := listen.Close()
		if err != nil {
			log.Panic(err)
			return
		}
	}()

	log.Printf("Listening on port 9092 \n")
	err = s.Serve(listen)
	if err != nil {
		log.Panic(err)
	}
}
