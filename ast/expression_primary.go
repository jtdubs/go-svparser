package ast

import (
	"fmt"
)

type ConstantPrimary interface {
	ConstantExpression
	isConstantPrimary()
}

type Primary interface {
	isPrimary()
}

type PrimaryLiteral interface {
	Primary
	ConstantPrimary
	isPrimaryLiteral()
}

type TimeLiteral struct {
	Token
	Number Number
	Unit   *TimeUnit
}

func (t *TimeLiteral) String() string {
	return fmt.Sprintf("TimeLiteral(%v, %v)", t.Number, t.Unit)
}

func (*TimeLiteral) isPrimaryLiteral()     {}
func (*TimeLiteral) isPrimary()            {}
func (*TimeLiteral) isConstantPrimary()    {}
func (*TimeLiteral) isConstantExpression() {}

type TimeUnitOption int

const (
	S TimeUnitOption = iota
	MS
	US
	NS
	PS
	FS
)

var timeUnitNames = map[TimeUnitOption]string{
	S:  "S",
	MS: "MS",
	US: "US",
	NS: "NS",
	PS: "PS",
	FS: "FS",
}

type TimeUnit struct {
	Token
	Op TimeUnitOption
}

func (t *TimeUnit) String() string {
	return fmt.Sprintf("TimeUnit(%v)", timeUnitNames[t.Op])
}
