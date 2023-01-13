package grammar

//
// A.9 General
//

//
// A.9.1 Attributes
//

/*
 * attribute_instance ::= (* attr_spec { , attr_spec } *)
 */

/*
 * attr_spec ::= attr_name [ = constant_expression ]
 */

/*
 * attr_name ::= identifier
 */

//
// A.9.2 Comments
//

/*
 * comment ::=
 *   one_line_comment
 *   | block_comment
 */

/*
 * one_line_comment ::= // comment_text \n
 */

//
// block_comment ::= /* comment_text */
//

/*
 * comment_text ::= { Any_ASCII_character }
 */

//
// A.9.3 Identifiers
//

/*
 * array_identifier ::= identifier
 */

/*
 * block_identifier ::= identifier
 */

/*
 * bin_identifier ::= identifier
 */

/*
 * c_identifier49 ::= [ a-zA-Z_ ] { [ a-zA-Z0-9_ ] }
 */

/*
 * cell_identifier ::= identifier
 */

/*
 * checker_identifier ::= identifier
 */

/*
 * class_identifier ::= identifier
 */

/*
 * class_variable_identifier ::= variable_identifier
 */

/*
 * clocking_identifier ::= identifier
 */

/*
 * config_identifier ::= identifier
 */

/*
 * const_identifier ::= identifier
 */

/*
 * constraint_identifier ::= identifier
 */

/*
 * covergroup_identifier ::= identifier
 */

/*
 * covergroup_variable_identifier ::= variable_identifier
 */

/*
 * cover_point_identifier ::= identifier
 */

/*
 * cross_identifier ::= identifier
 */

/*
 * dynamic_array_variable_identifier ::= variable_identifier
 */

/*
 * enum_identifier ::= identifier
 */

/*
 * escaped_identifier ::= \ {any_printable_ASCII_character_except_white_space} white_space
 */

/*
 * formal_identifier ::= identifier
 */

/*
 * formal_port_identifier ::= identifier
 */

/*
 * function_identifier ::= identifier
 */

/*
 * generate_block_identifier ::= identifier
 */

/*
 * genvar_identifier ::= identifier
 */

/*
 * hierarchical_array_identifier ::= hierarchical_identifier
 */

/*
 * hierarchical_block_identifier ::= hierarchical_identifier
 */

/*
 * hierarchical_event_identifier ::= hierarchical_identifier
 */

/*
 * hierarchical_identifier ::= [ $root . ] { identifier constant_bit_select . } identifier
 */

/*
 * hierarchical_net_identifier ::= hierarchical_identifier
 */

/*
 * hierarchical_parameter_identifier ::= hierarchical_identifier
 */

/*
 * hierarchical_property_identifier ::= hierarchical_identifier
 */

/*
 * hierarchical_sequence_identifier ::= hierarchical_identifier
 */

/*
 * hierarchical_task_identifier ::= hierarchical_identifier
 */

/*
 * hierarchical_tf_identifier ::= hierarchical_identifier
 */

/*
 * hierarchical_variable_identifier ::= hierarchical_identifier
 */

/*
 * identifier ::=
 *   simple_identifier
 *   | escaped_identifier
 */

/*
 * index_variable_identifier ::= identifier
 */

/*
 * interface_identifier ::= identifier
 */

/*
 * interface_instance_identifier ::= identifier
 */

/*
 * inout_port_identifier ::= identifier
 */

/*
 * input_port_identifier ::= identifier
 */

/*
 * instance_identifier ::= identifier
 */

/*
 * library_identifier ::= identifier
 */

/*
 * member_identifier ::= identifier
 */

/*
 * method_identifier ::= identifier
 */

/*
 * modport_identifier ::= identifier
 */

/*
 * module_identifier ::= identifier
 */

/*
 * net_identifier ::= identifier
 */

/*
 * net_type_identifier ::= identifier
 */

/*
 * output_port_identifier ::= identifier
 */

/*
 * package_identifier ::= identifier
 */

/*
 * package_scope ::=
 *   package_identifier ::
 *   | $unit ::
 */

/*
 * parameter_identifier ::= identifier
 */

/*
 * port_identifier ::= identifier
 */

/*
 * production_identifier ::= identifier
 */

/*
 * program_identifier ::= identifier
 */

/*
 * property_identifier ::= identifier
 */

/*
 * ps_class_identifier ::= [ package_scope ] class_identifier
 */

/*
 * ps_covergroup_identifier ::= [ package_scope ] covergroup_identifier
 */

/*
 * ps_checker_identifier ::= [ package_scope ] checker_identifier
 */

/*
 * ps_identifier ::= [ package_scope ] identifier
 */

/*
 * ps_or_hierarchical_array_identifier ::=
 *   [ implicit_class_handle . | class_scope | package_scope ] hierarchical_array_identifier
 */

/*
 * ps_or_hierarchical_net_identifier ::= [ package_scope ] net_identifier | hierarchical_net_identifier
 */

/*
 * ps_or_hierarchical_property_identifier ::=
 *   [ package_scope ] property_identifier
 *   | hierarchical_property_identifier
 */

/*
 * ps_or_hierarchical_sequence_identifier ::=
 *   [ package_scope ] sequence_identifier
 *   | hierarchical_sequence_identifier
 */

/*
 * ps_or_hierarchical_tf_identifier ::=
 *   [ package_scope ] tf_identifier
 *   | hierarchical_tf_identifier
 */

/*
 * ps_parameter_identifier ::=
 *   [ package_scope | class_scope ] parameter_identifier
 *   | { generate_block_identifier [ [ constant_expression ] ] . } parameter_identifier
 */

/*
 * ps_type_identifier ::= [ local ::43 | package_scope | class_scope ] type_identifier
 */

/*
 * sequence_identifier ::= identifier
 */

/*
 * signal_identifier ::= identifier
 */

/*
 * simple_identifier49 ::= [ a-zA-Z_ ] { [ a-zA-Z0-9_$ ] }
 */

/*
 * specparam_identifier ::= identifier
 */

/*
 * system_tf_identifier50 ::= $[ a-zA-Z0-9_$ ]{ [ a-zA-Z0-9_$ ] }
 */

/*
 * task_identifier ::= identifier
 */

/*
 * tf_identifier ::= identifier
 */

/*
 * terminal_identifier ::= identifier
 */

/*
 * topmodule_identifier ::= identifier
 */

/*
 * type_identifier ::= identifier
 */

/*
 * udp_identifier ::= identifier
 */

/*
 * variable_identifier ::= identifier
 */

//
// A.9.4 White space
//

/*
 * white_space ::= space | tab | newline | eof
 */
