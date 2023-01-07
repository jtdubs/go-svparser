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

func TestIdentifier(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		in        string
		want      any
		wantRest  string
		wantError bool
	}{
		{
			in:   "hello",
			want: &ast.SimpleIdentifier{Name: "hello"},
		},
		{
			in:   "_myVar1$",
			want: &ast.SimpleIdentifier{Name: "_myVar1$"},
		},
		{
			in:       "\\_myVar1$ ",
			want:     &ast.EscapedIdentifier{Name: "_myVar1$"},
			wantRest: " ",
		},
	}

	for _, tc := range testCases {
		c := runes.Cursor(tc.in)
		gotRest, got, err := Identifier(ctx, c)
		gotError := (err != nil)

		if gotError != tc.wantError {
			if tc.wantError {
				t.Errorf("Identifier(%q) = %v, want error", tc.in, got)
			} else {
				t.Errorf("Identifier(%q) unexpected error: %v", tc.in, err)
			}
			continue
		}

		if string(gotRest.Rest()) != tc.wantRest {
			t.Errorf("Identifier(%q) rest = %q, want %q", tc.in, string(gotRest.Rest()), tc.wantRest)
			continue
		}

		if diff := cmp.Diff(got, tc.want, cmpopts.IgnoreTypes(ast.Token{}, nom.Span[rune]{})); diff != "" {
			t.Errorf("Identifier(%q) = %v, want %v", tc.in, got, tc.want)
			continue
		}
	}
}
