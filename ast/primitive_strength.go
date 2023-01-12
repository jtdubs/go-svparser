package ast

import (
	"fmt"
)

type PulldownStrength struct {
	Token
	Zero *Strength0
	One  *Strength1
}

func (u *PulldownStrength) String() string {
	return fmt.Sprintf("PulldownStrength(%v, %v)", u.Zero, u.One)
}

type PullupStrength struct {
	Token
	Zero *Strength0
	One  *Strength1
}

func (u *PullupStrength) String() string {
	return fmt.Sprintf("PullupStrength(%v, %v)", u.Zero, u.One)
}
