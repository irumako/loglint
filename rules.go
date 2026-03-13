package loglint

import (
	"go/ast"
	"go/token"
	"strconv"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/go/analysis"
)

// Проверяет, что лог-сообщения начинаются со строчной буквы
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
