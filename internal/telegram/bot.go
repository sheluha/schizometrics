package telegram

import (
	"context"
	"log"

	tgbot "github.com/go-telegram/bot"
	"github.com/sheluha/schizometrics/internal/config"
)

type Bot struct {
	client *tgbot.Bot
}

func New(c config.Telegram, h *Handler) (*Bot, error) {
	client, err := tgbot.New(c.Token, tgbot.WithDefaultHandler(h.Handle))
	if err != nil {
		return nil, err
	}
	return &Bot{client}, nil
}

func (b *Bot) Start(ctx context.Context) {
	log.Println("Bot is running")
	b.client.Start(ctx)
}
