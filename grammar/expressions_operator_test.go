package grammar

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

func TestUnaryOperator(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		in        string
		want      any
		wantRest  string
		wantError bool
	}{
		{in: "~&", want: &ast.UnaryOperator{Op: ast.UnaryLogicalReductionNand}},
		{in: "~|", want: &ast.UnaryOperator{Op: ast.UnaryLogicalReductionNor}},
		{in: "~^", want: &ast.UnaryOperator{Op: ast.UnaryLogicalReductionXnor}},
		{in: "^~", want: &ast.UnaryOperator{Op: ast.UnaryLogicalReductionXnor}},
		{in: "+", want: &ast.UnaryOperator{Op: ast.UnaryPositive}},
		{in: "-", want: &ast.UnaryOperator{Op: ast.UnaryNegate}},
		{in: "!", want: &ast.UnaryOperator{Op: ast.UnaryLogicalNegation}},
		{in: "&", want: &ast.UnaryOperator{Op: ast.UnaryLogicalReductionAnd}},
		{in: "|", want: &ast.UnaryOperator{Op: ast.UnaryLogicalReductionOr}},
		{in: "^", want: &ast.UnaryOperator{Op: ast.UnaryLogicalReductionXor}},
		{in: "~", want: &ast.UnaryOperator{Op: ast.UnaryLogicalNegation}},
	}

	for _, tc := range testCases {
		c := runes.Cursor(tc.in)
		gotRest, got, err := UnaryOperator(ctx, c)
		gotError := (err != nil)

		if gotError != tc.wantError {
			if tc.wantError {
				t.Errorf("UnaryOperator(%q) = %v, want error", tc.in, got)
			} else {
				t.Errorf("UnaryOperator(%q) unexpected error: %v", tc.in, err)
			}
			continue
		}

		if string(gotRest.Rest()) != tc.wantRest {
			t.Errorf("UnaryOperator(%q) rest = %q, want %q", tc.in, string(gotRest.Rest()), tc.wantRest)
			continue
		}

		if diff := cmp.Diff(got, tc.want, cmpopts.IgnoreTypes(ast.Token{}, nom.Span[rune]{})); diff != "" {
			t.Errorf("UnaryOperator(%q) = %v, want %v", tc.in, got, tc.want)
			continue
		}
	}
}

func TestBinaryOperator(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		in        string
		want      any
		wantRest  string
		wantError bool
	}{
		{in: "<<<", want: &ast.BinaryOperator{Op: ast.BinaryArithmeticShiftLeft}},
		{in: ">>>", want: &ast.BinaryOperator{Op: ast.BinaryArithmeticShiftRight}},
		{in: "<->", want: &ast.BinaryOperator{Op: ast.BinaryLogicalIff}},
		{in: "===", want: &ast.BinaryOperator{Op: ast.BinaryCaseEquals}},
		{in: "!==", want: &ast.BinaryOperator{Op: ast.BinaryCaseNotEquals}},
		{in: "==?", want: &ast.BinaryOperator{Op: ast.BinaryWildcardEquals}},
		{in: "!=?", want: &ast.BinaryOperator{Op: ast.BinaryWildcardNotEquals}},
		{in: "->", want: &ast.BinaryOperator{Op: ast.BinaryLogicalImplies}},
		{in: "**", want: &ast.BinaryOperator{Op: ast.BinaryExp}},
		{in: "^~", want: &ast.BinaryOperator{Op: ast.BinaryBitwiseXnor}},
		{in: "~^", want: &ast.BinaryOperator{Op: ast.BinaryBitwiseXnor}},
		{in: "<<", want: &ast.BinaryOperator{Op: ast.BinaryLogicalShiftLeft}},
		{in: ">>", want: &ast.BinaryOperator{Op: ast.BinaryLogicalShiftRight}},
		{in: "&&", want: &ast.BinaryOperator{Op: ast.BinaryLogicalAnd}},
		{in: "||", want: &ast.BinaryOperator{Op: ast.BinaryLogicalOr}},
		{in: "<=", want: &ast.BinaryOperator{Op: ast.BinaryLessThanEqual}},
		{in: ">=", want: &ast.BinaryOperator{Op: ast.BinaryGreaterThanEqual}},
		{in: "!=", want: &ast.BinaryOperator{Op: ast.BinaryLogicalNotEquals}},
		{in: "==", want: &ast.BinaryOperator{Op: ast.BinaryLogicalEquals}},
		{in: "<", want: &ast.BinaryOperator{Op: ast.BinaryLessThan}},
		{in: ">", want: &ast.BinaryOperator{Op: ast.BinaryGreaterThan}},
		{in: "^", want: &ast.BinaryOperator{Op: ast.BinaryBitwiseXor}},
		{in: "+", want: &ast.BinaryOperator{Op: ast.BinaryAdd}},
		{in: "-", want: &ast.BinaryOperator{Op: ast.BinarySubtract}},
		{in: "*", want: &ast.BinaryOperator{Op: ast.BinaryMultiply}},
		{in: "/", want: &ast.BinaryOperator{Op: ast.BinaryDivide}},
		{in: "%", want: &ast.BinaryOperator{Op: ast.BinaryModulus}},
		{in: "&", want: &ast.BinaryOperator{Op: ast.BinaryBitwiseAnd}},
		{in: "|", want: &ast.BinaryOperator{Op: ast.BinaryBitwiseOr}},
	}

	for _, tc := range testCases {
		c := runes.Cursor(tc.in)
		gotRest, got, err := BinaryOperator(ctx, c)
		gotError := (err != nil)

		if gotError != tc.wantError {
			if tc.wantError {
				t.Errorf("BinaryOperator(%q) = %v, want error", tc.in, got)
			} else {
				t.Errorf("BinaryOperator(%q) unexpected error: %v", tc.in, err)
			}
			continue
		}

		if string(gotRest.Rest()) != tc.wantRest {
			t.Errorf("BinaryOperator(%q) rest = %q, want %q", tc.in, string(gotRest.Rest()), tc.wantRest)
			continue
		}

		if diff := cmp.Diff(got, tc.want, cmpopts.IgnoreTypes(ast.Token{}, nom.Span[rune]{})); diff != "" {
			t.Errorf("BinaryOperator(%q) = %v, want %v", tc.in, got, tc.want)
			continue
		}
	}
}
