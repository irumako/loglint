package loglint

import (
	"go/ast"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/go/analysis"
)

type ruleFunc func(pass *analysis.Pass, expr ast.Expr)

type rule struct {
	id    string
	check ruleFunc
}

var allRules = []rule{
	{id: "lowercase-first-letter", check: checkLowercaseFirstLetter},
	{id: "english-only", check: checkEnglishOnly},
	{id: "special-symbols", check: checkSpecialSymbolAndEmoji},
	{id: "sensitive-data", check: checkSensitiveData},
}

func checkLowercaseFirstLetter(pass *analysis.Pass, expr ast.Expr) {
	value, ok := getStringLiteralValue(expr)
	if !ok || value == "" {
		return
	}

	first, _ := utf8.DecodeRuneInString(value)
	if unicode.IsLower(first) {
		return
	}

	pass.Reportf(expr.Pos(), "message must start with a lowercase letter")
}

func checkEnglishOnly(pass *analysis.Pass, expr ast.Expr) {
	value, ok := getStringLiteralValue(expr)
	if !ok || value == "" {
		return
	}

	for _, r := range value {
		if unicode.IsLetter(r) && !unicode.In(r, unicode.Latin) {
			pass.Reportf(expr.Pos(), "message must be in english")

			return
		}
	}
}

func checkSpecialSymbolAndEmoji(pass *analysis.Pass, expr ast.Expr) {
	value, ok := getStringLiteralValue(expr)
	if !ok || value == "" {
		return
	}

	for _, r := range value {
		if isSpecialSymbol(r) || isEmoji(r) {
			pass.Reportf(expr.Pos(), "message should not contain special symbols or emojis")

			return
		}
	}
}

func checkSensitiveData(pass *analysis.Pass, expr ast.Expr) {
	value, ok := getStringLiteralValue(expr)
	if !ok {
		pass.Reportf(expr.Pos(), "message may contain potentially sensitive data")

		return
	}

	value = strings.ToLower(value)
	for _, keyword := range sensitiveDataKeywords {
		if strings.Contains(value, keyword) {
			pass.Reportf(expr.Pos(), "message may contain potentially sensitive data")

			return
		}
	}
}
