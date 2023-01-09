package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-nom/trace"
	"github.com/jtdubs/go-svparser/ast"
)

//
// A.9.4 White space
//

func Whitespace0(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], []ast.Whitespace, error) {
	return trace.Trace(fn.Many0(Whitespace))(ctx, start)
}

func Whitespace1(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], []ast.Whitespace, error) {
	return trace.Trace(fn.Many1(Whitespace))(ctx, start)
}

func Whitespace(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.Whitespace, error) {
	return trace.Trace(fn.Alt(
		to[ast.Whitespace](Comment),
		to[ast.Whitespace](Spaces),
	))(ctx, start)
}

func Spaces(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.Spaces, error) {
	res := &ast.Spaces{}
	return trace.Trace(bake(fn.Value(res, bindSpan(&res.Span, fn.Many1(whitespace)))))(ctx, start)
}

/*
 * white_space ::= space | tab | newline | eof
 */
var whitespace = runes.OneOf(" \t\r\n")
