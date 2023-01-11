package grammar

import (
	"testing"

	"github.com/jtdubs/go-svparser/ast"
)

func TestString(t *testing.T) {
	testCases := []testCase[*ast.StringLiteral]{
		{in: `""`, want: &ast.StringLiteral{Text: ""}},
		{in: `"hello"`, want: &ast.StringLiteral{Text: "hello"}},
		{in: `  "hello"`, want: &ast.StringLiteral{Text: "hello"}},
		{in: `"hello\nworld"`, want: &ast.StringLiteral{Text: "hello\nworld"}},
		{in: `"hello\n\"world\\\"world"`, want: &ast.StringLiteral{Text: "hello\n\"world\\\"world"}},
	}

	for _, tc := range testCases {
		validate(t, "StringLiteral", StringLiteral, tc)
	}
}
