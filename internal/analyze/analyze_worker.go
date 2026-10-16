package analyze

import (
	"context"

	"github.com/sheluha/schizometrics/internal/model"
)

type AnalyzeWorker struct {
	jobs     chan model.AnalysisJob
	analyzer MessageAnalyzer
}

func NewAnalyzeWorker(jobs chan model.AnalysisJob, analyzer MessageAnalyzer) *AnalyzeWorker {
	return &AnalyzeWorker{
		jobs:     jobs,
		analyzer: analyzer,
	}
}

func (aw *AnalyzeWorker) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-aw.jobs:
			if !ok {
				return
			}
			aw.analyzer.Analyze(ctx, job)
		}
	}
}
