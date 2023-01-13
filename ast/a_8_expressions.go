package ast

import (
	"fmt"
	"strconv"

	"github.com/jtdubs/go-nom"
)

//
// A.8 Expressions
//

//
// A.8.1 Concatenations
//

/*
 * concatenation ::=
 *   { expression { , expression } }
 */

/*
 * constant_concatenation ::=
 *   { constant_expression { , constant_expression } }
 */

/*
 * constant_multiple_concatenation ::= { constant_expression constant_concatenation }
 */

/*
 * module_path_concatenation ::= { module_path_expression { , module_path_expression } }
 */

/*
 * module_path_multiple_concatenation ::= { constant_expression module_path_concatenation }
 */

/*
 * multiple_concatenation ::= { expression concatenation }34
 */

/*
 * streaming_concatenation ::= { stream_operator [ slice_size ] stream_concatenation }
 */

/*
 * stream_operator ::= >> | <<
 */

/*
 * slice_size ::= simple_type | constant_expression
 */

/*
 * stream_concatenation ::= { stream_expression { , stream_expression } }
 */

/*
 * stream_expression ::= expression [ with [ array_range_expression ] ]
 */

/*
 * array_range_expression ::=
 *   expression
 *   | expression : expression
 *   | expression +: expression
 *   | expression -: expression
 */

/*
 * empty_unpacked_array_concatenation35 ::= { }
 */

//
// A.8.2 Subroutine calls
//

/*
 * constant_function_call ::= function_subroutine_call36
 */

/*
 * tf_call37 ::= ps_or_hierarchical_tf_identifier { attribute_instance } [ ( list_of_arguments ) ]
 */

/*
 * system_tf_call ::=
 *   system_tf_identifier [ ( list_of_arguments ) ]
 *   | system_tf_identifier ( data_type [ , expression ] )
 * | system_tf_identifier ( expression { , [ expression ] } [ , [ clocking_event ] ] )
 */

/*
 * subroutine_call ::=
 *   tf_call
 *   | system_tf_call
 *   | method_call
 *   | [ std :: ] randomize_call
 */

/*
 * function_subroutine_call ::= subroutine_call
 */

/*
 * list_of_arguments ::=
 *   [ expression ] { , [ expression ] } { , . identifier ( [ expression ] ) }
 *   | . identifier ( [ expression ] ) { , . identifier ( [ expression ] ) }
 */

/*
 * method_call ::= method_call_root . method_call_body
 */

/*
 * method_call_body ::=
 *   method_identifier { attribute_instance } [ ( list_of_arguments ) ]
 *   | built_in_method_call
 */

/*
 * built_in_method_call ::=
 *   array_manipulation_call
 *   | randomize_call
 */

/*
 * array_manipulation_call ::=
 *   array_method_name { attribute_instance }
 *   [ ( list_of_arguments ) ]
 *   [ with ( expression ) ]
 */

/*
 * randomize_call ::=
 *   randomize { attribute_instance }
 *   [ ( [ variable_identifier_list | null ] ) ]
 *   [ with [ ( [ identifier_list ] ) ] constraint_block ]38
 */

/*
 * method_call_root ::= primary | implicit_class_handle
 */

/*
 * array_method_name ::=
 *   method_identifier | unique | and | or | xor
 */

//
// A.8.3 Expressions
//

/*
 * inc_or_dec_expression ::=
 *   inc_or_dec_operator { attribute_instance } variable_lvalue
 *   | variable_lvalue { attribute_instance } inc_or_dec_operator
 */

/*
 * conditional_expression ::= cond_predicate ? { attribute_instance } expression : expression
 */

/*
 * constant_expression ::=
 *   constant_primary
 *   | unary_operator { attribute_instance } constant_primary
 *   | constant_expression binary_operator { attribute_instance } constant_expression
 *   | constant_expression ? { attribute_instance } constant_expression : constant_expression
 */
type ConstantExpression interface {
	isConstantExpression()
}

/*
 * constant_expression ::=
 *   ...
 *   | unary_operator { attribute_instance } constant_primary
 *   ...
 */
type ConstantUnaryExpression struct {
	Token
	Op      *UnaryOperator
	Attrs   []*AttributeInstance
	Primary ConstantPrimary
}

func (e *ConstantUnaryExpression) String() string {
	return fmt.Sprintf("ConstantUnaryExpression(%v, %v)", e.Op, e.Primary)
}

func (*ConstantUnaryExpression) isConstantExpression() {}

/*
 * constant_expression ::=
 *   ...
 *   | constant_expression binary_operator { attribute_instance } constant_expression
 *   ...
 */
type ConstantBinaryExpression struct {
	Token
	Op          *BinaryOperator
	Attrs       []*AttributeInstance
	Left, Right ConstantExpression
}

func (e *ConstantBinaryExpression) String() string {
	return fmt.Sprintf("ConstantBinaryExpression(%v, %v, %v)", e.Left, e.Op, e.Right)
}

func (*ConstantBinaryExpression) isConstantExpression() {}

/*
 * constant_expression ::=
 *   ...
 *   | constant_expression ? { attribute_instance } constant_expression : constant_expression
 *   ...
 */
type ConstantTernaryExpression struct {
	Token
	Attrs          []*AttributeInstance
	Cond, If, Else ConstantExpression
}

func (e *ConstantTernaryExpression) String() string {
	return fmt.Sprintf("ConstantTernaryExpression(%v, %v, %v)", e.Cond, e.If, e.Else)
}

func (*ConstantTernaryExpression) isConstantExpression() {}

/*
 * constant_mintypmax_expression ::=
 *   constant_expression
 *   | constant_expression : constant_expression : constant_expression
 */

/*
 * constant_param_expression ::=
 *   constant_mintypmax_expression | data_type | $
 */

/*
 * param_expression ::= mintypmax_expression | data_type | $
 */

/*
 * constant_range_expression ::=
 *   constant_expression
 *   | constant_part_select_range
 */

/*
 * constant_part_select_range ::=
 *   constant_range
 *   | constant_indexed_range
 */

/*
 * constant_range ::= constant_expression : constant_expression
 */

/*
 * constant_indexed_range ::=
 *   constant_expression +: constant_expression
 *   | constant_expression -: constant_expression
 */

/*
 * expression ::=
 *   primary
 *   | unary_operator { attribute_instance } primary
 *   | inc_or_dec_expression
 *   | ( operator_assignment )
 *   | expression binary_operator { attribute_instance } expression
 *   | conditional_expression
 *   | inside_expression
 *   | tagged_union_expression
 */

/*
 * tagged_union_expression ::=
 *   tagged member_identifier [ expression ]
 */

/*
 * inside_expression ::= expression inside { open_range_list }
 */

/*
 * value_range ::=
 *   expression
 *   | [ expression : expression ]
 */

/*
 * mintypmax_expression ::=
 *   expression
 *   | expression : expression : expression
 */

/*
 * module_path_conditional_expression ::= module_path_expression ? { attribute_instance }
 *   module_path_expression : module_path_expression
 */

/*
 * module_path_expression ::=
 *   module_path_primary
 *   | unary_module_path_operator { attribute_instance } module_path_primary
 *   | module_path_expression binary_module_path_operator { attribute_instance }
 *   module_path_expression
 *   | module_path_conditional_expression
 */

/*
 * module_path_mintypmax_expression ::=
 *   module_path_expression
 *   | module_path_expression : module_path_expression : module_path_expression
 */

/*
 * part_select_range ::= constant_range | indexed_range
 */

/*
 * indexed_range ::=
 *   expression +: constant_expression
 *   | expression -: constant_expression
 */

/*
 * genvar_expression ::= constant_expression
 */

//
// A.8.4 Primaries
//

/*
 * constant_primary ::=
 *   primary_literal
 *   | ps_parameter_identifier constant_select
 *   | specparam_identifier [ [ constant_range_expression ] ]
 *   | genvar_identifier39
 *   | formal_port_identifier constant_select
 *   | [ package_scope | class_scope ] enum_identifier
 *   | constant_concatenation [ [ constant_range_expression ] ]
 *   | constant_multiple_concatenation [ [ constant_range_expression ] ]
 *   | constant_function_call
 *   | constant_let_expression
 *   | ( constant_mintypmax_expression )
 *   | constant_cast
 *   | constant_assignment_pattern_expression
 *   | type_reference40
 *   | null
 */
type ConstantPrimary interface {
	ConstantExpression
	isConstantPrimary()
}

/*
 * module_path_primary ::=
 *   number
 *   | identifier
 *   | module_path_concatenation
 *   | module_path_multiple_concatenation
 *   | function_subroutine_call
 *   | ( module_path_mintypmax_expression )
 */

/*
 * primary ::=
 *   primary_literal
 *   | [ class_qualifier | package_scope ] hierarchical_identifier select
 *   | empty_unpacked_array_concatenation
 *   | concatenation [ [ range_expression ] ]
 *   | multiple_concatenation [ [ range_expression ] ]
 *   | function_subroutine_call
 *   | let_expression
 *   | ( mintypmax_expression )
 *   | cast
 *   | assignment_pattern_expression
 *   | streaming_concatenation
 *   | sequence_method_call
 *   | this41
 *   | $42
 *   | null
 */
type Primary interface {
	isPrimary()
}

/*
 * class_qualifier ::= [ local ::43 ] [ implicit_class_handle . | class_scope ]
 */

/*
 * range_expression ::=
 *   expression
 *   | part_select_range
 */

/*
 * primary_literal ::= number | time_literal | unbased_unsized_literal | string_literal
 */
type PrimaryLiteral interface {
	Primary
	ConstantPrimary
	isPrimaryLiteral()
}

/*
 * time_literal44 ::=
 *   unsigned_number time_unit
 *   | fixed_point_number time_unit
 */
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

/*
 * time_unit ::= s | ms | us | ns | ps | fs
 */
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

/*
 * implicit_class_handle41 ::= this | super | this . super
 */

/*
 * bit_select ::= { [ expression ] }
 */

/*
 * select ::=
 *   [ { . member_identifier bit_select } . member_identifier ] bit_select [ [ part_select_range ] ]
 */

/*
 * nonrange_select ::=
 *   [ { . member_identifier bit_select } . member_identifier ] bit_select
 */

/*
 * constant_bit_select ::= { [ constant_expression ] }
 */
type ConstantBitSelect struct {
	Token
	Exprs []ConstantExpression
}

func (e *ConstantBitSelect) String() string {
	return fmt.Sprintf("ConstantBitSelect(%v)", e.Exprs)
}

/*
 * constant_select ::=
 *   [ { . member_identifier constant_bit_select } . member_identifier ] constant_bit_select
 *   [ [ constant_part_select_range ] ]
 */

/*
 * constant_cast ::=
 *   casting_type ' ( constant_expression )
 */

/*
 * constant_let_expression ::= let_expression45
 */

/*
 * cast ::=
 *   casting_type ' ( expression )
 */

//
// A.8.5 Expression left-side values
//

/*
 * net_lvalue ::=
 *   ps_or_hierarchical_net_identifier constant_select
 *   | { net_lvalue { , net_lvalue } }
 *   | [ assignment_pattern_expression_type ] assignment_pattern_net_lvalue
 */

/*
 * variable_lvalue ::=
 *   [ implicit_class_handle . | package_scope ] hierarchical_variable_identifier select46
 *   | { variable_lvalue { , variable_lvalue } }
 *   | [ assignment_pattern_expression_type ] assignment_pattern_variable_lvalue
 *   | streaming_concatenation47
 */

/*
 * nonrange_variable_lvalue ::=
 *   [ implicit_class_handle . | package_scope ] hierarchical_variable_identifier nonrange_select
 */

//
// A.8.6 Operators
//

/*
 * unary_operator ::=
 *   + | - | ! | ~ | & | ~& | | | ~| | ^ | ~^ | ^~
 */
type UnaryOperatorType int

const (
	UnaryPositive UnaryOperatorType = iota
	UnaryNegate
	UnaryLogicalNegation
	UnaryLogicalReductionNot
	UnaryLogicalReductionAnd
	UnaryLogicalReductionNand
	UnaryLogicalReductionOr
	UnaryLogicalReductionNor
	UnaryLogicalReductionXor
	UnaryLogicalReductionXnor
)

var unaryOperatorNames = map[UnaryOperatorType]string{
	UnaryPositive:             "Positive",
	UnaryNegate:               "Negate",
	UnaryLogicalNegation:      "LogicalNegation",
	UnaryLogicalReductionNot:  "LogicalReductionNot",
	UnaryLogicalReductionAnd:  "LogicalReductionAnd",
	UnaryLogicalReductionNand: "LogicalReductionNand",
	UnaryLogicalReductionOr:   "LogicalReductionOr",
	UnaryLogicalReductionNor:  "LogicalReductionNor",
	UnaryLogicalReductionXor:  "LogicalReductionXor",
	UnaryLogicalReductionXnor: "LogicalReductionXnor",
}

type UnaryOperator struct {
	Token
	Op UnaryOperatorType
}

func (u *UnaryOperator) String() string {
	return fmt.Sprintf("UnaryOperator(%v)", unaryOperatorNames[u.Op])
}

/*
 * binary_operator ::=
 *   + | - | * | / | % | == | != | === | !== | ==? | !=? | && | || | **
 *   | < | <= | > | >= | & | | | ^ | ^~ | ~^ | >> | << | >>> | <<<
 *   | -> | <->
 */
type BinaryOperatorType int

const (
	BinaryAdd BinaryOperatorType = iota
	BinarySubtract
	BinaryMultiply
	BinaryDivide
	BinaryModulus
	BinaryExp
	BinaryBitwiseAnd
	BinaryBitwiseOr
	BinaryBitwiseXor
	BinaryBitwiseXnor
	BinaryLogicalShiftLeft
	BinaryLogicalShiftRight
	BinaryArithmeticShiftLeft
	BinaryArithmeticShiftRight
	BinaryLogicalAnd
	BinaryLogicalOr
	BinaryLogicalImplies
	BinaryLogicalIff
	BinaryLessThan
	BinaryLessThanEqual
	BinaryGreaterThan
	BinaryGreaterThanEqual
	BinaryCaseEquals
	BinaryCaseNotEquals
	BinaryLogicalEquals
	BinaryLogicalNotEquals
	BinaryWildcardEquals
	BinaryWildcardNotEquals
)

var binaryOperatorNames = map[BinaryOperatorType]string{
	BinaryAdd:                  "Add",
	BinarySubtract:             "Subtract",
	BinaryMultiply:             "Multiply",
	BinaryDivide:               "Divide",
	BinaryModulus:              "Modulus",
	BinaryExp:                  "Exp",
	BinaryBitwiseAnd:           "BitwiseAnd",
	BinaryBitwiseOr:            "BitwiseOr",
	BinaryBitwiseXor:           "BitwiseXor",
	BinaryBitwiseXnor:          "BitwiseXnor",
	BinaryLogicalShiftLeft:     "LogicalShiftLeft",
	BinaryLogicalShiftRight:    "LogicalShiftRight",
	BinaryArithmeticShiftLeft:  "ArithmeticShiftLeft",
	BinaryArithmeticShiftRight: "ArithmeticShiftRight",
	BinaryLogicalAnd:           "LogicalAnd",
	BinaryLogicalOr:            "LogicalOr",
	BinaryLogicalImplies:       "LogicalImplies",
	BinaryLogicalIff:           "LogicalIff",
	BinaryLessThan:             "LessThan",
	BinaryLessThanEqual:        "LessThanEqual",
	BinaryGreaterThan:          "GreaterThan",
	BinaryGreaterThanEqual:     "GreaterThanEqual",
	BinaryCaseEquals:           "CaseEquals",
	BinaryCaseNotEquals:        "CaseNotEquals",
	BinaryLogicalEquals:        "LogicalEquals",
	BinaryLogicalNotEquals:     "LogicalNotEquals",
	BinaryWildcardEquals:       "WildcardEquals",
	BinaryWildcardNotEquals:    "WildcardNotEquals",
}

type BinaryOperator struct {
	Token
	Op BinaryOperatorType
}

func (b *BinaryOperator) String() string {
	return fmt.Sprintf("BinaryOperator(%v)", binaryOperatorNames[b.Op])
}

/*
 * inc_or_dec_operator ::= ++ | --
 */
type IncOrDecOperatorType int

const (
	Inc IncOrDecOperatorType = iota
	Dec
)

var incOrDecOperatorNames = map[IncOrDecOperatorType]string{
	Inc: "Inc",
	Dec: "Dec",
}

type IncOrDecOperator struct {
	Token
	Op IncOrDecOperatorType
}

func (b *IncOrDecOperator) String() string {
	return fmt.Sprintf("IncOrDecOperator(%v)", incOrDecOperatorNames[b.Op])
}

/*
 * unary_module_path_operator ::=
 *   ! | ~ | & | ~& | | | ~| | ^ | ~^ | ^~
 */
type UnaryModulePathOperator struct {
	Token
	Op UnaryOperatorType
}

func (b *UnaryModulePathOperator) String() string {
	return fmt.Sprintf("UnaryModulePathOperator(%v)", unaryOperatorNames[b.Op])
}

/*
 * binary_module_path_operator ::=
 *   == | != | && | || | & | | | ^ | ^~ | ~^
 */
type BinaryModulePathOperator struct {
	Token
	Op BinaryOperatorType
}

func (b *BinaryModulePathOperator) String() string {
	return fmt.Sprintf("BinaryModulePathOperator(%v)", binaryOperatorNames[b.Op])
}

//
// A.8.7 Numbers
//

/*
 * number ::=
 *   integral_number
 *   | real_number
 */
type Number interface {
	PrimaryLiteral
	isNumber()
}

/*
 * integral_number ::=
 *   decimal_number
 *   | octal_number
 *   | binary_number
 *   | hex_number
 */
type IntegralNumber interface {
	Number
	isIntegralNumber()
}

/*
 * decimal_number ::=
 *   unsigned_number
 *   | [ size ] decimal_base unsigned_number
 *   | [ size ] decimal_base x_digit { _ }
 *   | [ size ] decimal_base z_digit { _ }
 */
type DecimalNumber interface {
	IntegralNumber
	isDecimalNumber()
}

/*
 * decimal_number ::=
 *   ...
 *   | [ size ] decimal_base unsigned_number
 *   ...
 */
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

func (*DecimalNumberUnsigned) isDecimalNumber()      {}
func (*DecimalNumberUnsigned) isIntegralNumber()     {}
func (*DecimalNumberUnsigned) isNumber()             {}
func (*DecimalNumberUnsigned) isPrimaryLiteral()     {}
func (*DecimalNumberUnsigned) isPrimary()            {}
func (*DecimalNumberUnsigned) isConstantPrimary()    {}
func (*DecimalNumberUnsigned) isConstantExpression() {}

/*
 * decimal_number ::=
 *   ...
 *   | [ size ] decimal_base x_digit { _ }
 *   ...
 */
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

func (*DecimalNumberX) isDecimalNumber()      {}
func (*DecimalNumberX) isIntegralNumber()     {}
func (*DecimalNumberX) isNumber()             {}
func (*DecimalNumberX) isPrimaryLiteral()     {}
func (*DecimalNumberX) isPrimary()            {}
func (*DecimalNumberX) isConstantPrimary()    {}
func (*DecimalNumberX) isConstantExpression() {}

/*
 * decimal_number ::=
 *   ...
 *   | [ size ] decimal_base z_digit { _ }
 *   ...
 */
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

func (*DecimalNumberZ) isDecimalNumber()      {}
func (*DecimalNumberZ) isIntegralNumber()     {}
func (*DecimalNumberZ) isNumber()             {}
func (*DecimalNumberZ) isPrimaryLiteral()     {}
func (*DecimalNumberZ) isPrimary()            {}
func (*DecimalNumberZ) isConstantPrimary()    {}
func (*DecimalNumberZ) isConstantExpression() {}

/*
 * binary_number ::= [ size ] binary_base binary_value
 */
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

func (*BinaryNumber) isIntegralNumber()     {}
func (*BinaryNumber) isNumber()             {}
func (*BinaryNumber) isPrimaryLiteral()     {}
func (*BinaryNumber) isPrimary()            {}
func (*BinaryNumber) isConstantPrimary()    {}
func (*BinaryNumber) isConstantExpression() {}

/*
 * octal_number ::= [ size ] octal_base octal_value
 */
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

func (*OctalNumber) isIntegralNumber()     {}
func (*OctalNumber) isNumber()             {}
func (*OctalNumber) isPrimaryLiteral()     {}
func (*OctalNumber) isPrimary()            {}
func (*OctalNumber) isConstantPrimary()    {}
func (*OctalNumber) isConstantExpression() {}

/*
 * hex_number ::= [ size ] hex_base hex_value
 */
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

func (*HexNumber) isIntegralNumber()     {}
func (*HexNumber) isNumber()             {}
func (*HexNumber) isPrimaryLiteral()     {}
func (*HexNumber) isPrimary()            {}
func (*HexNumber) isConstantPrimary()    {}
func (*HexNumber) isConstantExpression() {}

/*
 * sign ::= + | -
 */

/*
 * size ::= non_zero_unsigned_number
 */

/*
 * non_zero_unsigned_number33 ::= non_zero_decimal_digit { _ | decimal_digit}
 */

/*
 * real_number33 ::=
 *   fixed_point_number
 *   | unsigned_number [ . unsigned_number ] exp [ sign ] unsigned_number
 */
type RealNumber interface {
	isReal()
}

/*
 * real_number33 ::=
 *   ...
 *   | unsigned_number [ . unsigned_number ] exp [ sign ] unsigned_number
 *   ...
 */
type FloatingPointNumber struct {
	Token
	Value float64
}

func (d *FloatingPointNumber) String() string {
	return fmt.Sprintf("FloatingPointNumber(%v)", d.Value)
}

func (d *FloatingPointNumber) Bake() error {
	val, err := parseFloat(d.Span)
	if err != nil {
		return err
	}
	d.Value = val
	return nil
}

func (*FloatingPointNumber) isReal()               {}
func (*FloatingPointNumber) isNumber()             {}
func (*FloatingPointNumber) isPrimaryLiteral()     {}
func (*FloatingPointNumber) isPrimary()            {}
func (*FloatingPointNumber) isConstantPrimary()    {}
func (*FloatingPointNumber) isConstantExpression() {}

/*
 * fixed_point_number33 ::= unsigned_number . unsigned_number
 */
type FixedPointNumber struct {
	Token
	Value float64
}

func (d *FixedPointNumber) String() string {
	return fmt.Sprintf("FixedPointNumber(%v)", d.Value)
}

func (d *FixedPointNumber) Bake() error {
	val, err := parseFloat(d.Span)
	if err != nil {
		return err
	}
	d.Value = val
	return nil
}

func (*FixedPointNumber) isReal()               {}
func (*FixedPointNumber) isNumber()             {}
func (*FixedPointNumber) isPrimaryLiteral()     {}
func (*FixedPointNumber) isPrimary()            {}
func (*FixedPointNumber) isConstantPrimary()    {}
func (*FixedPointNumber) isConstantExpression() {}

/*
 * exp ::= e | E
 */

/*
 * unsigned_number33 ::= decimal_digit { _ | decimal_digit }
 */
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

func (*UnsignedNumber) isDecimalNumber()      {}
func (*UnsignedNumber) isIntegralNumber()     {}
func (*UnsignedNumber) isNumber()             {}
func (*UnsignedNumber) isPrimaryLiteral()     {}
func (*UnsignedNumber) isPrimary()            {}
func (*UnsignedNumber) isConstantPrimary()    {}
func (*UnsignedNumber) isConstantExpression() {}

/*
 * binary_value33 ::= binary_digit { _ | binary_digit }
 */

/*
 * octal_value33 ::= octal_digit { _ | octal_digit }
 */

/*
 * hex_value33 ::= hex_digit { _ | hex_digit }
 */

/*
 * decimal_base33 ::= '[s|S]d | '[s|S]D
 */

/*
 * binary_base33 ::= '[s|S]b | '[s|S]B
 */

/*
 * octal_base33 ::= '[s|S]o | '[s|S]O
 */

/*
 * hex_base33 ::= '[s|S]h | '[s|S]H
 */

/*
 * non_zero_decimal_digit ::= 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
 */

/*
 * decimal_digit ::= 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
 */

/*
 * binary_digit ::= x_digit | z_digit | 0 | 1
 */

/*
 * octal_digit ::= x_digit | z_digit | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7
 */

/*
 * hex_digit ::= x_digit | z_digit | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | a | b | c | d | e | f | A | B | C | D | E | F
 */

/*
 * x_digit ::= x | X
 */

/*
 * z_digit ::= z | Z | ?
 */

/*
 * unbased_unsized_literal ::= '0 | '1 | 'z_or_x 48
 */
type UnbasedUnsizedLiteral struct {
	Token
	Value rune
}

func (d *UnbasedUnsizedLiteral) String() string {
	return fmt.Sprintf("UnbasedUnsizedLiteral(%v)", d.Span)
}

func (d *UnbasedUnsizedLiteral) Bake() error {
	if len(d.Span.Value()) != 2 {
		return fmt.Errorf("invalid unsized literal: %q", d.Span.Value())
	}
	d.Value = []rune(d.Span.Value())[1]
	return nil
}

func (*UnbasedUnsizedLiteral) isPrimaryLiteral()     {}
func (*UnbasedUnsizedLiteral) isPrimary()            {}
func (*UnbasedUnsizedLiteral) isConstantPrimary()    {}
func (*UnbasedUnsizedLiteral) isConstantExpression() {}

//
// A.8.8 Strings
//

/*
 * string_literal ::= " { Any_ASCII_Characters } "
 */
type StringLiteral struct {
	Token
	Text string
}

func (s *StringLiteral) String() string {
	return fmt.Sprintf("StringLiteral(%q)", s.Text)
}

func (s *StringLiteral) Bake() error {
	text, err := strconv.Unquote(string(s.Span.Value()))
	if err != nil {
		return err
	}
	s.Text = text
	return nil
}

func (*StringLiteral) isPrimaryLiteral()     {}
func (*StringLiteral) isPrimary()            {}
func (*StringLiteral) isConstantPrimary()    {}
func (*StringLiteral) isConstantExpression() {}
