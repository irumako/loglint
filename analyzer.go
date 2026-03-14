// Package loglint provides a Go analysis pass for validating log usage.
package loglint

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/types/typeutil"

	"github.com/irumako/loglint/config"
)

// NewAnalyzer constructs the loglint analyzer.
func NewAnalyzer() *analysis.Analyzer {
	cfg, err := config.New()

	var rules []rule
	if err == nil {
		rules = loadActiveRules(cfg)
	}

	return &analysis.Analyzer{
		Name: "loglint",
		Doc:  "check log statements",
		Run:  run(rules, err),
	}
}

func run(rules []rule, initErr error) func(pass *analysis.Pass) (any, error) {
	return func(pass *analysis.Pass) (any, error) {
		if initErr != nil {
			return nil, initErr
		}

		for _, file := range pass.Files {
			ast.Inspect(file, func(node ast.Node) bool {
				call, ok := node.(*ast.CallExpr)
				if !ok {
					return true
				}

				fn := typeutil.StaticCallee(pass.TypesInfo, call)
				if fn == nil {
					return true
				}

				pos, ok := messageArgPosByFunction[fn.FullName()]
				if !ok {
					return true
				}

				if int(pos) >= len(call.Args) {
					return true
				}

				msgExpr := call.Args[pos]
				for _, rule := range rules {
					rule.check(pass, msgExpr)
				}

				return true
			})
		}

		return nil, nil
	}
}

func loadActiveRules(cfg *config.Config) []rule {
	disabled := make(map[string]struct{}, len(cfg.DisabledRules))
	for _, id := range cfg.DisabledRules {
		disabled[id] = struct{}{}
	}

	activeRules := make([]rule, 0, len(allRules))
	for _, rule := range allRules {
		if _, ok := disabled[rule.id]; ok {
			continue
		}

		activeRules = append(activeRules, rule)
	}

	return activeRules
}
