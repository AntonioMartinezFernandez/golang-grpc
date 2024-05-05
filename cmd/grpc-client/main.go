package main

import (
	"context"
	"fmt"

	"github.com/AntonioMartinezFernandez/golang-grpc/cmd/di"
	pb "github.com/AntonioMartinezFernandez/golang-grpc/internal/user/infra/protobuffer"
	"google.golang.org/grpc"
)

func main() {
	commonDi := di.InitCommonServices()

	commonDi.Logger.Info("Starting gRPC client...")

	conn, err := grpc.Dial("localhost:"+fmt.Sprint(commonDi.Config.GrpcPort), grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceClient := pb.NewUserServiceClient(conn)

	res, err := serviceClient.GetAllUsers(context.Background(), &pb.GetAllUsersRequest{})

	if err != nil {
		panic("users not found  " + err.Error())
	}

	fmt.Println(res.Users)
}
