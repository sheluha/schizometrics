package telegram

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

/*
todo
1) Кольцевой буффер для чата
2) Отправка в канал на обработку
*/

type Handler struct{}

func (h *Handler) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {

	if update.Message == nil {
		return
	}

	msg := update.Message

	log.Printf(
		"chat=%d user=%d message=%d text=%q",
		msg.Chat.ID,
		msg.From.ID,
		msg.ID,
		msg.Text,
	)

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: msg.Chat.ID,
		Text:   "Иди нахуй",
	})
	if err != nil {
		log.Print(err)
	}

}
