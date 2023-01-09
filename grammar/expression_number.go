package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-nom/trace"
	"github.com/jtdubs/go-svparser/ast"
)

//
// A.8.7 Numbers
//

/*
 * number ::=
 *   integral_number
 *   | real_number
 */
func Number(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.Number, error) {
	return trace.Trace(fn.Alt(
		to[ast.Number](RealNumber),
		to[ast.Number](IntegralNumber),
	))(ctx, start)
}

/*
 * integral_number ::=
 *   decimal_number
 *   | octal_number
 *   | binary_number
 *   | hex_number
 */
func IntegralNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.IntegralNumber, error) {
	return trace.Trace(fn.Alt(
		to[ast.IntegralNumber](octalNumber),
		to[ast.IntegralNumber](binaryNumber),
		to[ast.IntegralNumber](hexNumber),
		to[ast.IntegralNumber](DecimalNumber),
	))(ctx, start)
}

/*
 * decimal_number ::=
 *   unsigned_number
 *   | [ size ] decimal_base unsigned_number
 *   | [ size ] decimal_base x_digit { _ }
 *   | [ size ] decimal_base z_digit { _ }
 */
func DecimalNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.DecimalNumber, error) {
	return trace.Trace(fn.Alt(
		to[ast.DecimalNumber](decimalNumberX),
		to[ast.DecimalNumber](decimalNumberZ),
		to[ast.DecimalNumber](decimalNumberUnsigned),
		to[ast.DecimalNumber](UnsignedNumber),
	))(ctx, start)
}

/*
 * decimal_number ::=
 *   ...
 *   | [ size ] decimal_base z_digit { _ }
 *   ...
 */
func decimalNumberZ(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DecimalNumberZ, error) {
	res := &ast.DecimalNumberZ{}
	return tBindSeq(res, &res.Span,
		fn.Opt(bindSpan(&res.SizeT, size)),
		bindSpan(&res.BaseT, decimalBase),
		bindSpan(&res.Z, zDigit),
		fn.Discard(fn.Many0(runes.Rune('_'))),
	)(ctx, start)
}

/*
 * decimal_number ::=
 *   ...
 *   | [ size ] decimal_base x_digit { _ }
 *   ...
 */
func decimalNumberX(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DecimalNumberX, error) {
	res := &ast.DecimalNumberX{}
	return tBindSeq(res, &res.Span,
		fn.Opt(bindSpan(&res.SizeT, size)),
		bindSpan(&res.BaseT, decimalBase),
		bindSpan(&res.X, xDigit),
		fn.Discard(fn.Many0(runes.Rune('_'))),
	)(ctx, start)
}

/*
 * decimal_number ::=
 *   ...
 *   | [ size ] decimal_base unsigned_number
 *   ...
 */
func decimalNumberUnsigned(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DecimalNumberUnsigned, error) {
	res := &ast.DecimalNumberUnsigned{}
	return tBindSeq(res, &res.Span,
		fn.Opt(bindSpan(&res.SizeT, size)),
		bindSpan(&res.BaseT, decimalBase),
		bindSpan(&res.ValueT, unsignedNumber),
	)(ctx, start)
}

/*
 * binary_number ::= [ size ] binary_base binary_value
 */
func binaryNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.BinaryNumber, error) {
	res := &ast.BinaryNumber{}
	return tBindSeq(res, &res.Span,
		fn.Opt(bindSpan(&res.SizeT, size)),
		bindSpan(&res.BaseT, binaryBase),
		bindSpan(&res.ValueT, binaryValue),
	)(ctx, start)
}

/*
 * octal_number ::= [ size ] octal_base octal_value
 */
func octalNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.OctalNumber, error) {
	res := &ast.OctalNumber{}
	return tBindSeq(res, &res.Span,
		fn.Opt(bindSpan(&res.SizeT, size)),
		bindSpan(&res.BaseT, octalBase),
		bindSpan(&res.ValueT, octalValue),
	)(ctx, start)
}

/*
 * hex_number ::= [ size ] hex_base hex_value
 */
func hexNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.HexNumber, error) {
	res := &ast.HexNumber{}
	return tBindSeq(res, &res.Span,
		fn.Opt(bindSpan(&res.SizeT, size)),
		bindSpan(&res.BaseT, hexBase),
		bindSpan(&res.ValueT, hexValue),
	)(ctx, start)
}

/*
 * sign ::= + | -
 */
var sign = runes.Recognize(runes.OneOf("+-"))

/*
 * size ::= non_zero_unsigned_number
 */
func size(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return trace.Trace(nonZeroUnsignedNumber)(ctx, start)
}

/*
 * non_zero_unsigned_number ::= non_zero_decimal_digit { _ | decimal_digit}
 */
func nonZeroUnsignedNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tCons(
		nonZeroDecimalDigit,
		runes.Join(fn.Many0(decimalDigit_)),
	)(ctx, start)
}

/*
 * real_number ::=
 *   fixed_point_number
 *   | unsigned_number [ . unsigned_number ] exp [ sign ] unsigned_number
 */
func RealNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.RealNumber, error) {
	return trace.Trace(fn.Alt(
		to[ast.RealNumber](FloatingPointNumber),
		to[ast.RealNumber](FixedPointNumber),
	))(ctx, start)
}

/*
 * real_number ::=
 *   ...
 *   | unsigned_number [ . unsigned_number ] exp [ sign ] unsigned_number
 *   ...
 */
func FloatingPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.FloatingPointNumber, error) {
	res := &ast.FloatingPointNumber{}
	return tBind(res, &res.Span, floatingPointNumber)(ctx, start)
}

/*
 * real_number ::=
 *   ...
 *   | unsigned_number [ . unsigned_number ] exp [ sign ] unsigned_number
 *   ...
 */
func floatingPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tConcatSeq(
		unsignedNumber,
		fn.Opt(fn.Preceded(runes.Rune('.'), unsignedNumber)),
		exp,
		fn.Opt(sign),
		unsignedNumber,
	)(ctx, start)
}

/*
 * fixed_point_number ::= unsigned_number . unsigned_number
 */
func FixedPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.FixedPointNumber, error) {
	res := &ast.FixedPointNumber{}
	return tBind(res, &res.Span, fixedPointNumber)(ctx, start)
}

/*
 * fixed_point_number ::= unsigned_number . unsigned_number
 */
func fixedPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tConcatSeq(
		unsignedNumber,
		runes.Tag("."),
		unsignedNumber,
	)(ctx, start)
}

/*
 * exp ::= e | E
 */
var exp = runes.Recognize(runes.OneOf("eE"))

/*
 * unsigned_number ::= decimal_digit { _ | decimal_digit }
 */
func UnsignedNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.UnsignedNumber, error) {
	res := &ast.UnsignedNumber{}
	return tBind(res, &res.Span, unsignedNumber)(ctx, start)
}

/*
 * unsigned_number ::= decimal_digit { _ | decimal_digit }
 */
func unsignedNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tCons(
		decimalDigit,
		runes.Join(fn.Many0(decimalDigit_)),
	)(ctx, start)
}

/*
 * binary_value ::= binary_digit { _ | binary_digit }
 */
func binaryValue(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tCons(
		binaryDigit,
		runes.Join(fn.Many0(binaryDigit_)),
	)(ctx, start)
}

/*
 * octal_value ::= octal_digit { _ | octal_digit }
 */
func octalValue(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tCons(
		octalDigit,
		runes.Join(fn.Many0(octalDigit_)),
	)(ctx, start)
}

/*
 * hex_value ::= hex_digit { _ | hex_digit }
 */
func hexValue(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tCons(
		hexDigit,
		runes.Join(fn.Many0(hexDigit_)),
	)(ctx, start)
}

/*
 * decimal_base ::= '[s|S]d | '[s|S]D
 */
func decimalBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tJoinSeq(
		runes.Rune('\''),
		fn.Opt(runes.OneOf("sS")),
		runes.OneOf("dD"),
	)(ctx, start)
}

/*
 * binary_base ::= '[s|S]b | '[s|S]B
 */
func binaryBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tJoinSeq(
		runes.Rune('\''),
		fn.Opt(runes.OneOf("sS")),
		runes.OneOf("bB"),
	)(ctx, start)
}

/*
 * octal_base ::= '[s|S]o | '[s|S]O
 */
func octalBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tJoinSeq(
		runes.Rune('\''),
		fn.Opt(runes.OneOf("sS")),
		runes.OneOf("oO"),
	)(ctx, start)
}

/*
 * hex_base ::= '[s|S]h | '[s|S]H
 */
func hexBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tJoinSeq(
		runes.Rune('\''),
		fn.Opt(runes.OneOf("sS")),
		runes.OneOf("hH"),
	)(ctx, start)
}

/*
 * non_zero_decimal_digit ::= 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
 */
var nonZeroDecimalDigit = runes.OneOf("123456789")

/*
 * decimal_digit ::= 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
 */
var decimalDigit = runes.OneOf("0123456789")
var decimalDigit_ = runes.OneOf("0123456789_")

/*
 * binary_digit ::= x_digit | z_digit | 0 | 1
 */
var binaryDigit = runes.OneOf("01xXzZ?")
var binaryDigit_ = runes.OneOf("01xXzZ?_")

/*
 * octal_digit ::= x_digit | z_digit | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7
 */
var octalDigit = runes.OneOf("01234567xXzZ?")
var octalDigit_ = runes.OneOf("01234567xXzZ?_")

/*
 * hex_digit ::= x_digit | z_digit | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | a | b | c | d | e | f | A | B | C | D | E | F
 */
var hexDigit = runes.OneOf("0123456789abcdefABCDEFxXzZ?")
var hexDigit_ = runes.OneOf("0123456789abcdefABCDEFxXzZ?_")

/*
 * x_digit ::= x | X
 */
var xDigit = runes.OneOf("xX")

/*
 * z_digit ::= z | Z | ?
 */
var zDigit = runes.OneOf("zZ?")

/*
 * unbased_unsized_literal ::= '0 | '1 | 'z_or_x
 */
func UnbasedUnsizedLiteral(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.UnbasedUnsizedLiteral, error) {
	res := &ast.UnbasedUnsizedLiteral{}
	return tBind(res, &res.Span, unbasedUnsizedLiteral)(ctx, start)
}

/*
 * unbased_unsized_literal ::= '0 | '1 | 'z_or_x
 */
func unbasedUnsizedLiteral(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return tJoinSeq(runes.Rune('\''), runes.OneOf("01xXzZ"))(ctx, start)
}
