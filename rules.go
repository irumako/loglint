package loglint

import (
	"go/ast"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/go/analysis"
)

// Проверяет, что лог-сообщения начинаются со строчной буквы.
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

// Проверяет, что лог-сообщения только на английском языке.
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

// Проверяет, что лог-сообщения не содержат спецсимволы или эмодзи.
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

// Проверяет, что лог-сообщения не содержат чувствительные данные.
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
