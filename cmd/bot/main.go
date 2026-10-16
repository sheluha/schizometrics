package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sheluha/schizometrics/internal/analyze"
	"github.com/sheluha/schizometrics/internal/buffer"
	"github.com/sheluha/schizometrics/internal/config"
	"github.com/sheluha/schizometrics/internal/model"
	"github.com/sheluha/schizometrics/internal/telegram"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	jobs := make(chan model.AnalysisJob, 100)

	bm := buffer.NewBufferManager(jobs)
	handler := telegram.NewHandler(bm)

	bot, err := telegram.NewBot(cfg.Telegram, handler)
	if err != nil {
		log.Fatal("Cannot initialize telegram bot", err)
	}

	ma := analyze.MessageAnalyzer{}

	for range 4 {
		worker := analyze.NewAnalyzeWorker(jobs, ma)
		go worker.Run(ctx)
	}

	bot.Start(ctx)
}
