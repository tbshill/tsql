package tsql

import (
	"strings"
)

var _ Node = &MultiPartNameNode{}

type MultiPartNameNode struct {
	Server   string `json:"Server"`
	Database string `json:"Database"`
	Schema   string `json:"Schema"`
	Object   string `json:"Object"`
}

func (n *MultiPartNameNode) Parse(in string) (string, error) {
	//BUG: This doesn't parse all the parts of the name
	var tok, rem string
	var err error

	if tok, rem, err = ExpectIdentifier(in); err != nil {
		return in, err
	}

	n.Object = tok
	return rem, nil
}

func (n *MultiPartNameNode) String() string {
	if n == nil {
		return ""
	}

	parts := []string{
		sqBrackets(n.Server),
		sqBrackets(n.Database),
		sqBrackets(n.Schema),
		sqBrackets(n.Object),
	}

	min := 0
	for i, part := range parts {
		if part != "" {
			min = i
			break
		}
	}
	return strings.Join(parts[min:], ".")

}
