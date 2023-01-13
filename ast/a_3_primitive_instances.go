package ast

import "fmt"

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
type PulldownStrength struct {
	Token
	Zero *Strength0
	One  *Strength1
}

func (u *PulldownStrength) String() string {
	return fmt.Sprintf("PulldownStrength(%v, %v)", u.Zero, u.One)
}

/*
 * pullup_strength ::=
 *   ( strength0 , strength1 )
 *   | ( strength1 , strength0 )
 *   | ( strength1 )
 */
type PullupStrength struct {
	Token
	Zero *Strength0
	One  *Strength1
}

func (u *PullupStrength) String() string {
	return fmt.Sprintf("PullupStrength(%v, %v)", u.Zero, u.One)
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
type CMOSSwitchTypeOption int

const (
	CMOS CMOSSwitchTypeOption = iota
	RCMOS
)

var cmosSwitchTypeNames = map[CMOSSwitchTypeOption]string{
	CMOS:  "CMOS",
	RCMOS: "RCMOS",
}

type CMOSSwitchType struct {
	Token
	Type CMOSSwitchTypeOption
}

func (u *CMOSSwitchType) String() string {
	return fmt.Sprintf("CMOSSwitchType(%v)", cmosSwitchTypeNames[u.Type])
}

/*
 * enable_gatetype ::= bufif0 | bufif1 | notif0 | notif1
 */
type EnableGateTypeOption int

const (
	BUFIF0 EnableGateTypeOption = iota
	BUFIF1
	NOTIF0
	NOTIF1
)

var enableGateTypeNames = map[EnableGateTypeOption]string{
	BUFIF0: "BUFIF0",
	BUFIF1: "BUFIF1",
	NOTIF0: "NOTIF0",
	NOTIF1: "NOTIF1",
}

type EnableGateType struct {
	Token
	Type EnableGateTypeOption
}

func (u *EnableGateType) String() string {
	return fmt.Sprintf("EnableGateType(%v)", enableGateTypeNames[u.Type])
}

/*
 * mos_switchtype ::= nmos | pmos | rnmos | rpmos
 */
type MOSSwitchTypeOption int

const (
	NMOS MOSSwitchTypeOption = iota
	PMOS
	RNMOS
	RPMOS
)

var mosSwitchTypeNames = map[MOSSwitchTypeOption]string{
	NMOS:  "NMOS",
	PMOS:  "PMOS",
	RNMOS: "RNMOS",
	RPMOS: "RPMOS",
}

type MOSSwitchType struct {
	Token
	Type MOSSwitchTypeOption
}

func (u *MOSSwitchType) String() string {
	return fmt.Sprintf("MOSSwitchType(%v)", mosSwitchTypeNames[u.Type])
}

/*
 * n_input_gatetype ::= and | nand | or | nor | xor | xnor
 */
type NInputGateTypeOption int

const (
	InputGateAnd NInputGateTypeOption = iota
	InputGateNand
	InputGateOr
	InputGateNor
	InputGateXor
	InputGateXnor
)

var nInputGateTypeNames = map[NInputGateTypeOption]string{
	InputGateAnd:  "And",
	InputGateNand: "Nand",
	InputGateOr:   "Or",
	InputGateNor:  "Nor",
	InputGateXor:  "Xor",
	InputGateXnor: "Xnor",
}

type NInputGateType struct {
	Token
	Type NInputGateTypeOption
}

func (u *NInputGateType) String() string {
	return fmt.Sprintf("NInputGateType(%v)", nInputGateTypeNames[u.Type])
}

/*
 * n_output_gatetype ::= buf | not
 */
type NOutputGateTypeOption int

const (
	OutputGateBuf NOutputGateTypeOption = iota
	OutputGateNot
)

var nOutputGateTypeNames = map[NOutputGateTypeOption]string{
	OutputGateBuf: "Buf",
	OutputGateNot: "Not",
}

type NOutputGateType struct {
	Token
	Type NOutputGateTypeOption
}

func (u *NOutputGateType) String() string {
	return fmt.Sprintf("NOutputGateType(%v)", nOutputGateTypeNames[u.Type])
}

/*
 * pass_en_switchtype ::= tranif0 | tranif1 | rtranif1 | rtranif0
 */
type PassEnSwitchTypeOption int

const (
	TRANIF0 PassEnSwitchTypeOption = iota
	TRANIF1
	RTRANIF0
	RTRANIF1
)

var passEnSwitchTypeNames = map[PassEnSwitchTypeOption]string{
	TRANIF0:  "TRANIF0",
	TRANIF1:  "TRANIF1",
	RTRANIF0: "RTRANIF0",
	RTRANIF1: "RTRANIF1",
}

type PassEnSwitchType struct {
	Token
	Type PassEnSwitchTypeOption
}

func (u *PassEnSwitchType) String() string {
	return fmt.Sprintf("PassEnSwitchType(%v)", passEnSwitchTypeNames[u.Type])
}

/*
 * pass_switchtype ::= tran | rtran
 */
type PassSwitchTypeOption int

const (
	TRAN PassSwitchTypeOption = iota
	RTRAN
)

var passSwitchTypeNames = map[PassSwitchTypeOption]string{
	TRAN:  "TRAN",
	RTRAN: "RTRAN",
}

type PassSwitchType struct {
	Token
	Type PassSwitchTypeOption
}

func (u *PassSwitchType) String() string {
	return fmt.Sprintf("PassSwitchType(%v)", passSwitchTypeNames[u.Type])
}
