package grammar

import (
	"testing"

	"github.com/jtdubs/go-svparser/ast"
)

func TestIdentifier(t *testing.T) {
	testCases := []testCase[ast.Identifier]{
		{in: "hello", want: &ast.SimpleIdentifier{Name: "hello"}},
		{in: "_myVar1$", want: &ast.SimpleIdentifier{Name: "_myVar1$"}},
		{in: "\\_myVar1$ ", want: &ast.EscapedIdentifier{Name: "_myVar1$"}, wantRest: " "},
	}

	for _, tc := range testCases {
		validate(t, "Identifier", Identifier, tc)
	}
}
