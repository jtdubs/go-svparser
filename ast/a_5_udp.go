package ast

//
// A.5 UDP declaration and instantiation
//

//
// A.5.1 UDP declaration
//

/*
 * udp_nonansi_declaration ::=
 *   { attribute_instance } primitive udp_identifier ( udp_port_list ) ;
 */

/*
 * udp_ansi_declaration ::=
 *   { attribute_instance } primitive udp_identifier ( udp_declaration_port_list ) ;
 */

/*
 * udp_declaration ::=
 *   udp_nonansi_declaration udp_port_declaration { udp_port_declaration }
 *   udp_body
 *   endprimitive [ : udp_identifier ]
 *   | udp_ansi_declaration
 *   udp_body
 *   endprimitive [ : udp_identifier ]
 *   | extern udp_nonansi_declaration
 *   | extern udp_ansi_declaration
 *   | { attribute_instance } primitive udp_identifier ( .* ) ;
 *   { udp_port_declaration }
 *   udp_body
 *   endprimitive [ : udp_identifier ]
 */

//
// A.5.2 UDP ports
//

/*
 * udp_port_list ::= output_port_identifier , input_port_identifier { , input_port_identifier }
 */

/*
 * udp_declaration_port_list ::= udp_output_declaration , udp_input_declaration { , udp_input_declaration }
 */

/*
 * udp_port_declaration ::=
 *   udp_output_declaration ;
 *   | udp_input_declaration ;
 *   | udp_reg_declaration ;
 */

/*
 * udp_output_declaration ::=
 *   { attribute_instance } output port_identifier
 *   | { attribute_instance } output reg port_identifier [ = constant_expression ]
 */

/*
 * udp_input_declaration ::= { attribute_instance } input list_of_udp_port_identifiers
 */

/*
 * udp_reg_declaration ::= { attribute_instance } reg variable_identifier
 */

//
// A.5.3 UDP body
//

/*
 * udp_body ::= combinational_body | sequential_body
 */

/*
 * combinational_body ::= table combinational_entry { combinational_entry } endtable
 */

/*
 * combinational_entry ::= level_input_list : output_symbol ;
 */

/*
 * sequential_body ::= [ udp_initial_statement ] table sequential_entry { sequential_entry } endtable
 */

/*
 * udp_initial_statement ::= initial output_port_identifier = init_val ;
 */

/*
 * init_val ::= 1'b0 | 1'b1 | 1'bx | 1'bX | 1'B0 | 1'B1 | 1'Bx | 1'BX | 1 | 0
 */

/*
 * sequential_entry ::= seq_input_list : current_state : next_state ;
 */

/*
 * seq_input_list ::= level_input_list | edge_input_list
 */

/*
 * level_input_list ::= level_symbol { level_symbol }
 */

/*
 * edge_input_list ::= { level_symbol } edge_indicator { level_symbol }
 */

/*
 * edge_indicator ::= ( level_symbol level_symbol ) | edge_symbol
 */

/*
 * current_state ::= level_symbol
 */

/*
 * next_state ::= output_symbol | -
 */

/*
 * output_symbol ::= 0 | 1 | x | X
 */

/*
 * level_symbol ::= 0 | 1 | x | X | ? | b | B
 */

/*
 * edge_symbol ::= r | R | f | F | p | P | n | N | *
 */

//
// A.5.4 UDP instantiation
//

/*
 * udp_instantiation ::= udp_identifier [ drive_strength ] [ delay2 ] udp_instance { , udp_instance } ;
 */

/*
 * udp_instance ::= [ name_of_instance ] ( output_terminal , input_terminal { , input_terminal } )
 */
