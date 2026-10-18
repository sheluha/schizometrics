package analyze

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/sheluha/schizometrics/internal/llm"
	"github.com/sheluha/schizometrics/internal/model"
)

type MessageAnalyzer struct {
	client *llm.Client
}

func NewMessageAnalyzer(client *llm.Client) MessageAnalyzer {
	return MessageAnalyzer{
		client: client,
	}
}

func (ma MessageAnalyzer) Analyze(ctx context.Context, job model.AnalysisJob) {
	resp, err := ma.client.SendRequest(ctx, buildContext(job), llm.DefaultQuestions())
	if err != nil {
		log.Printf("Error sending request, %w", err)
		return
	}
	log.Println(resp)
}

func buildContext(job model.AnalysisJob) string {
	sb := strings.Builder{}
	sb.WriteString("Контекст чата: несколько участников переписываются. Нужно оценить только последнее сообщение автора с учетом соседних реплик. Формат - Автор:Сообщение;\n")
	for _, v := range job.Context {
		if v.Message == "" {
			continue
		}
		sb.WriteString(strconv.FormatInt(v.UserId, 10))
		sb.WriteString(":")
		sb.WriteString(v.Message)
		sb.WriteString(";")
	}
	sb.WriteString("\nСообщение автора: ")
	sb.WriteString(strconv.FormatInt(job.Target.UserId, 10))
	sb.WriteString(":")
	sb.WriteString(job.Target.Message)
	return sb.String()
}
