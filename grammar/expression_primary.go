package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

//
// A.8.4 Primaries
//

/*
 * constant_primary ::=
 *   primary_literal
 *   | ps_parameter_identifier constant_select
 *   | specparam_identifier [ [ constant_range_expression ] ]
 *   | genvar_identifier
 *   | formal_port_identifier constant_select
 *   | [ package_scope | class_scope ] enum_identifier
 *   | constant_concatenation [ [ constant_range_expression ] ]
 *   | constant_multiple_concatenation [ [ constant_range_expression ] ]
 *   | constant_function_call
 *   | constant_let_expression
 *   | ( constant_mintypmax_expression )
 *   | constant_cast
 *   | constant_assignment_pattern_expression
 *   | type_reference40
 *   | null
 */
func ConstantPrimary(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.ConstantPrimary, error) {
	return tAlt(
		to[ast.ConstantPrimary](PrimaryLiteral),
		to[ast.ConstantPrimary](GenvarIdentifier),
		// TODO(justindubs): the rest of the owl
	)(ctx, start)
}

/*
 * primary_literal ::= number | time_literal | unbased_unsized_literal | string_literal
 */
func PrimaryLiteral(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.PrimaryLiteral, error) {
	return tAlt(
		to[ast.PrimaryLiteral](UnbasedUnsizedLiteral),
		to[ast.PrimaryLiteral](StringLiteral),
		to[ast.PrimaryLiteral](TimeLiteral),
		to[ast.PrimaryLiteral](Number),
	)(ctx, start)
}

/*
 * time_literal ::=
 *   unsigned_number time_unit
 *   | fixed_point_number time_unit
 */
func TimeLiteral(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.TimeLiteral, error) {
	res := &ast.TimeLiteral{}
	return tBind(res, &res.Span,
		fn.Alt(
			fn.Seq(
				fn.Bind(&res.Number, to[ast.Number](UnsignedNumber)),
				fn.Bind(&res.Unit, TimeUnit),
			),
			fn.Seq(
				fn.Bind(&res.Number, to[ast.Number](FixedPointNumber)),
				fn.Bind(&res.Unit, TimeUnit),
			),
		),
	)(ctx, start)
}

/*
 * time_unit ::= s | ms | us | ns | ps | fs
 */
func TimeUnit(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.TimeUnit, error) {
	res := &ast.TimeUnit{}
	return tBind(res, &res.Span,
		fn.Bind(&res.Op,
			fn.Alt(
				fn.Value(ast.MS, runes.TagNoCase("ms")),
				fn.Value(ast.US, runes.TagNoCase("us")),
				fn.Value(ast.NS, runes.TagNoCase("ns")),
				fn.Value(ast.PS, runes.TagNoCase("ps")),
				fn.Value(ast.FS, runes.TagNoCase("fs")),
				fn.Value(ast.S, runes.TagNoCase("s")),
			),
		),
	)(ctx, start)
}

/*
 * constant_bit_select ::= { [ constant_expression ] }
 */
func ConstantBitSelect(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ConstantBitSelect, error) {
	res := &ast.ConstantBitSelect{}
	return tBind(res, &res.Span,
		bindValue(&res.Exprs, fn.Many0(
			fn.Surrounded(
				word(runes.Rune('[')),
				word(runes.Rune(']')),
				ConstantExpression,
			),
		)),
	)(ctx, start)
}
