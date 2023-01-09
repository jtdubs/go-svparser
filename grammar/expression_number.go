package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-nom/trace"
	"github.com/jtdubs/go-svparser/ast"
)

func Number(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.Number, error) {
	return trace.Trace(fn.Alt(
		to[ast.Number](RealNumber),
		to[ast.Number](IntegralNumber),
	))(ctx, start)
}

func IntegralNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.IntegralNumber, error) {
	return trace.Trace(fn.Alt(
		to[ast.IntegralNumber](octalNumber),
		to[ast.IntegralNumber](binaryNumber),
		to[ast.IntegralNumber](hexNumber),
		to[ast.IntegralNumber](DecimalNumber),
	))(ctx, start)
}

func DecimalNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.DecimalNumber, error) {
	return trace.Trace(fn.Alt(
		to[ast.DecimalNumber](decimalNumberX),
		to[ast.DecimalNumber](decimalNumberZ),
		to[ast.DecimalNumber](decimalNumberUnsigned),
		to[ast.DecimalNumber](UnsignedNumber),
	))(ctx, start)
}

func decimalNumberZ(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DecimalNumberZ, error) {
	res := &ast.DecimalNumberZ{}
	return tBindSeq(res, &res.Span,
		fn.Opt(bindSpan(&res.SizeT, size)),
		bindSpan(&res.BaseT, decimalBase),
		bindSpan(&res.Z, zDigit),
		fn.Discard(fn.Many0(runes.Rune('_'))),
	)(ctx, start)
}

func decimalNumberX(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DecimalNumberX, error) {
	res := &ast.DecimalNumberX{}
	return tBindSeq(res, &res.Span,
		fn.Opt(bindSpan(&res.SizeT, size)),
		bindSpan(&res.BaseT, decimalBase),
		bindSpan(&res.X, xDigit),
		fn.Discard(fn.Many0(runes.Rune('_'))),
	)(ctx, start)
}

func decimalNumberUnsigned(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DecimalNumberUnsigned, error) {
	res := &ast.DecimalNumberUnsigned{}
	return tBindSeq(res, &res.Span,
		fn.Opt(bindSpan(&res.SizeT, size)),
		bindSpan(&res.BaseT, decimalBase),
		bindSpan(&res.ValueT, unsignedNumber),
	)(ctx, start)
}

func binaryNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.BinaryNumber, error) {
	res := &ast.BinaryNumber{}
	return tBindSeq(res, &res.Span,
		fn.Opt(bindSpan(&res.SizeT, size)),
		bindSpan(&res.BaseT, binaryBase),
		bindSpan(&res.ValueT, binaryValue),
	)(ctx, start)
}

func octalNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.OctalNumber, error) {
	res := &ast.OctalNumber{}
	return tBindSeq(res, &res.Span,
		fn.Opt(bindSpan(&res.SizeT, size)),
		bindSpan(&res.BaseT, octalBase),
		bindSpan(&res.ValueT, octalValue),
	)(ctx, start)
}

func hexNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.HexNumber, error) {
	res := &ast.HexNumber{}
	return tBindSeq(res, &res.Span,
		fn.Opt(bindSpan(&res.SizeT, size)),
		bindSpan(&res.BaseT, hexBase),
		bindSpan(&res.ValueT, hexValue),
	)(ctx, start)
}

var sign = runes.Recognize(runes.OneOf("+-"))

func size(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return trace.Trace(nonZeroUnsignedNumber)(ctx, start)
}

func nonZeroUnsignedNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tCons(
		nonZeroDecimalDigit,
		runes.Join(fn.Many0(decimalDigit_)),
	)(ctx, start)
}

func RealNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.RealNumber, error) {
	return trace.Trace(fn.Alt(
		to[ast.RealNumber](FloatingPointNumber),
		to[ast.RealNumber](FixedPointNumber),
	))(ctx, start)
}

func FloatingPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.FloatingPointNumber, error) {
	res := &ast.FloatingPointNumber{}
	return tBind(res, &res.Span, floatingPointNumber)(ctx, start)
}

func floatingPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tConcatSeq(
		unsignedNumber,
		fn.Opt(fn.Preceded(runes.Rune('.'), unsignedNumber)),
		exp,
		fn.Opt(sign),
		unsignedNumber,
	)(ctx, start)
}

func FixedPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.FixedPointNumber, error) {
	res := &ast.FixedPointNumber{}
	return tBind(res, &res.Span, fixedPointNumber)(ctx, start)
}

func fixedPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tConcatSeq(
		unsignedNumber,
		runes.Tag("."),
		unsignedNumber,
	)(ctx, start)
}

var exp = runes.Recognize(runes.OneOf("eE"))

func UnsignedNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.UnsignedNumber, error) {
	res := &ast.UnsignedNumber{}
	return tBind(res, &res.Span, unsignedNumber)(ctx, start)
}

func unsignedNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tCons(
		decimalDigit,
		runes.Join(fn.Many0(decimalDigit_)),
	)(ctx, start)
}

func binaryValue(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tCons(
		binaryDigit,
		runes.Join(fn.Many0(binaryDigit_)),
	)(ctx, start)
}

func octalValue(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tCons(
		octalDigit,
		runes.Join(fn.Many0(octalDigit_)),
	)(ctx, start)
}

func hexValue(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tCons(
		hexDigit,
		runes.Join(fn.Many0(hexDigit_)),
	)(ctx, start)
}

func decimalBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tJoinSeq(
		runes.Rune('\''),
		fn.Opt(runes.OneOf("sS")),
		runes.OneOf("dD"),
	)(ctx, start)
}

func binaryBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tJoinSeq(
		runes.Rune('\''),
		fn.Opt(runes.OneOf("sS")),
		runes.OneOf("bB"),
	)(ctx, start)
}

func octalBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tJoinSeq(
		runes.Rune('\''),
		fn.Opt(runes.OneOf("sS")),
		runes.OneOf("oO"),
	)(ctx, start)
}

func hexBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tJoinSeq(
		runes.Rune('\''),
		fn.Opt(runes.OneOf("sS")),
		runes.OneOf("hH"),
	)(ctx, start)
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
	return tBind(res, &res.Span, unbasedUnsizedLiteral)(ctx, start)
}

func unbasedUnsizedLiteral(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tJoinSeq(runes.Rune('\''), runes.OneOf("01xXzZ"))(ctx, start)
}
