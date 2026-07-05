package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
	"syscall"
	"golang.org/x/sync/errgroup"
)

type server struct {
	serv *http.Server
}

func NewServer(port string, mux http.Handler) *server {
	return &server{
		&http.Server{
			Addr:port,
			Handler: mux,
		},
	}
}

func (s *server) Start(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		slog.Info("サーバースタート",slog.String("port:",s.serv.Addr))
		if err := s.serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("サーバーの起動に失敗しました: ", slog.String("port:", s.serv.Addr), slog.Any("error", err))
			return err
		}
		slog.Info("サーバーを閉じました", slog.String("port:", s.serv.Addr))
		return nil
	})

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.serv.Shutdown(shutdownCtx); err != nil {
		slog.Error("サーバーのシャットダウンに失敗しました", slog.String("port:", s.serv.Addr), slog.Any("error", err))
		return err
	}

	slog.Info("サーバーを正常にシャットダウンしました", slog.String("port:", s.serv.Addr))
	return eg.Wait()
}