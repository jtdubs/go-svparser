package ast

import (
	"fmt"

	"github.com/jtdubs/go-nom"
)

type Identifier interface {
	isIdentifier()
}

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

type ArrayIdentifier struct {
	Token
	ID Identifier
}

func (i *ArrayIdentifier) String() string {
	return fmt.Sprintf("ArrayIdentifier(%v)", i.ID)
}

type BlockIdentifier struct {
	Token
	ID Identifier
}

func (i *BlockIdentifier) String() string {
	return fmt.Sprintf("BlockIdentifier(%v)", i.ID)
}

type BinIdentifier struct {
	Token
	ID Identifier
}

func (i *BinIdentifier) String() string {
	return fmt.Sprintf("BinIdentifier(%v)", i.ID)
}

type CellIdentifier struct {
	Token
	ID Identifier
}

func (i *CellIdentifier) String() string {
	return fmt.Sprintf("CellIdentifier(%v)", i.ID)
}

type CheckerIdentifier struct {
	Token
	ID Identifier
}

func (i *CheckerIdentifier) String() string {
	return fmt.Sprintf("CheckerIdentifier(%v)", i.ID)
}

type ClassIdentifier struct {
	Token
	ID Identifier
}

func (i *ClassIdentifier) String() string {
	return fmt.Sprintf("ClassIdentifier(%v)", i.ID)
}

type ClockingIdentifier struct {
	Token
	ID Identifier
}

func (i *ClockingIdentifier) String() string {
	return fmt.Sprintf("ClockingIdentifier(%v)", i.ID)
}

type ConfigIdentifier struct {
	Token
	ID Identifier
}

func (i *ConfigIdentifier) String() string {
	return fmt.Sprintf("ConfigIdentifier(%v)", i.ID)
}

type ConstIdentifier struct {
	Token
	ID Identifier
}

func (i *ConstIdentifier) String() string {
	return fmt.Sprintf("ConstIdentifier(%v)", i.ID)
}

type ConstraintIdentifier struct {
	Token
	ID Identifier
}

func (i *ConstraintIdentifier) String() string {
	return fmt.Sprintf("ConstraintIdentifier(%v)", i.ID)
}

type CovergroupIdentifier struct {
	Token
	ID Identifier
}

func (i *CovergroupIdentifier) String() string {
	return fmt.Sprintf("CovergroupIdentifier(%v)", i.ID)
}

type CoverPointIdentifier struct {
	Token
	ID Identifier
}

func (i *CoverPointIdentifier) String() string {
	return fmt.Sprintf("CoverPointIdentifier(%v)", i.ID)
}

type CrossIdentifier struct {
	Token
	ID Identifier
}

func (i *CrossIdentifier) String() string {
	return fmt.Sprintf("CrossIdentifier(%v)", i.ID)
}

type EnumIdentifier struct {
	Token
	ID Identifier
}

func (i *EnumIdentifier) String() string {
	return fmt.Sprintf("EnumIdentifier(%v)", i.ID)
}

type FormalIdentifier struct {
	Token
	ID Identifier
}

func (i *FormalIdentifier) String() string {
	return fmt.Sprintf("FormalIdentifier(%v)", i.ID)
}

type FormalPortIdentifier struct {
	Token
	ID Identifier
}

func (i *FormalPortIdentifier) String() string {
	return fmt.Sprintf("FormalPortIdentifier(%v)", i.ID)
}

type FunctionIdentifier struct {
	Token
	ID Identifier
}

func (i *FunctionIdentifier) String() string {
	return fmt.Sprintf("FunctionIdentifier(%v)", i.ID)
}

type GenerateBlockIdentifier struct {
	Token
	ID Identifier
}

func (i *GenerateBlockIdentifier) String() string {
	return fmt.Sprintf("GenerateBlockIdentifier(%v)", i.ID)
}

type GenvarIdentifier struct {
	Token
	ID Identifier
}

func (i *GenvarIdentifier) String() string {
	return fmt.Sprintf("GenvarIdentifier(%v)", i.ID)
}

func (*GenvarIdentifier) isConstantPrimary()    {}
func (*GenvarIdentifier) isConstantExpression() {}

type IndexVariableIdentifier struct {
	Token
	ID Identifier
}

func (i *IndexVariableIdentifier) String() string {
	return fmt.Sprintf("IndexVariableIdentifier(%v)", i.ID)
}

type InterfaceIdentifier struct {
	Token
	ID Identifier
}

func (i *InterfaceIdentifier) String() string {
	return fmt.Sprintf("InterfaceIdentifier(%v)", i.ID)
}

type InterfaceInstanceIdentifier struct {
	Token
	ID Identifier
}

func (i *InterfaceInstanceIdentifier) String() string {
	return fmt.Sprintf("InterfaceInstanceIdentifier(%v)", i.ID)
}

type InoutPortIdentifier struct {
	Token
	ID Identifier
}

func (i *InoutPortIdentifier) String() string {
	return fmt.Sprintf("InoutPortIdentifier(%v)", i.ID)
}

type InputPortIdentifier struct {
	Token
	ID Identifier
}

func (i *InputPortIdentifier) String() string {
	return fmt.Sprintf("InputPortIdentifier(%v)", i.ID)
}

type InstanceIdentifier struct {
	Token
	ID Identifier
}

func (i *InstanceIdentifier) String() string {
	return fmt.Sprintf("InstanceIdentifier(%v)", i.ID)
}

type LibraryIdentifier struct {
	Token
	ID Identifier
}

func (i *LibraryIdentifier) String() string {
	return fmt.Sprintf("LibraryIdentifier(%v)", i.ID)
}

type MemberIdentifier struct {
	Token
	ID Identifier
}

func (i *MemberIdentifier) String() string {
	return fmt.Sprintf("MemberIdentifier(%v)", i.ID)
}

type MethodIdentifier struct {
	Token
	ID Identifier
}

func (i *MethodIdentifier) String() string {
	return fmt.Sprintf("MethodIdentifier(%v)", i.ID)
}

type ModportIdentifier struct {
	Token
	ID Identifier
}

func (i *ModportIdentifier) String() string {
	return fmt.Sprintf("ModportIdentifier(%v)", i.ID)
}

type ModuleIdentifier struct {
	Token
	ID Identifier
}

func (i *ModuleIdentifier) String() string {
	return fmt.Sprintf("ModuleIdentifier(%v)", i.ID)
}

type NetIdentifier struct {
	Token
	ID Identifier
}

func (i *NetIdentifier) String() string {
	return fmt.Sprintf("NetIdentifier(%v)", i.ID)
}

type NetTypeIdentifier struct {
	Token
	ID Identifier
}

func (i *NetTypeIdentifier) String() string {
	return fmt.Sprintf("NetTypeIdentifier(%v)", i.ID)
}

type OutputPortIdentifier struct {
	Token
	ID Identifier
}

func (i *OutputPortIdentifier) String() string {
	return fmt.Sprintf("OutputPortIdentifier(%v)", i.ID)
}

type PackageIdentifier struct {
	Token
	ID Identifier
}

func (i *PackageIdentifier) String() string {
	return fmt.Sprintf("PackageIdentifier(%v)", i.ID)
}

type ParameterIdentifier struct {
	Token
	ID Identifier
}

func (i *ParameterIdentifier) String() string {
	return fmt.Sprintf("ParameterIdentifier(%v)", i.ID)
}

type PortIdentifier struct {
	Token
	ID Identifier
}

func (i *PortIdentifier) String() string {
	return fmt.Sprintf("PortIdentifier(%v)", i.ID)
}

type ProductionIdentifier struct {
	Token
	ID Identifier
}

func (i *ProductionIdentifier) String() string {
	return fmt.Sprintf("ProductionIdentifier(%v)", i.ID)
}

type ProgramIdentifier struct {
	Token
	ID Identifier
}

func (i *ProgramIdentifier) String() string {
	return fmt.Sprintf("ProgramIdentifier(%v)", i.ID)
}

type PropertyIdentifier struct {
	Token
	ID Identifier
}

func (i *PropertyIdentifier) String() string {
	return fmt.Sprintf("PropertyIdentifier(%v)", i.ID)
}

type SequenceIdentifier struct {
	Token
	ID Identifier
}

func (i *SequenceIdentifier) String() string {
	return fmt.Sprintf("SequenceIdentifier(%v)", i.ID)
}

type SignalIdentifier struct {
	Token
	ID Identifier
}

func (i *SignalIdentifier) String() string {
	return fmt.Sprintf("SignalIdentifier(%v)", i.ID)
}

type SpecparamIdentifier struct {
	Token
	ID Identifier
}

func (i *SpecparamIdentifier) String() string {
	return fmt.Sprintf("SpecparamIdentifier(%v)", i.ID)
}

type TaskIdentifier struct {
	Token
	ID Identifier
}

func (i *TaskIdentifier) String() string {
	return fmt.Sprintf("TaskIdentifier(%v)", i.ID)
}

type TfIdentifier struct {
	Token
	ID Identifier
}

func (i *TfIdentifier) String() string {
	return fmt.Sprintf("TfIdentifier(%v)", i.ID)
}

type TerminalIdentifier struct {
	Token
	ID Identifier
}

func (i *TerminalIdentifier) String() string {
	return fmt.Sprintf("TerminalIdentifier(%v)", i.ID)
}

type TopmoduleIdentifier struct {
	Token
	ID Identifier
}

func (i *TopmoduleIdentifier) String() string {
	return fmt.Sprintf("TopmoduleIdentifier(%v)", i.ID)
}

type TypeIdentifier struct {
	Token
	ID Identifier
}

func (i *TypeIdentifier) String() string {
	return fmt.Sprintf("TypeIdentifier(%v)", i.ID)
}

type UdpIdentifier struct {
	Token
	ID Identifier
}

func (i *UdpIdentifier) String() string {
	return fmt.Sprintf("UdpIdentifier(%v)", i.ID)
}

type VariableIdentifier struct {
	Token
	ID Identifier
}

func (i *VariableIdentifier) String() string {
	return fmt.Sprintf("VariableIdentifier(%v)", i.ID)
}

type ClassVariableIdentifier struct {
	Token
	Var *VariableIdentifier
}

func (i *ClassVariableIdentifier) String() string {
	return fmt.Sprintf("ClassVariableIdentifier(%v)", i.Var)
}

type CovergroupVariableIdentifier struct {
	Token
	Var *VariableIdentifier
}

func (i *CovergroupVariableIdentifier) String() string {
	return fmt.Sprintf("CovergroupVariableIdentifier(%v)", i.Var)
}

type DynamicArrayVariableIdentifier struct {
	Token
	Var *VariableIdentifier
}

func (i *DynamicArrayVariableIdentifier) String() string {
	return fmt.Sprintf("DynamicArrayVariableIdentifier(%v)", i.Var)
}

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
