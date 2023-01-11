package ast

import (
	"fmt"

	"github.com/jtdubs/go-nom"
)

type IntegerType interface {
	isIntegerType()
}

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
	nom.Span[rune]
	Type IntegerAtomTypeOption
}

func (u *IntegerAtomType) String() string {
	return fmt.Sprintf("IntegerAtomType(%v)", integerAtomTypeNames[u.Type])
}

func (*IntegerAtomType) isIntegerType() {}

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
	nom.Span[rune]
	Type IntegerVectorTypeOption
}

func (u *IntegerVectorType) String() string {
	return fmt.Sprintf("IntegerVectorType(%v)", integerVectorTypeNames[u.Type])
}

func (*IntegerVectorType) isIntegerType() {}

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
	nom.Span[rune]
	Type NonIntegerTypeOption
}

func (u *NonIntegerType) String() string {
	return fmt.Sprintf("NonIntegerType(%v)", nonIntegerTypeNames[u.Type])
}

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
	nom.Span[rune]
	Type NetTypeOption
}

func (u *NetType) String() string {
	return fmt.Sprintf("NetType(%v)", netTypeNames[u.Type])
}

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
	nom.Span[rune]
	Type SigningOption
}

func (u *Signing) String() string {
	return fmt.Sprintf("Signing(%v)", signingNames[u.Type])
}

type DriveStrengthOption interface {
	isDriveStrengthOption()
}

type DriveStrength struct {
	nom.Span[rune]
	AT, BT nom.Span[rune]
	A, B   DriveStrengthOption
}

func (u *DriveStrength) String() string {
	return fmt.Sprintf("DriveStrength(%v, %v)", u.A, u.B)
}

type HighZ0 struct {
	nom.Span[rune]
}

func (*HighZ0) String() string {
	return "HighZ0"
}

func (*HighZ0) isDriveStrengthOption() {}

type HighZ1 struct {
	nom.Span[rune]
}

func (*HighZ1) String() string {
	return "HighZ1"
}

func (*HighZ1) isDriveStrengthOption() {}

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
	nom.Span[rune]
	Type Strength0Option
}

func (u *Strength0) String() string {
	return fmt.Sprintf("Strength0(%v)", strength0OptionNames[u.Type])
}

func (*Strength0) isDriveStrengthOption() {}

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
	nom.Span[rune]
	Type Strength1Option
}

func (u *Strength1) String() string {
	return fmt.Sprintf("Strength1(%v)", strength1OptionNames[u.Type])
}

func (*Strength1) isDriveStrengthOption() {}

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
	nom.Span[rune]
	TypeT nom.Span[rune]
	Type  ChargeStrengthOption
}

func (u *ChargeStrength) String() string {
	return fmt.Sprintf("ChargeStrength(%v)", chargeStrengthOptionNames[u.Type])
}
