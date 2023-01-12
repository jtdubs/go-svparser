package ast

import (
	"fmt"
)

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
