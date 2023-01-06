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

func TestString(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		in        string
		want      any
		wantRest  string
		wantError bool
	}{
		{
			in:   `""`,
			want: &ast.String{Text: ""},
		},
		{
			in:   `"hello"`,
			want: &ast.String{Text: "hello"},
		},
		{
			in:   `"hello\nworld"`,
			want: &ast.String{Text: "hello\nworld"},
		},
		{
			in:   `"hello\n\"world\\\""`,
			want: &ast.String{Text: "hello\n\"world\\\""},
		},
	}

	for _, tc := range testCases {
		c := runes.Cursor(tc.in)
		gotRest, got, err := String(ctx, c)
		gotError := (err != nil)

		if gotError != tc.wantError {
			if tc.wantError {
				t.Errorf("String(%q) = %v, want error", tc.in, got)
			} else {
				t.Errorf("String(%q) unexpected error: %v", tc.in, err)
			}
			continue
		}

		if string(gotRest.Rest()) != tc.wantRest {
			t.Errorf("String(%q) rest = %q, want %q", tc.in, string(gotRest.Rest()), tc.wantRest)
			continue
		}

		if diff := cmp.Diff(got, tc.want, cmpopts.IgnoreTypes(ast.Token{}, nom.Span[rune]{})); diff != "" {
			t.Errorf("String(%q) = %v, want %v", tc.in, got, tc.want)
			continue
		}
	}
}
