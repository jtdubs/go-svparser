package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

//
// A.3 Primitive instances
//

//
// A.3.1 Primitive instantiation and instances
//

/*
 * gate_instantiation ::=
 *   cmos_switchtype [delay3] cmos_switch_instance { , cmos_switch_instance } ;
 *   | enable_gatetype [drive_strength] [delay3] enable_gate_instance { , enable_gate_instance } ;
 *   | mos_switchtype [delay3] mos_switch_instance { , mos_switch_instance } ;
 *   | n_input_gatetype [drive_strength] [delay2] n_input_gate_instance { , n_input_gate_instance } ;
 *   | n_output_gatetype [drive_strength] [delay2] n_output_gate_instance
 *   { , n_output_gate_instance } ;
 *   | pass_en_switchtype [delay2] pass_enable_switch_instance { , pass_enable_switch_instance } ;
 *   | pass_switchtype pass_switch_instance { , pass_switch_instance } ;
 *   | pulldown [pulldown_strength] pull_gate_instance { , pull_gate_instance } ;
 *   | pullup [pullup_strength] pull_gate_instance { , pull_gate_instance } ;
 */

/*
 * cmos_switch_instance ::= [ name_of_instance ] ( output_terminal , input_terminal ,
 *   ncontrol_terminal , pcontrol_terminal )
 */

/*
 * enable_gate_instance ::= [ name_of_instance ] ( output_terminal , input_terminal , enable_terminal )
 */

/*
 * mos_switch_instance ::= [ name_of_instance ] ( output_terminal , input_terminal , enable_terminal )
 */

/*
 * n_input_gate_instance ::= [ name_of_instance ] ( output_terminal , input_terminal { , input_terminal } )
 */

/*
 * n_output_gate_instance ::= [ name_of_instance ] ( output_terminal { , output_terminal } ,
 *   input_terminal )
 */

/*
 * pass_switch_instance ::= [ name_of_instance ] ( inout_terminal , inout_terminal )
 */

/*
 * pass_enable_switch_instance ::= [ name_of_instance ] ( inout_terminal , inout_terminal ,
 *   enable_terminal )
 */

/*
 * pull_gate_instance ::= [ name_of_instance ] ( output_terminal )
 */

//
// A.3.2 Primitive strengths
//

/*
 * pulldown_strength ::=
 *   ( strength0 , strength1 )
 *   | ( strength1 , strength0 )
 *   | ( strength0 )
 */
func PulldownStrength(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.PulldownStrength, error) {
	res := &ast.PulldownStrength{}
	return top(
		token(res,
			parens(
				fn.Alt(
					phrase(
						bind(&res.Zero, Strength0),
						comma,
						bind(&res.One, Strength1)),
					phrase(
						bind(&res.One, Strength1),
						comma,
						bind(&res.Zero, Strength0)),
					bind(&res.Zero, Strength0),
				),
			),
		),
	)(ctx, start)
}

/*
 * pullup_strength ::=
 *   ( strength0 , strength1 )
 *   | ( strength1 , strength0 )
 *   | ( strength1 )
 */
func PullupStrength(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.PullupStrength, error) {
	res := &ast.PullupStrength{}
	return top(token(res,
		parens(
			fn.Alt(
				phrase(
					bind(&res.Zero, Strength0),
					comma,
					bind(&res.One, Strength1)),
				phrase(
					bind(&res.One, Strength1),
					comma,
					bind(&res.Zero, Strength0)),
				bind(&res.One, Strength1),
			),
		),
	))(ctx, start)
}

//
// A.3.3 Primitive terminals
//

/*
 * enable_terminal ::= expression
 */

/*
 * inout_terminal ::= net_lvalue
 */

/*
 * input_terminal ::= expression
 */

/*
 * ncontrol_terminal ::= expression
 */

/*
 * output_terminal ::= net_lvalue
 */

/*
 * pcontrol_terminal ::= expression
 */

//
// A.3.4 Primitive gate and switch types
//

/*
 * cmos_switchtype ::= cmos | rcmos
 */
func CMOSSwitchType(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.CMOSSwitchType, error) {
	res := &ast.CMOSSwitchType{}
	return top(
		token(res,
			word(
				bind(&res.Type,
					fn.Alt(
						fn.Value(ast.CMOS, runes.TagNoCase("cmos")),
						fn.Value(ast.RCMOS, runes.TagNoCase("rcmos")),
					),
				),
			),
		),
	)(ctx, start)
}

/*
 * enable_gatetype ::= bufif0 | bufif1 | notif0 | notif1
 */
func EnableGateType(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.EnableGateType, error) {
	res := &ast.EnableGateType{}
	return top(
		token(res,
			word(
				bind(&res.Type,
					fn.Alt(
						fn.Value(ast.BUFIF0, runes.TagNoCase("bufif0")),
						fn.Value(ast.BUFIF1, runes.TagNoCase("bufif1")),
						fn.Value(ast.NOTIF0, runes.TagNoCase("notif0")),
						fn.Value(ast.NOTIF1, runes.TagNoCase("notif1")),
					),
				),
			),
		),
	)(ctx, start)
}

/*
 * mos_switchtype ::= nmos | pmos | rnmos | rpmos
 */
func MOSSwitchType(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.MOSSwitchType, error) {
	res := &ast.MOSSwitchType{}
	return top(
		token(res,
			word(
				bind(&res.Type,
					fn.Alt(
						fn.Value(ast.NMOS, runes.TagNoCase("nmos")),
						fn.Value(ast.PMOS, runes.TagNoCase("pmos")),
						fn.Value(ast.RNMOS, runes.TagNoCase("rnmos")),
						fn.Value(ast.RPMOS, runes.TagNoCase("rpmos")),
					),
				),
			),
		),
	)(ctx, start)
}

/*
 * n_input_gatetype ::= and | nand | or | nor | xor | xnor
 */
func NInputGateType(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.NInputGateType, error) {
	res := &ast.NInputGateType{}
	return top(
		token(res,
			word(
				bind(&res.Type,
					fn.Alt(
						fn.Value(ast.InputGateAnd, runes.TagNoCase("and")),
						fn.Value(ast.InputGateNand, runes.TagNoCase("nand")),
						fn.Value(ast.InputGateOr, runes.TagNoCase("or")),
						fn.Value(ast.InputGateNor, runes.TagNoCase("nor")),
						fn.Value(ast.InputGateXor, runes.TagNoCase("xor")),
						fn.Value(ast.InputGateXnor, runes.TagNoCase("xnor")),
					),
				),
			),
		),
	)(ctx, start)
}

/*
 * n_output_gatetype ::= buf | not
 */
func NOutputGateType(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.NOutputGateType, error) {
	res := &ast.NOutputGateType{}
	return top(
		token(res,
			word(
				bind(&res.Type,
					fn.Alt(
						fn.Value(ast.OutputGateBuf, runes.TagNoCase("buf")),
						fn.Value(ast.OutputGateNot, runes.TagNoCase("not")),
					),
				),
			),
		),
	)(ctx, start)
}

/*
 * pass_en_switchtype ::= tranif0 | tranif1 | rtranif1 | rtranif0
 */
func PassEnSwitchType(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.PassEnSwitchType, error) {
	res := &ast.PassEnSwitchType{}
	return top(
		token(res,
			word(
				bind(&res.Type,
					fn.Alt(
						fn.Value(ast.TRANIF0, runes.TagNoCase("tranif0")),
						fn.Value(ast.TRANIF1, runes.TagNoCase("tranif1")),
						fn.Value(ast.RTRANIF0, runes.TagNoCase("rtranif0")),
						fn.Value(ast.RTRANIF1, runes.TagNoCase("rtranif1")),
					),
				),
			),
		),
	)(ctx, start)
}

/*
 * pass_switchtype ::= tran | rtran
 */
func PassSwitchType(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.PassSwitchType, error) {
	res := &ast.PassSwitchType{}
	return top(
		token(res,
			word(
				bind(&res.Type,
					fn.Alt(
						fn.Value(ast.TRAN, runes.TagNoCase("tran")),
						fn.Value(ast.RTRAN, runes.TagNoCase("rtran")),
					),
				),
			),
		),
	)(ctx, start)
}
