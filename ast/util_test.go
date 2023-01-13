package ast

import "testing"

func TestMaskedIntString(t *testing.T) {
	testCases := []struct {
		in   MaskedInt
		want string
	}{
		{
			in:   MaskedInt{Base: 16, Size: 32, V: 0x00000000, X: 0x00000000, Z: 0x00000000},
			want: "32'h00000000",
		},
		{
			in:   MaskedInt{Base: 16, Size: 32, V: 0x12345678, X: 0x00F00F00, Z: 0xF00F00F0},
			want: "32'hZ2XZ5XZ8",
		},
		{
			in:   MaskedInt{Base: 16, Size: 16, V: 0x12345678, X: 0x00F00F00, Z: 0xF00F00F0},
			want: "16'h5XZ8",
		},
		{
			in:   MaskedInt{Base: 2, Size: 8, V: 0b11011101, X: 0b01000010, Z: 0b00011000},
			want: "8'b1X0ZZ1X1",
		},
	}

	for _, tc := range testCases {
		if got := tc.in.String(); got != tc.want {
			t.Errorf("MaskedInt(%#v).String() = %q, want %q", tc.in, got, tc.want)
		}
	}
}
