package tsql

import (
	"testing"
)

func TestNCharDataTypeNode_Parse(t *testing.T) {
	node := &NCharDataTypeNode{}
	rem, err := node.Parse("NCHAR(10)")

	if err != nil {
		t.Errorf("NCharDataTypeNode.Parse returned an error with valid input: %v", err)
	}

	if rem != "" {
		t.Errorf("NCharDataTypeNode.Parse did not consume all characters")
	}

	if node.Capacity != 10 {
		t.Errorf("Node was not 10")
	}

}

func TestNCharDataTypeNode_String(t *testing.T) {
	node := NCharDataTypeNode{
		Capacity: 10,
	}

	if node.String() != "NCHAR(10)" {
		t.Errorf("NCharDataTypeNode does not return the correct string")
	}
}
