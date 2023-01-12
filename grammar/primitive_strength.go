package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-svparser/ast"
)

//
// A.3.2 Primitive strengths
//

/*
 * pulldown_strength ::=
 *   ( strength0 , strength1 )
 *   | ( strength1 , strength0 )
 *   | ( strength0 )
 */
func PulldownStrength(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.PulldownStrength, error) {
	res := &ast.PulldownStrength{}
	return top(
		token(res,
			parens(
				fn.Alt(
					phrase(
						bind(&res.Zero, Strength0),
						comma,
						bind(&res.One, Strength1)),
					phrase(
						bind(&res.One, Strength1),
						comma,
						bind(&res.Zero, Strength0)),
					bind(&res.Zero, Strength0),
				),
			),
		),
	)(ctx, start)
}

/*
 * pullup_strength ::=
 *   ( strength0 , strength1 )
 *   | ( strength1 , strength0 )
 *   | ( strength1 )
 */
func PullupStrength(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.PullupStrength, error) {
	res := &ast.PullupStrength{}
	return top(token(res,
		parens(
			fn.Alt(
				phrase(
					bind(&res.Zero, Strength0),
					comma,
					bind(&res.One, Strength1)),
				phrase(
					bind(&res.One, Strength1),
					comma,
					bind(&res.Zero, Strength0)),
				bind(&res.One, Strength1),
			),
		),
	))(ctx, start)
}
