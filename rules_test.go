package loglint

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"golang.org/x/tools/go/analysis"
)

func TestCheckLowercaseFirstLetter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expr     ast.Expr
		wantMsgs []string
	}{
		{
			name:     "lowercase literal",
			expr:     mustParseExpr(t, `"lowercase message"`),
			wantMsgs: nil,
		},
		{
			name:     "uppercase literal",
			expr:     mustParseExpr(t, `"Uppercase message"`),
			wantMsgs: []string{"message must start with a lowercase letter"},
		},
		{
			name:     "empty literal",
			expr:     mustParseExpr(t, `""`),
			wantMsgs: nil,
		},
		{
			name:     "leading whitespace",
			expr:     mustParseExpr(t, `"  lowercase after spaces"`),
			wantMsgs: []string{"message must start with a lowercase letter"},
		},
		{
			name:     "leading punctuation",
			expr:     mustParseExpr(t, `"!lowercase after punctuation"`),
			wantMsgs: []string{"message must start with a lowercase letter"},
		},
		{
			name:     "non string literal",
			expr:     mustParseExpr(t, `123`),
			wantMsgs: nil,
		},
		{
			name:     "non basic literal expression",
			expr:     mustParseExpr(t, `msg`),
			wantMsgs: nil,
		},
		{
			name:     "invalid quoted string",
			expr:     &ast.BasicLit{Kind: token.STRING, Value: `"unterminated`},
			wantMsgs: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := runLowercaseFirstLetterCheck(tt.expr)
			if len(got) != len(tt.wantMsgs) {
				t.Fatalf("got %d diagnostics, want %d: %#v", len(got), len(tt.wantMsgs), got)
			}

			for i, want := range tt.wantMsgs {
				if got[i].Message != want {
					t.Fatalf("diagnostic %d message = %q, want %q", i, got[i].Message, want)
				}
			}
		})
	}
}

func runLowercaseFirstLetterCheck(expr ast.Expr) []analysis.Diagnostic {
	diags := make([]analysis.Diagnostic, 0, 1)
	pass := &analysis.Pass{
		Report: func(d analysis.Diagnostic) {
			diags = append(diags, d)
		},
	}

	checkLowercaseFirstLetter(pass, expr)

	return diags
}

func mustParseExpr(t *testing.T, src string) ast.Expr {
	t.Helper()

	expr, err := parser.ParseExprFrom(token.NewFileSet(), "rules_test.go", src, 0)
	if err != nil {
		t.Fatalf("parse expr %q: %v", src, err)
	}

	return expr
}
