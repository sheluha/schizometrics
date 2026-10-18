package buffer

import (
	"context"
	"log"

	"github.com/sheluha/schizometrics/internal/model"
)

type Manager struct {
	chatBufferMap map[int64]*Ring[model.Message]
	c             chan<- model.AnalysisJob
}

func NewBufferManager(ch chan<- model.AnalysisJob) Manager {
	return Manager{
		chatBufferMap: make(map[int64]*Ring[model.Message]),
		c:             ch,
	}
}

func (b *Manager) ProcessMessage(ctx context.Context, m model.Message) {
	chatId := m.ChatId
	ring, ok := b.chatBufferMap[chatId]
	if !ok {
		log.Printf("Buffer created for chat %d not found", chatId)
		ring = NewRingBuffer[model.Message](10)
		b.chatBufferMap[chatId] = ring
	}

	job := model.AnalysisJob{
		Target:  m,
		Context: ring.AllValues(),
	}

	ring.Add(m)

	select {
	case b.c <- job:
	case <-ctx.Done():
		return
	}
}
