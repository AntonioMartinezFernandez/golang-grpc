package user_infra

import (
	"context"
	"fmt"

	pb "github.com/AntonioMartinezFernandez/golang-grpc/internal/user/infra/protobuffer"
)

type UserGrpcServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserGrpcServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	fmt.Println("getting user with id: ", req.Id)

	var user *pb.User

	for i := 0; i < len(users); i++ {
		if users[i].Id == req.Id {
			user = users[i]
			break
		}
	}

	return &pb.GetUserResponse{
		User: user,
	}, nil
}

func (s *UserGrpcServiceServer) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	fmt.Println("getting all users")

	return &pb.GetAllUsersResponse{
		Users: users,
	}, nil
}

var users []*pb.User = []*pb.User{
	{
		Id:      "1",
		Name:    "John",
		Age:     30,
		Enabled: true,
	},
	{
		Id:      "2",
		Name:    "Mick",
		Age:     39,
		Enabled: true,
	},
}
