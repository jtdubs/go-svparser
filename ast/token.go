package ast

import (
	"fmt"

	"github.com/jtdubs/go-nom"
)

type Token struct {
	nom.Span[rune]
}

func (t Token) String() string {
	return fmt.Sprintf("Token(%q)", string(t.Start.To(t.End)))
}

func (t Token) Value() string {
	return string(t.Start.To(t.End))
}
