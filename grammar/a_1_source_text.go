package grammar

//
// A.1 Source text
//

//
// A.1.1 Library source text
//

/*
 * library_text ::= { library_description }
 */

/*
 * library_description ::=
 *   library_declaration
 *   | include_statement
 *   | config_declaration
 *   | ;
 */

/*
 * library_declaration ::=
 *   library library_identifier file_path_spec { , file_path_spec }
 *   [ -incdir file_path_spec { , file_path_spec } ] ;
 */

/*
 * include_statement ::= include file_path_spec ;
 */

//
// A.1.2 SystemVerilog source text
//

/*
 * source_text ::= [ timeunits_declaration ] { description }
 */

/*
 * description ::=
 *   module_declaration
 *   | udp_declaration
 *   | interface_declaration
 *   | program_declaration
 *   | package_declaration
 *   | { attribute_instance } package_item
 *   | { attribute_instance } bind_directive
 *   | config_declaration
 */

/*
 * module_nonansi_header ::=
 *   { attribute_instance } module_keyword [ lifetime ] module_identifier
 *   { package_import_declaration } [ parameter_port_list ] list_of_ports ;
 */

/*
 * module_ansi_header ::=
 *   { attribute_instance } module_keyword [ lifetime ] module_identifier
 *   { package_import_declaration }1 [ parameter_port_list ] [ list_of_port_declarations ] ;
 */

/*
 * module_declaration ::=
 *   module_nonansi_header [ timeunits_declaration ] { module_item }
 *   endmodule [ : module_identifier ]
 *   | module_ansi_header [ timeunits_declaration ] { non_port_module_item }
 *   endmodule [ : module_identifier ]
 *   | { attribute_instance } module_keyword [ lifetime ] module_identifier ( .* ) ;
 *   [ timeunits_declaration ] { module_item } endmodule [ : module_identifier ]
 *   | extern module_nonansi_header
 *   | extern module_ansi_header
 */

/*
 * module_keyword ::= module | macromodule
 */

/*
 * interface_declaration ::=
 *   interface_nonansi_header [ timeunits_declaration ] { interface_item }
 *   endinterface [ : interface_identifier ]
 *   | interface_ansi_header [ timeunits_declaration ] { non_port_interface_item }
 *   endinterface [ : interface_identifier ]
 *   | { attribute_instance } interface interface_identifier ( .* ) ;
 *   [ timeunits_declaration ] { interface_item }
 *   endinterface [ : interface_identifier ]
 *   | extern interface_nonansi_header
 *   | extern interface_ansi_header
 */

/*
 * interface_nonansi_header ::=
 *   { attribute_instance } interface [ lifetime ] interface_identifier
 *   { package_import_declaration } [ parameter_port_list ] list_of_ports ;
 */

/*
 * interface_ansi_header ::=
 *   {attribute_instance } interface [ lifetime ] interface_identifier
 *   { package_import_declaration }1 [ parameter_port_list ] [ list_of_port_declarations ] ;
 */

/*
 * program_declaration ::=
 *   program_nonansi_header [ timeunits_declaration ] { program_item }
 *   endprogram [ : program_identifier ]
 *   | program_ansi_header [ timeunits_declaration ] { non_port_program_item }
 *   endprogram [ : program_identifier ]
 *   | { attribute_instance } program program_identifier ( .* ) ;
 *   [ timeunits_declaration ] { program_item }
 *   endprogram [ : program_identifier ]
 *   | extern program_nonansi_header
 *   | extern program_ansi_header
 */

/*
 * program_nonansi_header ::=
 *   { attribute_instance } program [ lifetime ] program_identifier
 *   { package_import_declaration } [ parameter_port_list ] list_of_ports ;
 */

/*
 * program_ansi_header ::=
 *   { attribute_instance } program [ lifetime ] program_identifier
 *   { package_import_declaration }1 [ parameter_port_list ] [ list_of_port_declarations ] ;
 */

/*
 * checker_declaration ::=
 *   checker checker_identifier [ ( [ checker_port_list ] ) ] ;
 *   { { attribute_instance } checker_or_generate_item }
 *   endchecker [ : checker_identifier ]
 */

/*
 * class_declaration ::=
 *   [ virtual ] class [ lifetime ] class_identifier [ parameter_port_list ]
 *   [ extends class_type [ ( list_of_arguments ) ] ]
 *   [ implements interface_class_type { , interface_class_type } ] ;
 *   { class_item }
 *   endclass [ : class_identifier]
 */

/*
 * interface_class_type ::= ps_class_identifier [ parameter_value_assignment ]
 */

/*
 * interface_class_declaration ::=
 *   interface class class_identifier [ parameter_port_list ]
 *   [ extends interface_class_type { , interface_class_type } ] ;
 *   { interface_class_item }
 *   endclass [ : class_identifier]
 */

/*
 * interface_class_item ::=
 *   type_declaration
 *   | { attribute_instance } interface_class_method
 *   | local_parameter_declaration ;
 *   | parameter_declaration7 ;
 *   | ;
 */

/*
 * interface_class_method ::=
 *   pure virtual method_prototype ;
 */

/*
 * package_declaration ::=
 *   { attribute_instance } package [ lifetime ] package_identifier ;
 *   [ timeunits_declaration ] { { attribute_instance } package_item }
 *   endpackage [ : package_identifier ]
 */

/*
 * timeunits_declaration ::=
 *   timeunit time_literal [ / time_literal ] ;
 *   | timeprecision time_literal ;
 *   | timeunit time_literal ; timeprecision time_literal ;
 *   | timeprecision time_literal ; timeunit time_literal ;
 */

//
// A.1.3 Module parameters and ports
//

/*
 * parameter_port_list ::=
 *   # ( list_of_param_assignments { , parameter_port_declaration } )
 *   | # ( parameter_port_declaration { , parameter_port_declaration } )
 *   | #( )
 */

/*
 * parameter_port_declaration ::=
 *   parameter_declaration
 *   | local_parameter_declaration
 *   | data_type list_of_param_assignments
 *   | type list_of_type_assignments
 */

/*
 * list_of_ports ::= ( port { , port } )
 */

/*
 * list_of_port_declarations2 ::=
 *   ( [ { attribute_instance} ansi_port_declaration { , { attribute_instance} ansi_port_declaration } ] )
 */

/*
 * port_declaration ::=
 *   { attribute_instance } inout_declaration
 *   | { attribute_instance } input_declaration
 *   | { attribute_instance } output_declaration
 *   | { attribute_instance } ref_declaration
 *   | { attribute_instance } interface_port_declaration
 */

/*
 * port ::=
 *   [ port_expression ]
 *   | . port_identifier ( [ port_expression ] )
 */

/*
 * port_expression ::=
 *   port_reference
 *   | { port_reference { , port_reference } }
 */

/*
 * port_reference ::=
 *   port_identifier constant_select
 */

/*
 * port_direction ::= input | output | inout | ref
 */

/*
 * net_port_header ::= [ port_direction ] net_port_type
 */

/*
 * variable_port_header ::= [ port_direction ] variable_port_type
 */

/*
 * interface_port_header ::=
 *   interface_identifier [ . modport_identifier ]
 *   | interface [ . modport_identifier ]
 */

/*
 * ansi_port_declaration ::=
 *   [ net_port_header | interface_port_header ] port_identifier { unpacked_dimension }
 *   [ = constant_expression ]
 *   | [ variable_port_header ] port_identifier { variable_dimension } [ = constant_expression ]
 *   | [ port_direction ] . port_identifier ( [ expression ] )
 */

//
// A.1.4 Module items
//

/*
 * elaboration_system_task ::=
 *   $fatal [ ( finish_number [, list_of_arguments ] ) ] ;
 *   | $error [ ( [ list_of_arguments ] ) ] ;
 *   | $warning [ ( [ list_of_arguments ] ) ] ;
 *   | $info [ ( [ list_of_arguments ] ) ] ;
 */

/*
 * finish_number ::= 0 | 1 | 2
 */

/*
 * module_common_item ::=
 *   module_or_generate_item_declaration
 *   | interface_instantiation
 *   | program_instantiation
 *   | assertion_item
 *   | bind_directive
 *   | continuous_assign
 *   | net_alias
 *   | initial_construct
 *   | final_construct
 *   | always_construct
 *   | loop_generate_construct
 *   | conditional_generate_construct
 *   | elaboration_system_task
 */

/*
 * module_item ::=
 *   port_declaration ;
 *   | non_port_module_item
 */

/*
 * module_or_generate_item ::=
 *   { attribute_instance } parameter_override
 *   | { attribute_instance } gate_instantiation
 *   | { attribute_instance } udp_instantiation
 *   | { attribute_instance } module_instantiation
 *   | { attribute_instance } module_common_item
 */

/*
 * module_or_generate_item_declaration ::=
 *   package_or_generate_item_declaration
 *   | genvar_declaration
 *   | clocking_declaration
 *   | default clocking clocking_identifier ;
 *   | default disable iff expression_or_dist ;
 */

/*
 * non_port_module_item ::=
 *   generate_region
 *   | module_or_generate_item
 *   | specify_block
 *   | { attribute_instance } specparam_declaration
 *   | program_declaration
 *   | module_declaration
 *   | interface_declaration
 *   | timeunits_declaration3
 */

/*
 * parameter_override ::= defparam list_of_defparam_assignments ;
 */

/*
 * bind_directive4 ::=
 *   bind bind_target_scope [: bind_target_instance_list] bind_instantiation ;
 *   | bind bind_target_instance bind_instantiation ;
 */

/*
 * bind_target_scope ::=
 *   module_identifier
 *   | interface_identifier
 */

/*
 * bind_target_instance ::=
 *   hierarchical_identifier constant_bit_select
 */

/*
 * bind_target_instance_list ::=
 *   bind_target_instance { , bind_target_instance }
 */

/*
 * bind_instantiation ::=
 *   program_instantiation
 *   | module_instantiation
 *   | interface_instantiation
 *   | checker_instantiation
 */

//
// A.1.5 Configuration source text
//

/*
 * config_declaration ::=
 *   config config_identifier ;
 *   { local_parameter_declaration ; }
 *   design_statement
 *   { config_rule_statement }
 *   endconfig [ : config_identifier ]
 */

/*
 * design_statement ::= design { [ library_identifier . ] cell_identifier } ;
 */

/*
 * config_rule_statement ::=
 *   default_clause liblist_clause ;
 *   | inst_clause liblist_clause ;
 *   | inst_clause use_clause ;
 *   | cell_clause liblist_clause ;
 *   | cell_clause use_clause ;
 */

/*
 * default_clause ::= default
 */

/*
 * inst_clause ::= instance inst_name
 */

/*
 * inst_name ::= topmodule_identifier { . instance_identifier }
 */

/*
 * cell_clause ::= cell [ library_identifier . ] cell_identifier
 */

/*
 * liblist_clause ::= liblist {library_identifier}
 */

/*
 * use_clause ::= use [ library_identifier . ] cell_identifier [ : config ]
 *   | use named_parameter_assignment { , named_parameter_assignment } [ : config ]
 *   | use [ library_identifier . ] cell_identifier named_parameter_assignment
 *   { , named_parameter_assignment } [ : config ]
 */

//
// A.1.6 Interface items
//

/*
 * interface_or_generate_item ::=
 *   { attribute_instance } module_common_item
 *   | { attribute_instance } extern_tf_declaration
 */

/*
 * extern_tf_declaration ::=
 *   extern method_prototype ;
 *   | extern forkjoin task_prototype ;
 */

/*
 * interface_item ::=
 *   port_declaration ;
 *   | non_port_interface_item
 */

/*
 * non_port_interface_item ::=
 *   generate_region
 *   | interface_or_generate_item
 *   | program_declaration
 *   | modport_declaration
 *   | interface_declaration
 *   | timeunits_declaration3

//
// A.1.7 Program items
//

/*
 * program_item ::=
 *   port_declaration ;
 *   | non_port_program_item
*/

/*
 * non_port_program_item ::=
 *   { attribute_instance } continuous_assign
 *   | { attribute_instance } module_or_generate_item_declaration
 *   | { attribute_instance } initial_construct
 *   | { attribute_instance } final_construct
 *   | { attribute_instance } concurrent_assertion_item
 *   | timeunits_declaration3
 *   | program_generate_item
 */

/*
 * program_generate_item5 ::=
 *   loop_generate_construct
 *   | conditional_generate_construct
 *   | generate_region
 *   | elaboration_system_task
 */

//
// A.1.8 Checker items
//

/*
 * checker_port_list ::=
 *   checker_port_item {, checker_port_item}
 */

/*
 * checker_port_item ::=
 *   { attribute_instance } [ checker_port_direction ] property_formal_type formal_port_identifier
 *   {variable_dimension} [ = property_actual_arg ]
 */

/*
 * checker_port_direction ::=
 *   input | output
 */

/*
 * checker_or_generate_item ::=
 *   checker_or_generate_item_declaration
 *   | initial_construct
 *   | always_construct
 *   | final_construct
 *   | assertion_item
 *   | continuous_assign
 *   | checker_generate_item
 */

/*
 * checker_or_generate_item_declaration ::=
 *   [ rand ] data_declaration
 *   | function_declaration
 *   | checker_declaration
 *   | assertion_item_declaration
 *   | covergroup_declaration
 *   | genvar_declaration
 *   | clocking_declaration
 *   | default clocking clocking_identifier ;
 *   | default disable iff expression_or_dist ;
 *   | ;
 */

/*
 * checker_generate_item6 ::=
 *   loop_generate_construct
 *   | conditional_generate_construct
 *   | generate_region
 *   | elaboration_system_task
 */

//
// A.1.9 Class items
//

/*
 * class_item ::=
 *   { attribute_instance } class_property
 *   | { attribute_instance } class_method
 *   | { attribute_instance } class_constraint
 *   | { attribute_instance } class_declaration
 *   | { attribute_instance } covergroup_declaration
 *   | local_parameter_declaration ;
 *   | parameter_declaration7 ;
 *   | ;
 */

/*
 * class_property ::=
 *   { property_qualifier } data_declaration
 *   | const { class_item_qualifier } data_type const_identifier [ = constant_expression ] ;
 */

/*
 * class_method ::=
 *   { method_qualifier } task_declaration
 *   | { method_qualifier } function_declaration
 *   | pure virtual { class_item_qualifier } method_prototype ;
 *   | extern { method_qualifier } method_prototype ;
 *   | { method_qualifier } class_constructor_declaration
 *   | extern { method_qualifier } class_constructor_prototype
 */

/*
 * class_constructor_prototype ::=
 *   function new [ ( [ tf_port_list ] ) ] ;
 */

/*
 * class_constraint ::=
 *   constraint_prototype
 *   | constraint_declaration
 */

/*
 * class_item_qualifier8 ::=
 *   static
 *   | protected
 *   | local
 */

/*
 * property_qualifier8 ::=
 *   random_qualifier
 *   | class_item_qualifier
 */

/*
 * random_qualifier8 ::=
 *   rand
 *   | randc
 */

/*
 * method_qualifier8 ::=
 *   [ pure ] virtual
 *   | class_item_qualifier
 */

/*
 * method_prototype ::=
 *   task_prototype
 *   | function_prototype
 */

/*
 * class_constructor_declaration ::=
 *   function [ class_scope ] new [ ( [ tf_port_list ] ) ] ;
 *   { block_item_declaration }
 *   [ super . new [ ( list_of_arguments ) ] ; ]
 *   { function_statement_or_null }
 *   endfunction [ : new ]
 */

//
// A.1.10 Constraints
//

/*
 * constraint_declaration ::= [ static ] constraint constraint_identifier constraint_block
 */

/*
 * constraint_block ::= { { constraint_block_item } }
 */

/*
 * constraint_block_item ::=
 *   solve solve_before_list before solve_before_list ;
 *   | constraint_expression
 */

/*
 * solve_before_list ::= constraint_primary { , constraint_primary }
 */

/*
 * constraint_primary ::= [ implicit_class_handle . | class_scope ] hierarchical_identifier select
 */

/*
 * constraint_expression ::=
 *   [ soft ] expression_or_dist ;
 *   | uniqueness_constraint ;
 *   | expression –> constraint_set
 *   | if ( expression ) constraint_set [ else constraint_set ]
 *   | foreach ( ps_or_hierarchical_array_identifier [ loop_variables ] ) constraint_set
 *   | disable soft constraint_primary ;
 */

/*
 * uniqueness_constraint ::=
 *   unique { open_range_list9 }
 */

/*
 * constraint_set ::=
 *   constraint_expression
 *   | { { constraint_expression } }
 */

/*
 * dist_list ::= dist_item { , dist_item }
 */

/*
 * dist_item ::= value_range [ dist_weight ]
 */

/*
 * dist_weight ::=
 *   := expression
 *   | :/ expression
 */

/*
 * constraint_prototype ::= [constraint_prototype_qualifier] [ static ] constraint constraint_identifier ;
 */

/*
 * constraint_prototype_qualifier ::= extern | pure
 */

/*
 * extern_constraint_declaration ::=
 *   [ static ] constraint class_scope constraint_identifier constraint_block
 */

/*
 * identifier_list ::= identifier { , identifier }
 */

//
// A.1.11 Package items
//

/*
 * package_item ::=
 *   package_or_generate_item_declaration
 *   | anonymous_program
 *   | package_export_declaration
 *   | timeunits_declaration3
 */

/*
 * package_or_generate_item_declaration ::=
 *   net_declaration
 *   | data_declaration
 *   | task_declaration
 *   | function_declaration
 *   | checker_declaration
 *   | dpi_import_export
 *   | extern_constraint_declaration
 *   | class_declaration
 *   | class_constructor_declaration
 *   | local_parameter_declaration ;
 *   | parameter_declaration ;
 *   | covergroup_declaration
 *   | assertion_item_declaration
 *   | ;
 */

/*
 * anonymous_program ::= program ; { anonymous_program_item } endprogram
 */

/*
 * anonymous_program_item ::=
 *   task_declaration
 *   | function_declaration
 *   | class_declaration
 *   | covergroup_declaration
 *   | class_constructor_declaration
 *   | ;
 */
