package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

//
// A.2.2.2 Strengths
//

/*
 * drive_strength ::=
 *   ( strength0 , strength1 )
 *   | ( strength1 , strength0 )
 *   | ( strength0 , highz1 )
 *   | ( strength1 , highz0 )
 *   | ( highz0 , strength1 )
 *   | ( highz1 , strength0 )
 */
func DriveStrength(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DriveStrength, error) {
	res := &ast.DriveStrength{}
	return tBind(res, &res.Span,
		parens(
			fn.Alt(
				fn.Seq(
					bindSpan(&res.AT, bindValue(&res.A, to[ast.DriveStrengthOption](Strength0))),
					fn.Discard(fn.Surrounded(Whitespace0, Whitespace0, runes.Rune(','))),
					bindSpan(&res.BT, bindValue(&res.B, to[ast.DriveStrengthOption](Strength1)))),
				fn.Seq(
					bindSpan(&res.AT, bindValue(&res.A, to[ast.DriveStrengthOption](Strength1))),
					fn.Discard(fn.Surrounded(Whitespace0, Whitespace0, runes.Rune(','))),
					bindSpan(&res.BT, bindValue(&res.B, to[ast.DriveStrengthOption](Strength0)))),
				fn.Seq(
					bindSpan(&res.AT, bindValue(&res.A, to[ast.DriveStrengthOption](Strength0))),
					fn.Discard(fn.Surrounded(Whitespace0, Whitespace0, runes.Rune(','))),
					bindSpan(&res.BT, bindValue(&res.B, to[ast.DriveStrengthOption](highZ1)))),
				fn.Seq(
					bindSpan(&res.AT, bindValue(&res.A, to[ast.DriveStrengthOption](Strength1))),
					fn.Discard(fn.Surrounded(Whitespace0, Whitespace0, runes.Rune(','))),
					bindSpan(&res.BT, bindValue(&res.B, to[ast.DriveStrengthOption](highZ0)))),
				fn.Seq(
					bindSpan(&res.AT, bindValue(&res.A, to[ast.DriveStrengthOption](highZ0))),
					fn.Discard(fn.Surrounded(Whitespace0, Whitespace0, runes.Rune(','))),
					bindSpan(&res.BT, bindValue(&res.B, to[ast.DriveStrengthOption](Strength1)))),
				fn.Seq(
					bindSpan(&res.AT, bindValue(&res.A, to[ast.DriveStrengthOption](highZ1))),
					fn.Discard(fn.Surrounded(Whitespace0, Whitespace0, runes.Rune(','))),
					bindSpan(&res.BT, bindValue(&res.B, to[ast.DriveStrengthOption](Strength0)))),
			),
		),
	)(ctx, start)
}

func highZ0(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.HighZ0, error) {
	res := &ast.HighZ0{}
	return tBind(res, &res.Span, runes.TagNoCase("highz0"))(ctx, start)
}

func highZ1(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.HighZ1, error) {
	res := &ast.HighZ1{}
	return tBind(res, &res.Span, runes.TagNoCase("highz1"))(ctx, start)
}

/*
 * strength0 ::= supply0 | strong0 | pull0 | weak0
 */
func Strength0(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.Strength0, error) {
	res := &ast.Strength0{}
	return tBind(res, &res.Span,
		bindValue(&res.Type,
			fn.Alt(
				fn.Value(ast.Supply0, runes.TagNoCase("supply0")),
				fn.Value(ast.Strong0, runes.TagNoCase("strong0")),
				fn.Value(ast.Pull0, runes.TagNoCase("pull0")),
				fn.Value(ast.Weak0, runes.TagNoCase("weak0")),
			),
		),
	)(ctx, start)
}

/*
 * strength1 ::= supply1 | strong1 | pull1 | weak1
 */
func Strength1(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.Strength1, error) {
	res := &ast.Strength1{}
	return tBind(res, &res.Span,
		bindValue(&res.Type,
			fn.Alt(
				fn.Value(ast.Supply1, runes.TagNoCase("supply1")),
				fn.Value(ast.Strong1, runes.TagNoCase("strong1")),
				fn.Value(ast.Pull1, runes.TagNoCase("pull1")),
				fn.Value(ast.Weak1, runes.TagNoCase("weak1")),
			),
		),
	)(ctx, start)
}

/*
 * charge_strength ::= ( small ) | ( medium ) | ( large )
 */
func ChargeStrength(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ChargeStrength, error) {
	res := &ast.ChargeStrength{}
	return tBind(res, &res.Span,
		parens(
			bindSpan(&res.TypeT,
				bindValue(&res.Type,
					fn.Alt(
						fn.Value(ast.Small, runes.TagNoCase("small")),
						fn.Value(ast.Medium, runes.TagNoCase("medium")),
						fn.Value(ast.Large, runes.TagNoCase("large")),
					),
				),
			),
		),
	)(ctx, start)
}
