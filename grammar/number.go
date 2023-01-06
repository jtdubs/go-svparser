package grammar

import (
	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

func Number(start nom.Cursor[rune]) (nom.Cursor[rune], ast.Number, error) {
	return nom.Alt(
		To[ast.Number](RealNumber),
		To[ast.Number](IntegralNumber),
	)(start)
}

func IntegralNumber(start nom.Cursor[rune]) (nom.Cursor[rune], ast.IntegralNumber, error) {
	return nom.Alt(
		To[ast.IntegralNumber](octalNumber),
		To[ast.IntegralNumber](binaryNumber),
		To[ast.IntegralNumber](hexNumber),
		To[ast.IntegralNumber](DecimalNumber),
	)(start)
}

func DecimalNumber(start nom.Cursor[rune]) (nom.Cursor[rune], ast.DecimalNumber, error) {
	return nom.Alt(
		To[ast.DecimalNumber](decimalNumberX),
		To[ast.DecimalNumber](decimalNumberZ),
		To[ast.DecimalNumber](decimalNumberUnsigned),
		To[ast.DecimalNumber](UnsignedNumber),
	)(start)
}

func decimalNumberZ(start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DecimalNumberZ, error) {
	res := &ast.DecimalNumberZ{}
	return Bake(nom.Value(res,
		Bind(&res.Token,
			nom.Seq(
				nom.Opt(Bind(&res.SizeT, size)),
				Bind(&res.BaseT, decimalBase),
				Bind(&res.Z, zDigit),
				nom.Discard(nom.Many0(runes.Rune('_'))),
			),
		),
	))(start)
}

func decimalNumberX(start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DecimalNumberX, error) {
	res := &ast.DecimalNumberX{}
	return Bake(nom.Value(res,
		Bind(&res.Token,
			nom.Seq(
				nom.Opt(Bind(&res.SizeT, size)),
				Bind(&res.BaseT, decimalBase),
				Bind(&res.X, xDigit),
				nom.Discard(nom.Many0(runes.Rune('_'))),
			),
		),
	))(start)
}

func decimalNumberUnsigned(start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DecimalNumberUnsigned, error) {
	res := &ast.DecimalNumberUnsigned{}
	return Bake(nom.Value(res,
		Bind(&res.Token,
			nom.Seq(
				nom.Opt(Bind(&res.SizeT, size)),
				Bind(&res.BaseT, decimalBase),
				Bind(&res.ValueT, unsignedNumber),
			),
		),
	))(start)
}

func binaryNumber(start nom.Cursor[rune]) (nom.Cursor[rune], *ast.BinaryNumber, error) {
	res := &ast.BinaryNumber{}
	return Bake(nom.Value(res,
		Bind(&res.Token,
			nom.Seq(
				nom.Opt(Bind(&res.SizeT, size)),
				Bind(&res.BaseT, binaryBase),
				Bind(&res.ValueT, binaryValue),
			),
		),
	))(start)
}

func octalNumber(start nom.Cursor[rune]) (nom.Cursor[rune], *ast.OctalNumber, error) {
	res := &ast.OctalNumber{}
	return Bake(nom.Value(res,
		Bind(&res.Token,
			nom.Seq(
				nom.Opt(Bind(&res.SizeT, size)),
				Bind(&res.BaseT, octalBase),
				Bind(&res.ValueT, octalValue),
			),
		),
	))(start)
}

func hexNumber(start nom.Cursor[rune]) (nom.Cursor[rune], *ast.HexNumber, error) {
	res := &ast.HexNumber{}
	return Bake(nom.Value(res,
		Bind(&res.Token,
			nom.Seq(
				nom.Opt(Bind(&res.SizeT, size)),
				Bind(&res.BaseT, hexBase),
				Bind(&res.ValueT, hexValue),
			),
		),
	))(start)
}

var sign = runes.Recognize(runes.OneOf("+-"))

func size(start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return nonZeroUnsignedNumber(start)
}

func nonZeroUnsignedNumber(start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Cons(
		nonZeroDecimalDigit,
		runes.Join(nom.Many0(nom.Alt(decimalDigit, runes.Rune('_')))),
	)(start)
}

func RealNumber(start nom.Cursor[rune]) (nom.Cursor[rune], ast.RealNumber, error) {
	return nom.Alt(
		To[ast.RealNumber](FloatingPointNumber),
		To[ast.RealNumber](FixedPointNumber),
	)(start)
}

func FloatingPointNumber(start nom.Cursor[rune]) (nom.Cursor[rune], *ast.FloatingPointNumber, error) {
	res := &ast.FloatingPointNumber{}
	return Bake(nom.Value(res, Bind(&res.Token, floatingPointNumber)))(start)
}

func floatingPointNumber(start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Concat(
		nom.Seq(
			unsignedNumber,
			nom.Opt(nom.Preceded(runes.Rune('.'), unsignedNumber)),
			exp,
			nom.Opt(sign),
			unsignedNumber,
		),
	)(start)
}

func FixedPointNumber(start nom.Cursor[rune]) (nom.Cursor[rune], *ast.FixedPointNumber, error) {
	res := &ast.FixedPointNumber{}
	return Bake(nom.Value(res, Bind(&res.Token, fixedPointNumber)))(start)
}

func fixedPointNumber(start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Concat(
		nom.Seq(
			unsignedNumber,
			runes.Tag("."),
			unsignedNumber,
		),
	)(start)
}

var exp = runes.Recognize(runes.OneOf("eE"))

func UnsignedNumber(start nom.Cursor[rune]) (nom.Cursor[rune], *ast.UnsignedNumber, error) {
	res := &ast.UnsignedNumber{}
	return Bake(nom.Value(res, Bind(&res.Token, unsignedNumber)))(start)
}

func unsignedNumber(start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Cons(
		decimalDigit,
		runes.Join(nom.Many0(nom.Alt(decimalDigit, runes.Rune('_')))),
	)(start)
}

func binaryValue(start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Cons(
		binaryDigit,
		runes.Join(nom.Many0(nom.Alt(binaryDigit, runes.Rune('_')))),
	)(start)
}

func octalValue(start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Cons(
		octalDigit,
		runes.Join(nom.Many0(nom.Alt(octalDigit, runes.Rune('_')))),
	)(start)
}

func hexValue(start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Cons(
		hexDigit,
		runes.Join(nom.Many0(nom.Alt(hexDigit, runes.Rune('_')))),
	)(start)
}

func decimalBase(start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Join(
		nom.Seq(
			runes.Rune('\''),
			nom.Opt(runes.OneOf("sS")),
			runes.OneOf("dD"),
		),
	)(start)
}

func binaryBase(start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Join(
		nom.Seq(
			runes.Rune('\''),
			nom.Opt(runes.OneOf("sS")),
			runes.OneOf("bB"),
		),
	)(start)
}

func octalBase(start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Join(
		nom.Seq(
			runes.Rune('\''),
			nom.Opt(runes.OneOf("sS")),
			runes.OneOf("oO"),
		),
	)(start)
}

func hexBase(start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Join(
		nom.Seq(
			runes.Rune('\''),
			nom.Opt(runes.OneOf("sS")),
			runes.OneOf("hH"),
		),
	)(start)
}

var nonZeroDecimalDigit = runes.OneOf("123456789")
var decimalDigit = runes.OneOf("0123456789")
var binaryDigit = runes.OneOf("01xXzZ?")
var octalDigit = runes.OneOf("01234567xXzZ?")
var hexDigit = runes.OneOf("0123456789abcdefABCDEFxXzZ?")
var xDigit = runes.OneOf("xX")
var zDigit = runes.OneOf("zZ?")

func UnbasedUnsizedLiteral(start nom.Cursor[rune]) (nom.Cursor[rune], *ast.UnbasedUnsizedLiteral, error) {
	res := &ast.UnbasedUnsizedLiteral{}
	return nom.Value(res, Bind(&res.Token, unbasedUnsizedLiteral))(start)
}

func unbasedUnsizedLiteral(start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.Join(nom.Seq(runes.Rune('\''), runes.OneOf("01xXzZ")))(start)
}
