package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

//
// A.3.4 Primitive gate and switch types
//

/*
 * cmos_switchtype ::= cmos | rcmos
 */
func CMOSSwitchType(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.CMOSSwitchType, error) {
	res := &ast.CMOSSwitchType{}
	return tBindPhrase(res,
		bindValue(&res.Type,
			fn.Alt(
				fn.Value(ast.CMOS, runes.TagNoCase("cmos")),
				fn.Value(ast.RCMOS, runes.TagNoCase("rcmos")),
			),
		),
	)(ctx, start)
}

/*
 * enable_gatetype ::= bufif0 | bufif1 | notif0 | notif1
 */
func EnableGateType(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.EnableGateType, error) {
	res := &ast.EnableGateType{}
	return tBindPhrase(res,
		bindValue(&res.Type,
			fn.Alt(
				fn.Value(ast.BUFIF0, runes.TagNoCase("bufif0")),
				fn.Value(ast.BUFIF1, runes.TagNoCase("bufif1")),
				fn.Value(ast.NOTIF0, runes.TagNoCase("notif0")),
				fn.Value(ast.NOTIF1, runes.TagNoCase("notif1")),
			),
		),
	)(ctx, start)
}

/*
 * mos_switchtype ::= nmos | pmos | rnmos | rpmos
 */
func MOSSwitchType(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.MOSSwitchType, error) {
	res := &ast.MOSSwitchType{}
	return tBindPhrase(res,
		bindValue(&res.Type,
			fn.Alt(
				fn.Value(ast.NMOS, runes.TagNoCase("nmos")),
				fn.Value(ast.PMOS, runes.TagNoCase("pmos")),
				fn.Value(ast.RNMOS, runes.TagNoCase("rnmos")),
				fn.Value(ast.RPMOS, runes.TagNoCase("rpmos")),
			),
		),
	)(ctx, start)
}

/*
 * n_input_gatetype ::= and | nand | or | nor | xor | xnor
 */
func NInputGateType(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.NInputGateType, error) {
	res := &ast.NInputGateType{}
	return tBindPhrase(res,
		bindValue(&res.Type,
			fn.Alt(
				fn.Value(ast.InputGateAnd, runes.TagNoCase("and")),
				fn.Value(ast.InputGateNand, runes.TagNoCase("nand")),
				fn.Value(ast.InputGateOr, runes.TagNoCase("or")),
				fn.Value(ast.InputGateNor, runes.TagNoCase("nor")),
				fn.Value(ast.InputGateXor, runes.TagNoCase("xor")),
				fn.Value(ast.InputGateXnor, runes.TagNoCase("xnor")),
			),
		),
	)(ctx, start)
}

/*
 * n_output_gatetype ::= buf | not
 */
func NOutputGateType(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.NOutputGateType, error) {
	res := &ast.NOutputGateType{}
	return tBindPhrase(res,
		bindValue(&res.Type,
			fn.Alt(
				fn.Value(ast.OutputGateBuf, runes.TagNoCase("buf")),
				fn.Value(ast.OutputGateNot, runes.TagNoCase("not")),
			),
		),
	)(ctx, start)
}

/*
 * pass_en_switchtype ::= tranif0 | tranif1 | rtranif1 | rtranif0
 */
func PassEnSwitchType(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.PassEnSwitchType, error) {
	res := &ast.PassEnSwitchType{}
	return tBindPhrase(res,
		bindValue(&res.Type,
			fn.Alt(
				fn.Value(ast.TRANIF0, runes.TagNoCase("tranif0")),
				fn.Value(ast.TRANIF1, runes.TagNoCase("tranif1")),
				fn.Value(ast.RTRANIF0, runes.TagNoCase("rtranif0")),
				fn.Value(ast.RTRANIF1, runes.TagNoCase("rtranif1")),
			),
		),
	)(ctx, start)
}

/*
 * pass_switchtype ::= tran | rtran
 */
func PassSwitchType(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.PassSwitchType, error) {
	res := &ast.PassSwitchType{}
	return tBindPhrase(res,
		bindValue(&res.Type,
			fn.Alt(
				fn.Value(ast.TRAN, runes.TagNoCase("tran")),
				fn.Value(ast.RTRAN, runes.TagNoCase("rtran")),
			),
		),
	)(ctx, start)
}
