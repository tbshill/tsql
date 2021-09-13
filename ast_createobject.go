package tsql

import parsec "github.com/tbshill/goparsec"

var _ Node = &CreateObjectNode{}

type CreateObjectNode struct {
	Object Node
}

func (n *CreateObjectNode) Parse(in string) (string, error) {
	var rem string
	var err error

	if _, rem, err = ExpectCREATELiteral(in); err != nil {
		return in, err
	}

	if _, rem, err = parsec.ExpectWhiteSpace(rem); err != nil {
		return in, err
	}

	possibleObjectNodes := []Node{
		&TableDefinitionNode{},
	}

	foundAMatch := false

	for _, node := range possibleObjectNodes {
		rem, err = node.Parse(rem)
		if err == nil {
			foundAMatch = true
			n.Object = node
			break
		}
	}

	if !foundAMatch {
		return in, err
	}

	return rem, nil
}

func (n *CreateObjectNode) String() string {
	if n.Object == nil {
		return "CREATE <No Object>"
	}
	return "CREATE " + n.Object.String()
}
