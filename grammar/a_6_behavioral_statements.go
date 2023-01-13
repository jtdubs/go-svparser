package grammar

//
// A.6 Behavioral statements
//

//
// A.6.1 Continuous assignment and net alias statements
//

/*
 * continuous_assign ::=
 *   assign [ drive_strength ] [ delay3 ] list_of_net_assignments ;
 *   | assign [ delay_control ] list_of_variable_assignments ;
 */

/*
 * continuous_assign ::=
 */

/*
 * list_of_net_assignments ::= net_assignment { , net_assignment }
 */

/*
 * list_of_variable_assignments ::= variable_assignment { , variable_assignment }
 */

/*
 * net_alias ::= alias net_lvalue = net_lvalue { = net_lvalue } ;
 */

/*
 * net_assignment ::= net_lvalue = expression
 */

//
// A.6.2 Procedural blocks and assignments
//

/*
 * initial_construct ::= initial statement_or_null
 */

/*
 * always_construct ::= always_keyword statement
 */

/*
 * always_keyword ::= always | always_comb | always_latch | always_ff
 */

/*
 * final_construct ::= final function_statement
 */

/*
 * blocking_assignment ::=
 *   variable_lvalue = delay_or_event_control expression
 *   | nonrange_variable_lvalue = dynamic_array_new
 *   | [ implicit_class_handle . | class_scope | package_scope ] hierarchical_variable_identifier
 *   select = class_new
 *   | operator_assignment
 */

/*
 * operator_assignment ::= variable_lvalue assignment_operator expression
 */

/*
 * assignment_operator ::=
 *   = | += | -= | *= | /= | %= | &= | |= | ^= | <<= | >>= | <<<= | >>>=
 */

/*
 * nonblocking_assignment ::=
 *   variable_lvalue <= [ delay_or_event_control ] expression
 */

/*
 * procedural_continuous_assignment ::=
 *   assign variable_assignment
 *   | deassign variable_lvalue
 *   | force variable_assignment
 *   | force net_assignment
 *   | release variable_lvalue
 *   | release net_lvalue
 */

/*
 * variable_assignment ::= variable_lvalue = expression
 */

//
// A.6.3 Parallel and sequential blocks
//

/*
 * action_block ::=
 *   statement_or_null
 *   | [ statement ] else statement_or_null
 */

/*
 * seq_block ::=
 *   begin [ : block_identifier ] { block_item_declaration } { statement_or_null }
 *   end [ : block_identifier ]
 */

/*
 * par_block ::=
 *   fork [ : block_identifier ] { block_item_declaration } { statement_or_null }
 *   join_keyword [ : block_identifier ]
 */

/*
 * join_keyword ::= join | join_any | join_none
 */

//
// A.6.4 Statements
//

/*
 * statement_or_null ::=
 *   statement
 *   | { attribute_instance } ;
 */

/*
 * statement ::= [ block_identifier : ] { attribute_instance } statement_item
 */

/*
 * statement_item ::=
 *   blocking_assignment ;
 *   | nonblocking_assignment ;
 *   | procedural_continuous_assignment ;
 *   | case_statement
 *   | conditional_statement
 *   | inc_or_dec_expression ;
 *   | subroutine_call_statement
 *   | disable_statement
 *   | event_trigger
 *   | loop_statement
 *   | jump_statement
 *   | par_block
 *   | procedural_timing_control_statement
 *   | seq_block
 *   | wait_statement
 *   | procedural_assertion_statement
 *   | clocking_drive ;
 *   | randsequence_statement
 *   | randcase_statement
 *   | expect_property_statement
 */

/*
 * function_statement ::= statement
 */

/*
 * function_statement_or_null ::=
 *   function_statement
 *   | { attribute_instance } ;
 */

/*
 * variable_identifier_list ::= variable_identifier { , variable_identifier }
 */

//
// A.6.5 Timing control statements
//

/*
 * procedural_timing_control_statement ::=
 *   procedural_timing_control statement_or_null
 */

/*
 * delay_or_event_control ::=
 *   delay_control
 *   | event_control
 *   | repeat ( expression ) event_control
 */

/*
 * delay_control ::=
 *   # delay_value
 *   | # ( mintypmax_expression )
 */

/*
 * event_control ::=
 *   @ hierarchical_event_identifier
 *   | @ ( event_expression )
 *   | @*
 *   | @ (*)
 *   | @ ps_or_hierarchical_sequence_identifier
 */

/*
 * event_expression31 ::=
 *   [ edge_identifier ] expression [ iff expression ]
 *   | sequence_instance [ iff expression ]
 *   | event_expression or event_expression
 *   | event_expression , event_expression
 *   | ( event_expression )
 */

/*
 * procedural_timing_control ::=
 *   delay_control
 *   | event_control
 *   | cycle_delay
 */

/*
 * jump_statement ::=
 *   return [ expression ] ;
 *   | break ;
 *   | continue ;
 */

/*
 * wait_statement ::=
 *   wait ( expression ) statement_or_null
 *   | wait fork ;
 *   | wait_order ( hierarchical_identifier { , hierarchical_identifier } ) action_block
 */

/*
 * event_trigger ::=
 *   -> hierarchical_event_identifier ;
 *   |->> [ delay_or_event_control ] hierarchical_event_identifier ;
 */

/*
 * disable_statement ::=
 *   disable hierarchical_task_identifier ;
 *   | disable hierarchical_block_identifier ;
 *   | disable fork ;
 */

//
// A.6.6 Conditional statements
//

/*
 * conditional_statement ::=
 *   [ unique_priority ] if ( cond_predicate ) statement_or_null
 *   {else if ( cond_predicate ) statement_or_null }
 *   [ else statement_or_null ]
 */

/*
 * unique_priority ::= unique | unique0 | priority
 */

/*
 * cond_predicate ::=
 *   expression_or_cond_pattern { &&& expression_or_cond_pattern }
 */

/*
 * expression_or_cond_pattern ::=
 *   expression | cond_pattern
 */

/*
 * cond_pattern ::= expression matches pattern
 */

//
// A.6.7 Case statements
//

/*
 * case_statement ::=
 *   [ unique_priority ] case_keyword ( case_expression )
 *   case_item { case_item } endcase
 *   | [ unique_priority ] case_keyword (case_expression )matches
 *   case_pattern_item { case_pattern_item } endcase
 *   | [ unique_priority ] case ( case_expression ) inside
 *   case_inside_item { case_inside_item } endcase
 */

/*
 * case_keyword ::= case | casez | casex
 */

/*
 * case_expression ::= expression
 */

/*
 * case_item ::=
 *   case_item_expression { , case_item_expression } : statement_or_null
 *   | default [ : ] statement_or_null
 */

/*
 * case_pattern_item ::=
 *   pattern [ &&& expression ] : statement_or_null
 *   | default [ : ] statement_or_null
 */

/*
 * case_inside_item ::=
 *   open_range_list : statement_or_null
 *   | default [ : ] statement_or_null
 */

/*
 * case_item_expression ::= expression
 */

/*
 * randcase_statement ::=
 *   randcase randcase_item { randcase_item } endcase
 */

/*
 * randcase_item ::= expression : statement_or_null
 */

/*
 * open_range_list ::= open_value_range { , open_value_range }
 */

/*
 * open_value_range ::= value_range25
 */

//
// A.6.7.1 Patterns
//

/*
 * pattern ::=
 *   . variable_identifier
 *   | .*
 *   | constant_expression
 *   | tagged member_identifier [ pattern ]
 *   | '{ pattern { , pattern } }
 *   | '{ member_identifier : pattern { , member_identifier : pattern } }
 */

/*
 * assignment_pattern ::=
 *   '{ expression { , expression } }
 *   | '{ structure_pattern_key : expression { , structure_pattern_key : expression } }
 *   | '{ array_pattern_key : expression { , array_pattern_key : expression } }
 *   | '{ constant_expression { expression { , expression } } }
 */

/*
 * structure_pattern_key ::= member_identifier | assignment_pattern_key
 */

/*
 * array_pattern_key ::= constant_expression | assignment_pattern_key
 */

/*
 * assignment_pattern_key ::= simple_type | default
 */

/*
 * assignment_pattern_expression ::=
 *   [ assignment_pattern_expression_type ] assignment_pattern
 */

/*
 * assignment_pattern_expression_type ::=
 *   ps_type_identifier
 *   | ps_parameter_identifier
 *   | integer_atom_type
 *   | type_reference
 */

/*
 * constant_assignment_pattern_expression32 ::= assignment_pattern_expression
 */

/*
 * assignment_pattern_net_lvalue ::=
 *   '{ net_lvalue {, net_lvalue } }
 */

/*
 * assignment_pattern_variable_lvalue ::=
 *   '{ variable_lvalue {, variable_lvalue } }
 */

//
// A.6.8 Looping statements
//

/*
 * loop_statement ::=
 *   forever statement_or_null
 *   | repeat ( expression ) statement_or_null
 *   | while ( expression ) statement_or_null
 *   | for ( [ for_initialization ] ; [ expression ] ; [ for_step ] )
 *   statement_or_null
 *   | do statement_or_null while ( expression ) ;
 *   | foreach ( ps_or_hierarchical_array_identifier [ loop_variables ] ) statement
 */

/*
 * for_initialization ::=
 *   list_of_variable_assignments
 *   | for_variable_declaration { , for_variable_declaration }
 */

/*
 * for_variable_declaration ::=
 *   [ var ] data_type variable_identifier = expression { , variable_identifier = expression }14
 */

/*
 * for_step ::= for_step_assignment { , for_step_assignment }
 */

/*
 * for_step_assignment ::=
 *   operator_assignment
 *   | inc_or_dec_expression
 *   | function_subroutine_call
 */

/*
 * loop_variables ::= [ index_variable_identifier ] { , [ index_variable_identifier ] }
 */

//
// A.6.9 Subroutine call statements
//

/*
 * subroutine_call_statement ::=
 *   subroutine_call ;
 *   | void ' ( function_subroutine_call ) ;
 */

//
// A.6.10 Assertion statements
//

/*
 * assertion_item ::=
 *   concurrent_assertion_item
 *   | deferred_immediate_assertion_item
 */

/*
 * deferred_immediate_assertion_item ::= [ block_identifier : ] deferred_immediate_assertion_statement
 */

/*
 * procedural_assertion_statement ::=
 *   concurrent_assertion_statement
 *   | immediate_assertion_statement
 *   | checker_instantiation
 */

/*
 * immediate_assertion_statement ::=
 *   simple_immediate_assertion_statement
 *   | deferred_immediate_assertion_statement
 */

/*
 * simple_immediate_assertion_statement ::=
 *   simple_immediate_assert_statement
 *   | simple_immediate_assume_statement
 *   | simple_immediate_cover_statement
 */

/*
 * simple_immediate_assert_statement ::=
 *   assert ( expression ) action_block
 */

/*
 * simple_immediate_assume_statement ::=
 *   assume ( expression ) action_block
 */

/*
 * simple_immediate_cover_statement ::=
 *   cover ( expression ) statement_or_null
 */

/*
 * deferred_immediate_assertion_statement ::=
 *   deferred_immediate_assert_statement
 *   | deferred_immediate_assume_statement
 *   | deferred_immediate_cover_statement
 */

/*
 * deferred_immediate_assert_statement ::=
 *   assert #0 ( expression ) action_block
 *   | assert final ( expression ) action_block
 */

/*
 * deferred_immediate_assume_statement ::=
 *   assume #0 ( expression ) action_block
 *   | assume final ( expression ) action_block
 */

/*
 * deferred_immediate_cover_statement ::=
 *   cover #0 ( expression ) statement_or_null
 *   | cover final ( expression ) statement_or_null
 */

//
// A.6.11 Clocking block
//

/*
 * clocking_declaration ::= [ default ] clocking [ clocking_identifier ] clocking_event ;
 *   { clocking_item }
 *   endclocking [ : clocking_identifier ]
 *   | global clocking [ clocking_identifier ] clocking_event ; endclocking [ : clocking_identifier ]
 */

/*
 * clocking_event ::=
 *   @ identifier
 *   | @ ( event_expression )
 */

/*
 * clocking_item ::=
 *   default default_skew ;
 *   | clocking_direction list_of_clocking_decl_assign ;
 *   | { attribute_instance } assertion_item_declaration
 */

/*
 * default_skew ::=
 *   input clocking_skew
 *   | output clocking_skew
 *   | input clocking_skew output clocking_skew
 */

/*
 * clocking_direction ::=
 *   input [ clocking_skew ]
 *   | output [ clocking_skew ]
 *   | input [ clocking_skew ] output [ clocking_skew ]
 *   | inout
 */

/*
 * list_of_clocking_decl_assign ::= clocking_decl_assign { , clocking_decl_assign }
 */

/*
 * clocking_decl_assign ::= signal_identifier [ = expression ]
 */

/*
 * clocking_skew ::=
 *   edge_identifier [ delay_control ]
 *   | delay_control
 */

/*
 * clocking_drive ::=
 *   clockvar_expression <= [ cycle_delay ] expression
 */

/*
 * cycle_delay ::=
 *   ## integral_number
 *   | ## identifier
 *   | ## ( expression )
 */

/*
 * clockvar ::= hierarchical_identifier
 */

/*
 * clockvar_expression ::= clockvar select
 */

//
// A.6.12 Randsequence
//

/*
 * randsequence_statement ::= randsequence ( [ production_identifier ] )
 *   production { production }
 *   endsequence
 */

/*
 * production ::= [ data_type_or_void ] production_identifier [ ( tf_port_list ) ] : rs_rule { | rs_rule } ;
 */

/*
 * rs_rule ::= rs_production_list [ := weight_specification [ rs_code_block ] ]
 */

/*
 * rs_production_list ::=
 *   rs_prod { rs_prod }
 *   | rand join [ ( expression ) ] production_item production_item { production_item }
 */

/*
 * weight_specification ::=
 *   integral_number
 *   | ps_identifier
 *   | ( expression )
 */

/*
 * rs_code_block ::= { { data_declaration } { statement_or_null } }
 */

/*
 * rs_prod ::=
 *   production_item
 *   | rs_code_block
 *   | rs_if_else
 *   | rs_repeat
 *   | rs_case
 */

/*
 * production_item ::= production_identifier [ ( list_of_arguments ) ]
 */

/*
 * rs_if_else ::= if ( expression ) production_item [ else production_item ]
 */

/*
 * rs_repeat ::= repeat ( expression ) production_item
 */

/*
 * rs_case ::= case ( case_expression ) rs_case_item { rs_case_item } endcase
 */

/*
 * rs_case_item ::=
 *   case_item_expression { , case_item_expression } : production_item ;
 *   | default [ : ] production_item ;
 */
