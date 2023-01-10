package grammar

import (
	"testing"

	"github.com/jtdubs/go-svparser/ast"
)

func TestDriveStrength(t *testing.T) {
	testCases := []testCase[*ast.DriveStrength]{
		{
			in:   "( pull0, pull1 )",
			want: &ast.DriveStrength{A: &ast.Strength0{Type: ast.Pull0}, B: &ast.Strength1{Type: ast.Pull1}},
		},
		{
			in:   "( supply0, weak1 )",
			want: &ast.DriveStrength{A: &ast.Strength0{Type: ast.Supply0}, B: &ast.Strength1{Type: ast.Weak1}},
		},
		{
			in:   "( highz0, strong1 )",
			want: &ast.DriveStrength{A: &ast.HighZ0{}, B: &ast.Strength1{Type: ast.Strong1}},
		},
	}

	for _, tc := range testCases {
		validate(t, "DriveStrength", DriveStrength, tc)
	}
}
