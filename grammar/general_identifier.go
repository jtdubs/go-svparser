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

func Identifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.Identifier, error) {
	return trace.Trace(cache.Cache(
		fn.Alt(
			to[ast.Identifier](simpleIdentifier),
			to[ast.Identifier](escapedIdentifier),
		),
	))(ctx, start)
}

func escapedIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.EscapedIdentifier, error) {
	res := &ast.EscapedIdentifier{}
	return tBindSeq(res, &res.Span,
		bindSpan(&res.SlashT, runes.Rune('\\')),
		bindSpan(&res.NameT, fn.Terminated(fn.Many1(asciiPrintNonWS), fn.Peek(fn.Alt(runes.Space)))),
	)(ctx, start)
}

func simpleIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.SimpleIdentifier, error) {
	res := &ast.SimpleIdentifier{}
	return tBind(res, &res.Span,
		fn.Preceded(alpha_, fn.Many0(alphanumeric_S)),
	)(ctx, start)
}

func CIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.CIdentifier, error) {
	res := &ast.CIdentifier{}
	return tBind(res, &res.Span, fn.Preceded(alpha_, fn.Many0(alphanumeric_)))(ctx, start)
}

func SystemTfIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.SystemTfIdentifier, error) {
	res := &ast.SystemTfIdentifier{}
	return tBind(res, &res.Span, fn.Preceded(runes.Rune('$'), fn.Many1(alphanumeric_S)))(ctx, start)
}

func ArrayIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ArrayIdentifier, error) {
	res := &ast.ArrayIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func BlockIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.BlockIdentifier, error) {
	res := &ast.BlockIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func BinIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.BinIdentifier, error) {
	res := &ast.BinIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func CellIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.CellIdentifier, error) {
	res := &ast.CellIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func CheckerIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.CheckerIdentifier, error) {
	res := &ast.CheckerIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func ClassIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ClassIdentifier, error) {
	res := &ast.ClassIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func ClockingIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ClockingIdentifier, error) {
	res := &ast.ClockingIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func ConfigIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ConfigIdentifier, error) {
	res := &ast.ConfigIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func ConstIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ConstIdentifier, error) {
	res := &ast.ConstIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func ConstraintIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ConstraintIdentifier, error) {
	res := &ast.ConstraintIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func CovergroupIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.CovergroupIdentifier, error) {
	res := &ast.CovergroupIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func CoverPointIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.CoverPointIdentifier, error) {
	res := &ast.CoverPointIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func CrossIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.CrossIdentifier, error) {
	res := &ast.CrossIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func EnumIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.EnumIdentifier, error) {
	res := &ast.EnumIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func FormalIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.FormalIdentifier, error) {
	res := &ast.FormalIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func FormalPortIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.FormalPortIdentifier, error) {
	res := &ast.FormalPortIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func FunctionIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.FunctionIdentifier, error) {
	res := &ast.FunctionIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func GenerateBlockIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.GenerateBlockIdentifier, error) {
	res := &ast.GenerateBlockIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func GenvarIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.GenvarIdentifier, error) {
	res := &ast.GenvarIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func IndexVariableIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.IndexVariableIdentifier, error) {
	res := &ast.IndexVariableIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func InterfaceIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.InterfaceIdentifier, error) {
	res := &ast.InterfaceIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func InterfaceInstanceIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.InterfaceInstanceIdentifier, error) {
	res := &ast.InterfaceInstanceIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func InoutPortIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.InoutPortIdentifier, error) {
	res := &ast.InoutPortIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func InstanceIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.InstanceIdentifier, error) {
	res := &ast.InstanceIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func LibraryIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.LibraryIdentifier, error) {
	res := &ast.LibraryIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func MemberIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.MemberIdentifier, error) {
	res := &ast.MemberIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func MethodIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.MethodIdentifier, error) {
	res := &ast.MethodIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func ModportIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ModportIdentifier, error) {
	res := &ast.ModportIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func ModuleIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ModuleIdentifier, error) {
	res := &ast.ModuleIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func NetIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.NetIdentifier, error) {
	res := &ast.NetIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func NetTypeIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.NetTypeIdentifier, error) {
	res := &ast.NetTypeIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func OutputPortIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.OutputPortIdentifier, error) {
	res := &ast.OutputPortIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func PackageIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.PackageIdentifier, error) {
	res := &ast.PackageIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func ParameterIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ParameterIdentifier, error) {
	res := &ast.ParameterIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func PortIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.PortIdentifier, error) {
	res := &ast.PortIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func ProductionIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ProductionIdentifier, error) {
	res := &ast.ProductionIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func ProgramIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ProgramIdentifier, error) {
	res := &ast.ProgramIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func PropertyIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.PropertyIdentifier, error) {
	res := &ast.PropertyIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func SequenceIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.SequenceIdentifier, error) {
	res := &ast.SequenceIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func SignalIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.SignalIdentifier, error) {
	res := &ast.SignalIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func SpecparamIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.SpecparamIdentifier, error) {
	res := &ast.SpecparamIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func TaskIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.TaskIdentifier, error) {
	res := &ast.TaskIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func TfIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.TfIdentifier, error) {
	res := &ast.TfIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func TerminalIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.TerminalIdentifier, error) {
	res := &ast.TerminalIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func TopmoduleIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.TopmoduleIdentifier, error) {
	res := &ast.TopmoduleIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func TypeIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.TypeIdentifier, error) {
	res := &ast.TypeIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func UdpIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.UdpIdentifier, error) {
	res := &ast.UdpIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func VariableIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.VariableIdentifier, error) {
	res := &ast.VariableIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.ID, Identifier))(ctx, start)
}

func ClassVariableIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.ClassVariableIdentifier, error) {
	res := &ast.ClassVariableIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.Var, VariableIdentifier))(ctx, start)
}

func CovergroupVariableIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.CovergroupVariableIdentifier, error) {
	res := &ast.CovergroupVariableIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.Var, VariableIdentifier))(ctx, start)
}

func DynamicArrayVariableIdentifier(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.DynamicArrayVariableIdentifier, error) {
	res := &ast.DynamicArrayVariableIdentifier{}
	return tBind(res, &res.Span, bindValue(&res.Var, VariableIdentifier))(ctx, start)
}

var asciiPrintNonWS = fn.Satisfy(func(r rune) bool {
	return r < 128 && unicode.IsPrint(r) && !unicode.IsSpace(r)
})

var alpha_ = runes.OneOf("abcdefghijklmnoprqstuvwxyzABCDEFGHIJKLMNOPRQSTUVWXYZ_")
var alphanumeric_ = runes.OneOf("abcdefghijklmnoprqstuvwxyzABCDEFGHIJKLMNOPRQSTUVWXYZ0123456789_")
var alphanumeric_S = runes.OneOf("abcdefghijklmnoprqstuvwxyzABCDEFGHIJKLMNOPRQSTUVWXYZ0123456789_$")
