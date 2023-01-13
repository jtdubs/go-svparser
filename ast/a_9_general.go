package ast

import (
	"fmt"

	"github.com/jtdubs/go-nom"
)

//
// A.9 General
//

//
// A.9.1 Attributes
//

/*
 * attribute_instance ::= (* attr_spec { , attr_spec } *)
 */
type AttributeInstance struct {
	Token
	Specs []*AttrSpec
}

/*
 * attr_spec ::= attr_name [ = constant_expression ]
 */
type AttrSpec struct {
	Token
	Name *AttrName
	Expr ConstantExpression
}

/*
 * attr_name ::= identifier
 */
type AttrName struct {
	Token
	ID Identifier
}

func (i *AttrName) String() string {
	return fmt.Sprintf("AttrName(%v)", i.ID)
}

//
// A.9.2 Comments
//

/*
 * comment ::=
 *   one_line_comment
 *   | block_comment
 */
type Comment interface {
	isComment()
}

/*
 * one_line_comment ::= // comment_text \n
 */
type OneLineComment struct {
	Token
	TextT nom.Span[rune]
	Text  string
}

func (c *OneLineComment) String() string {
	return fmt.Sprintf("OneLineComment(%q)", c.Text)
}

func (c *OneLineComment) Bake() error {
	c.Text = string(c.TextT.Value())
	return nil
}

func (*OneLineComment) isComment()    {}
func (*OneLineComment) isWhitespace() {}

// block_comment ::= /* comment_text */
type BlockComment struct {
	Token
	TextT nom.Span[rune]
	Text  string
}

func (c *BlockComment) String() string {
	return fmt.Sprintf("BlockComment(%q)", c.Text)
}

func (c *BlockComment) Bake() error {
	c.Text = string(c.TextT.Value())
	return nil
}

func (*BlockComment) isComment()    {}
func (*BlockComment) isWhitespace() {}

/*
 * comment_text ::= { Any_ASCII_character }
 */

//
// A.9.3 Identifiers
//

/*
 * array_identifier ::= identifier
 */
type ArrayIdentifier struct {
	Token
	ID Identifier
}

func (i *ArrayIdentifier) String() string {
	return fmt.Sprintf("ArrayIdentifier(%v)", i.ID)
}

/*
 * block_identifier ::= identifier
 */
type BlockIdentifier struct {
	Token
	ID Identifier
}

func (i *BlockIdentifier) String() string {
	return fmt.Sprintf("BlockIdentifier(%v)", i.ID)
}

/*
 * bin_identifier ::= identifier
 */
type BinIdentifier struct {
	Token
	ID Identifier
}

func (i *BinIdentifier) String() string {
	return fmt.Sprintf("BinIdentifier(%v)", i.ID)
}

/*
 * c_identifier49 ::= [ a-zA-Z_ ] { [ a-zA-Z0-9_ ] }
 */
type CIdentifier struct {
	Token
	Name string
}

func (i *CIdentifier) String() string {
	return fmt.Sprintf("CIdentifier(%v)", i.Name)
}

func (i *CIdentifier) Bake() error {
	i.Name = string(i.Span.Value())
	return assertNotKeyword(i.Name)
}

/*
 * cell_identifier ::= identifier
 */
type CellIdentifier struct {
	Token
	ID Identifier
}

func (i *CellIdentifier) String() string {
	return fmt.Sprintf("CellIdentifier(%v)", i.ID)
}

/*
 * checker_identifier ::= identifier
 */
type CheckerIdentifier struct {
	Token
	ID Identifier
}

func (i *CheckerIdentifier) String() string {
	return fmt.Sprintf("CheckerIdentifier(%v)", i.ID)
}

/*
 * class_identifier ::= identifier
 */
type ClassIdentifier struct {
	Token
	ID Identifier
}

func (i *ClassIdentifier) String() string {
	return fmt.Sprintf("ClassIdentifier(%v)", i.ID)
}

/*
 * class_variable_identifier ::= variable_identifier
 */
type ClassVariableIdentifier struct {
	Token
	Var *VariableIdentifier
}

func (i *ClassVariableIdentifier) String() string {
	return fmt.Sprintf("ClassVariableIdentifier(%v)", i.Var)
}

/*
 * clocking_identifier ::= identifier
 */
type ClockingIdentifier struct {
	Token
	ID Identifier
}

func (i *ClockingIdentifier) String() string {
	return fmt.Sprintf("ClockingIdentifier(%v)", i.ID)
}

/*
 * config_identifier ::= identifier
 */
type ConfigIdentifier struct {
	Token
	ID Identifier
}

func (i *ConfigIdentifier) String() string {
	return fmt.Sprintf("ConfigIdentifier(%v)", i.ID)
}

/*
 * const_identifier ::= identifier
 */
type ConstIdentifier struct {
	Token
	ID Identifier
}

func (i *ConstIdentifier) String() string {
	return fmt.Sprintf("ConstIdentifier(%v)", i.ID)
}

/*
 * constraint_identifier ::= identifier
 */
type ConstraintIdentifier struct {
	Token
	ID Identifier
}

func (i *ConstraintIdentifier) String() string {
	return fmt.Sprintf("ConstraintIdentifier(%v)", i.ID)
}

/*
 * covergroup_identifier ::= identifier
 */
type CovergroupIdentifier struct {
	Token
	ID Identifier
}

func (i *CovergroupIdentifier) String() string {
	return fmt.Sprintf("CovergroupIdentifier(%v)", i.ID)
}

/*
 * covergroup_variable_identifier ::= variable_identifier
 */
type CovergroupVariableIdentifier struct {
	Token
	Var *VariableIdentifier
}

func (i *CovergroupVariableIdentifier) String() string {
	return fmt.Sprintf("CovergroupVariableIdentifier(%v)", i.Var)
}

/*
 * cover_point_identifier ::= identifier
 */
type CoverPointIdentifier struct {
	Token
	ID Identifier
}

func (i *CoverPointIdentifier) String() string {
	return fmt.Sprintf("CoverPointIdentifier(%v)", i.ID)
}

/*
 * cross_identifier ::= identifier
 */
type CrossIdentifier struct {
	Token
	ID Identifier
}

func (i *CrossIdentifier) String() string {
	return fmt.Sprintf("CrossIdentifier(%v)", i.ID)
}

/*
 * dynamic_array_variable_identifier ::= variable_identifier
 */
type DynamicArrayVariableIdentifier struct {
	Token
	Var *VariableIdentifier
}

func (i *DynamicArrayVariableIdentifier) String() string {
	return fmt.Sprintf("DynamicArrayVariableIdentifier(%v)", i.Var)
}

/*
 * enum_identifier ::= identifier
 */
type EnumIdentifier struct {
	Token
	ID Identifier
}

func (i *EnumIdentifier) String() string {
	return fmt.Sprintf("EnumIdentifier(%v)", i.ID)
}

/*
 * escaped_identifier ::= \ {any_printable_ASCII_character_except_white_space} white_space
 */
type EscapedIdentifier struct {
	Token
	SlashT, NameT nom.Span[rune]
	Name          string
}

func (i *EscapedIdentifier) String() string {
	return fmt.Sprintf("EscapedIdentifier(%v)", i.Name)
}

func (i *EscapedIdentifier) Bake() error {
	i.Name = string(i.NameT.Value())
	return assertNotKeyword(i.Name)
}

func (*EscapedIdentifier) isIdentifier() {}

/*
 * formal_identifier ::= identifier
 */
type FormalIdentifier struct {
	Token
	ID Identifier
}

func (i *FormalIdentifier) String() string {
	return fmt.Sprintf("FormalIdentifier(%v)", i.ID)
}

/*
 * formal_port_identifier ::= identifier
 */
type FormalPortIdentifier struct {
	Token
	ID Identifier
}

func (i *FormalPortIdentifier) String() string {
	return fmt.Sprintf("FormalPortIdentifier(%v)", i.ID)
}

/*
 * function_identifier ::= identifier
 */
type FunctionIdentifier struct {
	Token
	ID Identifier
}

func (i *FunctionIdentifier) String() string {
	return fmt.Sprintf("FunctionIdentifier(%v)", i.ID)
}

/*
 * generate_block_identifier ::= identifier
 */
type GenerateBlockIdentifier struct {
	Token
	ID Identifier
}

func (i *GenerateBlockIdentifier) String() string {
	return fmt.Sprintf("GenerateBlockIdentifier(%v)", i.ID)
}

/*
 * genvar_identifier ::= identifier
 */
type GenvarIdentifier struct {
	Token
	ID Identifier
}

func (i *GenvarIdentifier) String() string {
	return fmt.Sprintf("GenvarIdentifier(%v)", i.ID)
}

func (*GenvarIdentifier) isConstantPrimary()    {}
func (*GenvarIdentifier) isConstantExpression() {}

/*
 * hierarchical_array_identifier ::= hierarchical_identifier
 */
type HierarchicalIdentifier struct {
	Token
	RootT nom.Span[rune]
	Root  bool
	Parts []*HierarchicalIdentifierPart
}

func (h *HierarchicalIdentifier) Bake() error {
	if len(h.RootT.Value()) > 0 {
		h.Root = true
	}
	return nil
}

func (i *HierarchicalIdentifier) String() string {
	return fmt.Sprintf("HierarchicalIdentifier(Root=%v, %v)", i.Root, i.Parts)
}

type HierarchicalIdentifierPart struct {
	Token
	ID   Identifier
	Bits *ConstantBitSelect
}

func (i *HierarchicalIdentifierPart) String() string {
	return fmt.Sprintf("HierarchicalIdentifierPart(%v, %v)", i.ID, i.Bits)
}

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

type Identifier interface {
	isIdentifier()
}

/*
 * index_variable_identifier ::= identifier
 */
type IndexVariableIdentifier struct {
	Token
	ID Identifier
}

func (i *IndexVariableIdentifier) String() string {
	return fmt.Sprintf("IndexVariableIdentifier(%v)", i.ID)
}

/*
 * interface_identifier ::= identifier
 */
type InterfaceIdentifier struct {
	Token
	ID Identifier
}

func (i *InterfaceIdentifier) String() string {
	return fmt.Sprintf("InterfaceIdentifier(%v)", i.ID)
}

/*
 * interface_instance_identifier ::= identifier
 */
type InterfaceInstanceIdentifier struct {
	Token
	ID Identifier
}

func (i *InterfaceInstanceIdentifier) String() string {
	return fmt.Sprintf("InterfaceInstanceIdentifier(%v)", i.ID)
}

/*
 * inout_port_identifier ::= identifier
 */
type InoutPortIdentifier struct {
	Token
	ID Identifier
}

func (i *InoutPortIdentifier) String() string {
	return fmt.Sprintf("InoutPortIdentifier(%v)", i.ID)
}

/*
 * input_port_identifier ::= identifier
 */
type InputPortIdentifier struct {
	Token
	ID Identifier
}

func (i *InputPortIdentifier) String() string {
	return fmt.Sprintf("InputPortIdentifier(%v)", i.ID)
}

/*
 * instance_identifier ::= identifier
 */
type InstanceIdentifier struct {
	Token
	ID Identifier
}

func (i *InstanceIdentifier) String() string {
	return fmt.Sprintf("InstanceIdentifier(%v)", i.ID)
}

/*
 * library_identifier ::= identifier
 */
type LibraryIdentifier struct {
	Token
	ID Identifier
}

func (i *LibraryIdentifier) String() string {
	return fmt.Sprintf("LibraryIdentifier(%v)", i.ID)
}

/*
 * member_identifier ::= identifier
 */
type MemberIdentifier struct {
	Token
	ID Identifier
}

func (i *MemberIdentifier) String() string {
	return fmt.Sprintf("MemberIdentifier(%v)", i.ID)
}

/*
 * method_identifier ::= identifier
 */
type MethodIdentifier struct {
	Token
	ID Identifier
}

func (i *MethodIdentifier) String() string {
	return fmt.Sprintf("MethodIdentifier(%v)", i.ID)
}

/*
 * modport_identifier ::= identifier
 */
type ModportIdentifier struct {
	Token
	ID Identifier
}

func (i *ModportIdentifier) String() string {
	return fmt.Sprintf("ModportIdentifier(%v)", i.ID)
}

/*
 * module_identifier ::= identifier
 */
type ModuleIdentifier struct {
	Token
	ID Identifier
}

func (i *ModuleIdentifier) String() string {
	return fmt.Sprintf("ModuleIdentifier(%v)", i.ID)
}

/*
 * net_identifier ::= identifier
 */
type NetIdentifier struct {
	Token
	ID Identifier
}

func (i *NetIdentifier) String() string {
	return fmt.Sprintf("NetIdentifier(%v)", i.ID)
}

/*
 * net_type_identifier ::= identifier
 */
type NetTypeIdentifier struct {
	Token
	ID Identifier
}

func (i *NetTypeIdentifier) String() string {
	return fmt.Sprintf("NetTypeIdentifier(%v)", i.ID)
}

/*
 * output_port_identifier ::= identifier
 */
type OutputPortIdentifier struct {
	Token
	ID Identifier
}

func (i *OutputPortIdentifier) String() string {
	return fmt.Sprintf("OutputPortIdentifier(%v)", i.ID)
}

/*
 * package_identifier ::= identifier
 */
type PackageIdentifier struct {
	Token
	ID Identifier
}

func (i *PackageIdentifier) String() string {
	return fmt.Sprintf("PackageIdentifier(%v)", i.ID)
}

/*
 * package_scope ::=
 *   package_identifier ::
 *   | $unit ::
 */
type PackageScope interface {
	isPackageScope()
}

/*
 * package_scope ::=
 *   ...
 *   package_identifier ::
 *   ...
 */
type IdentifierPackageScope struct {
	Token
	ID *PackageIdentifier
}

func (i *IdentifierPackageScope) String() string {
	return fmt.Sprintf("IdentifierPackageScope(%v)", i.ID)
}

func (*IdentifierPackageScope) isPackageScope() {}

/*
 * package_scope ::=
 *   ...
 *   | $unit ::
 *   ...
 */
type UnitPackageScope struct {
	Token
}

func (i *UnitPackageScope) String() string {
	return "UnitPackageScope()"
}

func (*UnitPackageScope) isPackageScope() {}

/*
 * parameter_identifier ::= identifier
 */
type ParameterIdentifier struct {
	Token
	ID Identifier
}

func (i *ParameterIdentifier) String() string {
	return fmt.Sprintf("ParameterIdentifier(%v)", i.ID)
}

/*
 * port_identifier ::= identifier
 */
type PortIdentifier struct {
	Token
	ID Identifier
}

func (i *PortIdentifier) String() string {
	return fmt.Sprintf("PortIdentifier(%v)", i.ID)
}

/*
 * production_identifier ::= identifier
 */
type ProductionIdentifier struct {
	Token
	ID Identifier
}

func (i *ProductionIdentifier) String() string {
	return fmt.Sprintf("ProductionIdentifier(%v)", i.ID)
}

/*
 * program_identifier ::= identifier
 */
type ProgramIdentifier struct {
	Token
	ID Identifier
}

func (i *ProgramIdentifier) String() string {
	return fmt.Sprintf("ProgramIdentifier(%v)", i.ID)
}

/*
 * property_identifier ::= identifier
 */
type PropertyIdentifier struct {
	Token
	ID Identifier
}

func (i *PropertyIdentifier) String() string {
	return fmt.Sprintf("PropertyIdentifier(%v)", i.ID)
}

/*
 * ps_class_identifier ::= [ package_scope ] class_identifier
 */
type PsClassIdentifier struct {
	Token
	Scope PackageScope
	ID    *ClassIdentifier
}

func (i *PsClassIdentifier) String() string {
	return fmt.Sprintf("PsClassIdentifier(%v, %v)", i.Scope, i.ID)
}

/*
 * ps_covergroup_identifier ::= [ package_scope ] covergroup_identifier
 */
type PsCovergroupIdentifier struct {
	Token
	Scope PackageScope
	ID    *CovergroupIdentifier
}

func (i *PsCovergroupIdentifier) String() string {
	return fmt.Sprintf("PsCovergroupIdentifier(%v, %v)", i.Scope, i.ID)
}

/*
 * ps_checker_identifier ::= [ package_scope ] checker_identifier
 */
type PsCheckerIdentifier struct {
	Token
	Scope PackageScope
	ID    *CheckerIdentifier
}

func (i *PsCheckerIdentifier) String() string {
	return fmt.Sprintf("PsCheckerIdentifier(%v, %v)", i.Scope, i.ID)
}

/*
 * ps_identifier ::= [ package_scope ] identifier
 */
type PsIdentifier struct {
	Token
	Scope PackageScope
	ID    Identifier
}

func (i *PsIdentifier) String() string {
	return fmt.Sprintf("PsIdentifier(%v, %v)", i.Scope, i.ID)
}

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
type SequenceIdentifier struct {
	Token
	ID Identifier
}

func (i *SequenceIdentifier) String() string {
	return fmt.Sprintf("SequenceIdentifier(%v)", i.ID)
}

/*
 * signal_identifier ::= identifier
 */
type SignalIdentifier struct {
	Token
	ID Identifier
}

func (i *SignalIdentifier) String() string {
	return fmt.Sprintf("SignalIdentifier(%v)", i.ID)
}

/*
 * simple_identifier49 ::= [ a-zA-Z_ ] { [ a-zA-Z0-9_$ ] }
 */
type SimpleIdentifier struct {
	Token
	Name string
}

func (i *SimpleIdentifier) String() string {
	return fmt.Sprintf("SimpleIdentifier(%v)", i.Name)
}

func (i *SimpleIdentifier) Bake() error {
	i.Name = string(i.Span.Value())
	return assertNotKeyword(i.Name)
}

func (*SimpleIdentifier) isIdentifier() {}

/*
 * specparam_identifier ::= identifier
 */
type SpecparamIdentifier struct {
	Token
	ID Identifier
}

func (i *SpecparamIdentifier) String() string {
	return fmt.Sprintf("SpecparamIdentifier(%v)", i.ID)
}

/*
 * system_tf_identifier50 ::= $[ a-zA-Z0-9_$ ]{ [ a-zA-Z0-9_$ ] }
 */
type SystemTfIdentifier struct {
	Token
	Name string
}

func (i *SystemTfIdentifier) String() string {
	return fmt.Sprintf("SystemTfIdentifier(%v)", i.Name)
}

func (i *SystemTfIdentifier) Bake() error {
	i.Name = string(i.Span.Value())
	return nil
}

/*
 * task_identifier ::= identifier
 */
type TaskIdentifier struct {
	Token
	ID Identifier
}

func (i *TaskIdentifier) String() string {
	return fmt.Sprintf("TaskIdentifier(%v)", i.ID)
}

/*
 * tf_identifier ::= identifier
 */
type TfIdentifier struct {
	Token
	ID Identifier
}

func (i *TfIdentifier) String() string {
	return fmt.Sprintf("TfIdentifier(%v)", i.ID)
}

/*
 * terminal_identifier ::= identifier
 */
type TerminalIdentifier struct {
	Token
	ID Identifier
}

func (i *TerminalIdentifier) String() string {
	return fmt.Sprintf("TerminalIdentifier(%v)", i.ID)
}

/*
 * topmodule_identifier ::= identifier
 */
type TopmoduleIdentifier struct {
	Token
	ID Identifier
}

func (i *TopmoduleIdentifier) String() string {
	return fmt.Sprintf("TopmoduleIdentifier(%v)", i.ID)
}

/*
 * type_identifier ::= identifier
 */
type TypeIdentifier struct {
	Token
	ID Identifier
}

func (i *TypeIdentifier) String() string {
	return fmt.Sprintf("TypeIdentifier(%v)", i.ID)
}

/*
 * udp_identifier ::= identifier
 */
type UdpIdentifier struct {
	Token
	ID Identifier
}

func (i *UdpIdentifier) String() string {
	return fmt.Sprintf("UdpIdentifier(%v)", i.ID)
}

/*
 * variable_identifier ::= identifier
 */
type VariableIdentifier struct {
	Token
	ID Identifier
}

func (i *VariableIdentifier) String() string {
	return fmt.Sprintf("VariableIdentifier(%v)", i.ID)
}

//
// A.9.4 White space
//

/*
 * white_space ::= space | tab | newline | eof
 */
type Whitespace interface {
	isWhitespace()
}

type Spaces struct {
	nom.Span[rune]
	Text string
}

func (c *Spaces) String() string {
	return fmt.Sprintf("Spaces(%q)", c.Text)
}

func (c *Spaces) Bake() error {
	c.Text = string(c.Span.Value())
	return nil
}

func (*Spaces) isWhitespace() {}
