package rules

import (
	"cmp"
	"log"
	"slices"

	"github.com/sheluha/schizometrics/internal/llm"
)

type Engine struct {
	rules []Rule
}

func NewRulesEngine(rules []Rule) *Engine {
	return &Engine{
		rules: rules,
	}
}

func (e *Engine) RunRules(dto map[llm.Feature]llm.Answer) (bool, Rule) {
	actions := make([]Rule, 0)
	for _, rule := range e.rules {
		if !rule.Enabled {
			continue
		}
		var sum float64
		for k, v := range rule.Weights {
			sum += *dto[k].Noul * v
		}
		if sum >= rule.Threshold {
			actions = append(actions, rule)
		}
		log.Printf("%v sum = %v", rule.ID, sum)
	}
	if len(actions) == 0 {
		return false, Rule{}
	}

	return true, slices.MaxFunc(actions, func(a, b Rule) int {
		return cmp.Compare(a.Priority, b.Priority)
	})
}
