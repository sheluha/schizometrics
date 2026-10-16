package analyze

import (
	"context"
	"log"

	"github.com/sheluha/schizometrics/internal/model"
)

type MessageAnalyzer struct {
}

func (MessageAnalyzer) Analyze(ctx context.Context, job model.AnalysisJob) {
	log.Printf("Analyzing job %d", job)
}
