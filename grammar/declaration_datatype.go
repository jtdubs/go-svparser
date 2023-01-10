package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

//
// A.2.2 Declaration data types
//

//
// A.2.2.1 Net and variable types
//

/*
 * casting_type ::= simple_type | constant_primary | signing | string | const
 */

/*
 * data_type ::=
 *   integer_vector_type [ signing ] { packed_dimension }
 *   | integer_atom_type [ signing ]
 *   | non_integer_type
 *   | struct_union [ packed [ signing ] ] { struct_union_member { struct_union_member } }
 *   { packed_dimension }13
 *   | enum [ enum_base_type ] { enum_name_declaration { , enum_name_declaration } }
 *   { packed_dimension }
 *   | string
 *   | chandle
 *   | virtual [ interface ] interface_identifier [ parameter_value_assignment ] [ . modport_identifier ]
 *   | [ class_scope | package_scope ] type_identifier { packed_dimension }
 *   | class_type
 *   | event
 *   | ps_covergroup_identifier
 *   | type_reference
 */

/*
 * data_type_or_implicit ::=
 *   data_type
 *   | implicit_data_type
 */

/*
 * implicit_data_type ::= [ signing ] { packed_dimension }
 */

/*
 * enum_base_type ::=
 *   integer_atom_type [ signing ]
 *   | integer_vector_type [ signing ] [ packed_dimension ]
 *   | type_identifier [ packed_dimension ]15
 */

/*
 * enum_name_declaration ::=
 *   enum_identifier [ [ integral_number [ : integral_number ] ] ] [ = constant_expression ]
 */

/*
 * class_scope ::= class_type ::
 */

/*
 * class_type ::=
 *   ps_class_identifier [ parameter_value_assignment ]
 *   { :: class_identifier [ parameter_value_assignment ] }
 */

/*
 * integer_type ::= integer_vector_type | integer_atom_type
 */

/*
 * integer_atom_type ::= byte | shortint | int | longint | integer | time
 */

/*
 * integer_vector_type ::= bit | logic | reg
 */

/*
 * non_integer_type ::= shortreal | real | realtime
 */

/*
 * net_type ::= supply0 | supply1 | tri | triand | trior | trireg| tri0 | tri1 | uwire| wire | wand | wor
 */

/*
 * net_port_type ::=
 *   [ net_type ] data_type_or_implicit
 *   | net_type_identifier
 *   | interconnect implicit_data_type
 */

/*
 * variable_port_type ::= var_data_type
 */

/*
 * var_data_type ::= data_type | var data_type_or_implicit
 */

/*
 * signing ::= signed | unsigned
 */

/*
 * simple_type ::= integer_type | non_integer_type | ps_type_identifier | ps_parameter_identifier
 */

/*
 * struct_union_member16 ::=
 *   { attribute_instance } [random_qualifier] data_type_or_void list_of_variable_decl_assignments ;
 */

/*
 * data_type_or_void ::= data_type | void
 */

/*
 * struct_union ::= struct | union [ tagged ]
 */

/*
 * type_reference ::=
 *   type ( expression17 )
 *   | type ( data_type )
 */

//
// A.2.2.2 Strengths
//

/*
 * drive_strength ::=
 *   ( strength0 , strength1 )
 *   | ( strength1 , strength0 )
 *   | ( strength0 , highz1 )
 *   | ( strength1 , highz0 )
 *   | ( highz0 , strength1 )
 *   | ( highz1 , strength0 )
 */
func DriveStrength(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DriveStrength, error) {
	res := &ast.DriveStrength{}
	return tBind(res, &res.Span,
		parens(
			fn.Alt(
				fn.Seq(
					bindSpan(&res.AT, bindValue(&res.A, to[ast.DriveStrengthOption](Strength0))),
					fn.Discard(fn.Surrounded(Whitespace0, Whitespace0, runes.Rune(','))),
					bindSpan(&res.BT, bindValue(&res.B, to[ast.DriveStrengthOption](Strength1)))),
				fn.Seq(
					bindSpan(&res.AT, bindValue(&res.A, to[ast.DriveStrengthOption](Strength1))),
					fn.Discard(fn.Surrounded(Whitespace0, Whitespace0, runes.Rune(','))),
					bindSpan(&res.BT, bindValue(&res.B, to[ast.DriveStrengthOption](Strength0)))),
				fn.Seq(
					bindSpan(&res.AT, bindValue(&res.A, to[ast.DriveStrengthOption](Strength0))),
					fn.Discard(fn.Surrounded(Whitespace0, Whitespace0, runes.Rune(','))),
					bindSpan(&res.BT, bindValue(&res.B, to[ast.DriveStrengthOption](highZ1)))),
				fn.Seq(
					bindSpan(&res.AT, bindValue(&res.A, to[ast.DriveStrengthOption](Strength1))),
					fn.Discard(fn.Surrounded(Whitespace0, Whitespace0, runes.Rune(','))),
					bindSpan(&res.BT, bindValue(&res.B, to[ast.DriveStrengthOption](highZ0)))),
				fn.Seq(
					bindSpan(&res.AT, bindValue(&res.A, to[ast.DriveStrengthOption](highZ0))),
					fn.Discard(fn.Surrounded(Whitespace0, Whitespace0, runes.Rune(','))),
					bindSpan(&res.BT, bindValue(&res.B, to[ast.DriveStrengthOption](Strength1)))),
				fn.Seq(
					bindSpan(&res.AT, bindValue(&res.A, to[ast.DriveStrengthOption](highZ1))),
					fn.Discard(fn.Surrounded(Whitespace0, Whitespace0, runes.Rune(','))),
					bindSpan(&res.BT, bindValue(&res.B, to[ast.DriveStrengthOption](Strength0)))),
			),
		),
	)(ctx, start)
}

func highZ0(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.HighZ0, error) {
	res := &ast.HighZ0{}
	return tBind(res, &res.Span, runes.TagNoCase("highz0"))(ctx, start)
}

func highZ1(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.HighZ1, error) {
	res := &ast.HighZ1{}
	return tBind(res, &res.Span, runes.TagNoCase("highz1"))(ctx, start)
}

/*
 * strength0 ::= supply0 | strong0 | pull0 | weak0
 */
func Strength0(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.Strength0, error) {
	res := &ast.Strength0{}
	return tBind(res, &res.Span,
		bindValue(&res.Type,
			fn.Alt(
				fn.Value(ast.Supply0, runes.TagNoCase("supply0")),
				fn.Value(ast.Strong0, runes.TagNoCase("strong0")),
				fn.Value(ast.Pull0, runes.TagNoCase("pull0")),
				fn.Value(ast.Weak0, runes.TagNoCase("weak0")),
			),
		),
	)(ctx, start)
}

/*
 * strength1 ::= supply1 | strong1 | pull1 | weak1
 */
func Strength1(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.Strength1, error) {
	res := &ast.Strength1{}
	return tBind(res, &res.Span,
		bindValue(&res.Type,
			fn.Alt(
				fn.Value(ast.Supply1, runes.TagNoCase("supply1")),
				fn.Value(ast.Strong1, runes.TagNoCase("strong1")),
				fn.Value(ast.Pull1, runes.TagNoCase("pull1")),
				fn.Value(ast.Weak1, runes.TagNoCase("weak1")),
			),
		),
	)(ctx, start)
}

/*
 * charge_strength ::= ( small ) | ( medium ) | ( large )
 */
func ChargeStrength(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ChargeStrength, error) {
	res := &ast.ChargeStrength{}
	return tBind(res, &res.Span,
		parens(
			bindSpan(&res.TypeT,
				bindValue(&res.Type,
					fn.Alt(
						fn.Value(ast.Small, runes.TagNoCase("small")),
						fn.Value(ast.Medium, runes.TagNoCase("medium")),
						fn.Value(ast.Large, runes.TagNoCase("large")),
					),
				),
			),
		),
	)(ctx, start)
}

//
// A.2.2.3 Delays
//

/*
 * delay3 ::= # delay_value | # ( mintypmax_expression [ , mintypmax_expression [ ,
 *   mintypmax_expression ] ] )
 */

/*
 * delay2 ::= # delay_value | # ( mintypmax_expression [ , mintypmax_expression ] )
 */

/*
 * delay_value ::=
 *   unsigned_number
 *   | real_number
 *   | ps_identifier
 *   | time_literal
 *   | 1step
 */
