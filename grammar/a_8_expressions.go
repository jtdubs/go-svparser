package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

//
// A.8 Expressions
//

//
// A.8.1 Concatenations
//

/*
 * concatenation ::=
 *   { expression { , expression } }
 */

/*
 * constant_concatenation ::=
 *   { constant_expression { , constant_expression } }
 */

/*
 * constant_multiple_concatenation ::= { constant_expression constant_concatenation }
 */

/*
 * module_path_concatenation ::= { module_path_expression { , module_path_expression } }
 */

/*
 * module_path_multiple_concatenation ::= { constant_expression module_path_concatenation }
 */

/*
 * multiple_concatenation ::= { expression concatenation }34
 */

/*
 * streaming_concatenation ::= { stream_operator [ slice_size ] stream_concatenation }
 */

/*
 * stream_operator ::= >> | <<
 */

/*
 * slice_size ::= simple_type | constant_expression
 */

/*
 * stream_concatenation ::= { stream_expression { , stream_expression } }
 */

/*
 * stream_expression ::= expression [ with [ array_range_expression ] ]
 */

/*
 * array_range_expression ::=
 *   expression
 *   | expression : expression
 *   | expression +: expression
 *   | expression -: expression
 */

/*
 * empty_unpacked_array_concatenation35 ::= { }
 */

//
// A.8.2 Subroutine calls
//

/*
 * constant_function_call ::= function_subroutine_call36
 */

/*
 * tf_call37 ::= ps_or_hierarchical_tf_identifier { attribute_instance } [ ( list_of_arguments ) ]
 */

/*
 * system_tf_call ::=
 *   system_tf_identifier [ ( list_of_arguments ) ]
 *   | system_tf_identifier ( data_type [ , expression ] )
 * | system_tf_identifier ( expression { , [ expression ] } [ , [ clocking_event ] ] )
 */

/*
 * subroutine_call ::=
 *   tf_call
 *   | system_tf_call
 *   | method_call
 *   | [ std :: ] randomize_call
 */

/*
 * function_subroutine_call ::= subroutine_call
 */

/*
 * list_of_arguments ::=
 *   [ expression ] { , [ expression ] } { , . identifier ( [ expression ] ) }
 *   | . identifier ( [ expression ] ) { , . identifier ( [ expression ] ) }
 */

/*
 * method_call ::= method_call_root . method_call_body
 */

/*
 * method_call_body ::=
 *   method_identifier { attribute_instance } [ ( list_of_arguments ) ]
 *   | built_in_method_call
 */

/*
 * built_in_method_call ::=
 *   array_manipulation_call
 *   | randomize_call
 */

/*
 * array_manipulation_call ::=
 *   array_method_name { attribute_instance }
 *   [ ( list_of_arguments ) ]
 *   [ with ( expression ) ]
 */

/*
 * randomize_call ::=
 *   randomize { attribute_instance }
 *   [ ( [ variable_identifier_list | null ] ) ]
 *   [ with [ ( [ identifier_list ] ) ] constraint_block ]38
 */

/*
 * method_call_root ::= primary | implicit_class_handle
 */

/*
 * array_method_name ::=
 *   method_identifier | unique | and | or | xor
 */

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

//
// A.8.4 Primaries
//

/*
 * constant_primary ::=
 *   primary_literal
 *   | ps_parameter_identifier constant_select
 *   | specparam_identifier [ [ constant_range_expression ] ]
 *   | genvar_identifier39
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
	return top(
		fn.Alt(
			to[ast.ConstantPrimary](PrimaryLiteral),
			to[ast.ConstantPrimary](GenvarIdentifier),
			// TODO(justindubs): the rest of the owl
		),
	)(ctx, start)
}

/*
 * module_path_primary ::=
 *   number
 *   | identifier
 *   | module_path_concatenation
 *   | module_path_multiple_concatenation
 *   | function_subroutine_call
 *   | ( module_path_mintypmax_expression )
 */

/*
 * primary ::=
 *   primary_literal
 *   | [ class_qualifier | package_scope ] hierarchical_identifier select
 *   | empty_unpacked_array_concatenation
 *   | concatenation [ [ range_expression ] ]
 *   | multiple_concatenation [ [ range_expression ] ]
 *   | function_subroutine_call
 *   | let_expression
 *   | ( mintypmax_expression )
 *   | cast
 *   | assignment_pattern_expression
 *   | streaming_concatenation
 *   | sequence_method_call
 *   | this41
 *   | $42
 *   | null
 */

/*
 * class_qualifier ::= [ local ::43 ] [ implicit_class_handle . | class_scope ]
 */

/*
 * range_expression ::=
 *   expression
 *   | part_select_range
 */

/*
 * primary_literal ::= number | time_literal | unbased_unsized_literal | string_literal
 */
func PrimaryLiteral(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.PrimaryLiteral, error) {
	return top(
		fn.Alt(
			to[ast.PrimaryLiteral](UnbasedUnsizedLiteral),
			to[ast.PrimaryLiteral](StringLiteral),
			to[ast.PrimaryLiteral](TimeLiteral),
			to[ast.PrimaryLiteral](Number),
		),
	)(ctx, start)
}

/*
 * time_literal44 ::=
 *   unsigned_number time_unit
 *   | fixed_point_number time_unit
 */
func TimeLiteral(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.TimeLiteral, error) {
	res := &ast.TimeLiteral{}
	return top(
		token(res,
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
		),
	)(ctx, start)
}

/*
 * time_unit ::= s | ms | us | ns | ps | fs
 */
func TimeUnit(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.TimeUnit, error) {
	res := &ast.TimeUnit{}
	return top(
		token(res,
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
		),
	)(ctx, start)
}

/*
 * implicit_class_handle41 ::= this | super | this . super
 */

/*
 * bit_select ::= { [ expression ] }
 */

/*
 * select ::=
 *   [ { . member_identifier bit_select } . member_identifier ] bit_select [ [ part_select_range ] ]
 */

/*
 * nonrange_select ::=
 *   [ { . member_identifier bit_select } . member_identifier ] bit_select
 */

/*
 * constant_bit_select ::= { [ constant_expression ] }
 */
func ConstantBitSelect(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ConstantBitSelect, error) {
	res := &ast.ConstantBitSelect{}
	return top(
		token(res,
			bind(&res.Exprs,
				fn.Many0(brackets(ConstantExpression)),
			),
		),
	)(ctx, start)
}

/*
 * constant_select ::=
 *   [ { . member_identifier constant_bit_select } . member_identifier ] constant_bit_select
 *   [ [ constant_part_select_range ] ]
 */

/*
 * constant_cast ::=
 *   casting_type ' ( constant_expression )
 */

/*
 * constant_let_expression ::= let_expression45
 */

/*
 * cast ::=
 *   casting_type ' ( expression )
 */

//
// A.8.5 Expression left-side values
//

/*
 * net_lvalue ::=
 *   ps_or_hierarchical_net_identifier constant_select
 *   | { net_lvalue { , net_lvalue } }
 *   | [ assignment_pattern_expression_type ] assignment_pattern_net_lvalue
 */

/*
 * variable_lvalue ::=
 *   [ implicit_class_handle . | package_scope ] hierarchical_variable_identifier select46
 *   | { variable_lvalue { , variable_lvalue } }
 *   | [ assignment_pattern_expression_type ] assignment_pattern_variable_lvalue
 *   | streaming_concatenation47
 */

/*
 * nonrange_variable_lvalue ::=
 *   [ implicit_class_handle . | package_scope ] hierarchical_variable_identifier nonrange_select
 */

//
// A.8.6 Operators
//

/*
 * unary_operator ::=
 *   + | - | ! | ~ | & | ~& | | | ~| | ^ | ~^ | ^~
 */
func UnaryOperator(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.UnaryOperator, error) {
	res := &ast.UnaryOperator{}
	return top(
		token(res,
			word(
				bind(&res.Op,
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
	return top(
		token(res,
			word(
				bind(&res.Op,
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
			),
		),
	)(ctx, start)
}

/*
 * inc_or_dec_operator ::= ++ | --
 */
func IncOrDecOperator(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.IncOrDecOperator, error) {
	res := &ast.IncOrDecOperator{}
	return top(
		token(res,
			word(
				bind(&res.Op,
					fn.Alt(
						fn.Value(ast.Inc, runes.Tag("++")),
						fn.Value(ast.Dec, runes.Tag("--")),
					),
				),
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
	return top(
		token(res,
			word(
				bind(&res.Op,
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
	return top(
		token(res,
			word(
				bind(&res.Op,
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
			),
		),
	)(ctx, start)
}

//
// A.8.7 Numbers
//

/*
 * number ::=
 *   integral_number
 *   | real_number
 */
func Number(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.Number, error) {
	return top(
		fn.Alt(
			to[ast.Number](RealNumber),
			to[ast.Number](IntegralNumber),
		),
	)(ctx, start)
}

/*
 * integral_number ::=
 *   decimal_number
 *   | octal_number
 *   | binary_number
 *   | hex_number
 */
func IntegralNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.IntegralNumber, error) {
	return top(
		fn.Alt(
			to[ast.IntegralNumber](octalNumber),
			to[ast.IntegralNumber](binaryNumber),
			to[ast.IntegralNumber](hexNumber),
			to[ast.IntegralNumber](DecimalNumber),
		),
	)(ctx, start)
}

/*
 * decimal_number ::=
 *   unsigned_number
 *   | [ size ] decimal_base unsigned_number
 *   | [ size ] decimal_base x_digit { _ }
 *   | [ size ] decimal_base z_digit { _ }
 */
func DecimalNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.DecimalNumber, error) {
	return top(
		fn.Alt(
			to[ast.DecimalNumber](decimalNumberX),
			to[ast.DecimalNumber](decimalNumberZ),
			to[ast.DecimalNumber](decimalNumberUnsigned),
			to[ast.DecimalNumber](UnsignedNumber),
		),
	)(ctx, start)
}

/*
 * decimal_number ::=
 *   ...
 *   | [ size ] decimal_base z_digit { _ }
 *   ...
 */
func decimalNumberZ(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DecimalNumberZ, error) {
	res := &ast.DecimalNumberZ{}
	return top(
		token(res,
			phrase(
				fn.Opt(bindSpan(&res.SizeT, size)),
				bindSpan(&res.BaseT, decimalBase),
				bindSpan(&res.Z, zDigit),
				fn.Discard(fn.Many0(word(runes.Rune('_')))),
			),
		),
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
	return top(
		token(res,
			phrase(
				fn.Opt(bindSpan(&res.SizeT, size)),
				bindSpan(&res.BaseT, decimalBase),
				bindSpan(&res.X, xDigit),
				fn.Discard(fn.Many0(word(runes.Rune('_')))),
			),
		),
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
	return top(
		token(res,
			phrase(
				fn.Opt(bindSpan(&res.SizeT, size)),
				bindSpan(&res.BaseT, decimalBase),
				bindSpan(&res.ValueT, unsignedNumber),
			),
		),
	)(ctx, start)
}

/*
 * binary_number ::= [ size ] binary_base binary_value
 */
func binaryNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.BinaryNumber, error) {
	res := &ast.BinaryNumber{}
	return top(
		token(res,
			phrase(
				fn.Opt(bindSpan(&res.SizeT, size)),
				bindSpan(&res.BaseT, binaryBase),
				bindSpan(&res.ValueT, binaryValue),
			),
		),
	)(ctx, start)
}

/*
 * octal_number ::= [ size ] octal_base octal_value
 */
func octalNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.OctalNumber, error) {
	res := &ast.OctalNumber{}
	return top(
		token(res,
			phrase(
				fn.Opt(bindSpan(&res.SizeT, size)),
				bindSpan(&res.BaseT, octalBase),
				bindSpan(&res.ValueT, octalValue),
			),
		),
	)(ctx, start)
}

/*
 * hex_number ::= [ size ] hex_base hex_value
 */
func hexNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.HexNumber, error) {
	res := &ast.HexNumber{}
	return top(
		token(res,
			phrase(
				fn.Opt(bindSpan(&res.SizeT, size)),
				bindSpan(&res.BaseT, hexBase),
				bindSpan(&res.ValueT, hexValue),
			),
		),
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
	return top(nonZeroUnsignedNumber)(ctx, start)
}

/*
 * non_zero_unsigned_number33 ::= non_zero_decimal_digit { _ | decimal_digit}
 */
func nonZeroUnsignedNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return top(
		runes.Cons(
			nonZeroDecimalDigit,
			runes.Join(fn.Many0(decimalDigit_)),
		),
	)(ctx, start)
}

/*
 * real_number33 ::=
 *   fixed_point_number
 *   | unsigned_number [ . unsigned_number ] exp [ sign ] unsigned_number
 */
func RealNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.RealNumber, error) {
	return top(
		fn.Alt(
			to[ast.RealNumber](FloatingPointNumber),
			to[ast.RealNumber](FixedPointNumber),
		),
	)(ctx, start)
}

/*
 * real_number ::=
 *   ...
 *   | unsigned_number [ . unsigned_number ] exp [ sign ] unsigned_number
 *   ...
 */
func FloatingPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.FloatingPointNumber, error) {
	res := &ast.FloatingPointNumber{}
	return top(
		token(res, floatingPointNumber),
	)(ctx, start)
}

/*
 * real_number ::=
 *   ...
 *   | unsigned_number [ . unsigned_number ] exp [ sign ] unsigned_number
 *   ...
 */
func floatingPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return top(
		concat(
			unsignedNumber,
			fn.Opt(fn.Preceded(runes.Rune('.'), unsignedNumber)),
			exp,
			fn.Opt(sign),
			unsignedNumber,
		),
	)(ctx, start)
}

/*
 * fixed_point_number33 ::= unsigned_number . unsigned_number
 */
func FixedPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.FixedPointNumber, error) {
	res := &ast.FixedPointNumber{}
	return top(
		token(res, fixedPointNumber),
	)(ctx, start)
}

func fixedPointNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return top(
		concat(
			unsignedNumber,
			runes.Tag("."),
			unsignedNumber,
		),
	)(ctx, start)
}

/*
 * exp ::= e | E
 */
var exp = runes.Recognize(runes.OneOf("eE"))

/*
 * unsigned_number33 ::= decimal_digit { _ | decimal_digit }
 */
func UnsignedNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.UnsignedNumber, error) {
	res := &ast.UnsignedNumber{}
	return top(
		// TODO(justindubs): capture whitespace
		word(
			token(res, unsignedNumber),
		),
	)(ctx, start)
}

func unsignedNumber(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return top(
		runes.Cons(
			decimalDigit,
			runes.Join(fn.Many0(decimalDigit_)),
		),
	)(ctx, start)
}

/*
 * binary_value33 ::= binary_digit { _ | binary_digit }
 */
func binaryValue(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return top(
		runes.Cons(
			binaryDigit,
			runes.Join(fn.Many0(binaryDigit_)),
		),
	)(ctx, start)
}

/*
 * octal_value33 ::= octal_digit { _ | octal_digit }
 */
func octalValue(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return top(
		runes.Cons(
			octalDigit,
			runes.Join(fn.Many0(octalDigit_)),
		),
	)(ctx, start)
}

/*
 * hex_value33 ::= hex_digit { _ | hex_digit }
 */
func hexValue(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return top(
		runes.Cons(
			hexDigit,
			runes.Join(fn.Many0(hexDigit_)),
		),
	)(ctx, start)
}

/*
 * decimal_base33 ::= '[s|S]d | '[s|S]D
 */
func decimalBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return top(
		join(
			runes.Rune('\''),
			fn.Opt(runes.OneOf("sS")),
			runes.OneOf("dD"),
		),
	)(ctx, start)
}

/*
 * binary_base33 ::= '[s|S]b | '[s|S]B
 */
func binaryBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return top(
		join(
			runes.Rune('\''),
			fn.Opt(runes.OneOf("sS")),
			runes.OneOf("bB"),
		),
	)(ctx, start)
}

/*
 * octal_base33 ::= '[s|S]o | '[s|S]O
 */
func octalBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return top(
		join(
			runes.Rune('\''),
			fn.Opt(runes.OneOf("sS")),
			runes.OneOf("oO"),
		),
	)(ctx, start)
}

/*
 * hex_base33 ::= '[s|S]h | '[s|S]H
 */
func hexBase(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return top(
		join(
			runes.Rune('\''),
			fn.Opt(runes.OneOf("sS")),
			runes.OneOf("hH"),
		),
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
 * unbased_unsized_literal ::= '0 | '1 | 'z_or_x 48
 */
func UnbasedUnsizedLiteral(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.UnbasedUnsizedLiteral, error) {
	res := &ast.UnbasedUnsizedLiteral{}
	return top(
		token(res, unbasedUnsizedLiteral),
	)(ctx, start)
}

func unbasedUnsizedLiteral(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], string, error) {
	return top(
		word(
			join(
				runes.Rune('\''),
				runes.OneOf("01xXzZ"),
			),
		),
	)(ctx, start)
}

//
// A.8.8 Strings
//

/*
 * string_literal ::= " { Any_ASCII_Characters } "
 */
func StringLiteral(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.StringLiteral, error) {
	res := &ast.StringLiteral{}
	return top(
		// TODO(justindubs): capture whitespace
		word(
			token(res,
				fn.Surrounded(runes.Rune('"'), runes.Rune('"'), stringContents),
			),
		),
	)(ctx, start)
}

func stringContents(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], []rune, error) {
	return fn.Many0(
		fn.Alt(
			fn.Preceded(fn.Satisfy(func(r rune) bool { return r == '\\' }), fn.Any[rune]),
			fn.Satisfy(func(r rune) bool { return r != '\\' && r != '"' }),
		),
	)(ctx, start)
}
