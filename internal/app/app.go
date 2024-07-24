package app

import (
	"github.com/adametsderschopfer/gRPC-auth-service/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	// todo refactor this (storage path)
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TODO: Инициализировать хранилище
	// TODO: Инициализировать auth service

	grpcApp := grpcapp.NewApp(log, grpcPort)

	return &App{GRPCSrv: grpcApp}
}
