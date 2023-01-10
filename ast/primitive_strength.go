package ast

import (
	"fmt"

	"github.com/jtdubs/go-nom"
)

type PulldownStrength struct {
	nom.Span[rune]
	ZeroT, OneT nom.Span[rune]
	Zero        *Strength0
	One         *Strength1
}

func (u *PulldownStrength) String() string {
	return fmt.Sprintf("PulldownStrength(%v, %v)", u.Zero, u.One)
}

type PullupStrength struct {
	nom.Span[rune]
	ZeroT, OneT nom.Span[rune]
	Zero        *Strength0
	One         *Strength1
}

func (u *PullupStrength) String() string {
	return fmt.Sprintf("PullupStrength(%v, %v)", u.Zero, u.One)
}
