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

func TestCIdentifier(t *testing.T) {
	testCases := []testCase[*ast.CIdentifier]{
		{in: "hello", want: &ast.CIdentifier{Name: "hello"}},
		{in: "_myVar1", want: &ast.CIdentifier{Name: "_myVar1"}},
	}

	for _, tc := range testCases {
		validate(t, "CIdentifier", CIdentifier, tc)
	}
}

func TestTaskIdentifier(t *testing.T) {
	testCases := []testCase[*ast.TaskIdentifier]{
		{in: "hello", want: &ast.TaskIdentifier{ID: &ast.SimpleIdentifier{Name: "hello"}}},
		{in: "_myVar1", want: &ast.TaskIdentifier{ID: &ast.SimpleIdentifier{Name: "_myVar1"}}},
	}

	for _, tc := range testCases {
		validate(t, "TaskIdentifier", TaskIdentifier, tc)
	}
}
