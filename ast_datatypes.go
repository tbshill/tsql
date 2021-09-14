package tsql

import (
	. "github.com/tbshill/goparsec"
	"strconv"
)

type (
	IntDataTypeNode            struct{}
	BigIntDataTypeNode         struct{}
	BitDataTyleNode            struct{}
	DecimalDataTypeNode        struct{}
	MoneyDataTypeNode          struct{}
	NumericDataTypeNode        struct{}
	SmallIntDataTypeNode       struct{}
	SmallMoneyDataTypeNode     struct{}
	TinyIntDataTypeNode        struct{}
	FloatDataTypeNode          struct{}
	RealDataTypeNode           struct{}
	DateDataTypeNode           struct{}
	DatetimeDataTypeNode       struct{}
	Datetime2DataTypeNode      struct{}
	DatetimeOffsetDataTypeNode struct{}
	SmallDatetimeDataTypeNode  struct{}
	TimeDataTypeNode           struct{}
	CharDataTypeNode           struct{}
	VarcharDataTypeNode        struct {
		Capacity int
	}
	TextDataTypeNode     struct{}
	NcharDataTypeNode    struct{}
	NtextDataTypeNode    struct{}
	NVarcharDataTypeNode struct {
		Capacity int
	}
	BinaryDataTypeNode           struct{}
	ImageDataTypeNode            struct{}
	VarBinaryDataTypeNode        struct{}
	RowVersionDataTypeNode       struct{}
	HierarchyIDDataTypeNode      struct{}
	UniqueIdentifierDataTypeNode struct{}
	XMLDataTypeNode              struct{}
)

var (
	_ Node = &IntDataTypeNode{}
	_      = &BigIntDataTypeNode{}
	_      = &BitDataTyleNode{}
	_      = &DecimalDataTypeNode{}
	_      = &MoneyDataTypeNode{}
	_      = &NumericDataTypeNode{}
	_      = &SmallIntDataTypeNode{}
	_      = &SmallMoneyDataTypeNode{}
	_      = &TinyIntDataTypeNode{}
	_      = &FloatDataTypeNode{}
	_      = &RealDataTypeNode{}
	_      = &DateDataTypeNode{}
	_      = &DatetimeDataTypeNode{}
	_      = &Datetime2DataTypeNode{}
	_      = &DatetimeOffsetDataTypeNode{}
	_      = &SmallDatetimeDataTypeNode{}
	_      = &TimeDataTypeNode{}
	_      = &VarcharDataTypeNode{}
	_      = &CharDataTypeNode{}
	_      = &TextDataTypeNode{}
	_      = &NcharDataTypeNode{}
	_      = &NtextDataTypeNode{}
	_      = &NVarcharDataTypeNode{}
	_      = &BinaryDataTypeNode{}
	_      = &ImageDataTypeNode{}
	_      = &VarBinaryDataTypeNode{}
	_      = &RowVersionDataTypeNode{}
	_      = &HierarchyIDDataTypeNode{}
	_      = &UniqueIdentifierDataTypeNode{}
	_      = &XMLDataTypeNode{}
)

func (i *IntDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectINTLiteral(in)
	return rem, err
}

func (i *IntDataTypeNode) String() string {
	return INTLiteral
}

func (i *BigIntDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectBIGINTLiteral(in)
	return rem, err
}

func (i *BigIntDataTypeNode) String() string {
	return BIGINTLiteral
}

func (i *BitDataTyleNode) Parse(in string) (string, error) {
	_, rem, err := ExpectBITLiteral(in)
	return rem, err
}

func (i *BitDataTyleNode) String() string {
	return BITLiteral
}

func (i *DecimalDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectDECIMALLiteral(in)
	return rem, err
}

func (i *DecimalDataTypeNode) String() string {
	return DECIMALLiteral
}

func (i *MoneyDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectMONEYLiteral(in)
	return rem, err
}

func (i *MoneyDataTypeNode) String() string {
	return MONEYLiteral
}

func (i *NumericDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectNUMERICLiteral(in)
	return rem, err
}

func (i *NumericDataTypeNode) String() string {
	return NUMERICLiteral
}

func (i *SmallIntDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectSMALLINTLiteral(in)
	return rem, err
}

func (i *SmallIntDataTypeNode) String() string {
	return SMALLINTLiteral
}

func (i *SmallMoneyDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectSMALLMONEYLiteral(in)
	return rem, err
}

func (i *SmallMoneyDataTypeNode) String() string {
	return SMALLMONEYLiteral
}

func (i *TinyIntDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectTINYINTLiteral(in)
	return rem, err
}

func (i *TinyIntDataTypeNode) String() string {
	return TINYINTLiteral
}

func (i *FloatDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectFLOATLiteral(in)
	return rem, err
}

func (i *FloatDataTypeNode) String() string {
	return FLOATLiteral
}

func (i *RealDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectREALLiteral(in)
	return rem, err
}

func (i *RealDataTypeNode) String() string {
	return REALLiteral
}

func (i *DateDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectDATELiteral(in)
	return rem, err
}

func (i *DateDataTypeNode) String() string {
	return DATELiteral
}

func (i *DatetimeDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectDATETIMELiteral(in)
	return rem, err
}

func (i *DatetimeDataTypeNode) String() string {
	return DATETIMELiteral
}

func (i *Datetime2DataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectDATETIME2Literal(in)
	return rem, err
}

func (i *Datetime2DataTypeNode) String() string {
	return DATETIME2Literal
}

func (i *DatetimeOffsetDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectDATETIMEOFFSETLiteral(in)
	return rem, err
}

func (i *DatetimeOffsetDataTypeNode) String() string {
	return DATETIMEOFFSETLiteral
}

func (i *SmallDatetimeDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectSMALLDATETIMELiteral(in)
	return rem, err
}

func (i *SmallDatetimeDataTypeNode) String() string {
	return SMALLDATETIMELiteral
}

func (i *TimeDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectTIMELiteral(in)
	return rem, err
}

func (i *TimeDataTypeNode) String() string {
	return TIMELiteral
}

func (i *CharDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectCHARLiteral(in)
	return rem, err
}

func (i *CharDataTypeNode) String() string {
	return CHARLiteral
}

func (i *VarcharDataTypeNode) Parse(in string) (string, error) {
	var tok, rem string
	var err error

	// VARCHAR ( <END>
	start := And(ExpectVARCHARLiteral,
		ExpectOptionalWhitespace,
		ExpectParenLeft,
		ExpectOptionalWhitespace)

	if _, rem, err = start(in); err != nil {
		return in, err
	}

	if tok, rem, err = ExpectInteger(rem); err != nil {
		return in, err
	}

	if i.Capacity, err = strconv.Atoi(tok); err != nil {
		return in, err
	}

	// <start> )
	end := And(ExpectOptionalWhitespace,
		ExpectParenRight)

	if _, rem, err = end(rem); err != nil {
		return in, err
	}

	return rem, err
}

func (i *VarcharDataTypeNode) String() string {
	capStr := strconv.Itoa(i.Capacity)
	return VARCHARLiteral + "(" + capStr + ")"
}

func (i *TextDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectTEXTLiteral(in)
	return rem, err
}

func (i *TextDataTypeNode) String() string {
	return TEXTLiteral
}

func (i *NcharDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectNCHARLiteral(in)
	return rem, err
}

func (i *NcharDataTypeNode) String() string {
	return NCHARLiteral
}

func (i *NtextDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectNTEXTLiteral(in)
	return rem, err
}

func (i *NtextDataTypeNode) String() string {
	return NTEXTLiteral
}

func (i *NVarcharDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectNVARCHARLiteral(in)
	return rem, err
}

func (i *NVarcharDataTypeNode) String() string {
	return NVARCHARLiteral
}

func (i *BinaryDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectBINARYLiteral(in)
	return rem, err
}

func (i *BinaryDataTypeNode) String() string {
	return BINARYLiteral
}

func (i *ImageDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectIMAGELiteral(in)
	return rem, err
}

func (i *ImageDataTypeNode) String() string {
	return IMAGELiteral
}

func (i *VarBinaryDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectVARBINARYLiteral(in)
	return rem, err
}

func (i *VarBinaryDataTypeNode) String() string {
	return VARBINARYLiteral
}

func (i *RowVersionDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectROWVERSIONLiteral(in)
	return rem, err
}

func (i *RowVersionDataTypeNode) String() string {
	return ROWVERSIONLiteral
}

func (i *HierarchyIDDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectHIERARCHYIDLiteral(in)
	return rem, err
}

func (i *HierarchyIDDataTypeNode) String() string {
	return HIERARCHYIDLiteral
}

func (i *UniqueIdentifierDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectUNIQUEIDENTIFIERLiteral(in)
	return rem, err
}

func (i *UniqueIdentifierDataTypeNode) String() string {
	return UNIQUEIDENTIFIERLiteral
}

// SKIPPING SQL_ VARIANT

func (i *XMLDataTypeNode) Parse(in string) (string, error) {
	_, rem, err := ExpectXMLLiteral(in)
	return rem, err
}

func (i *XMLDataTypeNode) String() string {
	return XMLLiteral
}
