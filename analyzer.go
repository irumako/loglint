// Package loglint provides a Go analysis pass for validating log usage.
package loglint

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// NewAnalyzer constructs the loglint analyzer.
func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "loglint",
		Doc:  "check log statements",
		Run:  run,
	}
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(_ ast.Node) bool {
			return true
		})
	}

	return nil, nil
}
