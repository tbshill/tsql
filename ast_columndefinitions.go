package tsql

import (
	"strings"
)

var _ Node = &ColumnDefinitionsNode{}

// ColumnDefinitionsNode is a root node for all the column definitions
type ColumnDefinitionsNode struct {
	Definitions []*ColumnDefinitionNode
}

func (n *ColumnDefinitionsNode) Parse(in string) (string, error) {
	var rem string
	var err error
	rem = in

	newColumn := func (in string) (string, error){
		rem := in
		if _, rem, err = ExpectOptionalWhitespace(rem); err != nil {
			return in, err
		}
		columnDefinition := new(ColumnDefinitionNode)
		rem, err = columnDefinition.Parse(rem)
		if err != nil {
			return in, err
		}
		if _, rem, err = ExpectOptionalWhitespace(rem); err != nil {
			return in, err
		}
		n.Definitions = append(n.Definitions, columnDefinition)
		return rem, nil
	}

	if rem, err = newColumn(rem); err != nil {
		return in, err
	}

	for {
		if rem[0] == ')' {
			break
		}
		if _, rem, err = ExpectComma(rem); err != nil {
			return in, err
		}
		if rem, err = newColumn(rem); err != nil {
			return in, err
		}
		if _, rem, err = ExpectOptionalWhitespace(rem); err != nil {
			return in, err
		}
	}
	return rem, nil
}

func (n *ColumnDefinitionsNode) String() string {
	strs := make([]string, len(n.Definitions))
	for i, def := range n.Definitions {
		strs[i] = def.String()
	}
	return strings.Join(strs, ",\n")
}
