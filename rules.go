package loglint

import (
	"go/ast"
	"go/token"
	"strconv"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/go/analysis"
)

// Проверяет, что лог-сообщения начинаются со строчной буквы.
func checkLowercaseFirstLetter(pass *analysis.Pass, expr ast.Expr) {
	lit, ok := expr.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return
	}

	value, err := strconv.Unquote(lit.Value)
	if err != nil || value == "" {
		return
	}

	first, _ := utf8.DecodeRuneInString(value)
	if unicode.IsLower(first) {
		return
	}

	pass.Reportf(expr.Pos(), "message must start with a lowercase letter")
}

// Проверяет, что лог-сообщения только на английском языке.
func checkEnglishOnly(pass *analysis.Pass, expr ast.Expr) {
	lit, ok := expr.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return
	}

	value, err := strconv.Unquote(lit.Value)
	if err != nil || value == "" {
		return
	}

	for _, r := range value {
		if unicode.IsLetter(r) && !unicode.In(r, unicode.Latin) {
			pass.Reportf(expr.Pos(), "message must be in english")

			return
		}
	}
}
