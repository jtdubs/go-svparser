package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-nom/trace"
	"github.com/jtdubs/go-svparser/ast"
)

func UnaryOperator(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.UnaryOperator, error) {
	res := &ast.UnaryOperator{}
	return trace.Trace(fn.Value(res, BindSpan(&res.Span,
		BindValue(&res.Op,
			fn.Alt(
				fn.Value(ast.UnaryNand, runes.Tag("~&")),
				fn.Value(ast.UnaryNor, runes.Tag("~|")),
				fn.Value(ast.UnaryXnor, runes.Tag("~^")),
				fn.Value(ast.UnaryXnor, runes.Tag("^~")),
				fn.Value(ast.UnaryPlus, runes.Tag("+")),
				fn.Value(ast.UnaryMinus, runes.Tag("-")),
				fn.Value(ast.UnaryLogicalNot, runes.Tag("!")),
				fn.Value(ast.UnaryAnd, runes.Tag("&")),
				fn.Value(ast.UnaryOr, runes.Tag("|")),
				fn.Value(ast.UnaryXor, runes.Tag("^")),
				fn.Value(ast.UnaryBinaryNot, runes.Tag("~")),
			),
		),
	)))(ctx, start)
}
