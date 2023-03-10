package ast

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jtdubs/go-nom"
)

type Token struct {
	nom.Span[rune]
}

func (t *Token) GetSpan() nom.Span[rune] {
	return t.Span
}

func (t *Token) SetSpan(s nom.Span[rune]) {
	t.Span = s
}

type HasSpan interface {
	GetSpan() nom.Span[rune]
	SetSpan(nom.Span[rune])
}

type Bakeable interface {
	Bake() error
}

func assertNotKeyword(s string) error {
	if lower := strings.ToLower(s); keywords[lower] {
		return fmt.Errorf("reserved keyword '%q'", lower)
	}
	return nil
}

var keywords = map[string]bool{
	"accept_on":           true,
	"alias":               true,
	"always":              true,
	"always_comb":         true,
	"always_ff":           true,
	"always_latch":        true,
	"and":                 true,
	"assert":              true,
	"assign":              true,
	"assume":              true,
	"automatic":           true,
	"before":              true,
	"begin":               true,
	"bind":                true,
	"bins":                true,
	"binsof":              true,
	"bit":                 true,
	"break":               true,
	"buf":                 true,
	"bufif0":              true,
	"bufif1":              true,
	"byte":                true,
	"case":                true,
	"casex":               true,
	"casez":               true,
	"cell":                true,
	"chandle":             true,
	"checker":             true,
	"class":               true,
	"clocking":            true,
	"cmos":                true,
	"config":              true,
	"const":               true,
	"constraint":          true,
	"context":             true,
	"continue":            true,
	"cover":               true,
	"covergroup":          true,
	"coverpoint":          true,
	"cross":               true,
	"deassign":            true,
	"default":             true,
	"defparam":            true,
	"design":              true,
	"disable":             true,
	"dist":                true,
	"do":                  true,
	"edge":                true,
	"else":                true,
	"end":                 true,
	"endcase":             true,
	"endchecker":          true,
	"endclass":            true,
	"endclocking":         true,
	"endconfig":           true,
	"endfunction":         true,
	"endgenerate":         true,
	"endgroup":            true,
	"endinterface":        true,
	"endmodule":           true,
	"endpackage":          true,
	"endprimitive":        true,
	"endprogram":          true,
	"endproperty":         true,
	"endspecify":          true,
	"endsequence":         true,
	"endtable":            true,
	"endtask":             true,
	"enum":                true,
	"event":               true,
	"eventually":          true,
	"expect":              true,
	"export":              true,
	"extends":             true,
	"extern":              true,
	"final":               true,
	"first_match":         true,
	"for":                 true,
	"force":               true,
	"foreach":             true,
	"forever":             true,
	"fork":                true,
	"forkjoin":            true,
	"function":            true,
	"generate":            true,
	"genvar":              true,
	"global":              true,
	"highz0":              true,
	"highz1":              true,
	"if":                  true,
	"iff":                 true,
	"ifnone":              true,
	"ignore_bins":         true,
	"illegal_bins":        true,
	"implements":          true,
	"implies":             true,
	"import":              true,
	"incdir":              true,
	"include":             true,
	"initial":             true,
	"inout":               true,
	"input":               true,
	"inside":              true,
	"instance":            true,
	"int":                 true,
	"integer":             true,
	"interconnect":        true,
	"interface":           true,
	"intersect":           true,
	"join":                true,
	"join_any":            true,
	"join_none":           true,
	"large":               true,
	"let":                 true,
	"liblist":             true,
	"library":             true,
	"local":               true,
	"localparam":          true,
	"logic":               true,
	"longint":             true,
	"macromodule":         true,
	"matches":             true,
	"medium":              true,
	"modport":             true,
	"module":              true,
	"nand":                true,
	"negedge":             true,
	"nettype":             true,
	"new":                 true,
	"nexttime":            true,
	"nmos":                true,
	"nor":                 true,
	"noshowcancelled":     true,
	"not":                 true,
	"notif0":              true,
	"notif1":              true,
	"null":                true,
	"or":                  true,
	"output":              true,
	"package":             true,
	"packed":              true,
	"parameter":           true,
	"pmos":                true,
	"posedge":             true,
	"primitive":           true,
	"priority":            true,
	"program":             true,
	"property":            true,
	"protected":           true,
	"pull0":               true,
	"pull1":               true,
	"pulldown":            true,
	"pullup":              true,
	"pulsestyle_ondetect": true,
	"pulsestyle_onevent":  true,
	"pure":                true,
	"rand":                true,
	"randc":               true,
	"randcase":            true,
	"randsequence":        true,
	"rcmos":               true,
	"real":                true,
	"realtime":            true,
	"ref":                 true,
	"reg":                 true,
	"reject_on":           true,
	"release":             true,
	"repeat":              true,
	"restrict":            true,
	"return":              true,
	"rnmos":               true,
	"rpmos":               true,
	"rtran":               true,
	"rtranif0":            true,
	"rtranif1":            true,
	"s_always":            true,
	"s_eventually":        true,
	"s_nexttime":          true,
	"s_until":             true,
	"s_until_with":        true,
	"scalared":            true,
	"sequence":            true,
	"shortint":            true,
	"shortreal":           true,
	"showcancelled":       true,
	"signed":              true,
	"small":               true,
	"soft":                true,
	"solve":               true,
	"specify":             true,
	"specparam":           true,
	"static":              true,
	"string":              true,
	"strong":              true,
	"strong0":             true,
	"strong1":             true,
	"struct":              true,
	"super":               true,
	"supply0":             true,
	"supply1":             true,
	"sync_accept_on":      true,
	"sync_reject_on":      true,
	"table":               true,
	"tagged":              true,
	"task":                true,
	"this":                true,
	"throughout":          true,
	"time":                true,
	"timeprecision":       true,
	"timeunit":            true,
	"tran":                true,
	"tranif0":             true,
	"tranif1":             true,
	"tri":                 true,
	"tri0":                true,
	"tri1":                true,
	"triand":              true,
	"trior":               true,
	"trireg":              true,
	"type":                true,
	"typedef":             true,
	"union":               true,
	"unique":              true,
	"unique0":             true,
	"unsigned":            true,
	"until":               true,
	"until_with":          true,
	"untyped":             true,
	"use":                 true,
	"uwire":               true,
	"var":                 true,
	"vectored":            true,
	"virtual":             true,
	"void":                true,
	"wait":                true,
	"wait_order":          true,
	"wand":                true,
	"weak":                true,
	"weak0":               true,
	"weak1":               true,
	"while":               true,
	"wildcard":            true,
	"wire":                true,
	"with":                true,
	"within":              true,
	"wor":                 true,
	"xnor":                true,
	"xor":                 true,
}

func parseUint(t nom.Span[rune], base, size int) (uint64, error) {
	s := strings.ReplaceAll(string(t.Value()), "_", "")
	return strconv.ParseUint(s, base, size)
}

func parseFloat(t nom.Span[rune]) (float64, error) {
	s := strings.ReplaceAll(string(t.Value()), "_", "")
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
