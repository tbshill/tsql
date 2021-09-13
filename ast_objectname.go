package tsql

var _ Node = &ObjectNameNode{}

type ObjectNameNode struct {
	Name string
}

func (n *ObjectNameNode) Parse(in string) (string, error) {
	var tok, rem string
	var err error

	//BUG: Doesn't parse square bracket wrapped names
	if tok, rem, err = ExpectIdentifier(in); err != nil {
		return in, err
	}

	n.Name = tok
	return rem, nil
}

func (n *ObjectNameNode) String() string {
	return sqBrackets(n.Name)
}
