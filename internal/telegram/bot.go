package telegram

import (
	"context"
	"log"

	tgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/sheluha/schizometrics/internal/config"
)

type Bot struct {
	client *tgbot.Bot
}

func NewBot(c config.Telegram, h *Handler) (*Bot, error) {
	client, err := tgbot.New(
		c.Token,
		tgbot.WithDefaultHandler(h.Handle),
		tgbot.WithWorkers(1),
		tgbot.WithNotAsyncHandlers(),
	)
	if err != nil {
		return nil, err
	}
	return &Bot{client}, nil
}

func (b *Bot) Start(ctx context.Context) {
	log.Println("Bot is running")
	b.client.Start(ctx)
}

func (b *Bot) SendSticker(ctx context.Context, chatId int64, s string, messageId int) error {
	_, err := b.client.SendSticker(ctx, &tgbot.SendStickerParams{
		ChatID: chatId,
		Sticker: &models.InputFileString{
			Data: s,
		},
		ReplyParameters: &models.ReplyParameters{
			MessageID: messageId,
		},
	})
	return err
}
