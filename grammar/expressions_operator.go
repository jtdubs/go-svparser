package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

func UnaryOperator(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.UnaryOperator, error) {
	res := &ast.UnaryOperator{}
	return nom.Value(res, BindSpan(&res.Span,
		BindValue(&res.Op,
			nom.Alt(
				nom.Value(ast.UnaryNand, runes.Tag("~&")),
				nom.Value(ast.UnaryNor, runes.Tag("~|")),
				nom.Value(ast.UnaryXnor, runes.Tag("~^")),
				nom.Value(ast.UnaryXnor, runes.Tag("^~")),
				nom.Value(ast.UnaryPlus, runes.Tag("+")),
				nom.Value(ast.UnaryMinus, runes.Tag("-")),
				nom.Value(ast.UnaryLogicalNot, runes.Tag("!")),
				nom.Value(ast.UnaryAnd, runes.Tag("&")),
				nom.Value(ast.UnaryOr, runes.Tag("|")),
				nom.Value(ast.UnaryXor, runes.Tag("^")),
				nom.Value(ast.UnaryBinaryNot, runes.Tag("~")),
			),
		),
	))(ctx, start)
}
