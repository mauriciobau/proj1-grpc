package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/mauriciobau/proj1-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Coud not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	//AddUser(client)
	//AddUserVerbose(client)
	AddUsers(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Mauricio",
		Email: "mauricio@email.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Coud not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Mauricio",
		Email: "mauricio@email.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Coud not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Coud not receive the msg: %v", err)
		}
		fmt.Println("Status:", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "1",
			Name:  "Mauricio 1",
			Email: "mauricio1@mail.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "Mauricio 2",
			Email: "mauricio2@mail.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "Mauricio 3",
			Email: "mauricio3@mail.com",
		},
		&pb.User{
			Id:    "4",
			Name:  "Mauricio 4",
			Email: "mauricio4@mail.com",
		},
		&pb.User{
			Id:    "5",
			Name:  "Mauricio 5",
			Email: "mauricio5@mail.com",
		},
		&pb.User{
			Id:    "6",
			Name:  "Mauricio 6",
			Email: "mauricio6@mail.com",
		},
		&pb.User{
			Id:    "7",
			Name:  "Mauricio 7",
			Email: "mauricio7@mail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}
