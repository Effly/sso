package grpcapp

import (
	"fmt"
	authgrpc "github.com/Effly/sso/internal/grpc/auth"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewApp(
	log *slog.Logger,
	authService authgrpc.Auth,
	port int,
) *App {
	grpcServer := grpc.NewServer()

	authgrpc.Register(grpcServer, authService)

	return &App{
		log:        log,
		port:       port,
		gRPCServer: grpcServer,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grcapp.Run" //операция для лога

	log := a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port),
	)

	log.Info("start grpc server")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", slog.String("addr", l.Addr().String()))

	if err = a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grcapp.Stop"

	a.log.With(slog.String("op", op)).Info("stop grpc server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
