package ast

//
// A.4 Instantiations
//

//
// A.4.1 Instantiation
//

//
// A.4.1.1 Module instantiation
//

/*
 * module_instantiation ::=
 *   module_identifier [ parameter_value_assignment ] hierarchical_instance { , hierarchical_instance } ;
 */

/*
 * parameter_value_assignment ::= # ( [ list_of_parameter_assignments ] )
 */

/*
 * list_of_parameter_assignments ::=
 *   ordered_parameter_assignment { , ordered_parameter_assignment }
 *   | named_parameter_assignment { , named_parameter_assignment }
 */

/*
 * ordered_parameter_assignment ::= param_expression
 */

/*
 * named_parameter_assignment ::= . parameter_identifier ( [ param_expression ] )
 */

/*
 * hierarchical_instance ::= name_of_instance ( [ list_of_port_connections ] )
 */

/*
 * name_of_instance ::= instance_identifier { unpacked_dimension }
 */

/*
 * list_of_port_connections29 ::=
 *   ordered_port_connection { , ordered_port_connection }
 *   | named_port_connection { , named_port_connection }
 */

/*
 * ordered_port_connection ::= { attribute_instance } [ expression ]
 */

/*
 * named_port_connection ::=
 *   { attribute_instance } . port_identifier [ ( [ expression ] ) ]
 *   | { attribute_instance } .*
 */

//
// A.4.1.2 Interface instantiation
//

/*
 * interface_instantiation ::=
 *   interface_identifier [ parameter_value_assignment ] hierarchical_instance { , hierarchical_instance } ;
 */

//
// A.4.1.3 Program instantiation
//

/*
 * program_instantiation ::=
 *   program_identifier [ parameter_value_assignment ] hierarchical_instance { , hierarchical_instance } ;
 */

//
// A.4.1.4 Checker instantiation
//

/*
 * checker_instantiation ::=
 *   ps_checker_identifier name_of_instance ( [list_of_checker_port_connections] ) ;
 */

/*
 * list_of_checker_port_connections29 ::=
 *   ordered_checker_port_connection { , ordered_checker_port_connection }
 *   | named_checker_port_connection { , named_checker_port_connection }
 */

/*
 * ordered_checker_port_connection ::= { attribute_instance } [ property_actual_arg ]
 */

/*
 * named_checker_port_connection ::=
 *   { attribute_instance } . formal_port_identifier [ ( [ property_actual_arg ] ) ]
 *   | { attribute_instance } .*
 */

//
// A.4.2 Generated instantiation
//

/*
 * generate_region ::=
 *   generate { generate_item } endgenerate
 */

/*
 * loop_generate_construct ::=
 *   for ( genvar_initialization ; genvar_expression ; genvar_iteration )
 *   generate_block
 */

/*
 * genvar_initialization ::=
 *   [ genvar ] genvar_identifier = constant_expression
 */

/*
 * genvar_iteration ::=
 *   genvar_identifier assignment_operator genvar_expression
 *   | inc_or_dec_operator genvar_identifier
 *   | genvar_identifier inc_or_dec_operator
 */

/*
 * conditional_generate_construct ::=
 *   if_generate_construct
 *   | case_generate_construct
 */

/*
 * if_generate_construct ::=
 *   if ( constant_expression ) generate_block [ else generate_block ]
 */

/*
 * case_generate_construct ::=
 *   case ( constant_expression ) case_generate_item { case_generate_item } endcase
 */

/*
 * case_generate_item ::=
 *   constant_expression { , constant_expression } : generate_block
 *   | default [ : ] generate_block
 */

/*
 * generate_block ::=
 *   generate_item
 *   | [ generate_block_identifier : ] begin [ : generate_block_identifier ]
 *   { generate_item }
 *   end [ : generate_block_identifier ]
 */

/*
 * generate_item30 ::=
 *   module_or_generate_item
 *   | interface_or_generate_item
 *   | checker_or_generate_item
 */
