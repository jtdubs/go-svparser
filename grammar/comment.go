package grammar

import (
	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

func Comment(start nom.Cursor[rune]) (nom.Cursor[rune], ast.Comment, error) {
	return nom.Alt(
		To[ast.Comment](BlockComment),
		To[ast.Comment](OneLineComment),
	)(start)
}

func BlockComment(start nom.Cursor[rune]) (nom.Cursor[rune], *ast.BlockComment, error) {
	res := &ast.BlockComment{}
	return Bake(nom.Value(res,
		Bind(&res.Token,
			nom.Seq(
				Bind(&res.StartT, runes.Tag("/*")),
				Bind(&res.TextT,
					runes.Join(
						nom.First(
							nom.ManyTill(
								nom.Any[rune],
								nom.Peek(runes.Tag("*/")),
							),
						),
					),
				),
				Bind(&res.EndT, runes.Tag("*/")),
			),
		),
	))(start)
}

func OneLineComment(start nom.Cursor[rune]) (nom.Cursor[rune], *ast.OneLineComment, error) {
	res := &ast.OneLineComment{}
	return Bake(nom.Value(res,
		Bind(&res.Token,
			nom.Seq(
				Bind(&res.StartT, runes.Tag("//")),
				Bind(&res.TextT,
					runes.Join(
						nom.First(
							nom.ManyTill(
								nom.Any[rune],
								nom.Peek(runes.Newline()),
							),
						),
					),
				),
				Bind(&res.EndT, runes.Newline()),
			),
		),
	))(start)
}
