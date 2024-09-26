package main

import (
	"grpc-calculator/internal/app"
	"grpc-calculator/internal/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func main() {
	config := config.MustLoad()

	logger := setupLogger(config.Env)
	logger.Info("Application started")

	application := app.New(logger, config.GRPC.Port)

	go func() {
		application.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
	logger.Info("Gracefully stopped")
}

func setupLogger(env string) *logrus.Logger {
	logger := logrus.New()

	switch env {
	case envDev:
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
		logger.SetLevel(logrus.DebugLevel)
		logger.SetOutput(os.Stdout)
	case envProd:
		logger.SetFormatter(&logrus.JSONFormatter{})
		logger.SetLevel(logrus.InfoLevel)
		logger.SetOutput(os.Stdout)
	default:
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
		logger.SetLevel(logrus.WarnLevel)
		logger.SetOutput(os.Stdout)
	}
	return logger
}
