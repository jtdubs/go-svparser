package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

func Whitespace0(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], []ast.Whitespace, error) {
	return nom.Many0(Whitespace)(ctx, start)
}

func Whitespace1(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], []ast.Whitespace, error) {
	return nom.Many1(Whitespace)(ctx, start)
}

func Whitespace(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.Whitespace, error) {
	return nom.Alt(
		To[ast.Whitespace](Comment),
		To[ast.Whitespace](Spaces),
	)(ctx, start)
}

func Spaces(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.Spaces, error) {
	res := &ast.Spaces{}
	return Bake(nom.Value(res, BindSpan(&res.Span, runes.Space0)))(ctx, start)
}
