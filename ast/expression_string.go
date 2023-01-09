package ast

import (
	"fmt"
	"strconv"

	"github.com/jtdubs/go-nom"
)

type String struct {
	nom.Span[rune]
	Text string
}

func (s *String) String() string {
	return fmt.Sprintf("String(%q)", s.Text)
}

func (s *String) Bake() error {
	text, err := strconv.Unquote(string(s.Span.Value()))
	if err != nil {
		return err
	}
	s.Text = text
	return nil
}
