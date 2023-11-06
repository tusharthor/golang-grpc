package main

import (
	"context"
	"fmt"
	"log"
	"net"
	totality_pb "totality_grpc_assign/proto"

	"google.golang.org/grpc"
)

type userServiceServer struct {
	data map[int32]*totality_pb.UserResponse
	totality_pb.UnimplementedUserServiceProtoServer
	// totality_pb.UserServiceProtoServer
}

// get user by id
func (s *userServiceServer) GetUserByIdProto(ctx context.Context, req *totality_pb.UserRequest) (*totality_pb.UserResponse, error) {
	// var user models.User
	user, ok := s.data[int32(req.UserId)]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

// get users list by id list
func (s *userServiceServer) GetUsersList(ctx context.Context, req *totality_pb.UsersListRequest) (*totality_pb.UsersListResponse, error) {
	//slice for user response
	var users []*totality_pb.UserResponse
	for _, userID := range req.UserIds {
		user, ok := s.data[int32(userID)]
		if !ok {
			return nil, fmt.Errorf("user not found for id: %d", userID)
		}

		//if found then append in the slice
		users = append(users, user)
	}

	return &totality_pb.UsersListResponse{Users: users}, nil
}
func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//create grpc server
	grpcServer := grpc.NewServer()
	data := map[int32]*totality_pb.UserResponse{
		1: {UserId: 1, Fname: "Steve", City: "LA", Phone: 9158906526, Height: 5.8, Married: true},
		2: {UserId: 2, Fname: "Harvey", City: "LV", Phone: 9158906526, Height: 5.9, Married: false},
	}

	totality_pb.RegisterUserServiceProtoServer(grpcServer, &userServiceServer{data: data})

	log.Println("started grpc sever")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
