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
