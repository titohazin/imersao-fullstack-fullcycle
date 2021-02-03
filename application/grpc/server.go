package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"github.com/titohazin/imersao-fullstack-fullcycle/application/grpc/pb"
	"github.com/titohazin/imersao-fullstack-fullcycle/application/usecase"
	"github.com/titohazin/imersao-fullstack-fullcycle/repo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// StartGrpcServer StartGrpcServer
func StartGrpcServer(database *gorm.DB, port int) {

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pixRepository := repo.PixKeyRepositoryDb{Db: database}
	pixUseCase := usecase.PixUseCase{PixKeyRepository: &pixRepository}
	pixGrpcService := NewPixGrpcService(pixUseCase)
	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}
