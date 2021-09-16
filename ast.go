package tsql

import "errors"

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrUnfinished     = errors.New("implementation is not complete")
)

type Node interface {
	String() string
	Parse(in string) (string, error)
}

func sqBrackets(s string) string {
	if s == "" {
		return s
	}
	return "[" + s + "]"
}

func ParseStatement(in string) (Node, error) {
	p := new(StatementNode)
	_, err := p.Parse(in)
	if err != nil {
		return nil, err
	}

	return p, nil
}
