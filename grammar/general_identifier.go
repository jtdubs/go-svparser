package grammar

import (
	"context"
	"unicode"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/cache"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-nom/trace"
	"github.com/jtdubs/go-svparser/ast"
)

//
// A.9.3 Identifiers
//

/*
 * array_identifier ::= identifier
 */
func ArrayIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ArrayIdentifier, error) {
	res := &ast.ArrayIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * block_identifier ::= identifier
 */
func BlockIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.BlockIdentifier, error) {
	res := &ast.BlockIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * bin_identifier ::= identifier
 */
func BinIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.BinIdentifier, error) {
	res := &ast.BinIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * c_identifier ::= [ a-zA-Z_ ] { [ a-zA-Z0-9_ ] }
 */
func CIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.CIdentifier, error) {
	res := &ast.CIdentifier{}
	return tBind(res, &res.Span, fn.Preceded(alpha_, fn.Many0(alphanumeric_)))(ctx, start)
}

/*
 * cell_identifier ::= identifier
 */
func CellIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.CellIdentifier, error) {
	res := &ast.CellIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * checker_identifier ::= identifier
 */
func CheckerIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.CheckerIdentifier, error) {
	res := &ast.CheckerIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * class_identifier ::= identifier
 */
func ClassIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ClassIdentifier, error) {
	res := &ast.ClassIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * class_variable_identifier ::= variable_identifier
 */
func ClassVariableIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ClassVariableIdentifier, error) {
	res := &ast.ClassVariableIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.Var, VariableIdentifier))(ctx, start)
}

/*
 * clocking_identifier ::= identifier
 */
func ClockingIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ClockingIdentifier, error) {
	res := &ast.ClockingIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * config_identifier ::= identifier
 */
func ConfigIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ConfigIdentifier, error) {
	res := &ast.ConfigIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * const_identifier ::= identifier
 */
func ConstIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ConstIdentifier, error) {
	res := &ast.ConstIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * constraint_identifier ::= identifier
 */
func ConstraintIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ConstraintIdentifier, error) {
	res := &ast.ConstraintIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * covergroup_identifier ::= identifier
 */
func CovergroupIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.CovergroupIdentifier, error) {
	res := &ast.CovergroupIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * covergroup_variable_identifier ::= variable_identifier
 */
func CovergroupVariableIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.CovergroupVariableIdentifier, error) {
	res := &ast.CovergroupVariableIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.Var, VariableIdentifier))(ctx, start)
}

/*
 * cover_point_identifier ::= identifier
 */
func CoverPointIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.CoverPointIdentifier, error) {
	res := &ast.CoverPointIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * cross_identifier ::= identifier
 */
func CrossIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.CrossIdentifier, error) {
	res := &ast.CrossIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * dynamic_array_variable_identifier ::= variable_identifier
 */
func DynamicArrayVariableIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DynamicArrayVariableIdentifier, error) {
	res := &ast.DynamicArrayVariableIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.Var, VariableIdentifier))(ctx, start)
}

/*
 * enum_identifier ::= identifier
 */
func EnumIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.EnumIdentifier, error) {
	res := &ast.EnumIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * escaped_identifier ::= \ {any_printable_ASCII_character_except_white_space} white_space
 */
func escapedIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.EscapedIdentifier, error) {
	res := &ast.EscapedIdentifier{}
	return tBindSeq(res, &res.Span,
		bindSpan(&res.SlashT, runes.Rune('\\')),
		bindSpan(&res.NameT, fn.Terminated(fn.Many1(asciiPrintNonWS), fn.Peek(fn.Alt(runes.Space)))),
	)(ctx, start)
}

/*
 * formal_identifier ::= identifier
 */
func FormalIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.FormalIdentifier, error) {
	res := &ast.FormalIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * formal_port_identifier ::= identifier
 */
func FormalPortIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.FormalPortIdentifier, error) {
	res := &ast.FormalPortIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * function_identifier ::= identifier
 */
func FunctionIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.FunctionIdentifier, error) {
	res := &ast.FunctionIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * generate_block_identifier ::= identifier
 */
func GenerateBlockIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.GenerateBlockIdentifier, error) {
	res := &ast.GenerateBlockIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * genvar_identifier ::= identifier
 */
func GenvarIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.GenvarIdentifier, error) {
	res := &ast.GenvarIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * identifier ::=
 *   simple_identifier
 * | escaped_identifier
 */
func Identifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.Identifier, error) {
	return trace.Trace(cache.Cache(
		fn.Alt(
			to[ast.Identifier](simpleIdentifier),
			to[ast.Identifier](escapedIdentifier),
		),
	))(ctx, start)
}

/*
 * index_variable_identifier ::= identifier
 */
func IndexVariableIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.IndexVariableIdentifier, error) {
	res := &ast.IndexVariableIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * interface_identifier ::= identifier
 */
func InterfaceIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.InterfaceIdentifier, error) {
	res := &ast.InterfaceIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * interface_instance_identifier ::= identifier
 */
func InterfaceInstanceIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.InterfaceInstanceIdentifier, error) {
	res := &ast.InterfaceInstanceIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * inout_port_identifier ::= identifier
 */
func InoutPortIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.InoutPortIdentifier, error) {
	res := &ast.InoutPortIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * input_port_identifier ::= identifier
 */
func InputPortIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.InputPortIdentifier, error) {
	res := &ast.InputPortIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * instance_identifier ::= identifier
 */
func InstanceIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.InstanceIdentifier, error) {
	res := &ast.InstanceIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * library_identifier ::= identifier
 */
func LibraryIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.LibraryIdentifier, error) {
	res := &ast.LibraryIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * member_identifier ::= identifier
 */
func MemberIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.MemberIdentifier, error) {
	res := &ast.MemberIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * method_identifier ::= identifier
 */
func MethodIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.MethodIdentifier, error) {
	res := &ast.MethodIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * modport_identifier ::= identifier
 */
func ModportIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ModportIdentifier, error) {
	res := &ast.ModportIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * module_identifier ::= identifier
 */
func ModuleIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ModuleIdentifier, error) {
	res := &ast.ModuleIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * net_identifier ::= identifier
 */
func NetIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.NetIdentifier, error) {
	res := &ast.NetIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * net_type_identifier ::= identifier
 */
func NetTypeIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.NetTypeIdentifier, error) {
	res := &ast.NetTypeIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * output_port_identifier ::= identifier
 */
func OutputPortIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.OutputPortIdentifier, error) {
	res := &ast.OutputPortIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * package_identifier ::= identifier
 */
func PackageIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.PackageIdentifier, error) {
	res := &ast.PackageIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * parameter_identifier ::= identifier
 */
func ParameterIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ParameterIdentifier, error) {
	res := &ast.ParameterIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * port_identifier ::= identifier
 */
func PortIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.PortIdentifier, error) {
	res := &ast.PortIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * production_identifier ::= identifier
 */
func ProductionIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ProductionIdentifier, error) {
	res := &ast.ProductionIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * program_identifier ::= identifier
 */
func ProgramIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ProgramIdentifier, error) {
	res := &ast.ProgramIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * property_identifier ::= identifier
 */
func PropertyIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.PropertyIdentifier, error) {
	res := &ast.PropertyIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * sequence_identifier ::= identifier
 */
func SequenceIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.SequenceIdentifier, error) {
	res := &ast.SequenceIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * signal_identifier ::= identifier
 */
func SignalIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.SignalIdentifier, error) {
	res := &ast.SignalIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * simple_identifier ::= [ a-zA-Z_ ] { [ a-zA-Z0-9_$ ] }
 */
func simpleIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.SimpleIdentifier, error) {
	res := &ast.SimpleIdentifier{}
	return tBind(res, &res.Span,
		fn.Preceded(alpha_, fn.Many0(alphanumeric_S)),
	)(ctx, start)
}

/*
 * specparam_identifier ::= identifier
 */
func SpecparamIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.SpecparamIdentifier, error) {
	res := &ast.SpecparamIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * system_tf_identifier ::= $[ a-zA-Z0-9_$ ]{ [ a-zA-Z0-9_$ ] }
 */
func SystemTfIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.SystemTfIdentifier, error) {
	res := &ast.SystemTfIdentifier{}
	return tBind(res, &res.Span, fn.Preceded(runes.Rune('$'), fn.Many1(alphanumeric_S)))(ctx, start)
}

/*
 * task_identifier ::= identifier
 */
func TaskIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.TaskIdentifier, error) {
	res := &ast.TaskIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * tf_identifier ::= identifier
 */
func TfIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.TfIdentifier, error) {
	res := &ast.TfIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * terminal_identifier ::= identifier
 */
func TerminalIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.TerminalIdentifier, error) {
	res := &ast.TerminalIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * topmodule_identifier ::= identifier
 */
func TopmoduleIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.TopmoduleIdentifier, error) {
	res := &ast.TopmoduleIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * type_identifier ::= identifier
 */
func TypeIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.TypeIdentifier, error) {
	res := &ast.TypeIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * udp_identifier ::= identifier
 */
func UdpIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.UdpIdentifier, error) {
	res := &ast.UdpIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

/*
 * variable_identifier ::= identifier
 */
func VariableIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.VariableIdentifier, error) {
	res := &ast.VariableIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

var asciiPrintNonWS = fn.Satisfy(func(r rune) bool {
	return r < 128 && unicode.IsPrint(r) && !unicode.IsSpace(r)
})

var alpha_ = runes.OneOf("abcdefghijklmnoprqstuvwxyzABCDEFGHIJKLMNOPRQSTUVWXYZ_")
var alphanumeric_ = runes.OneOf("abcdefghijklmnoprqstuvwxyzABCDEFGHIJKLMNOPRQSTUVWXYZ0123456789_")
var alphanumeric_S = runes.OneOf("abcdefghijklmnoprqstuvwxyzABCDEFGHIJKLMNOPRQSTUVWXYZ0123456789_$")
