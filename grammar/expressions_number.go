package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/cache"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-nom/trace"
	"github.com/jtdubs/go-svparser/ast"
)

func Number(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.Number, error) {
	return trace.Trace(fn.Alt(
		To[ast.Number](RealNumber),
		To[ast.Number](IntegralNumber),
	))(ctx, start)
}

func IntegralNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.IntegralNumber, error) {
	return trace.Trace(fn.Alt(
		To[ast.IntegralNumber](octalNumber),
		To[ast.IntegralNumber](binaryNumber),
		To[ast.IntegralNumber](hexNumber),
		To[ast.IntegralNumber](DecimalNumber),
	))(ctx, start)
}

func DecimalNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.DecimalNumber, error) {
	return trace.Trace(fn.Alt(
		To[ast.DecimalNumber](decimalNumberX),
		To[ast.DecimalNumber](decimalNumberZ),
		To[ast.DecimalNumber](decimalNumberUnsigned),
		To[ast.DecimalNumber](UnsignedNumber),
	))(ctx, start)
}

func decimalNumberZ(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DecimalNumberZ, error) {
	res := &ast.DecimalNumberZ{}
	return trace.Trace(Bake(fn.Value(res,
		BindSpan(&res.Span,
			fn.Seq(
				fn.Opt(BindSpan(&res.SizeT, size)),
				BindSpan(&res.BaseT, decimalBase),
				BindSpan(&res.Z, zDigit),
				fn.Discard(fn.Many0(runes.Rune('_'))),
			),
		),
	)))(ctx, start)
}

func decimalNumberX(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DecimalNumberX, error) {
	res := &ast.DecimalNumberX{}
	return trace.Trace(Bake(fn.Value(res,
		BindSpan(&res.Span,
			fn.Seq(
				fn.Opt(BindSpan(&res.SizeT, size)),
				BindSpan(&res.BaseT, decimalBase),
				BindSpan(&res.X, xDigit),
				fn.Discard(fn.Many0(runes.Rune('_'))),
			),
		),
	)))(ctx, start)
}

func decimalNumberUnsigned(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DecimalNumberUnsigned, error) {
	res := &ast.DecimalNumberUnsigned{}
	return trace.Trace(Bake(fn.Value(res,
		BindSpan(&res.Span,
			fn.Seq(
				fn.Opt(BindSpan(&res.SizeT, size)),
				BindSpan(&res.BaseT, decimalBase),
				BindSpan(&res.ValueT, unsignedNumber),
			),
		),
	)))(ctx, start)
}

func binaryNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.BinaryNumber, error) {
	res := &ast.BinaryNumber{}
	return trace.Trace(Bake(fn.Value(res,
		BindSpan(&res.Span,
			fn.Seq(
				fn.Opt(BindSpan(&res.SizeT, size)),
				BindSpan(&res.BaseT, binaryBase),
				BindSpan(&res.ValueT, binaryValue),
			),
		),
	)))(ctx, start)
}

func octalNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.OctalNumber, error) {
	res := &ast.OctalNumber{}
	return trace.Trace(Bake(fn.Value(res,
		BindSpan(&res.Span,
			fn.Seq(
				fn.Opt(BindSpan(&res.SizeT, size)),
				BindSpan(&res.BaseT, octalBase),
				BindSpan(&res.ValueT, octalValue),
			),
		),
	)))(ctx, start)
}

func hexNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.HexNumber, error) {
	res := &ast.HexNumber{}
	return trace.Trace(Bake(fn.Value(res,
		BindSpan(&res.Span,
			fn.Seq(
				fn.Opt(BindSpan(&res.SizeT, size)),
				BindSpan(&res.BaseT, hexBase),
				BindSpan(&res.ValueT, hexValue),
			),
		),
	)))(ctx, start)
}

var sign = runes.Recognize(runes.OneOf("+-"))

func size(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return trace.Trace(nonZeroUnsignedNumber)(ctx, start)
}

func nonZeroUnsignedNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return trace.Trace(runes.Cons(
		nonZeroDecimalDigit,
		runes.Join(fn.Many0(decimalDigit_)),
	))(ctx, start)
}

func RealNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.RealNumber, error) {
	return trace.Trace(fn.Alt(
		To[ast.RealNumber](FloatingPointNumber),
		To[ast.RealNumber](FixedPointNumber),
	))(ctx, start)
}

func FloatingPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.FloatingPointNumber, error) {
	res := &ast.FloatingPointNumber{}
	return trace.Trace(Bake(fn.Value(res, BindSpan(&res.Span, floatingPointNumber))))(ctx, start)
}

func floatingPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return trace.Trace(runes.Concat(
		fn.Seq(
			unsignedNumber,
			fn.Opt(fn.Preceded(runes.Rune('.'), unsignedNumber)),
			exp,
			fn.Opt(sign),
			unsignedNumber,
		),
	))(ctx, start)
}

func FixedPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.FixedPointNumber, error) {
	res := &ast.FixedPointNumber{}
	return trace.Trace(Bake(fn.Value(res, BindSpan(&res.Span, fixedPointNumber))))(ctx, start)
}

func fixedPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return trace.Trace(runes.Concat(
		fn.Seq(
			unsignedNumber,
			runes.Tag("."),
			unsignedNumber,
		),
	))(ctx, start)
}

var exp = runes.Recognize(runes.OneOf("eE"))

func UnsignedNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.UnsignedNumber, error) {
	res := &ast.UnsignedNumber{}
	return trace.Trace(Bake(fn.Value(res, BindSpan(&res.Span, unsignedNumber))))(ctx, start)
}

func unsignedNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return trace.Trace(cache.Cache(runes.Cons(
		decimalDigit,
		runes.Join(fn.Many0(decimalDigit_)),
	)))(ctx, start)
}

func binaryValue(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return trace.Trace(runes.Cons(
		binaryDigit,
		runes.Join(fn.Many0(binaryDigit_)),
	))(ctx, start)
}

func octalValue(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return trace.Trace(runes.Cons(
		octalDigit,
		runes.Join(fn.Many0(octalDigit_)),
	))(ctx, start)
}

func hexValue(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return trace.Trace(runes.Cons(
		hexDigit,
		runes.Join(fn.Many0(hexDigit_)),
	))(ctx, start)
}

func decimalBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return trace.Trace(runes.Join(
		fn.Seq(
			runes.Rune('\''),
			fn.Opt(runes.OneOf("sS")),
			runes.OneOf("dD"),
		),
	))(ctx, start)
}

func binaryBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return trace.Trace(runes.Join(
		fn.Seq(
			runes.Rune('\''),
			fn.Opt(runes.OneOf("sS")),
			runes.OneOf("bB"),
		),
	))(ctx, start)
}

func octalBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return trace.Trace(runes.Join(
		fn.Seq(
			runes.Rune('\''),
			fn.Opt(runes.OneOf("sS")),
			runes.OneOf("oO"),
		),
	))(ctx, start)
}

func hexBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return trace.Trace(runes.Join(
		fn.Seq(
			runes.Rune('\''),
			fn.Opt(runes.OneOf("sS")),
			runes.OneOf("hH"),
		),
	))(ctx, start)
}

var nonZeroDecimalDigit = runes.OneOf("123456789")
var decimalDigit = runes.OneOf("0123456789")
var decimalDigit_ = runes.OneOf("0123456789_")
var binaryDigit = runes.OneOf("01xXzZ?")
var binaryDigit_ = runes.OneOf("01xXzZ?_")
var octalDigit = runes.OneOf("01234567xXzZ?")
var octalDigit_ = runes.OneOf("01234567xXzZ?_")
var hexDigit = runes.OneOf("0123456789abcdefABCDEFxXzZ?")
var hexDigit_ = runes.OneOf("0123456789abcdefABCDEFxXzZ?_")
var xDigit = runes.OneOf("xX")
var zDigit = runes.OneOf("zZ?")

func UnbasedUnsizedLiteral(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.UnbasedUnsizedLiteral, error) {
	res := &ast.UnbasedUnsizedLiteral{}
	return trace.Trace(fn.Value(res, BindSpan(&res.Span, unbasedUnsizedLiteral)))(ctx, start)
}

func unbasedUnsizedLiteral(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return trace.Trace(runes.Join(fn.Seq(runes.Rune('\''), runes.OneOf("01xXzZ"))))(ctx, start)
}
