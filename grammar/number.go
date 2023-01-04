package grammar

import (
	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

func Number(c nom.Cursor[rune]) (nom.Cursor[rune], ast.Number, error) {
	return nom.Alt(
		To[ast.Number](RealNumber),
		To[ast.Number](IntegralNumber),
	)(c)
}

func IntegralNumber(c nom.Cursor[rune]) (nom.Cursor[rune], ast.IntegralNumber, error) {
	return nom.Alt(
		To[ast.IntegralNumber](octalNumber),
		To[ast.IntegralNumber](binaryNumber),
		To[ast.IntegralNumber](hexNumber),
		To[ast.IntegralNumber](DecimalNumber),
	)(c)
}

func DecimalNumber(c nom.Cursor[rune]) (nom.Cursor[rune], ast.DecimalNumber, error) {
	return nom.Alt(
		To[ast.DecimalNumber](decimalNumberX),
		To[ast.DecimalNumber](decimalNumberZ),
		To[ast.DecimalNumber](decimalNumberUnsigned),
		To[ast.DecimalNumber](UnsignedNumber),
	)(c)
}

func decimalNumberZ(c nom.Cursor[rune]) (nom.Cursor[rune], *ast.DecimalNumberZ, error) {
	res := &ast.DecimalNumberZ{}
	return Bake(nom.Value(
		BindToken(
			nom.Seq(
				nom.Opt(BindToken(size, &res.SizeT)),
				BindToken(decimalBase, &res.BaseT),
				BindToken(zDigit, &res.Z),
				nom.Discard(nom.Many0(runes.Rune('_'))),
			),
			&res.Token,
		),
		res,
	))(c)
}

func decimalNumberX(c nom.Cursor[rune]) (nom.Cursor[rune], *ast.DecimalNumberX, error) {
	res := &ast.DecimalNumberX{}
	return Bake(nom.Value(
		BindToken(
			nom.Seq(
				nom.Opt(BindToken(size, &res.SizeT)),
				BindToken(decimalBase, &res.BaseT),
				BindToken(xDigit, &res.X),
				nom.Discard(nom.Many0(runes.Rune('_'))),
			),
			&res.Token,
		),
		res,
	))(c)
}

func decimalNumberUnsigned(c nom.Cursor[rune]) (nom.Cursor[rune], *ast.DecimalNumberUnsigned, error) {
	res := &ast.DecimalNumberUnsigned{}
	return Bake(nom.Value(
		BindToken(
			nom.Seq(
				nom.Opt(BindToken(size, &res.SizeT)),
				BindToken(decimalBase, &res.BaseT),
				BindToken(unsignedNumber, &res.ValueT),
			),
			&res.Token,
		),
		res,
	))(c)
}

func binaryNumber(c nom.Cursor[rune]) (nom.Cursor[rune], *ast.BinaryNumber, error) {
	res := &ast.BinaryNumber{}
	return Bake(nom.Value(
		BindToken(
			nom.Seq(
				nom.Opt(BindToken(size, &res.SizeT)),
				BindToken(binaryBase, &res.BaseT),
				BindToken(binaryValue, &res.ValueT),
			),
			&res.Token,
		),
		res,
	))(c)
}

func octalNumber(c nom.Cursor[rune]) (nom.Cursor[rune], *ast.OctalNumber, error) {
	res := &ast.OctalNumber{}
	return Bake(nom.Value(
		BindToken(
			nom.Seq(
				nom.Opt(BindToken(size, &res.SizeT)),
				BindToken(octalBase, &res.BaseT),
				BindToken(octalValue, &res.ValueT),
			),
			&res.Token,
		),
		res,
	))(c)
}

func hexNumber(c nom.Cursor[rune]) (nom.Cursor[rune], *ast.HexNumber, error) {
	res := &ast.HexNumber{}
	return Bake(nom.Value(
		BindToken(
			nom.Seq(
				nom.Opt(BindToken(size, &res.SizeT)),
				BindToken(hexBase, &res.BaseT),
				BindToken(hexValue, &res.ValueT),
			),
			&res.Token,
		),
		res,
	))(c)
}

var sign = runes.OneOf("+-")

func size(c nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return nonZeroUnsignedNumber(c)
}

func nonZeroUnsignedNumber(c nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.RecognizeSeq2(
		nonZeroDecimalDigit,
		nom.Many0(nom.Alt(decimalDigit, runes.Rune('_'))),
	)(c)
}

func RealNumber(c nom.Cursor[rune]) (nom.Cursor[rune], ast.RealNumber, error) {
	return nom.Alt(
		To[ast.RealNumber](FloatingPointNumber),
		To[ast.RealNumber](FixedPointNumber),
	)(c)
}

func FloatingPointNumber(c nom.Cursor[rune]) (nom.Cursor[rune], *ast.FloatingPointNumber, error) {
	res := &ast.FloatingPointNumber{}
	return Bake(nom.Value(BindToken(floatingPointNumber, &res.Token), res))(c)
}

func floatingPointNumber(c nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.RecognizeSeq5(
		UnsignedNumber,
		nom.Opt(nom.Preceded(runes.Rune('.'), UnsignedNumber)),
		exp,
		nom.Opt(sign),
		unsignedNumber,
	)(c)
}

func FixedPointNumber(c nom.Cursor[rune]) (nom.Cursor[rune], *ast.FixedPointNumber, error) {
	res := &ast.FixedPointNumber{}
	return Bake(nom.Value(BindToken(fixedPointNumber, &res.Token), res))(c)
}

func fixedPointNumber(c nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.RecognizeSeq3(
		unsignedNumber,
		runes.Rune('.'),
		unsignedNumber,
	)(c)
}

var exp = runes.OneOf("eE")

func UnsignedNumber(c nom.Cursor[rune]) (nom.Cursor[rune], *ast.UnsignedNumber, error) {
	res := &ast.UnsignedNumber{}
	return Bake(nom.Value(BindToken(unsignedNumber, &res.Token), res))(c)
}

func unsignedNumber(c nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.RecognizeSeq2(
		decimalDigit,
		nom.Many0(nom.Alt(decimalDigit, runes.Rune('_'))),
	)(c)
}

func binaryValue(c nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.RecognizeSeq2(
		binaryDigit,
		nom.Many0(nom.Alt(binaryDigit, runes.Rune('_'))),
	)(c)
}

func octalValue(c nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.RecognizeSeq2(
		octalDigit,
		nom.Many0(nom.Alt(octalDigit, runes.Rune('_'))),
	)(c)
}

func hexValue(c nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.RecognizeSeq2(
		hexDigit,
		nom.Many0(nom.Alt(hexDigit, runes.Rune('_'))),
	)(c)
}

func decimalBase(c nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.RecognizeSeq3(
		runes.Rune('\''),
		nom.Opt(runes.OneOf("sS")),
		runes.OneOf("dD"),
	)(c)
}

func binaryBase(c nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.RecognizeSeq3(
		runes.Rune('\''),
		nom.Opt(runes.OneOf("sS")),
		runes.OneOf("bB"),
	)(c)
}

func octalBase(c nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.RecognizeSeq3(
		runes.Rune('\''),
		nom.Opt(runes.OneOf("sS")),
		runes.OneOf("oO"),
	)(c)
}

func hexBase(c nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.RecognizeSeq3(
		runes.Rune('\''),
		nom.Opt(runes.OneOf("sS")),
		runes.OneOf("hH"),
	)(c)
}

var nonZeroDecimalDigit = runes.OneOf("123456789")
var decimalDigit = runes.OneOf("0123456789")
var binaryDigit = runes.OneOf("01xXzZ?")
var octalDigit = runes.OneOf("01234567xXzZ?")
var hexDigit = runes.OneOf("0123456789abcdefABCDEFxXzZ?")
var xDigit = runes.OneOf("xX")
var zDigit = runes.OneOf("zZ?")

func UnbasedUnsizedLiteral(c nom.Cursor[rune]) (nom.Cursor[rune], *ast.UnbasedUnsizedLiteral, error) {
	res := &ast.UnbasedUnsizedLiteral{}
	return nom.Value(BindToken(unbasedUnsizedLiteral, &res.Token), res)(c)
}

func unbasedUnsizedLiteral(c nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return runes.RecognizeSeq(runes.Rune('\''), runes.OneOf("01xXzZ"))(c)
}
