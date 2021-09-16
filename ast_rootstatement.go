package tsql

import (
	"fmt"
)

var _ Node = &StatementNode{}

type StatementNode struct {
	Statement Node
}

func (n *StatementNode) Parse(in string) (string, error) {
	var err error
	var rem string
	startingStatements := []Node{
		&CreateObjectNode{},
	}

	for _, p := range startingStatements{
		if rem, err = p.Parse(in); err == nil {
			n.Statement = p
			return rem, nil
		} else {
			fmt.Println(err)
		}
	}

	return in, fmt.Errorf("Could not find a starting statement")
}

func (n *StatementNode) String() string {
	return n.Statement.String()
}
