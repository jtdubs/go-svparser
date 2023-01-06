package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

func Number(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.Number, error) {
	return nom.Alt(
		To[ast.Number](RealNumber),
		To[ast.Number](IntegralNumber),
	)(ctx, start)
}

func IntegralNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.IntegralNumber, error) {
	return nom.Alt(
		To[ast.IntegralNumber](octalNumber),
		To[ast.IntegralNumber](binaryNumber),
		To[ast.IntegralNumber](hexNumber),
		To[ast.IntegralNumber](DecimalNumber),
	)(ctx, start)
}

func DecimalNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.DecimalNumber, error) {
	return nom.Alt(
		To[ast.DecimalNumber](decimalNumberX),
		To[ast.DecimalNumber](decimalNumberZ),
		To[ast.DecimalNumber](decimalNumberUnsigned),
		To[ast.DecimalNumber](UnsignedNumber),
	)(ctx, start)
}

func decimalNumberZ(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DecimalNumberZ, error) {
	res := &ast.DecimalNumberZ{}
	return Bake(nom.Value(res,
		BindSpan(&res.Span,
			nom.Seq(
				nom.Opt(BindSpan(&res.SizeT, size)),
				BindSpan(&res.BaseT, decimalBase),
				BindSpan(&res.Z, zDigit),
				nom.Discard(nom.Many0(runes.Rune('_'))),
			),
		),
	))(ctx, start)
}

func decimalNumberX(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DecimalNumberX, error) {
	res := &ast.DecimalNumberX{}
	return Bake(nom.Value(res,
		BindSpan(&res.Span,
			nom.Seq(
				nom.Opt(BindSpan(&res.SizeT, size)),
				BindSpan(&res.BaseT, decimalBase),
				BindSpan(&res.X, xDigit),
				nom.Discard(nom.Many0(runes.Rune('_'))),
			),
		),
	))(ctx, start)
}

func decimalNumberUnsigned(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DecimalNumberUnsigned, error) {
	res := &ast.DecimalNumberUnsigned{}
	return Bake(nom.Value(res,
		BindSpan(&res.Span,
			nom.Seq(
				nom.Opt(BindSpan(&res.SizeT, size)),
				BindSpan(&res.BaseT, decimalBase),
				BindSpan(&res.ValueT, unsignedNumber),
			),
		),
	))(ctx, start)
}

func binaryNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.BinaryNumber, error) {
	res := &ast.BinaryNumber{}
	return Bake(nom.Value(res,
		BindSpan(&res.Span,
			nom.Seq(
				nom.Opt(BindSpan(&res.SizeT, size)),
				BindSpan(&res.BaseT, binaryBase),
				BindSpan(&res.ValueT, binaryValue),
			),
		),
	))(ctx, start)
}

func octalNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.OctalNumber, error) {
	res := &ast.OctalNumber{}
	return Bake(nom.Value(res,
		BindSpan(&res.Span,
			nom.Seq(
				nom.Opt(BindSpan(&res.SizeT, size)),
				BindSpan(&res.BaseT, octalBase),
				BindSpan(&res.ValueT, octalValue),
			),
		),
	))(ctx, start)
}

func hexNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.HexNumber, error) {
	res := &ast.HexNumber{}
	return Bake(nom.Value(res,
		BindSpan(&res.Span,
			nom.Seq(
				nom.Opt(BindSpan(&res.SizeT, size)),
				BindSpan(&res.BaseT, hexBase),
				BindSpan(&res.ValueT, hexValue),
			),
		),
	))(ctx, start)
}

var sign = runes.Recognize(runes.OneOf("+-"))

func size(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return nonZeroUnsignedNumber(ctx, start)
}

func nonZeroUnsignedNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Cons(
		nonZeroDecimalDigit,
		runes.Join(nom.Many0(nom.Alt(decimalDigit, runes.Rune('_')))),
	)(ctx, start)
}

func RealNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.RealNumber, error) {
	return nom.Alt(
		To[ast.RealNumber](FloatingPointNumber),
		To[ast.RealNumber](FixedPointNumber),
	)(ctx, start)
}

func FloatingPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.FloatingPointNumber, error) {
	res := &ast.FloatingPointNumber{}
	return Bake(nom.Value(res, BindSpan(&res.Span, floatingPointNumber)))(ctx, start)
}

func floatingPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Concat(
		nom.Seq(
			unsignedNumber,
			nom.Opt(nom.Preceded(runes.Rune('.'), unsignedNumber)),
			exp,
			nom.Opt(sign),
			unsignedNumber,
		),
	)(ctx, start)
}

func FixedPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.FixedPointNumber, error) {
	res := &ast.FixedPointNumber{}
	return Bake(nom.Value(res, BindSpan(&res.Span, fixedPointNumber)))(ctx, start)
}

func fixedPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Concat(
		nom.Seq(
			unsignedNumber,
			runes.Tag("."),
			unsignedNumber,
		),
	)(ctx, start)
}

var exp = runes.Recognize(runes.OneOf("eE"))

func UnsignedNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.UnsignedNumber, error) {
	res := &ast.UnsignedNumber{}
	return Bake(nom.Value(res, BindSpan(&res.Span, unsignedNumber)))(ctx, start)
}

func unsignedNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Cons(
		decimalDigit,
		runes.Join(nom.Many0(nom.Alt(decimalDigit, runes.Rune('_')))),
	)(ctx, start)
}

func binaryValue(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Cons(
		binaryDigit,
		runes.Join(nom.Many0(nom.Alt(binaryDigit, runes.Rune('_')))),
	)(ctx, start)
}

func octalValue(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Cons(
		octalDigit,
		runes.Join(nom.Many0(nom.Alt(octalDigit, runes.Rune('_')))),
	)(ctx, start)
}

func hexValue(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Cons(
		hexDigit,
		runes.Join(nom.Many0(nom.Alt(hexDigit, runes.Rune('_')))),
	)(ctx, start)
}

func decimalBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Join(
		nom.Seq(
			runes.Rune('\''),
			nom.Opt(runes.OneOf("sS")),
			runes.OneOf("dD"),
		),
	)(ctx, start)
}

func binaryBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Join(
		nom.Seq(
			runes.Rune('\''),
			nom.Opt(runes.OneOf("sS")),
			runes.OneOf("bB"),
		),
	)(ctx, start)
}

func octalBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Join(
		nom.Seq(
			runes.Rune('\''),
			nom.Opt(runes.OneOf("sS")),
			runes.OneOf("oO"),
		),
	)(ctx, start)
}

func hexBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Join(
		nom.Seq(
			runes.Rune('\''),
			nom.Opt(runes.OneOf("sS")),
			runes.OneOf("hH"),
		),
	)(ctx, start)
}

var nonZeroDecimalDigit = runes.OneOf("123456789")
var decimalDigit = runes.OneOf("0123456789")
var binaryDigit = runes.OneOf("01xXzZ?")
var octalDigit = runes.OneOf("01234567xXzZ?")
var hexDigit = runes.OneOf("0123456789abcdefABCDEFxXzZ?")
var xDigit = runes.OneOf("xX")
var zDigit = runes.OneOf("zZ?")

func UnbasedUnsizedLiteral(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.UnbasedUnsizedLiteral, error) {
	res := &ast.UnbasedUnsizedLiteral{}
	return nom.Value(res, BindSpan(&res.Span, unbasedUnsizedLiteral))(ctx, start)
}

func unbasedUnsizedLiteral(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Join(nom.Seq(runes.Rune('\''), runes.OneOf("01xXzZ")))(ctx, start)
}
