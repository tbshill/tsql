package tsql

import "fmt"

var _ Node = &ColumnDefinitionNode{}

// ColumnDefinitionNode represents the definition of a column. Typically, this will be in a CREATE TABLE Statement
type ColumnDefinitionNode struct {
	Name ObjectNameNode
	DataType Node
}

func (c *ColumnDefinitionNode) Parse(in string) (string, error) {
	var rem string
	var err error
	if rem, err = c.Name.Parse(in); err != nil {
			return in, err
	}

	if _, rem, err = ExpectOptionalWhitespace(rem); err != nil{
		return in, err
	}

	datatypes := []Node{
		&IntDataTypeNode{},
		&BigIntDataTypeNode{},
		&BitDataTypeNode{},
		&DecimalDataTypeNode{},
		&MoneyDataTypeNode{},
		&NumericDataTypeNode{},
		&SmallIntDataTypeNode{},
		&SmallMoneyDataTypeNode{},
		&TinyIntDataTypeNode{},
		&FloatDataTypeNode{},
		&RealDataTypeNode{},
		&DateDataTypeNode{},
		&DatetimeDataTypeNode{},
		&Datetime2DataTypeNode{},
		&DatetimeOffsetDataTypeNode{},
		&SmallDatetimeDataTypeNode{},
		&TimeDataTypeNode{},
		&VarcharDataTypeNode{},
		&CharDataTypeNode{},
		&TextDataTypeNode{},
		&NCharDataTypeNode{},
		&NTextDataTypeNode{},
		&NVarcharDataTypeNode{},
		&BinaryDataTypeNode{},
		&ImageDataTypeNode{},
		&VarBinaryDataTypeNode{},
		&RowVersionDataTypeNode{},
		&HierarchyIDDataTypeNode{},
		&UniqueIdentifierDataTypeNode{},
		&XMLDataTypeNode{},
    }

	tmpRem := ""
	for _, datatype := range datatypes {
		 if tmpRem, err = datatype.Parse(rem); err == nil {
			 c.DataType = datatype
			 rem = tmpRem
			 break
		 }
	}

	if c.DataType == nil {
		return in, fmt.Errorf("could not find a datatype for the column")
	}
	return rem, nil
}

func (c *ColumnDefinitionNode) String() string {
	return c.Name.String() + " " +c.DataType.String()
}
