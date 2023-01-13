package grammar

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

/*
 * time_literal44 ::=
 *   unsigned_number time_unit
 *   | fixed_point_number time_unit
 */

/*
 * time_unit ::= s | ms | us | ns | ps | fs
 */

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

/*
 * binary_operator ::=
 *   + | - | * | / | % | == | != | === | !== | ==? | !=? | && | || | **
 *   | < | <= | > | >= | & | | | ^ | ^~ | ~^ | >> | << | >>> | <<<
 *   | -> | <->
 */

/*
 * inc_or_dec_operator ::= ++ | --
 */

/*
 * unary_module_path_operator ::=
 *   ! | ~ | & | ~& | | | ~| | ^ | ~^ | ^~
 */

/*
 * binary_module_path_operator ::=
 *   == | != | && | || | & | | | ^ | ^~ | ~^
 */

//
// A.8.7 Numbers
//

/*
 * number ::=
 *   integral_number
 *   | real_number
 */

/*
 * integral_number ::=
 *   decimal_number
 *   | octal_number
 *   | binary_number
 *   | hex_number
 */

/*
 * decimal_number ::=
 *   unsigned_number
 *   | [ size ] decimal_base unsigned_number
 *   | [ size ] decimal_base x_digit { _ }
 *   | [ size ] decimal_base z_digit { _ }
 */

/*
 * binary_number ::= [ size ] binary_base binary_value
 */

/*
 * octal_number ::= [ size ] octal_base octal_value
 */

/*
 * hex_number ::= [ size ] hex_base hex_value
 */

/*
 * sign ::= + | -
 */

/*
 * size ::= non_zero_unsigned_number
 */

/*
 * non_zero_unsigned_number33 ::= non_zero_decimal_digit { _ | decimal_digit}
 */

/*
 * real_number33 ::=
 *   fixed_point_number
 *   | unsigned_number [ . unsigned_number ] exp [ sign ] unsigned_number
 */

/*
 * fixed_point_number33 ::= unsigned_number . unsigned_number
 */

/*
 * exp ::= e | E
 */

/*
 * unsigned_number33 ::= decimal_digit { _ | decimal_digit }
 */

/*
 * binary_value33 ::= binary_digit { _ | binary_digit }
 */

/*
 * octal_value33 ::= octal_digit { _ | octal_digit }
 */

/*
 * hex_value33 ::= hex_digit { _ | hex_digit }
 */

/*
 * decimal_base33 ::= '[s|S]d | '[s|S]D
 */

/*
 * binary_base33 ::= '[s|S]b | '[s|S]B
 */

/*
 * octal_base33 ::= '[s|S]o | '[s|S]O
 */

/*
 * hex_base33 ::= '[s|S]h | '[s|S]H
 */

/*
 * non_zero_decimal_digit ::= 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
 */

/*
 * decimal_digit ::= 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
 */

/*
 * binary_digit ::= x_digit | z_digit | 0 | 1
 */

/*
 * octal_digit ::= x_digit | z_digit | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7
 */

/*
 * hex_digit ::= x_digit | z_digit | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | a | b | c | d | e | f | A | B | C | D | E | F
 */

/*
 * x_digit ::= x | X
 */

/*
 * z_digit ::= z | Z | ?
 */

/*
 * unbased_unsized_literal ::= '0 | '1 | 'z_or_x 48
 */

//
// A.8.8 Strings
//

/*
 * string_literal ::= " { Any_ASCII_Characters } "
 */
