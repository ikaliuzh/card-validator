package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"

	"google.golang.org/grpc"

	api "github.com/ikaliuzh/card-validator/api/proto"
	"github.com/ikaliuzh/card-validator/internal/server"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	var env string
	var port int
	flag.StringVar(&env, "env", envProd, "deployment environment")
	flag.IntVar(&port, "port", 8080, "port to listen on")

	logger := setupLogger(env)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	cardValidatorServer := server.New(
		server.WithLog(logger),
	)

	api.RegisterCardValidatorServer(grpcServer, cardValidatorServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		panic(fmt.Errorf("unknown env %q", env))
	}

	return logger
}
