package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

//
// A.8.3 Expressions
//

/*
 * inc_or_dec_expression ::=
 *   inc_or_dec_operator { attribute_instance } variable_lvalue
 *   | variable_lvalue { attribute_instance } inc_or_dec_operator
 */

/*
 * conditional_expression ::= cond_predicate ? { attribute_instance } expression : expression
 */

/*
 * constant_expression ::=
 *   constant_primary
 *   | unary_operator { attribute_instance } constant_primary
 *   | constant_expression binary_operator { attribute_instance } constant_expression
 *   | constant_expression ? { attribute_instance } constant_expression : constant_expression
 */
func ConstantExpression(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.ConstantExpression, error) {
	return top(func(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.ConstantExpression, error) {
		end, expr, err := fn.Alt(
			to[ast.ConstantExpression](ConstantPrimary),
			to[ast.ConstantExpression](constantUnaryExpression),
		)(ctx, start)

		if err != nil {
			return start, nil, err
		}

		return fn.Alt(
			to[ast.ConstantExpression](constantBinaryExpression(expr)),
			to[ast.ConstantExpression](constantTernaryExpression(expr)),
			fn.Success[rune](expr),
		)(ctx, end)
	})(ctx, start)
}

/*
 * constant_expression ::=
 *   ...
 *   | unary_operator { attribute_instance } constant_primary
 *   ...
 */
func constantUnaryExpression(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ConstantUnaryExpression, error) {
	res := &ast.ConstantUnaryExpression{}
	return top(
		token(res,
			phrase(
				fn.Bind(&res.Op, UnaryOperator),
				fn.Opt(fn.Bind(&res.Attrs, fn.Many1(AttributeInstance))),
				fn.Bind(&res.Primary, ConstantPrimary),
			),
		),
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
	return top(
		token(res,
			phrase(
				fn.Bind(&res.Op, BinaryOperator),
				fn.Opt(fn.Bind(&res.Attrs, fn.Many1(AttributeInstance))),
				fn.Bind(&res.Right, ConstantExpression),
			),
		),
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
	return top(
		token(res,
			phrase(
				fn.Bind(&res.Cond, ConstantExpression),
				fn.Discard(runes.Rune('?')),
				fn.Opt(fn.Bind(&res.Attrs, fn.Many1(AttributeInstance))),
				fn.Bind(&res.If, ConstantExpression),
				fn.Discard(runes.Rune(':')),
				fn.Bind(&res.Else, ConstantExpression),
			),
		),
	)
}

/*
 * constant_mintypmax_expression ::=
 *   constant_expression
 *   | constant_expression : constant_expression : constant_expression
 */

/*
 * constant_param_expression ::=
 *   constant_mintypmax_expression | data_type | $
 */

/*
 * param_expression ::= mintypmax_expression | data_type | $
 */

/*
 * constant_range_expression ::=
 *   constant_expression
 *   | constant_part_select_range
 */

/*
 * constant_part_select_range ::=
 *   constant_range
 *   | constant_indexed_range
 */

/*
 * constant_range ::= constant_expression : constant_expression
 */

/*
 * constant_indexed_range ::=
 *   constant_expression +: constant_expression
 *   | constant_expression -: constant_expression
 */

/*
 * expression ::=
 *   primary
 *   | unary_operator { attribute_instance } primary
 *   | inc_or_dec_expression
 *   | ( operator_assignment )
 *   | expression binary_operator { attribute_instance } expression
 *   | conditional_expression
 *   | inside_expression
 *   | tagged_union_expression
 */

/*
 * tagged_union_expression ::=
 *   tagged member_identifier [ expression ]
 */

/*
 * inside_expression ::= expression inside { open_range_list }
 */

/*
 * value_range ::=
 *   expression
 *   | [ expression : expression ]
 */

/*
 * mintypmax_expression ::=
 *   expression
 *   | expression : expression : expression
 */

/*
 * module_path_conditional_expression ::= module_path_expression ? { attribute_instance }
 *   module_path_expression : module_path_expression
 */

/*
 * module_path_expression ::=
 *   module_path_primary
 *   | unary_module_path_operator { attribute_instance } module_path_primary
 *   | module_path_expression binary_module_path_operator { attribute_instance }
 *   module_path_expression
 *   | module_path_conditional_expression
 */

/*
 * module_path_mintypmax_expression ::=
 *   module_path_expression
 *   | module_path_expression : module_path_expression : module_path_expression
 */

/*
 * part_select_range ::= constant_range | indexed_range
 */

/*
 * indexed_range ::=
 *   expression +: constant_expression
 *   | expression -: constant_expression
 */

/*
 * genvar_expression ::= constant_expression
 */
