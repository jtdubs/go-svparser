package ast

import (
	"fmt"
	"strconv"

	"github.com/jtdubs/go-nom"
)

type StringLiteral struct {
	nom.Span[rune]
	Text string
}

func (s *StringLiteral) String() string {
	return fmt.Sprintf("StringLiteral(%q)", s.Text)
}

func (s *StringLiteral) Bake() error {
	text, err := strconv.Unquote(string(s.Span.Value()))
	if err != nil {
		return err
	}
	s.Text = text
	return nil
}

func (*StringLiteral) isPrimaryLiteral()     {}
func (*StringLiteral) isPrimary()            {}
func (*StringLiteral) isConstantPrimary()    {}
func (*StringLiteral) isConstantExpression() {}
