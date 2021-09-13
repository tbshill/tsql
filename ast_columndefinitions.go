package tsql

var _ Node = &ColumnDefinitionsNode{}

// ColumnDefinitionsNode is a root node for all the column definitions
type ColumnDefinitionsNode struct {
	Definitions []*ColumnDefinitionNode
}

func (n *ColumnDefinitionsNode) Parse(in string) (string, error) {
	var rem, tmpRem string
	var err error
	rem = in

	for {
		columnDefinition := new(ColumnDefinitionNode)
		tmpRem, err = columnDefinition.Parse(rem)
		if err != nil {
			return rem, nil
		}
		rem = tmpRem
		n.Definitions = append(n.Definitions, columnDefinition)
	}
}

func (n *ColumnDefinitionsNode) String() string {
	return "[<Column Definitions>]"
}
