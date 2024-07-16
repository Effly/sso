package app

import (
	grpcapp "github.com/Effly/sso/internal/app/grpc"
	"github.com/Effly/sso/internal/services/auth"
	"github.com/Effly/sso/internal/storage/sqlite"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.NewApp(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
