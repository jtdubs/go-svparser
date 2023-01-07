package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-nom/trace"
	"github.com/jtdubs/go-svparser/ast"
)

func Comment(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.Comment, error) {
	return trace.Trace(fn.Alt(
		to[ast.Comment](BlockComment),
		to[ast.Comment](OneLineComment),
	))(ctx, start)
}

func BlockComment(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.BlockComment, error) {
	res := &ast.BlockComment{}
	return tBindSeq(res, &res.Span,
		bindSpan(&res.StartT, runes.Tag("/*")),
		bindSpan(&res.TextT,
			runes.Join(
				fn.First(
					fn.ManyTill(
						fn.Any[rune],
						fn.Peek(runes.Tag("*/")),
					),
				),
			),
		),
		bindSpan(&res.EndT, runes.Tag("*/")),
	)(ctx, start)
}

func OneLineComment(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.OneLineComment, error) {
	res := &ast.OneLineComment{}
	return tBindSeq(res, &res.Span,
		bindSpan(&res.StartT, runes.Tag("//")),
		bindSpan(&res.TextT,
			runes.Join(
				fn.First(
					fn.ManyTill(
						fn.Any[rune],
						fn.Peek(runes.Newline),
					),
				),
			),
		),
		bindSpan(&res.EndT, runes.Newline),
	)(ctx, start)
}
