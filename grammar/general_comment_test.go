package grammar

import (
	"testing"

	"github.com/jtdubs/go-svparser/ast"
)

func TestComment(t *testing.T) {
	testCases := []testCase[ast.Comment]{
		{in: "// hello world\n", want: &ast.OneLineComment{Text: " hello world"}},
		{in: "/* hello\nworld */", want: &ast.BlockComment{Text: " hello\nworld "}},
	}

	for _, tc := range testCases {
		validate(t, "Comment", Comment, tc)
	}
}
