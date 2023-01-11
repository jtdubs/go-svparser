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
func IntegerType(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.IntegerType, error) {
	return tAlt(
		to[ast.IntegerType](IntegerVectorType),
		to[ast.IntegerType](IntegerAtomType),
	)(ctx, start)
}

/*
 * integer_atom_type ::= byte | shortint | int | longint | integer | time
 */
func IntegerAtomType(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.IntegerAtomType, error) {
	res := &ast.IntegerAtomType{}
	return word(tBind(res, &res.Span,
		bindValue(&res.Type,
			fn.Alt(
				fn.Value(ast.Byte, runes.TagNoCase("byte")),
				fn.Value(ast.ShortInt, runes.TagNoCase("shortint")),
				fn.Value(ast.Integer, runes.TagNoCase("integer")),
				fn.Value(ast.Int, runes.TagNoCase("int")),
				fn.Value(ast.LongInt, runes.TagNoCase("longint")),
				fn.Value(ast.Time, runes.TagNoCase("time")),
			),
		),
	))(ctx, start)
}

/*
 * integer_vector_type ::= bit | logic | reg
 */
func IntegerVectorType(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.IntegerVectorType, error) {
	res := &ast.IntegerVectorType{}
	return word(tBind(res, &res.Span,
		bindValue(&res.Type,
			fn.Alt(
				fn.Value(ast.Bit, runes.TagNoCase("bit")),
				fn.Value(ast.Logic, runes.TagNoCase("logic")),
				fn.Value(ast.Reg, runes.TagNoCase("reg")),
			),
		),
	))(ctx, start)
}

/*
 * non_integer_type ::= shortreal | real | realtime
 */
func NonIntegerType(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.NonIntegerType, error) {
	res := &ast.NonIntegerType{}
	return word(tBind(res, &res.Span,
		bindValue(&res.Type,
			fn.Alt(
				fn.Value(ast.ShortReal, runes.TagNoCase("shortreal")),
				fn.Value(ast.RealTime, runes.TagNoCase("realtime")),
				fn.Value(ast.Real, runes.TagNoCase("real")),
			),
		),
	))(ctx, start)
}

/*
 * net_type ::= supply0 | supply1 | tri | triand | trior | trireg| tri0 | tri1 | uwire| wire | wand | wor
 */
func NetType(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.NetType, error) {
	res := &ast.NetType{}
	return word(tBind(res, &res.Span,
		bindValue(&res.Type,
			fn.Alt(
				fn.Value(ast.NetSupply0, runes.TagNoCase("supply0")),
				fn.Value(ast.NetSupply1, runes.TagNoCase("supply1")),
				fn.Value(ast.NetTri, runes.TagNoCase("tri")),
				fn.Value(ast.NetTriAnd, runes.TagNoCase("triand")),
				fn.Value(ast.NetTriOr, runes.TagNoCase("trior")),
				fn.Value(ast.NetTriReg, runes.TagNoCase("trireg")),
				fn.Value(ast.NetTri0, runes.TagNoCase("tri0")),
				fn.Value(ast.NetTri1, runes.TagNoCase("tri1")),
				fn.Value(ast.NetUWire, runes.TagNoCase("uwire")),
				fn.Value(ast.NetWire, runes.TagNoCase("wire")),
				fn.Value(ast.NetWAnd, runes.TagNoCase("wand")),
				fn.Value(ast.NetWOr, runes.TagNoCase("wor")),
			),
		),
	))(ctx, start)
}

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
func Signing(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.Signing, error) {
	res := &ast.Signing{}
	return word(tBind(res, &res.Span,
		bindValue(&res.Type,
			fn.Alt(
				fn.Value(ast.Signed, runes.TagNoCase("signed")),
				fn.Value(ast.Unsigned, runes.TagNoCase("unsigned")),
			),
		),
	))(ctx, start)
}

/*
 * simple_type ::= integer_type | non_integer_type | ps_type_identifier | ps_parameter_identifier
 */
func SimpleType() {
	_ = fn.Discard(IntegerType)
	_ = fn.Discard(NonIntegerType)
	// PsTypeIdentifier()
	// PsParameterIdentifier()
}

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
				phrase(
					bindSpan(&res.AT, bindValue(&res.A, to[ast.DriveStrengthOption](Strength0))),
					fn.Discard(runes.Rune(',')),
					bindSpan(&res.BT, bindValue(&res.B, to[ast.DriveStrengthOption](Strength1)))),
				phrase(
					bindSpan(&res.AT, bindValue(&res.A, to[ast.DriveStrengthOption](Strength1))),
					fn.Discard(runes.Rune(',')),
					bindSpan(&res.BT, bindValue(&res.B, to[ast.DriveStrengthOption](Strength0)))),
				phrase(
					bindSpan(&res.AT, bindValue(&res.A, to[ast.DriveStrengthOption](Strength0))),
					fn.Discard(runes.Rune(',')),
					bindSpan(&res.BT, bindValue(&res.B, to[ast.DriveStrengthOption](highZ1)))),
				phrase(
					bindSpan(&res.AT, bindValue(&res.A, to[ast.DriveStrengthOption](Strength1))),
					fn.Discard(runes.Rune(',')),
					bindSpan(&res.BT, bindValue(&res.B, to[ast.DriveStrengthOption](highZ0)))),
				phrase(
					bindSpan(&res.AT, bindValue(&res.A, to[ast.DriveStrengthOption](highZ0))),
					fn.Discard(runes.Rune(',')),
					bindSpan(&res.BT, bindValue(&res.B, to[ast.DriveStrengthOption](Strength1)))),
				phrase(
					bindSpan(&res.AT, bindValue(&res.A, to[ast.DriveStrengthOption](highZ1))),
					fn.Discard(runes.Rune(',')),
					bindSpan(&res.BT, bindValue(&res.B, to[ast.DriveStrengthOption](Strength0)))),
			),
		),
	)(ctx, start)
}

func highZ0(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.HighZ0, error) {
	res := &ast.HighZ0{}
	return word(tBind(res, &res.Span, runes.TagNoCase("highz0")))(ctx, start)
}

func highZ1(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.HighZ1, error) {
	res := &ast.HighZ1{}
	return word(tBind(res, &res.Span, runes.TagNoCase("highz1")))(ctx, start)
}

/*
 * strength0 ::= supply0 | strong0 | pull0 | weak0
 */
func Strength0(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.Strength0, error) {
	res := &ast.Strength0{}
	return word(tBind(res, &res.Span,
		bindValue(&res.Type,
			fn.Alt(
				fn.Value(ast.StrengthSupply0, runes.TagNoCase("supply0")),
				fn.Value(ast.StrengthStrong0, runes.TagNoCase("strong0")),
				fn.Value(ast.StrengthPull0, runes.TagNoCase("pull0")),
				fn.Value(ast.StrengthWeak0, runes.TagNoCase("weak0")),
			),
		),
	))(ctx, start)
}

/*
 * strength1 ::= supply1 | strong1 | pull1 | weak1
 */
func Strength1(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.Strength1, error) {
	res := &ast.Strength1{}
	return word(tBind(res, &res.Span,
		bindValue(&res.Type,
			fn.Alt(
				fn.Value(ast.StrengthSupply1, runes.TagNoCase("supply1")),
				fn.Value(ast.StrengthStrong1, runes.TagNoCase("strong1")),
				fn.Value(ast.StrengthPull1, runes.TagNoCase("pull1")),
				fn.Value(ast.StrengthWeak1, runes.TagNoCase("weak1")),
			),
		),
	))(ctx, start)
}

/*
 * charge_strength ::= ( small ) | ( medium ) | ( large )
 */
func ChargeStrength(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ChargeStrength, error) {
	res := &ast.ChargeStrength{}
	return word(tBind(res, &res.Span,
		parens(
			bindSpan(&res.TypeT,
				bindValue(&res.Type,
					fn.Alt(
						fn.Value(ast.ChargeSmall, runes.TagNoCase("small")),
						fn.Value(ast.ChargeMedium, runes.TagNoCase("medium")),
						fn.Value(ast.ChargeLarge, runes.TagNoCase("large")),
					),
				),
			),
		),
	))(ctx, start)
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
