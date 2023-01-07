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
		To[ast.Comment](BlockComment),
		To[ast.Comment](OneLineComment),
	))(ctx, start)
}

func BlockComment(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.BlockComment, error) {
	res := &ast.BlockComment{}
	return TBindSeq(res, &res.Span,
		BindSpan(&res.StartT, runes.Tag("/*")),
		BindSpan(&res.TextT,
			runes.Join(
				fn.First(
					fn.ManyTill(
						fn.Any[rune],
						fn.Peek(runes.Tag("*/")),
					),
				),
			),
		),
		BindSpan(&res.EndT, runes.Tag("*/")),
	)(ctx, start)
}

func OneLineComment(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.OneLineComment, error) {
	res := &ast.OneLineComment{}
	return TBindSeq(res, &res.Span,
		BindSpan(&res.StartT, runes.Tag("//")),
		BindSpan(&res.TextT,
			runes.Join(
				fn.First(
					fn.ManyTill(
						fn.Any[rune],
						fn.Peek(runes.Newline),
					),
				),
			),
		),
		BindSpan(&res.EndT, runes.Newline),
	)(ctx, start)
}
