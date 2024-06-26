package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"

	"google.golang.org/grpc"

	api "github.com/ikaliuzh/card-validator/gen/proto"
	"github.com/ikaliuzh/card-validator/internal/server"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	var env string
	var host string
	var port int
	flag.StringVar(&env, "env", envProd, "deployment environment")
	flag.StringVar(&host, "host", "0.0.0.0", "host to listen on")
	flag.IntVar(&port, "port", 8080, "port to listen on")
	flag.Parse()

	logger := setupLogger(env)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	cardValidatorServer := server.New(
		server.WithLog(logger),
	)

	api.RegisterCardValidatorServer(grpcServer, cardValidatorServer)
	logger.Info("starting server", slog.Int("port", port), slog.String("env", env))
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
		log.Fatalf("unknown env %q", env)
	}

	return logger
}
