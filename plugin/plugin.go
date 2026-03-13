// Package plugin exposes the loglint analyzers for golangci-lint integration.
package plugin

import (
	"golang.org/x/tools/go/analysis"

	"github.com/irumako/loglint"
)

// New returns the analyzers exposed by the loglint plugin.
func New(_ any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		loglint.NewAnalyzer(),
	}, nil
}
