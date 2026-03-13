// Package loglint provides a Go analysis pass for validating log usage.
package loglint

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/types/typeutil"
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

			// Правила
			checkLowercaseFirstLetter(pass, msgExpr)
			checkEnglishOnly(pass, msgExpr)
			checkSpecialSymbolAndEmoji(pass, msgExpr)

			return true
		})
	}

	return nil, nil
}
