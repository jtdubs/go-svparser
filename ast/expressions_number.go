package ast

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jtdubs/go-nom"
)

type Number interface {
	isNumber()
}

type IntegralNumber interface {
	Number
	isIntegralNumber()
}

type DecimalNumber interface {
	IntegralNumber
	isDecimalNumber()
}

type DecimalNumberUnsigned struct {
	Token
	SizeT, BaseT, ValueT nom.Span[rune]
	Size                 uint
	Value                uint64
}

func (d *DecimalNumberUnsigned) String() string {
	return fmt.Sprintf("DecimalNumberUnsigned(%v, %v)", d.Size, d.Value)
}

func (d *DecimalNumberUnsigned) Bake() error {
	size, err := parseUint(d.SizeT, 10, 32)
	if err != nil {
		return err
	}
	d.Size = uint(size)
	d.Value, err = parseUint(d.ValueT, 10, int(d.Size))
	return err
}

func (*DecimalNumberUnsigned) isDecimalNumber()  {}
func (*DecimalNumberUnsigned) isIntegralNumber() {}
func (*DecimalNumberUnsigned) isNumber()         {}

type DecimalNumberX struct {
	Token
	SizeT, BaseT, X nom.Span[rune]
	Size            uint
}

func (d *DecimalNumberX) String() string {
	return fmt.Sprintf("DecimalNumberX(%v, %v)", d.Size, d.X)
}

func (d *DecimalNumberX) Bake() error {
	size, err := parseUint(d.SizeT, 10, 32)
	if err != nil {
		return err
	}
	d.Size = uint(size)
	return nil
}

func (*DecimalNumberX) isDecimalNumber()  {}
func (*DecimalNumberX) isIntegralNumber() {}
func (*DecimalNumberX) isNumber()         {}

type DecimalNumberZ struct {
	Token
	SizeT, BaseT, Z nom.Span[rune]
	Size            uint
}

func (d *DecimalNumberZ) String() string {
	return fmt.Sprintf("DecimalNumberZ(%v, %v)", d.SizeT, d.Z)
}

func (d *DecimalNumberZ) Bake() error {
	size, err := parseUint(d.SizeT, 10, 32)
	if err != nil {
		return err
	}
	d.Size = uint(size)
	return nil
}

func (*DecimalNumberZ) isDecimalNumber()  {}
func (*DecimalNumberZ) isIntegralNumber() {}
func (*DecimalNumberZ) isNumber()         {}

type BinaryNumber struct {
	Token
	SizeT, BaseT, ValueT nom.Span[rune]
	Value                MaskedInt
}

func (d *BinaryNumber) String() string {
	return fmt.Sprintf("BinaryNumber(%v)", d.Value)
}

func (d *BinaryNumber) Bake() error {
	size, err := parseUint(d.SizeT, 10, 32)
	if err != nil {
		return err
	}
	d.Value, err = NewMaskedInt(d.ValueT, 2, int(size))
	return err
}

func (*BinaryNumber) isIntegralNumber() {}
func (*BinaryNumber) isNumber()         {}

type OctalNumber struct {
	Token
	SizeT, BaseT, ValueT nom.Span[rune]
	Value                MaskedInt
}

func (d *OctalNumber) String() string {
	return fmt.Sprintf("OctalNumber(%v)", d.Value)
}

func (d *OctalNumber) Bake() error {
	size, err := parseUint(d.SizeT, 10, 32)
	if err != nil {
		return err
	}
	d.Value, err = NewMaskedInt(d.ValueT, 8, int(size))
	return err
}

func (*OctalNumber) isIntegralNumber() {}
func (*OctalNumber) isNumber()         {}

type HexNumber struct {
	Token
	SizeT, BaseT, ValueT nom.Span[rune]
	Value                MaskedInt
}

func (d *HexNumber) String() string {
	return fmt.Sprintf("HexNumber(%v)", d.Value)
}

func (d *HexNumber) Bake() error {
	size, err := parseUint(d.SizeT, 10, 32)
	if err != nil {
		return err
	}
	d.Value, err = NewMaskedInt(d.ValueT, 16, int(size))
	return err
}

func (*HexNumber) isIntegralNumber() {}
func (*HexNumber) isNumber()         {}

type RealNumber interface {
	isReal()
}

type FloatingPointNumber struct {
	Token
	Value float64
}

func (d *FloatingPointNumber) String() string {
	return fmt.Sprintf("FloatingPointNumber(%v)", d.Value)
}

func (d *FloatingPointNumber) Bake() error {
	val, err := parseFloat(d.Token)
	if err != nil {
		return err
	}
	d.Value = val
	return nil
}

func (*FloatingPointNumber) isReal()   {}
func (*FloatingPointNumber) isNumber() {}

type FixedPointNumber struct {
	Token
	Value float64
}

func (d *FixedPointNumber) String() string {
	return fmt.Sprintf("FixedPointNumber(%v)", d.Value)
}

func (d *FixedPointNumber) Bake() error {
	val, err := parseFloat(d.Token)
	if err != nil {
		return err
	}
	d.Value = val
	return nil
}

func (*FixedPointNumber) isReal()   {}
func (*FixedPointNumber) isNumber() {}

type UnsignedNumber struct {
	Token
	Value uint64
}

func (d *UnsignedNumber) String() string {
	return fmt.Sprintf("UnsignedNumber(%v)", d.Value)
}

func (d *UnsignedNumber) Bake() error {
	val, err := parseUint(d.Span, 10, 64)
	if err != nil {
		return err
	}
	d.Value = val
	return nil
}

func (*UnsignedNumber) isDecimalNumber()  {}
func (*UnsignedNumber) isIntegralNumber() {}
func (*UnsignedNumber) isNumber()         {}

type UnbasedUnsizedLiteral struct {
	Token
}

func (d *UnbasedUnsizedLiteral) String() string {
	return fmt.Sprintf("UnbasedUnsizedLiteral(%v)", d.Token)
}

func parseUint(t nom.Span[rune], base, size int) (uint64, error) {
	s := strings.ReplaceAll(string(t.Value()), "_", "")
	return strconv.ParseUint(s, base, size)
}

func parseFloat(t Token) (float64, error) {
	s := strings.ReplaceAll(t.Value(), "_", "")
	return strconv.ParseFloat(s, 64)
}

type MaskedInt struct {
	Base, Size uint
	V, X, Z    uint64
}

func NewMaskedInt(t nom.Span[rune], base, size int) (result MaskedInt, err error) {
	s := strings.ReplaceAll(string(t.Value()), "_", "")

	var max rune
	switch base {
	case 2:
		max = '1'
	case 8:
		max = '7'
	case 16:
		max = 'F'
	}

	var vs, xs, zs strings.Builder
	for _, r := range s {
		switch r {
		case 'x', 'X':
			vs.WriteRune('0')
			xs.WriteRune(max)
			zs.WriteRune('0')
		case 'z', 'Z', '?':
			vs.WriteRune('0')
			xs.WriteRune('0')
			zs.WriteRune(max)
		default:
			vs.WriteRune(r)
			xs.WriteRune('0')
			zs.WriteRune('0')
		}
	}

	result.Base = uint(base)
	result.Size = uint(size)
	if result.V, err = strconv.ParseUint(vs.String(), base, size); err != nil {
		return
	}
	if result.X, err = strconv.ParseUint(xs.String(), base, size); err != nil {
		return
	}
	if result.Z, err = strconv.ParseUint(zs.String(), base, size); err != nil {
		return
	}
	return
}

func (m MaskedInt) String() string {
	var (
		width    uint
		mask     uint64
		baseChar rune
	)
	switch m.Base {
	case 2:
		width, mask, baseChar = 1, 1, 'b'
	case 8:
		width, mask, baseChar = 3, 7, 'o'
	case 16:
		width, mask, baseChar = 4, 15, 'h'
	default:
		return "unbaked"
	}

	chars := m.Size / width
	if m.Size%width == 0 {
		chars = chars - 1
	}

	var s strings.Builder
	s.WriteString(fmt.Sprintf("%v'%v", m.Size, string(baseChar)))
	for shift := int(chars * width); shift >= 0; shift = shift - int(width) {
		v, x, z := (m.V>>shift)&mask, (m.X>>shift)&mask, (m.Z>>shift)&mask
		if x != 0 {
			s.WriteRune('X')
		} else if z != 0 {
			s.WriteRune('Z')
		} else {
			s.WriteRune([]rune("0123456789abcdef")[v])
		}
	}
	return s.String()
}
