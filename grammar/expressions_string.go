package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

func String(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.String, error) {
	res := &ast.String{}
	return Bake(nom.Value(res, BindSpan(&res.Span, nom.Surrounded(runes.Rune('"'), runes.Rune('"'), stringContents))))(ctx, start)
}

func stringContents(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], []rune, error) {
	return nom.Many0(
		nom.Alt(
			nom.Preceded(nom.Satisfy(func(r rune) bool { return r == '\\' }), nom.Any[rune]),
			nom.Satisfy(func(r rune) bool { return r != '\\' && r != '"' }),
		),
	)(ctx, start)
}
