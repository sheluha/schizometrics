package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sheluha/schizometrics/internal/config"
	"github.com/sheluha/schizometrics/internal/telegram"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	bot, err := telegram.New(cfg.Telegram, &telegram.Handler{})
	if err != nil {
		log.Fatal("Cannot initialize telegram bot", err)
	}

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

	bot.Start(ctx)
}
