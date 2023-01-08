package grammar

import (
	"testing"

	"github.com/jtdubs/go-svparser/ast"
)

func TestWhitespace(t *testing.T) {
	testCases := []testCase[ast.Whitespace]{
		{in: "// hello\nworld", want: &ast.OneLineComment{Text: " hello"}, wantRest: "world"},
		{in: "  // hello\nworld", want: &ast.Spaces{Text: "  "}, wantRest: "// hello\nworld"},
	}

	for _, tc := range testCases {
		validate(t, "Whitespace", Whitespace, tc)
	}
}
