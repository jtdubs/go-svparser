package ast

import (
	"fmt"

	"github.com/jtdubs/go-nom"
)

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
	Supply0 Strength0Option = iota
	Strong0
	Pull0
	Weak0
)

var strength0OptionNames = map[Strength0Option]string{
	Supply0: "Supply0",
	Strong0: "Strong0",
	Pull0:   "Pull0",
	Weak0:   "Weak0",
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
	Supply1 Strength1Option = iota
	Strong1
	Pull1
	Weak1
)

var strength1OptionNames = map[Strength1Option]string{
	Supply1: "Supply1",
	Strong1: "Strong1",
	Pull1:   "Pull1",
	Weak1:   "Weak1",
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
	Small ChargeStrengthOption = iota
	Medium
	Large
)

var chargeStrengthOptionNames = map[ChargeStrengthOption]string{
	Small:  "Small",
	Medium: "Medium",
	Large:  "Large",
}

type ChargeStrength struct {
	nom.Span[rune]
	TypeT nom.Span[rune]
	Type  ChargeStrengthOption
}

func (u *ChargeStrength) String() string {
	return fmt.Sprintf("ChargeStrength(%v)", chargeStrengthOptionNames[u.Type])
}
