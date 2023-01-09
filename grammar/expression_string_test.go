package grammar

import (
	"testing"

	"github.com/jtdubs/go-svparser/ast"
)

func TestString(t *testing.T) {
	testCases := []testCase[*ast.String]{
		{in: `""`, want: &ast.String{Text: ""}},
		{in: `"hello"`, want: &ast.String{Text: "hello"}},
		{in: `"hello\nworld"`, want: &ast.String{Text: "hello\nworld"}},
		{in: `"hello\n\"world\\\"world"`, want: &ast.String{Text: "hello\n\"world\\\"world"}},
	}

	for _, tc := range testCases {
		validate(t, "String", String, tc)
	}
}
