package main

import (
	"context"
	"fmt"
	"log"
	pb "totality_grpc_assign/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect in client: %v", err)
	}

	defer conn.Close()

	client := pb.NewUserServiceProtoClient(conn)

	//by id
	usedID := 1

	req := &pb.UserRequest{
		UserId: int64(usedID),
	}

	user, err := client.GetUserByIdProto(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to fetch user: %v", err)
	}

	fmt.Printf("\nPrinting User Info for ID: %v \n%s", usedID, user)

	//list
	fmt.Println("\n\nUsers list")
	userIDs := []int64{1, 2}

	reqList := &pb.UsersListRequest{
		UserIds: userIDs,
	}

	usersList, err := client.GetUsersList(context.Background(), reqList)
	if err != nil {
		log.Fatalf(" failed to fetch users list: %v ", err)
	}

	for _, users := range usersList.Users {
		fmt.Printf("%s", users)
		fmt.Println()
	}
}
