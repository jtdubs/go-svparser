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

func TestNumber(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		in        string
		want      any
		wantRest  string
		wantError bool
	}{
		{
			in:   "123",
			want: &ast.UnsignedNumber{Value: 123},
		},
		{
			in:   "32'd42",
			want: &ast.DecimalNumberUnsigned{Size: 32, Value: 42},
		},
		{
			in:   "32'b101",
			want: &ast.BinaryNumber{Value: ast.MaskedInt{Size: 32, Base: 2, V: 5}},
		},
		{
			in:   "32'hF0",
			want: &ast.HexNumber{Value: ast.MaskedInt{Size: 32, Base: 16, V: 240}},
		},
		{
			in:   "32'o77",
			want: &ast.OctalNumber{Value: ast.MaskedInt{Size: 32, Base: 8, V: 63}},
		},
		{
			in:   "32'd100_002",
			want: &ast.DecimalNumberUnsigned{Size: 32, Value: 100002},
		},
		{
			in:   "123.456",
			want: &ast.FixedPointNumber{Value: 123.456},
		},
		{
			in:   "123_456.789",
			want: &ast.FixedPointNumber{Value: 123456.789},
		},
		{
			in:   "123.456e7",
			want: &ast.FloatingPointNumber{Value: 123.456e7},
		},
		{
			in:   "123.456e-2",
			want: &ast.FloatingPointNumber{Value: 123.456e-2},
		},
		{
			in:   "32'hFX0F8ZZ2",
			want: &ast.HexNumber{Value: ast.MaskedInt{Size: 32, Base: 16, V: 0xF00F8002, X: 0x0F000000, Z: 0x00000FF0}},
		},
	}

	for _, tc := range testCases {
		c := runes.Cursor(tc.in)
		gotRest, got, err := Number(ctx, c)
		gotError := (err != nil)

		if gotError != tc.wantError {
			if tc.wantError {
				t.Errorf("Number(%q) = %v, want error", tc.in, got)
			} else {
				t.Errorf("Number(%q) unexpected error: %v", tc.in, err)
			}
			continue
		}

		if string(gotRest.Rest()) != tc.wantRest {
			t.Errorf("Number(%q) rest = %q, want %q", tc.in, string(gotRest.Rest()), tc.wantRest)
			continue
		}

		if diff := cmp.Diff(got, tc.want, cmpopts.IgnoreTypes(ast.Token{}, nom.Span[rune]{})); diff != "" {
			t.Errorf("Number(%q) = %v, want %v", tc.in, got, tc.want)
			continue
		}
	}
}
