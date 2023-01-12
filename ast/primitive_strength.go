package ast

import (
	"fmt"

	"github.com/jtdubs/go-nom"
)

type PulldownStrength struct {
	Token
	ZeroT, OneT nom.Span[rune]
	Zero        *Strength0
	One         *Strength1
}

func (u *PulldownStrength) String() string {
	return fmt.Sprintf("PulldownStrength(%v, %v)", u.Zero, u.One)
}

type PullupStrength struct {
	Token
	ZeroT, OneT nom.Span[rune]
	Zero        *Strength0
	One         *Strength1
}

func (u *PullupStrength) String() string {
	return fmt.Sprintf("PullupStrength(%v, %v)", u.Zero, u.One)
}
