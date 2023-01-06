package grammar

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

func TestWhitespace(t *testing.T) {
	testCases := []struct {
		in        string
		want      any
		wantRest  string
		wantError bool
	}{
		{
			in:       "// hello\nworld",
			want:     &ast.OneLineComment{Text: " hello"},
			wantRest: "world",
		},
		{
			in:       "  // hello\nworld",
			want:     &ast.Spaces{Text: "  "},
			wantRest: "// hello\nworld",
		},
	}

	for _, tc := range testCases {
		c := runes.Cursor(tc.in)
		gotRest, got, err := Whitespace(c)
		gotError := (err != nil)

		if gotError != tc.wantError {
			if tc.wantError {
				t.Errorf("Whitespace(%q) = %v, want error", tc.in, got)
			} else {
				t.Errorf("Whitespace(%q) unexpected error: %v", tc.in, err)
			}
			continue
		}

		if string(gotRest.Rest()) != tc.wantRest {
			t.Errorf("Whitespace(%q) rest = %q, want %q", tc.in, string(gotRest.Rest()), tc.wantRest)
			continue
		}

		if diff := cmp.Diff(got, tc.want, cmpopts.IgnoreTypes(ast.Token{}, nom.Span[rune]{})); diff != "" {
			t.Errorf("Whitespace(%q) = %v, want %v", tc.in, got, tc.want)
			continue
		}
	}
}
