package grammar

import (
	"context"
	"unicode"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/cache"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-nom/trace"
	"github.com/jtdubs/go-svparser/ast"
)

func Identifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.Identifier, error) {
	return trace.Trace(cache.Cache(
		fn.Alt(
			to[ast.Identifier](simpleIdentifier),
			to[ast.Identifier](escapedIdentifier),
		),
	))(ctx, start)
}

func escapedIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.EscapedIdentifier, error) {
	res := &ast.EscapedIdentifier{}
	return tBindSeq(res, &res.Span,
		bindSpan(&res.SlashT, runes.Rune('\\')),
		bindSpan(&res.NameT, fn.Terminated(fn.Many1(asciiPrintNonWS), fn.Peek(fn.Alt(runes.Space)))),
	)(ctx, start)
}

func simpleIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.SimpleIdentifier, error) {
	res := &ast.SimpleIdentifier{}
	return tBind(res, &res.Span,
		fn.Preceded(alpha_, fn.Many0(alphanumeric_S)),
	)(ctx, start)
}

var asciiPrintNonWS = fn.Satisfy(func(r rune) bool {
	return r < 128 && unicode.IsPrint(r) && !unicode.IsSpace(r)
})

var alpha_ = runes.OneOf("abcdefghijklmnoprqstuvwxyzABCDEFGHIJKLMNOPRQSTUVWXYZ_")
var alphanumeric_S = runes.OneOf("abcdefghijklmnoprqstuvwxyzABCDEFGHIJKLMNOPRQSTUVWXYZ0123456789_$")
