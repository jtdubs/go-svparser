package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

func Comment(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.Comment, error) {
	return nom.Alt(
		To[ast.Comment](BlockComment),
		To[ast.Comment](OneLineComment),
	)(ctx, start)
}

func BlockComment(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.BlockComment, error) {
	res := &ast.BlockComment{}
	return Bake(nom.Value(res,
		BindSpan(&res.Span,
			nom.Seq(
				BindSpan(&res.StartT, runes.Tag("/*")),
				BindSpan(&res.TextT,
					runes.Join(
						nom.First(
							nom.ManyTill(
								nom.Any[rune],
								nom.Peek(runes.Tag("*/")),
							),
						),
					),
				),
				BindSpan(&res.EndT, runes.Tag("*/")),
			),
		),
	))(ctx, start)
}

func OneLineComment(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.OneLineComment, error) {
	res := &ast.OneLineComment{}
	return Bake(nom.Value(res,
		BindSpan(&res.Span,
			nom.Seq(
				BindSpan(&res.StartT, runes.Tag("//")),
				BindSpan(&res.TextT,
					runes.Join(
						nom.First(
							nom.ManyTill(
								nom.Any[rune],
								nom.Peek(runes.Newline),
							),
						),
					),
				),
				BindSpan(&res.EndT, runes.Newline),
			),
		),
	))(ctx, start)
}
