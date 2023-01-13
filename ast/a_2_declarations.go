package ast

import (
	"fmt"

	"github.com/jtdubs/go-nom"
)

//
// A.2 Declarations
//

//
// A.2.1 Declaration types
//

//
// A.2.1.1 Module parameter declarations
//

/*
 * local_parameter_declaration ::=
 *   localparam data_type_or_implicit list_of_param_assignments
 *   | localparam type list_of_type_assignments
 */

/*
 * parameter_declaration ::=
 *   parameter data_type_or_implicit list_of_param_assignments
 *   | parameter type list_of_type_assignments
 */

/*
 * specparam_declaration ::=
 *   specparam [ packed_dimension ] list_of_specparam_assignments ;
 */

//
// A.2.1.2 Port declarations
//

/*
 * inout_declaration ::=
 *   inout net_port_type list_of_port_identifiers
 */

/*
 * input_declaration ::=
 *   input net_port_type list_of_port_identifiers
 *   | input variable_port_type list_of_variable_identifiers
 */

/*
 * output_declaration ::=
 *   output net_port_type list_of_port_identifiers
 *   | output variable_port_type list_of_variable_port_identifiers
 */

/*
 * interface_port_declaration ::=
 *   interface_identifier list_of_interface_identifiers
 *   | interface_identifier . modport_identifier list_of_interface_identifiers
 */

/*
 * ref_declaration ::= ref variable_port_type list_of_variable_identifiers
 */

//
// A.2.1.3 Type declarations
//

/*
 * data_declaration ::=
 *   [ const ] [ var ] [ lifetime ] data_type_or_implicit list_of_variable_decl_assignments ;10
 *   | type_declaration
 *   | package_import_declaration11
 *   | net_type_declaration
 */

/*
 * package_import_declaration ::=
 *   import package_import_item { , package_import_item } ;
 */

/*
 * package_import_item ::=
 *   package_identifier :: identifier
 *   | package_identifier :: *
 */

/*
 * package_export_declaration ::=
 *   export *::* ;
 *   | export package_import_item { , package_import_item } ;
 */

/*
 * genvar_declaration ::= genvar list_of_genvar_identifiers ;
 */

/*
 * net_declaration12 ::=
 *   net_type [ drive_strength | charge_strength ] [ vectored | scalared ]
 *   data_type_or_implicit [ delay3 ] list_of_net_decl_assignments ;
 *   | net_type_identifier [ delay_control ]
 *   list_of_net_decl_assignments ;
 *   | interconnect implicit_data_type [ # delay_value ]
 *   net_identifier { unpacked_dimension }
 *   [ , net_identifier { unpacked_dimension }] ;
 */

/*
 * type_declaration ::=
 *   typedef data_type type_identifier { variable_dimension } ;
 *   | typedef interface_instance_identifier constant_bit_select . type_identifier type_identifier ;
 *   | typedef [ enum | struct | union | class | interface class ] type_identifier ;
 */

/*
 * net_type_declaration ::=
 *   nettype data_type net_type_identifier
 *   [ with [ package_scope | class_scope ] tf_identifier ] ;
 *   | nettype [ package_scope | class_scope ] net_type_identifier net_type_identifier ;
 */

/*
 * lifetime ::= static | automatic
 */

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
 *   | type_reference14
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
type IntegerType interface {
	isIntegerType()
}

/*
 * integer_atom_type ::= byte | shortint | int | longint | integer | time
 */
type IntegerAtomTypeOption int

const (
	Byte IntegerAtomTypeOption = iota
	ShortInt
	Int
	LongInt
	Integer
	Time
)

var integerAtomTypeNames = map[IntegerAtomTypeOption]string{
	Byte:     "Byte",
	ShortInt: "ShortInt",
	Int:      "Int",
	LongInt:  "LongInt",
	Integer:  "Integer",
	Time:     "Time",
}

type IntegerAtomType struct {
	Token
	Type IntegerAtomTypeOption
}

func (u *IntegerAtomType) String() string {
	return fmt.Sprintf("IntegerAtomType(%v)", integerAtomTypeNames[u.Type])
}

func (*IntegerAtomType) isIntegerType() {}

/*
 * integer_vector_type ::= bit | logic | reg
 */
type IntegerVectorTypeOption int

const (
	Bit IntegerVectorTypeOption = iota
	Logic
	Reg
)

var integerVectorTypeNames = map[IntegerVectorTypeOption]string{
	Bit:   "Bit",
	Logic: "Logic",
	Reg:   "Reg",
}

type IntegerVectorType struct {
	Token
	Type IntegerVectorTypeOption
}

func (u *IntegerVectorType) String() string {
	return fmt.Sprintf("IntegerVectorType(%v)", integerVectorTypeNames[u.Type])
}

func (*IntegerVectorType) isIntegerType() {}

/*
 * non_integer_type ::= shortreal | real | realtime
 */
type NonIntegerTypeOption int

const (
	ShortReal NonIntegerTypeOption = iota
	Real
	RealTime
)

var nonIntegerTypeNames = map[NonIntegerTypeOption]string{
	ShortReal: "ShortReal",
	Real:      "Real",
	RealTime:  "RealTime",
}

type NonIntegerType struct {
	Token
	Type NonIntegerTypeOption
}

func (u *NonIntegerType) String() string {
	return fmt.Sprintf("NonIntegerType(%v)", nonIntegerTypeNames[u.Type])
}

/*
 * net_type ::= supply0 | supply1 | tri | triand | trior | trireg| tri0 | tri1 | uwire| wire | wand | wor
 */
type NetTypeOption int

const (
	NetSupply0 NetTypeOption = iota
	NetSupply1
	NetTri
	NetTriAnd
	NetTriOr
	NetTriReg
	NetTri0
	NetTri1
	NetUWire
	NetWire
	NetWAnd
	NetWOr
)

var netTypeNames = map[NetTypeOption]string{
	NetSupply0: "Supply0",
	NetSupply1: "Supply1",
	NetTri:     "NetTri",
	NetTriAnd:  "NetTriAnd",
	NetTriOr:   "NetTriOr",
	NetTriReg:  "NetTriReg",
	NetTri0:    "NetTri0",
	NetTri1:    "NetTri1",
	NetUWire:   "NetUWire",
	NetWire:    "NetWire",
	NetWAnd:    "NetWAnd",
	NetWOr:     "NetWOr",
}

type NetType struct {
	Token
	Type NetTypeOption
}

func (u *NetType) String() string {
	return fmt.Sprintf("NetType(%v)", netTypeNames[u.Type])
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
type SigningOption int

const (
	Signed SigningOption = iota
	Unsigned
)

var signingNames = map[SigningOption]string{
	Signed:   "Signed",
	Unsigned: "Unsigned",
}

type Signing struct {
	Token
	Type SigningOption
}

func (u *Signing) String() string {
	return fmt.Sprintf("Signing(%v)", signingNames[u.Type])
}

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
type DriveStrengthOption interface {
	isDriveStrengthOption()
}

type DriveStrength struct {
	Token
	AT, BT nom.Span[rune]
	A, B   DriveStrengthOption
}

func (u *DriveStrength) String() string {
	return fmt.Sprintf("DriveStrength(%v, %v)", u.A, u.B)
}

type HighZ0 struct {
	Token
}

func (*HighZ0) String() string {
	return "HighZ0"
}

func (*HighZ0) isDriveStrengthOption() {}

type HighZ1 struct {
	Token
}

func (*HighZ1) String() string {
	return "HighZ1"
}

func (*HighZ1) isDriveStrengthOption() {}

/*
 * strength0 ::= supply0 | strong0 | pull0 | weak0
 */
type Strength0Option int

const (
	StrengthSupply0 Strength0Option = iota
	StrengthStrong0
	StrengthPull0
	StrengthWeak0
)

var strength0OptionNames = map[Strength0Option]string{
	StrengthSupply0: "Supply0",
	StrengthStrong0: "Strong0",
	StrengthPull0:   "Pull0",
	StrengthWeak0:   "Weak0",
}

type Strength0 struct {
	Token
	Type Strength0Option
}

func (u *Strength0) String() string {
	return fmt.Sprintf("Strength0(%v)", strength0OptionNames[u.Type])
}

func (*Strength0) isDriveStrengthOption() {}

/*
 * strength1 ::= supply1 | strong1 | pull1 | weak1
 */
type Strength1Option int

const (
	StrengthSupply1 Strength1Option = iota
	StrengthStrong1
	StrengthPull1
	StrengthWeak1
)

var strength1OptionNames = map[Strength1Option]string{
	StrengthSupply1: "Supply1",
	StrengthStrong1: "Strong1",
	StrengthPull1:   "Pull1",
	StrengthWeak1:   "Weak1",
}

type Strength1 struct {
	Token
	Type Strength1Option
}

func (u *Strength1) String() string {
	return fmt.Sprintf("Strength1(%v)", strength1OptionNames[u.Type])
}

func (*Strength1) isDriveStrengthOption() {}

/*
 * charge_strength ::= ( small ) | ( medium ) | ( large )
 */
type ChargeStrengthOption int

const (
	ChargeSmall ChargeStrengthOption = iota
	ChargeMedium
	ChargeLarge
)

var chargeStrengthOptionNames = map[ChargeStrengthOption]string{
	ChargeSmall:  "Small",
	ChargeMedium: "Medium",
	ChargeLarge:  "Large",
}

type ChargeStrength struct {
	Token
	TypeT nom.Span[rune]
	Type  ChargeStrengthOption
}

func (u *ChargeStrength) String() string {
	return fmt.Sprintf("ChargeStrength(%v)", chargeStrengthOptionNames[u.Type])
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

//
// A.2.3 Declaration lists
//

/*
 * list_of_defparam_assignments ::= defparam_assignment { , defparam_assignment }
 */

/*
 * list_of_genvar_identifiers ::= genvar_identifier { , genvar_identifier }
 */

/*
 * list_of_interface_identifiers ::= interface_identifier { unpacked_dimension }
 *   { , interface_identifier { unpacked_dimension } }
 */

/*
 * list_of_net_decl_assignments ::= net_decl_assignment { , net_decl_assignment }
 */

/*
 * list_of_param_assignments ::= param_assignment { , param_assignment }
 */

/*
 * list_of_port_identifiers ::= port_identifier { unpacked_dimension }
 *   { , port_identifier { unpacked_dimension } }
 */

/*
 * list_of_udp_port_identifiers ::= port_identifier { , port_identifier }
 */

/*
 * list_of_specparam_assignments ::= specparam_assignment { , specparam_assignment }
 */

/*
 * list_of_tf_variable_identifiers ::= port_identifier { variable_dimension } [ = expression ]
 *   { , port_identifier { variable_dimension } [ = expression ] }
 */

/*
 * list_of_type_assignments ::= type_assignment { , type_assignment }
 */

/*
 * list_of_variable_decl_assignments ::= variable_decl_assignment { , variable_decl_assignment }
 */

/*
 * list_of_variable_identifiers ::= variable_identifier { variable_dimension }
 *   { , variable_identifier { variable_dimension } }
 */

/*
 * list_of_variable_port_identifiers ::= port_identifier { variable_dimension } [ = constant_expression ]
 *   { , port_identifier { variable_dimension } [ = constant_expression ] }
 */

//
// A.2.4 Declaration assignments
//

/*
 * defparam_assignment ::= hierarchical_parameter_identifier = constant_mintypmax_expression
 */

/*
 * net_decl_assignment ::= net_identifier { unpacked_dimension } [ = expression ]
 */

/*
 * param_assignment ::=
 *   parameter_identifier { unpacked_dimension } [ = constant_param_expression ]18
 */

/*
 * specparam_assignment ::=
 *   specparam_identifier = constant_mintypmax_expression
 *   | pulse_control_specparam
 */

/*
 * type_assignment ::=
 *   type_identifier [ = data_type ]18
 */

/*
 * pulse_control_specparam ::=
 *   PATHPULSE$ = ( reject_limit_value [ , error_limit_value ] )
 *   | PATHPULSE$specify_input_terminal_descriptor$specify_output_terminal_descriptor
 *   = ( reject_limit_value [ , error_limit_value ] )
 */

/*
 * error_limit_value ::= limit_value
 */

/*
 * reject_limit_value ::= limit_value
 */

/*
 * limit_value ::= constant_mintypmax_expression
 */

/*
 * variable_decl_assignment ::=
 *   variable_identifier { variable_dimension } [ = expression ]
 *   | dynamic_array_variable_identifier unsized_dimension { variable_dimension }
 *   [ = dynamic_array_new ]
 *   | class_variable_identifier [ = class_new ]
 */

/*
 * class_new19 ::=
 *   [ class_scope ] new [ ( list_of_arguments ) ]
 *   | new expression
 */

/*
 * dynamic_array_new ::= new [ expression ] [ ( expression ) ]
 */

//
// A.2.5 Declaration ranges
//

/*
 * unpacked_dimension ::=
 *   [ constant_range ]
 *   | [ constant_expression ]
 */

/*
 * packed_dimension20 ::=
 *   [ constant_range ]
 *   | unsized_dimension
 */

/*
 * associative_dimension ::=
 *   [ data_type ]
 *   | [ * ]
 */

/*
 * variable_dimension ::=
 *   unsized_dimension
 *   | unpacked_dimension
 *   | associative_dimension
 *   | queue_dimension
 */

/*
 * queue_dimension ::= [ $ [ : constant_expression ] ]
 */

/*
 * unsized_dimension ::= [ ]
 */

//
// A.2.6 Function declarations
//

/*
 * function_data_type_or_implicit ::=
 *   data_type_or_void
 *   | implicit_data_type
 */

/*
 * function_declaration ::= function [ lifetime ] function_body_declaration
 */

/*
 * function_body_declaration ::=
 *   function_data_type_or_implicit
 *   [ interface_identifier . | class_scope ] function_identifier ;
 *   { tf_item_declaration }
 *   { function_statement_or_null }
 *   endfunction [ : function_identifier ]
 *   | function_data_type_or_implicit
 *   [ interface_identifier . | class_scope ] function_identifier ( [ tf_port_list ] ) ;
 *   { block_item_declaration }
 *   { function_statement_or_null }
 *   endfunction [ : function_identifier ]
 */

/*
 * function_prototype ::= function data_type_or_void function_identifier [ ( [ tf_port_list ] ) ]
 */

/*
 * dpi_import_export ::=
 *   import dpi_spec_string [ dpi_function_import_property ] [ c_identifier = ] dpi_function_proto ;
 *   | import dpi_spec_string [ dpi_task_import_property ] [ c_identifier = ] dpi_task_proto ;
 *   | export dpi_spec_string [ c_identifier = ] function function_identifier ;
 *   | export dpi_spec_string [ c_identifier = ] task task_identifier ;
 */

/*
 * dpi_spec_string ::= "DPI-C" | "DPI"
 */

/*
 * dpi_function_import_property ::= context | pure
 */

/*
 * dpi_task_import_property ::= context
 */

/*
 * dpi_function_proto21,22 ::= function_prototype
 */

/*
 * dpi_task_proto22 ::= task_prototype
 */

//
// A.2.7 Task declarations
//

/*
 * task_declaration ::= task [ lifetime ] task_body_declaration
 */

/*
 * task_body_declaration ::=
 *   [ interface_identifier . | class_scope ] task_identifier ;
 *   { tf_item_declaration }
 *   { statement_or_null }
 *   endtask [ : task_identifier ]
 *   | [ interface_identifier . | class_scope ] task_identifier ( [ tf_port_list ] ) ;
 *   { block_item_declaration }
 *   { statement_or_null }
 *   endtask [ : task_identifier ]
 */

/*
 * tf_item_declaration ::=
 *   block_item_declaration
 *   | tf_port_declaration
 */

/*
 * tf_port_list ::=
 *   tf_port_item { , tf_port_item }
 */

/*
 * tf_port_item23 ::=
 *   { attribute_instance }
 *   [ tf_port_direction ] [ var ] data_type_or_implicit
 *   [ port_identifier { variable_dimension } [ = expression ] ]
 */

/*
 * tf_port_direction ::= port_direction | const ref
 */

/*
 * tf_port_declaration ::=
 *   { attribute_instance } tf_port_direction [ var ] data_type_or_implicit list_of_tf_variable_identifiers ;
 */

/*
 * task_prototype ::= task task_identifier [ ( [ tf_port_list ] ) ]
 */

//
// A.2.8 Block item declarations
//

/*
 * block_item_declaration ::=
 *   { attribute_instance } data_declaration
 *   | { attribute_instance } local_parameter_declaration ;
 *   | { attribute_instance } parameter_declaration ;
 *   | { attribute_instance } let_declaration
 */

//
// A.2.9 Interface declarations
//

/*
 * modport_declaration ::= modport modport_item { , modport_item } ;
 */

/*
 * modport_item ::= modport_identifier ( modport_ports_declaration { , modport_ports_declaration } )
 */

/*
 * modport_ports_declaration ::=
 *   { attribute_instance } modport_simple_ports_declaration
 *   | { attribute_instance } modport_tf_ports_declaration
 *   | { attribute_instance } modport_clocking_declaration
 */

/*
 * modport_clocking_declaration ::= clocking clocking_identifier
 */

/*
 * modport_simple_ports_declaration ::=
 *   port_direction modport_simple_port { , modport_simple_port }
 */

/*
 * modport_simple_port ::=
 *   port_identifier
 *   | . port_identifier ( [ expression ] )
 */

/*
 * modport_tf_ports_declaration ::=
 *   import_export modport_tf_port { , modport_tf_port }
 */

/*
 * modport_tf_port ::=
 *   method_prototype
 *   | tf_identifier
 */

/*
 * import_export ::= import | export
 */

//
// A.2.10 Assertion declarations
//

/*
 * concurrent_assertion_item ::=
 *   [ block_identifier : ] concurrent_assertion_statement
 *   | checker_instantiation
 */

/*
 * concurrent_assertion_statement ::=
 *   assert_property_statement
 *   | assume_property_statement
 *   | cover_property_statement
 *   | cover_sequence_statement
 *   | restrict_property_statement
 */

/*
 * assert_property_statement ::=
 *   assert property ( property_spec ) action_block
 */

/*
 * assume_property_statement ::=
 *   assume property ( property_spec ) action_block
 */

/*
 * cover_property_statement::=
 *   cover property ( property_spec ) statement_or_null
 */

/*
 * expect_property_statement ::=
 *   expect ( property_spec ) action_block
 */

/*
 * cover_sequence_statement ::=
 *   cover sequence ( [clocking_event ] [ disable iff ( expression_or_dist ) ]
 *   sequence_expr ) statement_or_null
 */

/*
 * restrict_property_statement ::=
 *   restrict property ( property_spec ) ;
 */

/*
 * property_instance ::=
 *   ps_or_hierarchical_property_identifier [ ( [ property_list_of_arguments ] ) ]
 */

/*
 * property_list_of_arguments ::=
 *   [property_actual_arg] { , [property_actual_arg] } { , . identifier ( [property_actual_arg] ) }
 *   | . identifier ( [property_actual_arg] ) { , . identifier ( [property_actual_arg] ) }
 */

/*
 * property_actual_arg ::=
 *   property_expr
 *   | sequence_actual_arg
 */

/*
 * assertion_item_declaration ::=
 *   property_declaration
 *   | sequence_declaration
 *   | let_declaration
 */

/*
 * property_declaration ::=
 *   property property_identifier [ ( [ property_port_list ] ) ] ;
 *   { assertion_variable_declaration }
 *   property_spec [ ; ]
 *   endproperty [ : property_identifier ]
 */

/*
 * property_port_list ::=
 *   property_port_item {, property_port_item}
 */

/*
 * property_port_item ::=
 *   { attribute_instance } [ local [ property_lvar_port_direction ] ] property_formal_type
 *   formal_port_identifier {variable_dimension} [ = property_actual_arg ]
 */

/*
 * property_lvar_port_direction ::= input
 */

/*
 * property_formal_type ::=
 *   sequence_formal_type
 *   | property
 */

/*
 * property_spec ::=
 *   [ clocking_event ] [ disable iff ( expression_or_dist ) ] property_expr
 */

/*
 * property_expr ::=
 *   sequence_expr
 *   | strong ( sequence_expr )
 *   | weak ( sequence_expr )
 *   | ( property_expr )
 *   | not property_expr
 *   | property_expr or property_expr
 *   | property_expr and property_expr
 *   | sequence_expr |-> property_expr
 *   | sequence_expr |=> property_expr
 *   | if ( expression_or_dist ) property_expr [ else property_expr ]
 *   | case ( expression_or_dist ) property_case_item { property_case_item } endcase
 *   | sequence_expr #-# property_expr
 *   | sequence_expr #=# property_expr
 *   | nexttime property_expr
 *   | nexttime [ constant _expression ] property_expr
 *   | s_nexttime property_expr
 *   | s_nexttime [ constant_expression ] property_expr
 *   | always property_expr
 *   | always [ cycle_delay_const_range_expression ] property_expr
 *   | s_always [ constant_range] property_expr
 *   | s_eventually property_expr
 *   | eventually [ constant_range ] property_expr
 *   | s_eventually [ cycle_delay_const_range_expression ] property_expr
 *   | property_expr until property_expr
 *   | property_expr s_until property_expr
 *   | property_expr until_with property_expr
 *   | property_expr s_until_with property_expr
 *   | property_expr implies property_expr
 *   | property_expr iff property_expr
 *   | accept_on ( expression_or_dist ) property_expr
 *   | reject_on ( expression_or_dist ) property_expr
 *   | sync_accept_on ( expression_or_dist ) property_expr
 *   | sync_reject_on ( expression_or_dist ) property_expr
 *   | property_instance
 *   | clocking_event property_expr
 */

/*
 * property_case_item ::=
 *   expression_or_dist { , expression_or_dist } : property_expr ;
 *   | default [ : ] property_expr ;
 */

/*
 * sequence_declaration ::=
 *   sequence sequence_identifier [ ( [ sequence_port_list ] ) ] ;
 *   { assertion_variable_declaration }
 *   sequence_expr [ ; ]
 *   endsequence [ : sequence_identifier ]
 */

/*
 * sequence_port_list ::=
 *   sequence_port_item {, sequence_port_item}
 */

/*
 * sequence_port_item ::=
 *   { attribute_instance } [ local [ sequence_lvar_port_direction ] ] sequence_formal_type
 *   formal_port_identifier {variable_dimension} [ = sequence_actual_arg ]
 */

/*
 * sequence_lvar_port_direction ::= input | inout | output
 */

/*
 * sequence_formal_type ::=
 *   data_type_or_implicit
 *   | sequence
 *   | untyped
 */

/*
 * sequence_expr ::=
 *   cycle_delay_range sequence_expr { cycle_delay_range sequence_expr }
 *   | sequence_expr cycle_delay_range sequence_expr { cycle_delay_range sequence_expr }
 *   | expression_or_dist [ boolean_abbrev ]
 *   | sequence_instance [ sequence_abbrev ]
 *   | ( sequence_expr {, sequence_match_item } ) [ sequence_abbrev ]
 *   | sequence_expr and sequence_expr
 *   | sequence_expr intersect sequence_expr
 *   | sequence_expr or sequence_expr
 *   | first_match ( sequence_expr {, sequence_match_item} )
 *   | expression_or_dist throughout sequence_expr
 *   | sequence_expr within sequence_expr
 *   | clocking_event sequence_expr
 */

/*
 * cycle_delay_range ::=
 *   ## constant_primary
 *   | ## [ cycle_delay_const_range_expression ]
 *   | ##[*]
 *   | ##[+]
 */

/*
 * sequence_method_call ::=
 *   sequence_instance . method_identifier
 */

/*
 * sequence_match_item ::=
 *   operator_assignment
 *   | inc_or_dec_expression
 *   | subroutine_call
 */

/*
 * sequence_instance ::=
 *   ps_or_hierarchical_sequence_identifier [ ( [ sequence_list_of_arguments ] ) ]
 */

/*
 * sequence_list_of_arguments ::=
 *   [sequence_actual_arg] { , [sequence_actual_arg] } { , . identifier ( [sequence_actual_arg] ) }
 *   | . identifier ( [sequence_actual_arg] ) { , . identifier ( [sequence_actual_arg] ) }
 */

/*
 * sequence_actual_arg ::=
 *   event_expression
 *   | sequence_expr
 */

/*
 * boolean_abbrev ::=
 *   consecutive_repetition
 *   | non_consecutive_repetition
 *   | goto_repetition
 */

/*
 * sequence_abbrev ::= consecutive_repetition
 */

/*
 * consecutive_repetition ::=
 *   [* const_or_range_expression ]
 *   | [*]
 *   | [+]
 */

/*
 * non_consecutive_repetition ::= [= const_or_range_expression ]
 */

/*
 * goto_repetition ::= [-> const_or_range_expression ]
 */

/*
 * const_or_range_expression ::=
 *   constant_expression
 *   | cycle_delay_const_range_expression
 */

/*
 * cycle_delay_const_range_expression ::=
 *   constant_expression : constant_expression
 *   | constant_expression : $
 */

/*
 * expression_or_dist ::= expression [ dist { dist_list } ]
 */

/*
 * assertion_variable_declaration ::=
 *   var_data_type list_of_variable_decl_assignments ;
 */

//
// A.2.11 Covergroup declarations
//

/*
 * covergroup_declaration ::=
 *   covergroup covergroup_identifier [ ( [ tf_port_list ] ) ] [ coverage_event ] ;
 *   { coverage_spec_or_option }
 *   endgroup [ : covergroup_identifier ]
 */

/*
 * coverage_spec_or_option ::=
 *   { attribute_instance } coverage_spec
 *   | { attribute_instance } coverage_option ;
 */

/*
 * coverage_option ::=
 *   option.member_identifier = expression
 *   | type_option.member_identifier = constant_expression
 */

/*
 * coverage_spec ::=
 *   cover_point
 *   | cover_cross
 */

/*
 * coverage_event ::=
 *   clocking_event
 *   | with function sample ( [ tf_port_list ] )
 *   | @@( block_event_expression )
 */

/*
 * block_event_expression ::=
 *   block_event_expression or block_event_expression
 *   | begin hierarchical_btf_identifier
 *   | end hierarchical_btf_identifier
 */

/*
 * hierarchical_btf_identifier ::=
 *   hierarchical_tf_identifier
 *   | hierarchical_block_identifier
 *   | [ hierarchical_identifier. | class_scope ] method_identifier
 */

/*
 * cover_point ::=
 *   [ [ data_type_or_implicit ] cover_point_identifier : ] coverpoint expression [ iff ( expression ) ]
 *   bins_or_empty
 */

/*
 * bins_or_empty ::=
 *   { {attribute_instance} { bins_or_options ; } }
 *   | ;
 */

/*
 * bins_or_options ::=
 *   coverage_option
 *   | [ wildcard ] bins_keyword bin_identifier [ [ [ covergroup_expression ] ] ] =
 *   { covergroup_range_list } [ with ( with_covergroup_expression ) ]
 *   [ iff ( expression ) ]
 *   | [ wildcard ] bins_keyword bin_identifier [ [ [ covergroup_expression ] ] ] =
 *   cover_point_identifier with ( with_covergroup_expression ) [ iff ( expression ) ]
 *   | [ wildcard ] bins_keyword bin_identifier [ [ [ covergroup_expression ] ] ] =
 *   set_covergroup_expression [ iff ( expression ) ]
 *   | [ wildcard] bins_keyword bin_identifier [ [ ] ] = trans_list [ iff ( expression ) ]
 *   | bins_keyword bin_identifier [ [ [ covergroup_expression ] ] ] = default [ iff ( expression ) ]
 *   | bins_keyword bin_identifier = default sequence [ iff ( expression ) ]
 */

/*
 * bins_keyword ::= bins | illegal_bins | ignore_bins
 */

/*
 * trans_list ::= ( trans_set ) { , ( trans_set ) }
 */

/*
 * trans_set ::= trans_range_list { => trans_range_list }
 */

/*
 * trans_range_list ::=
 *   trans_item
 *   | trans_item [* repeat_range ]
 *   | trans_item [–> repeat_range ]
 *   | trans_item [= repeat_range ]
 */

/*
 * trans_item ::= covergroup_range_list
 */

/*
 * repeat_range ::=
 *   covergroup_expression
 *   | covergroup_expression : covergroup_expression
 */

/*
 * cover_cross ::=
 *   [ cross_identifier : ] cross list_of_cross_items [ iff ( expression ) ] cross_body
 */

/*
 * list_of_cross_items ::= cross_item , cross_item { , cross_item }
 */

/*
 * cross_item ::=
 *   cover_point_identifier
 *   | variable_identifier
 */

/*
 * cross_body ::=
 *   { { cross_body_item ; } }
 *   | ;
 */

/*
 * cross_body_item ::=
 *   function_declaraton
 *   | bins_selection_or_option ;
 */

/*
 * bins_selection_or_option ::=
 *   { attribute_instance } coverage_option
 *   | { attribute_instance } bins_selection
 */

/*
 * bins_selection ::= bins_keyword bin_identifier = select_expression [ iff ( expression ) ]
 */

/*
 * select_expression24 ::=
 *   select_condition
 *   | ! select_condition
 *   | select_expression && select_expression
 *   | select_expression || select_expression
 *   | ( select_expression )
 *   | select_expression with ( with_covergroup_expression ) [ matches integer_covergroup_expression ]
 *   | cross_identifier
 *   | cross_set_expression [ matches integer_covergroup_expression ]
 */

/*
 * select_condition ::= binsof ( bins_expression ) [ intersect { covergroup_range_list } ]
 */

/*
 * bins_expression ::=
 *   variable_identifier
 *   | cover_point_identifier [ . bin_identifier ]
 */

/*
 * covergroup_range_list ::= covergroup_value_range { , covergroup_value_range }
 */

/*
 * covergroup_value_range ::=
 *   covergroup_expression
 *   | [ covergroup_expression : covergroup_expression ]25
 */

/*
 * with_covergroup_expression ::= covergroup_expression26
 */

/*
 * set_covergroup_expression ::= covergroup_expression27
 */

/*
 * integer_covergroup_expression ::= covergroup_expression
 */

/*
 * cross_set_expression ::= covergroup_expression
 */

/*
 * covergroup_expression ::= expression28
 */

//
// A.2.12 Let declarations
//

/*
 * let_declaration ::=
 *   let let_identifier [ ( [ let_port_list ] ) ] = expression ;
 */

/*
 * let_identifier ::=
 *   identifier
 */

/*
 * let_port_list ::=
 *   let_port_item {, let_port_item}
 */

/*
 * let_port_item ::=
 *   { attribute_instance } let_formal_type formal_port_identifier { variable_dimension } [ = expression ]
 */

/*
 * let_formal_type ::=
 *   data_type_or_implicit
 *   | untyped
 */

/*
 * let_expression ::=
 *   [ package_scope ] let_identifier [ ( [ let_list_of_arguments ] ) ]
 */

/*
 * let_list_of_arguments ::=
 *   [ let_actual_arg ] {, [ let_actual_arg ] } {, . identifier ( [ let_actual_arg ] ) }
 *   | . identifier ( [ let_actual_arg ] ) { , . identifier ( [ let_actual_arg ] ) }
 */

/*
 * let_actual_arg ::=
 *   expression
 */
