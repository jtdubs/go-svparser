package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

//
// A.8.6 Operators
//

/*
 * unary_operator ::=
 *   + | - | ! | ~ | & | ~& | | | ~| | ^ | ~^ | ^~
 */
func UnaryOperator(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.UnaryOperator, error) {
	res := &ast.UnaryOperator{}
	return tBindPhrase(res,
		bindValue(&res.Op,
			fn.Alt(
				fn.Value(ast.UnaryLogicalReductionNand, runes.Tag("~&")),
				fn.Value(ast.UnaryLogicalReductionNor, runes.Tag("~|")),
				fn.Value(ast.UnaryLogicalReductionXnor, runes.Tag("~^")),
				fn.Value(ast.UnaryLogicalReductionXnor, runes.Tag("^~")),
				fn.Value(ast.UnaryPositive, runes.Tag("+")),
				fn.Value(ast.UnaryNegate, runes.Tag("-")),
				fn.Value(ast.UnaryLogicalNegation, runes.Tag("!")),
				fn.Value(ast.UnaryLogicalReductionAnd, runes.Tag("&")),
				fn.Value(ast.UnaryLogicalReductionOr, runes.Tag("|")),
				fn.Value(ast.UnaryLogicalReductionXor, runes.Tag("^")),
				fn.Value(ast.UnaryLogicalReductionNot, runes.Tag("~")),
			),
		),
	)(ctx, start)
}

/*
 * binary_operator ::=
 *   + | - | * | / | % | == | != | === | !== | ==? | !=? | && | || | **
 *   | < | <= | > | >= | & | | | ^ | ^~ | ~^ | >> | << | >>> | <<<
 *   | -> | <->
 */
func BinaryOperator(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.BinaryOperator, error) {
	res := &ast.BinaryOperator{}
	return tBindPhrase(res,
		bindValue(&res.Op,
			fn.Alt(
				fn.Value(ast.BinaryArithmeticShiftLeft, runes.Tag("<<<")),
				fn.Value(ast.BinaryArithmeticShiftRight, runes.Tag(">>>")),
				fn.Value(ast.BinaryLogicalIff, runes.Tag("<->")),
				fn.Value(ast.BinaryCaseEquals, runes.Tag("===")),
				fn.Value(ast.BinaryCaseNotEquals, runes.Tag("!==")),
				fn.Value(ast.BinaryWildcardEquals, runes.Tag("==?")),
				fn.Value(ast.BinaryWildcardNotEquals, runes.Tag("!=?")),
				fn.Value(ast.BinaryLogicalImplies, runes.Tag("->")),
				fn.Value(ast.BinaryExp, runes.Tag("**")),
				fn.Value(ast.BinaryBitwiseXnor, runes.Tag("^~")),
				fn.Value(ast.BinaryBitwiseXnor, runes.Tag("~^")),
				fn.Value(ast.BinaryLogicalShiftLeft, runes.Tag("<<")),
				fn.Value(ast.BinaryLogicalShiftRight, runes.Tag(">>")),
				fn.Value(ast.BinaryLogicalAnd, runes.Tag("&&")),
				fn.Value(ast.BinaryLogicalOr, runes.Tag("||")),
				fn.Value(ast.BinaryLessThanEqual, runes.Tag("<=")),
				fn.Value(ast.BinaryGreaterThanEqual, runes.Tag(">=")),
				fn.Value(ast.BinaryLogicalNotEquals, runes.Tag("!=")),
				fn.Value(ast.BinaryLogicalEquals, runes.Tag("==")),
				fn.Value(ast.BinaryLessThan, runes.Tag("<")),
				fn.Value(ast.BinaryGreaterThan, runes.Tag(">")),
				fn.Value(ast.BinaryBitwiseXor, runes.Tag("^")),
				fn.Value(ast.BinaryAdd, runes.Tag("+")),
				fn.Value(ast.BinarySubtract, runes.Tag("-")),
				fn.Value(ast.BinaryMultiply, runes.Tag("*")),
				fn.Value(ast.BinaryDivide, runes.Tag("/")),
				fn.Value(ast.BinaryModulus, runes.Tag("%")),
				fn.Value(ast.BinaryBitwiseAnd, runes.Tag("&")),
				fn.Value(ast.BinaryBitwiseOr, runes.Tag("|")),
			),
		),
	)(ctx, start)
}

/*
 * inc_or_dec_operator ::= ++ | --
 */
func IncOrDecOperator(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.IncOrDecOperator, error) {
	res := &ast.IncOrDecOperator{}
	return tBindPhrase(res,
		bindValue(&res.Op,
			fn.Alt(
				fn.Value(ast.Inc, runes.Tag("++")),
				fn.Value(ast.Dec, runes.Tag("--")),
			),
		),
	)(ctx, start)
}

/*
 * unary_module_path_operator ::=
 *   ! | ~ | & | ~& | | | ~| | ^ | ~^ | ^~
 */
func UnaryModulePathOperator(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.UnaryModulePathOperator, error) {
	res := &ast.UnaryModulePathOperator{}
	return tBindPhrase(res,
		bindValue(&res.Op,
			fn.Alt(
				fn.Value(ast.UnaryLogicalReductionNand, runes.Tag("~&")),
				fn.Value(ast.UnaryLogicalReductionNor, runes.Tag("~|")),
				fn.Value(ast.UnaryLogicalReductionXnor, runes.Tag("~^")),
				fn.Value(ast.UnaryLogicalReductionXnor, runes.Tag("^~")),
				fn.Value(ast.UnaryLogicalNegation, runes.Tag("!")),
				fn.Value(ast.UnaryLogicalReductionAnd, runes.Tag("&")),
				fn.Value(ast.UnaryLogicalReductionOr, runes.Tag("|")),
				fn.Value(ast.UnaryLogicalReductionXor, runes.Tag("^")),
				fn.Value(ast.UnaryLogicalReductionNot, runes.Tag("~")),
			),
		),
	)(ctx, start)
}

/*
 * binary_module_path_operator ::=
 *   == | != | && | || | & | | | ^ | ^~ | ~^
 */
func BinaryModulePathOperator(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.BinaryModulePathOperator, error) {
	res := &ast.BinaryModulePathOperator{}
	return tBindPhrase(res,
		bindValue(&res.Op,
			fn.Alt(
				fn.Value(ast.BinaryBitwiseXnor, runes.Tag("^~")),
				fn.Value(ast.BinaryBitwiseXnor, runes.Tag("~^")),
				fn.Value(ast.BinaryLogicalAnd, runes.Tag("&&")),
				fn.Value(ast.BinaryLogicalOr, runes.Tag("||")),
				fn.Value(ast.BinaryLogicalNotEquals, runes.Tag("!=")),
				fn.Value(ast.BinaryLogicalEquals, runes.Tag("==")),
				fn.Value(ast.BinaryBitwiseXor, runes.Tag("^")),
				fn.Value(ast.BinaryBitwiseAnd, runes.Tag("&")),
				fn.Value(ast.BinaryBitwiseOr, runes.Tag("|")),
			),
		),
	)(ctx, start)
}
