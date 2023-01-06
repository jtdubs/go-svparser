package grammar

import (
	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

func Whitespace0(start nom.Cursor[rune]) (nom.Cursor[rune], []ast.Whitespace, error) {
	return nom.Many0(Whitespace)(start)
}

func Whitespace1(start nom.Cursor[rune]) (nom.Cursor[rune], []ast.Whitespace, error) {
	return nom.Many1(Whitespace)(start)
}

func Whitespace(start nom.Cursor[rune]) (nom.Cursor[rune], ast.Whitespace, error) {
	return nom.Alt(
		To[ast.Whitespace](Comment),
		To[ast.Whitespace](Spaces),
	)(start)
}

func Spaces(start nom.Cursor[rune]) (nom.Cursor[rune], *ast.Spaces, error) {
	res := &ast.Spaces{}
	return Bake(nom.Value(res, BindSpan(&res.Span, runes.Space0)))(start)
}
