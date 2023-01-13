package grammar

import (
	"testing"

	"github.com/jtdubs/go-svparser/ast"
)

func TestAttributeInstance(t *testing.T) {
	testCases := []testCase[*ast.AttributeInstance]{
		{
			in: "\t(* foo , bar = 42 *)",
			want: &ast.AttributeInstance{
				Specs: []*ast.AttrSpec{
					{
						Name: &ast.AttrName{ID: &ast.SimpleIdentifier{Name: "foo"}},
					},
					{
						Name: &ast.AttrName{ID: &ast.SimpleIdentifier{Name: "bar"}},
						Expr: &ast.UnsignedNumber{Value: 42},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		validate(t, "AttributeInstance", AttributeInstance, tc)
	}
}

func TestAttrSpec(t *testing.T) {
	testCases := []testCase[*ast.AttrSpec]{
		{
			in:   "foo",
			want: &ast.AttrSpec{Name: &ast.AttrName{ID: &ast.SimpleIdentifier{Name: "foo"}}},
		},
		{
			in:   "\tfoo = 42",
			want: &ast.AttrSpec{Name: &ast.AttrName{ID: &ast.SimpleIdentifier{Name: "foo"}}, Expr: &ast.UnsignedNumber{Value: 42}},
		},
	}

	for _, tc := range testCases {
		validate(t, "AttrSpec", AttrSpec, tc)
	}
}

func TestComment(t *testing.T) {
	testCases := []testCase[ast.Comment]{
		{in: "// hello world\n", want: &ast.OneLineComment{Text: " hello world"}},
		{in: "/* hello\nworld */", want: &ast.BlockComment{Text: " hello\nworld "}},
	}

	for _, tc := range testCases {
		validate(t, "Comment", Comment, tc)
	}
}

func TestIdentifier(t *testing.T) {
	testCases := []testCase[ast.Identifier]{
		{in: "hello", want: &ast.SimpleIdentifier{Name: "hello"}},
		{in: "\thello", want: &ast.SimpleIdentifier{Name: "hello"}},
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
		{in: "\nhello", want: &ast.CIdentifier{Name: "hello"}},
		{in: "_myVar1", want: &ast.CIdentifier{Name: "_myVar1"}},
	}

	for _, tc := range testCases {
		validate(t, "CIdentifier", CIdentifier, tc)
	}
}

func TestTaskIdentifier(t *testing.T) {
	testCases := []testCase[*ast.TaskIdentifier]{
		{in: "hello", want: &ast.TaskIdentifier{ID: &ast.SimpleIdentifier{Name: "hello"}}},
		{in: "  hello", want: &ast.TaskIdentifier{ID: &ast.SimpleIdentifier{Name: "hello"}}},
		{in: "_myVar1", want: &ast.TaskIdentifier{ID: &ast.SimpleIdentifier{Name: "_myVar1"}}},
	}

	for _, tc := range testCases {
		validate(t, "TaskIdentifier", TaskIdentifier, tc)
	}
}

func TestArrayIdentifier(t *testing.T) {
	testCases := []testCase[*ast.ArrayIdentifier]{
		{in: "hello", want: &ast.ArrayIdentifier{ID: &ast.SimpleIdentifier{Name: "hello"}}},
		{in: "_myVar1", want: &ast.ArrayIdentifier{ID: &ast.SimpleIdentifier{Name: "_myVar1"}}},
	}

	for _, tc := range testCases {
		validate(t, "ArrayIdentifier", ArrayIdentifier, tc)
	}
}

func TestBlockIdentifier(t *testing.T) {
	testCases := []testCase[*ast.BlockIdentifier]{
		{in: "hello", want: &ast.BlockIdentifier{ID: &ast.SimpleIdentifier{Name: "hello"}}},
		{in: "_myVar1", want: &ast.BlockIdentifier{ID: &ast.SimpleIdentifier{Name: "_myVar1"}}},
	}

	for _, tc := range testCases {
		validate(t, "BlockIdentifier", BlockIdentifier, tc)
	}
}

func TestBinIdentifier(t *testing.T) {
	testCases := []testCase[*ast.BinIdentifier]{
		{in: "hello", want: &ast.BinIdentifier{ID: &ast.SimpleIdentifier{Name: "hello"}}},
		{in: "_myVar1", want: &ast.BinIdentifier{ID: &ast.SimpleIdentifier{Name: "_myVar1"}}},
	}

	for _, tc := range testCases {
		validate(t, "BinIdentifier", BinIdentifier, tc)
	}
}

func TestHierarchicalIdentifier(t *testing.T) {
	testCases := []testCase[*ast.HierarchicalIdentifier]{
		{
			in: "$root.foo[42].bar",
			want: &ast.HierarchicalIdentifier{
				Root: true,
				Parts: []*ast.HierarchicalIdentifierPart{
					{
						ID:   &ast.SimpleIdentifier{Name: "foo"},
						Bits: &ast.ConstantBitSelect{Exprs: []ast.ConstantExpression{&ast.UnsignedNumber{Value: 42}}},
					},
					{
						ID: &ast.SimpleIdentifier{Name: "bar"},
					},
				},
			},
		},
		{
			in: "foo[42].bar",
			want: &ast.HierarchicalIdentifier{
				Parts: []*ast.HierarchicalIdentifierPart{
					{
						ID:   &ast.SimpleIdentifier{Name: "foo"},
						Bits: &ast.ConstantBitSelect{Exprs: []ast.ConstantExpression{&ast.UnsignedNumber{Value: 42}}},
					},
					{
						ID: &ast.SimpleIdentifier{Name: "bar"},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		validate(t, "HierarchicalIdentifier", HierarchicalIdentifier, tc)
	}
}
