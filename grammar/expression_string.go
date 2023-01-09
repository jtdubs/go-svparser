package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-nom/trace"
	"github.com/jtdubs/go-svparser/ast"
)

func String(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.String, error) {
	res := &ast.String{}
	return tBind(res, &res.Span, fn.Surrounded(runes.Rune('"'), runes.Rune('"'), stringContents))(ctx, start)
}

func stringContents(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], []rune, error) {
	return trace.Trace(fn.Many0(
		fn.Alt(
			fn.Preceded(fn.Satisfy(func(r rune) bool { return r == '\\' }), fn.Any[rune]),
			fn.Satisfy(func(r rune) bool { return r != '\\' && r != '"' }),
		),
	))(ctx, start)
}
