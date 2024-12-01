package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"nirvana/internal/app/nirvana"
	"nirvana/internal/app/nirvana/repository"
	"nirvana/internal/app/nirvana/usecases"
	"nirvana/internal/config"
	nirvana2 "nirvana/pkg/api/nirvana"
)

func main() {
	ctx := context.Background()
	dbConfig := config.NewDBConfig()

	db, err := config.ConnectDB(ctx, dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	exceptionRepo := repository.NewRepository(db)
	createUseCase := usecases.NewCreateExceptionUseCase(exceptionRepo)
	checkUseCase := usecases.NewCheckExceptionUseCase(exceptionRepo)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	exceptionService := nirvana.NewService(
		createUseCase,
		checkUseCase,
	)
	nirvana2.RegisterNirvanaServer(grpcServer, exceptionService)

	reflection.Register(grpcServer)

	log.Println("gRPC server is running on port 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
