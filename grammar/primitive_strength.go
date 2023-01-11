package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
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
	return tBind(res, &res.Span,
		parens(
			fn.Alt(
				phrase(
					bindSpan(&res.ZeroT, bindValue(&res.Zero, Strength0)),
					fn.Discard(runes.Rune(',')),
					bindSpan(&res.OneT, bindValue(&res.One, Strength1))),
				phrase(
					bindSpan(&res.OneT, bindValue(&res.One, Strength1)),
					fn.Discard(runes.Rune(',')),
					bindSpan(&res.ZeroT, bindValue(&res.Zero, Strength0))),
				bindSpan(&res.ZeroT, bindValue(&res.Zero, Strength0)),
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
	return tBind(res, &res.Span,
		parens(
			fn.Alt(
				phrase(
					bindSpan(&res.ZeroT, bindValue(&res.Zero, Strength0)),
					fn.Discard(runes.Rune(',')),
					bindSpan(&res.OneT, bindValue(&res.One, Strength1))),
				phrase(
					bindSpan(&res.OneT, bindValue(&res.One, Strength1)),
					fn.Discard(runes.Rune(',')),
					bindSpan(&res.ZeroT, bindValue(&res.Zero, Strength0))),
				bindSpan(&res.ZeroT, bindValue(&res.One, Strength1)),
			),
		),
	)(ctx, start)
}
