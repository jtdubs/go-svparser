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
	end, expr, err := trace.Trace(fn.Alt(
		to[ast.ConstantExpression](ConstantPrimary),
		to[ast.ConstantExpression](constantUnaryExpression),
	))(ctx, start)

	if err != nil {
		return start, nil, err
	}

	return trace.Trace(fn.Alt(
		to[ast.ConstantExpression](constantBinaryExpression(expr)),
		to[ast.ConstantExpression](constantTernaryExpression(expr)),
		fn.Success[rune](expr),
	))(ctx, end)
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
func constantBinaryExpression(left ast.ConstantExpression) nom.ParseFn[rune, *ast.ConstantBinaryExpression] {
	res := &ast.ConstantBinaryExpression{Left: left}
	return tBindPhrase(res, &res.Span,
		fn.Bind(&res.Op, BinaryOperator),
		fn.Opt(fn.Bind(&res.Attrs, fn.Many1(AttributeInstance))),
		fn.Bind(&res.Right, ConstantExpression),
	)
}

/*
 * constant_expression ::=
 *   ...
 *   | constant_expression ? { attribute_instance } constant_expression : constant_expression
 *   ...
 */
func constantTernaryExpression(cond ast.ConstantExpression) nom.ParseFn[rune, *ast.ConstantTernaryExpression] {
	res := &ast.ConstantTernaryExpression{Cond: cond}
	return tBindPhrase(res, &res.Span,
		fn.Bind(&res.Cond, ConstantExpression),
		fn.Discard(runes.Rune('?')),
		fn.Opt(fn.Bind(&res.Attrs, fn.Many1(AttributeInstance))),
		fn.Bind(&res.If, ConstantExpression),
		fn.Discard(runes.Rune(':')),
		fn.Bind(&res.Else, ConstantExpression),
	)
}
