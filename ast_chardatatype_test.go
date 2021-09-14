package tsql

import (
	"testing"
)

func TestCharDataTypeNode_Parse(t *testing.T) {
	node := &CharDataTypeNode{}
	rem, err := node.Parse("CHAR(10)")

	if err != nil {
		t.Errorf("CharDataTypeNode.Parse returned an error with valid input: %v", err)
	}

	if rem != "" {
		t.Errorf("CharDataTypeNode.Parse did not consume all characters")
	}

	if node.Capacity != 10 {
		t.Errorf("Node was not 10")
	}

}

func TestCharDataTypeNode_String(t *testing.T) {
	node := CharDataTypeNode{
		Capacity: 10,
	}

	if node.String() != "CHAR(10)" {
		t.Errorf("CharDataTypeNode does not return the correct string")
	}
}
