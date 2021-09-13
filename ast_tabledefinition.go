package tsql

var _ Node = &TableDefinitionNode{}

// TableDefinitionNode represents the definition of a table. This the parent node for CREATE TABLE
type TableDefinitionNode struct {
	Name MultiPartNameNode
}

func (n *TableDefinitionNode) Parse(in string) (string, error) {
	var rem string
	var err error

	if _, rem, err = ExpectTABLELiteral(in); err != nil {
		return in, err
	}

	if _, rem, err = ExpectOptionalWhitespace(rem); err != nil {
		return in, err
	}

	if rem, err = n.Name.Parse(rem); err != nil {
		return in, err
	}

	if _, rem, err = ExpectOptionalWhitespace(rem); err != nil {
		return in, err
	}

	if _, rem, err = ExpectParenLeft(rem); err != nil {
		return in, err
	}

	return rem, nil
}

func (n *TableDefinitionNode) String() string {
	return "TABLE " + n.Name.String() + " ("
}
