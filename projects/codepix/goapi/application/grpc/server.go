package grpc 

import (
	"gorm.io/gorm"
	"google.golang.org/gprc"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	pixUseCase := usecase.PixUseCase(PixKeyRepository: pixRepository)
	PixGrpcService := NewPixGrpcService(pixUseCase)
	pb.RegisterPixServiceServer(grpcServer, PixGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}

	log.Printf("gRPC server has benn started on port %d", port)

	err = grpcServer.Server(listener)

	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}
