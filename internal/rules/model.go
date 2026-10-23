package rules

import (
	"github.com/sheluha/schizometrics/internal/llm"
)

type Rule struct {
	ID        string                  `yaml:"id"`
	Enabled   bool                    `yaml:"enabled"`
	Priority  int                     `yaml:"priority"`
	Weights   map[llm.Feature]float64 `yaml:"weights"`
	Threshold float64                 `yaml:"threshold"`
	Sticker   string                  `yaml:"sticker"`
}
