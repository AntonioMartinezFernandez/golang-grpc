package main

import (
	"fmt"
	"net"

	"github.com/AntonioMartinezFernandez/golang-grpc/cmd/di"
	user_infra "github.com/AntonioMartinezFernandez/golang-grpc/internal/user/infra"
	pb "github.com/AntonioMartinezFernandez/golang-grpc/internal/user/infra/protobuffer"
	"google.golang.org/grpc"
)

func main() {
	commonDi := di.InitCommonServices()
	commonDi.Logger.Info("Starting gRPC server...", "port", commonDi.Config.GrpcPort)

	listner, err := net.Listen("tcp", ":"+fmt.Sprint(commonDi.Config.GrpcPort))

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterUserServiceServer(serv, &user_infra.UserGrpcServiceServer{})
	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}
}
