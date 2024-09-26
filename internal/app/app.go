package app

import (
	grpcapp "grpc-calculator/internal/app/grpc"

	"github.com/sirupsen/logrus"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	logger *logrus.Logger,
	port int,
) *App {
	grpcApp := grpcapp.New(logger, port)
	return &App{
		GRPCServer: grpcApp,
	}
}
