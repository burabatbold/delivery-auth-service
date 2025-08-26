package grpc

import (
	"fmt"
	"log"
	"net"

	adminGrpcService "github.com/burabatbold/delivery-auth-service/grpc/admin"
	protos "github.com/burabatbold/delivery-auth-service/grpc/protos/admin-service/grpc/protos"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func StartServer() {

	log.Println("Starting gRPC server on port", viper.GetInt("grpc.port"))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("grpc.port")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	adminService := adminGrpcService.NewAdminService()

	protos.RegisterAdminServiceServer(s, adminService)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
