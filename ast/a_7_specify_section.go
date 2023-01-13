package ast

//
// A.7 Specify section
//

//
// A.7.1 Specify block declaration
//

/*
 * specify_block ::= specify { specify_item } endspecify
 */

/*
 * specify_item ::=
 *   specparam_declaration
 *   | pulsestyle_declaration
 *   | showcancelled_declaration
 *   | path_declaration
 *   | system_timing_check
 */

/*
 * pulsestyle_declaration ::=
 *   pulsestyle_onevent list_of_path_outputs ;
 *   | pulsestyle_ondetect list_of_path_outputs ;
 */

/*
 * showcancelled_declaration ::=
 *   showcancelled list_of_path_outputs ;
 *   | noshowcancelled list_of_path_outputs ;
 */

//
// A.7.2 Specify path declarations
//

/*
 * path_declaration ::=
 *   simple_path_declaration ;
 *   | edge_sensitive_path_declaration ;
 *   | state_dependent_path_declaration ;
 */

/*
 * simple_path_declaration ::=
 *   parallel_path_description = path_delay_value
 *   | full_path_description = path_delay_value
 */

/*
 * parallel_path_description ::=
 *   ( specify_input_terminal_descriptor [ polarity_operator ] => specify_output_terminal_descriptor )
 */

/*
 * full_path_description ::=
 *   ( list_of_path_inputs [ polarity_operator ] *> list_of_path_outputs )
 */

/*
 * list_of_path_inputs ::=
 *   specify_input_terminal_descriptor { , specify_input_terminal_descriptor }
 */

/*
 * list_of_path_outputs ::=
 *   specify_output_terminal_descriptor { , specify_output_terminal_descriptor }
 */

//
// A.7.3 Specify block terminals
//

/*
 * specify_input_terminal_descriptor ::=
 *   input_identifier [ [ constant_range_expression ] ]
 */

/*
 * specify_output_terminal_descriptor ::=
 *   output_identifier [ [ constant_range_expression ] ]
 */

/*
 * input_identifier ::= input_port_identifier | inout_port_identifier | interface_identifier.port_identifier
 */

/*
 * output_identifier ::= output_port_identifier | inout_port_identifier | interface_identifier.port_identifier
 */

//
// A.7.4 Specify path delays
//

/*
 * path_delay_value ::=
 *   list_of_path_delay_expressions
 *   | ( list_of_path_delay_expressions )
 */

/*
 * list_of_path_delay_expressions ::=
 *   t_path_delay_expression
 *   | trise_path_delay_expression , tfall_path_delay_expression
 *   | trise_path_delay_expression , tfall_path_delay_expression , tz_path_delay_expression
 *   | t01_path_delay_expression , t10_path_delay_expression , t0z_path_delay_expression ,
 *   tz1_path_delay_expression , t1z_path_delay_expression , tz0_path_delay_expression
 *   | t01_path_delay_expression , t10_path_delay_expression , t0z_path_delay_expression ,
 *   tz1_path_delay_expression , t1z_path_delay_expression , tz0_path_delay_expression ,
 *   t0x_path_delay_expression , tx1_path_delay_expression , t1x_path_delay_expression ,
 *   tx0_path_delay_expression , txz_path_delay_expression , tzx_path_delay_expression
 */

/*
 * t_path_delay_expression ::= path_delay_expression
 */

/*
 * trise_path_delay_expression ::= path_delay_expression
 */

/*
 * tfall_path_delay_expression ::= path_delay_expression
 */

/*
 * tz_path_delay_expression ::= path_delay_expression
 */

/*
 * t01_path_delay_expression ::= path_delay_expression
 */

/*
 * t10_path_delay_expression ::= path_delay_expression
 */

/*
 * t0z_path_delay_expression ::= path_delay_expression
 */

/*
 * tz1_path_delay_expression ::= path_delay_expression
 */

/*
 * t1z_path_delay_expression ::= path_delay_expression
 */

/*
 * tz0_path_delay_expression ::= path_delay_expression
 */

/*
 * t0x_path_delay_expression ::= path_delay_expression
 */

/*
 * tx1_path_delay_expression ::= path_delay_expression
 */

/*
 * t1x_path_delay_expression ::= path_delay_expression
 */

/*
 * tx0_path_delay_expression ::= path_delay_expression
 */

/*
 * txz_path_delay_expression ::= path_delay_expression
 */

/*
 * tzx_path_delay_expression ::= path_delay_expression
 */

/*
 * path_delay_expression ::= constant_mintypmax_expression
 */

/*
 * edge_sensitive_path_declaration ::=
 *   parallel_edge_sensitive_path_description = path_delay_value
 *   | full_edge_sensitive_path_description = path_delay_value
 */

/*
 * parallel_edge_sensitive_path_description ::=
 *   ( [ edge_identifier ] specify_input_terminal_descriptor [ polarity_operator ] =>
 *   ( specify_output_terminal_descriptor [ polarity_operator ] : data_source_expression ) )
 */

/*
 * full_edge_sensitive_path_description ::=
 *   ( [ edge_identifier ] list_of_path_inputs [ polarity_operator ] *>
 *   ( list_of_path_outputs [ polarity_operator ] : data_source_expression ) )
 */

/*
 * data_source_expression ::= expression
 */

/*
 * edge_identifier ::= posedge | negedge | edge
 */

/*
 * state_dependent_path_declaration ::=
 *   if ( module_path_expression ) simple_path_declaration
 *   | if ( module_path_expression ) edge_sensitive_path_declaration
 *   | ifnone simple_path_declaration
 */

/*
 * polarity_operator ::= + | -
 */

//
// A.7.5 System timing checks
//

//
// A.7.5.1 System timing check commands
//

/*
 * system_timing_check ::=
 *   $setup_timing_check
 *   | $hold_timing_check
 *   | $setuphold_timing_check
 *   | $recovery_timing_check
 *   | $removal_timing_check
 *   | $recrem_timing_check
 *   | $skew_timing_check
 *   | $timeskew_timing_check
 *   | $fullskew_timing_check
 *   | $period_timing_check
 *   | $width_timing_check
 *   | $nochange_timing_check
 */

/*
 * $setup_timing_check ::=
 *   $setup ( data_event , reference_event , timing_check_limit [ , [ notifier ] ] ) ;
 */

/*
 * $hold_timing_check ::=
 *   $hold ( reference_event , data_event , timing_check_limit [ , [ notifier ] ] ) ;
 */

/*
 * $setuphold_timing_check ::=
 *   $setuphold ( reference_event , data_event , timing_check_limit , timing_check_limit
 *   [ , [ notifier ] [ , [ timestamp_condition ] [ , [ timecheck_condition ]
 *   [ , [ delayed_reference ] [ , [ delayed_data ] ] ] ] ] ] ) ;
 */

/*
 * $recovery_timing_check ::=
 *   $recovery ( reference_event , data_event , timing_check_limit [ , [ notifier ] ] ) ;
 */

/*
 * $removal_timing_check ::=
 *   $removal ( reference_event , data_event , timing_check_limit [ , [ notifier ] ] ) ;
 */

/*
 * $recrem_timing_check ::=
 *   $recrem ( reference_event , data_event , timing_check_limit , timing_check_limit
 *   [ , [ notifier ] [ , [ timestamp_condition ] [ , [ timecheck_condition ]
 *   [ , [ delayed_reference ] [ , [ delayed_data ] ] ] ] ] ] ) ;
 */

/*
 * $skew_timing_check ::=
 *   $skew ( reference_event , data_event , timing_check_limit [ , [ notifier ] ] ) ;
 */

/*
 * $timeskew_timing_check ::=
 *   $timeskew ( reference_event , data_event , timing_check_limit
 *   [ , [ notifier ] [ , [ event_based_flag ] [ , [ remain_active_flag ] ] ] ] ) ;
 */

/*
 * $fullskew_timing_check ::=
 *   $fullskew ( reference_event , data_event , timing_check_limit , timing_check_limit
 *   [ , [ notifier ] [ , [ event_based_flag ] [ , [ remain_active_flag ] ] ] ] ) ;
 */

/*
 * $period_timing_check ::=
 *   $period ( controlled_reference_event , timing_check_limit [ , [ notifier ] ] ) ;
 */

/*
 * $width_timing_check ::=
 *   $width ( controlled_reference_event , timing_check_limit , threshold [ , [ notifier ] ] ) ;
 */

/*
 * $nochange_timing_check ::=
 *   $nochange ( reference_event , data_event , start_edge_offset , end_edge_offset [ , [ notifier ] ] );
 */

//
// A.7.5.2 System timing check command arguments
//

/*
 * timecheck_condition ::= mintypmax_expression
 */

/*
 * controlled_reference_event ::= controlled_timing_check_event
 */

/*
 * data_event ::= timing_check_event
 */

/*
 * delayed_data ::=
 *   terminal_identifier
 *   | terminal_identifier [ constant_mintypmax_expression ]
 */

/*
 * delayed_reference ::=
 *   terminal_identifier
 *   | terminal_identifier [ constant_mintypmax_expression ]
 */

/*
 * end_edge_offset ::= mintypmax_expression
 */

/*
 * event_based_flag ::= constant_expression
 */

/*
 * notifier ::= variable_identifier
 */

/*
 * reference_event ::= timing_check_event
 */

/*
 * remain_active_flag ::= constant_mintypmax_expression
 */

/*
 * timestamp_condition ::= mintypmax_expression
 */

/*
 * start_edge_offset ::= mintypmax_expression
 */

/*
 * threshold ::= constant_expression
 */

/*
 * timing_check_limit ::= expression
 */

//
// A.7.5.3 System timing check event definitions
//

/*
 * timing_check_event ::=
 *   [timing_check_event_control] specify_terminal_descriptor [ &&& timing_check_condition ]
 */

/*
 * controlled_timing_check_event ::=
 *   timing_check_event_control specify_terminal_descriptor [ &&& timing_check_condition ]
 */

/*
 * timing_check_event_control ::=
 *   posedge
 *   | negedge
 *   | edge
 *   | edge_control_specifier
 */

/*
 * specify_terminal_descriptor ::=
 *   specify_input_terminal_descriptor
 *   | specify_output_terminal_descriptor
 */

/*
 * edge_control_specifier ::= edge [ edge_descriptor { , edge_descriptor } ]
 */

/*
 * edge_descriptor33 ::= 01 | 10 | z_or_x zero_or_one | zero_or_one z_or_x
 */

/*
 * zero_or_one ::= 0 | 1
 */

/*
 * z_or_x ::= x | X | z | Z
 */

/*
 * timing_check_condition ::=
 *   scalar_timing_check_condition
 *   | ( scalar_timing_check_condition )
 */

/*
 * scalar_timing_check_condition ::=
 *   expression
 *   | ~ expression
 *   | expression == scalar_constant
 *   | expression === scalar_constant
 *   | expression != scalar_constant
 *   | expression !== scalar_constant
 */

/*
 * scalar_constant ::= 1'b0 | 1'b1 | 1'B0 | 1'B1 | 'b0 | 'b1 | 'B0 | 'B1 | 1 | 0
 */
