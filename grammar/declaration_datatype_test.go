package grammar

import (
	"testing"

	"github.com/jtdubs/go-svparser/ast"
)

func TestDriveStrength(t *testing.T) {
	testCases := []testCase[*ast.DriveStrength]{
		{
			in:   "( pull0, pull1 )",
			want: &ast.DriveStrength{A: &ast.Strength0{Type: ast.StrengthPull0}, B: &ast.Strength1{Type: ast.StrengthPull1}},
		},
		{
			in:   "( supply0, weak1 )",
			want: &ast.DriveStrength{A: &ast.Strength0{Type: ast.StrengthSupply0}, B: &ast.Strength1{Type: ast.StrengthWeak1}},
		},
		{
			in:   "  ( highz0, strong1 )",
			want: &ast.DriveStrength{A: &ast.HighZ0{}, B: &ast.Strength1{Type: ast.StrengthStrong1}},
		},
	}

	for _, tc := range testCases {
		validate(t, "DriveStrength", DriveStrength, tc)
	}
}
