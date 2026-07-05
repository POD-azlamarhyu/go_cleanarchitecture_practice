package main

import (
	"log/slog"
	"context"
	"cleanarchitecture-practice/src/config"
	"fmt"
	"runtime"
	"cleanarchitecture-practice/src/infrastructure/database"
	"cleanarchitecture-practice/src/infrastructure/server"
)

func run(ctx context.Context, cfg *config.Config)  {
	db := database.NewDammyDB(ctx, cfg)
	defer db.Close()

	goServer := server.NewServer(cfg.Server, server.NewMux())
	goServer.Start(ctx)
}

func main(){
	slog.Info("Starting the main function for Application...",
	slog.String("version", "0.1.0"),
	slog.String("Language", fmt.Sprintf("Go %s", runtime.Version())),
	slog.String("OS", fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)),
	)
	ctx := context.Background()
	cfg, err := config.GetConfig()
	if err != nil {
		slog.Error("failed to get config", "error", err)
		return
	}

	run(ctx, cfg)
}