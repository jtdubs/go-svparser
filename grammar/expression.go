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
// A.8.3 Expressions
//

/*
 * constant_expression ::=
 *   constant_primary
 *   | unary_operator { attribute_instance } constant_primary
 *   | constant_expression binary_operator { attribute_instance } constant_expression
 *   | constant_expression ? { attribute_instance } constant_expression : constant_expression
 */
func ConstantExpression(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.ConstantExpression, error) {
	return trace.Trace(fn.Alt(
		to[ast.ConstantExpression](constantUnaryExpression),
		to[ast.ConstantExpression](constantBinaryExpression),
		to[ast.ConstantExpression](constantTernaryExpression),
		to[ast.ConstantExpression](ConstantPrimary),
	))(ctx, start)
}

/*
 * constant_expression ::=
 *   ...
 *   | unary_operator { attribute_instance } constant_primary
 *   ...
 */
func constantUnaryExpression(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ConstantUnaryExpression, error) {
	res := &ast.ConstantUnaryExpression{}
	return tBindPhrase(res, &res.Span,
		fn.Bind(&res.Op, UnaryOperator),
		fn.Opt(fn.Bind(&res.Attrs, fn.Many1(AttributeInstance))),
		fn.Bind(&res.Primary, ConstantPrimary),
	)(ctx, start)
}

/*
 * constant_expression ::=
 *   ...
 *   | constant_expression binary_operator { attribute_instance } constant_expression
 *   ...
 */
func constantBinaryExpression(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ConstantBinaryExpression, error) {
	res := &ast.ConstantBinaryExpression{}
	return tBindPhrase(res, &res.Span,
		fn.Bind(&res.Left, ConstantExpression),
		fn.Bind(&res.Op, BinaryOperator),
		fn.Opt(fn.Bind(&res.Attrs, fn.Many1(AttributeInstance))),
		fn.Bind(&res.Right, ConstantExpression),
	)(ctx, start)
}

/*
 * constant_expression ::=
 *   ...
 *   | constant_expression ? { attribute_instance } constant_expression : constant_expression
 *   ...
 */
func constantTernaryExpression(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ConstantTernaryExpression, error) {
	res := &ast.ConstantTernaryExpression{}
	return tBindPhrase(res, &res.Span,
		fn.Bind(&res.Cond, ConstantExpression),
		fn.Discard(runes.Rune('?')),
		fn.Opt(fn.Bind(&res.Attrs, fn.Many1(AttributeInstance))),
		fn.Bind(&res.If, ConstantExpression),
		fn.Discard(runes.Rune(':')),
		fn.Bind(&res.Else, ConstantExpression),
	)(ctx, start)
}
