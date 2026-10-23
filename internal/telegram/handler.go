package telegram

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/sheluha/schizometrics/internal/buffer"
	"github.com/sheluha/schizometrics/internal/model"
)

type Handler struct {
	bm buffer.Manager
}

func NewHandler(bm buffer.Manager) *Handler {
	return &Handler{
		bm,
	}
}

func (h *Handler) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {

	if update.Message == nil {
		return
	}
	msg := update.Message
	if msg.Chat.Type != "group" && msg.Chat.Type != "supergroup" {
		return
	}
	if msg.From == nil || msg.From.IsBot {
		return
	}
	text := msg.Text
	if text == "" {
		return
	}

	log.Printf(
		"chat=%d user=%d message=%d text=%q",
		msg.Chat.ID,
		msg.From.ID,
		msg.ID,
		text,
	)
	h.bm.ProcessMessage(ctx, msg.Chat.ID, msg.ID, model.Message{
		UserId:  msg.From.ID,
		Message: text,
	})

}
