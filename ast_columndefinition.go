package tsql

var _ Node = &ColumnDefinitionNode{}

// ColumnDefinitionNode represents the definition of a column. Typically, this will be in a CREATE TABLE Statement
type ColumnDefinitionNode struct {
	Name ObjectNameNode
}

func (c *ColumnDefinitionNode) Parse(in string) (string, error) {
	return in, ErrNotImplemented
}

func (c *ColumnDefinitionNode) String() string {
	return "<Column Definition>"
}
