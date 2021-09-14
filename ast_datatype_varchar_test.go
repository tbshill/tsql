package tsql

import (
	"testing"
)

func TestVarcharDataTypeNode_Parse(t *testing.T) {
	node := &VarcharDataTypeNode{}
	rem, err := node.Parse("VARCHAR(10)")

	if err != nil {
		t.Errorf("VarcharDataTypeNode.Parse returned an error with valid input: %v", err)
	}

	if rem != "" {
		t.Errorf("VarcharDataTypeNode.Parse did not consume all characters")
	}

	if node.Capacity != 10 {
		t.Errorf("Node was not 10")
	}

	rem, err = node.Parse("VARCHAR(MAX)")
	if err != nil {
		t.Errorf("VarcharDataTypeNode.Parse returned an error with valid input: %v", err)
	}

	if rem != "" {
		t.Errorf("VarcharDataTypeNode.Parse did not consume all characters")
	}

	if node.Capacity != MaxVarcharCapacity {
		t.Errorf("Node was not MAX")
	}
}

func TestVarcharDataTypeNode_String(t *testing.T) {
	node := VarcharDataTypeNode{
		Capacity: 10,
	}

	if node.String() != "VARCHAR(10)" {
		t.Errorf("VarcharDataTypeNode does not return the correct string")
	}

	node = VarcharDataTypeNode{
		Capacity: MaxVarcharCapacity,
	}

	if node.String() != "VARCHAR(MAX)" {
		t.Errorf("VarcharDataTypeNode does not return the correct string")
	}
}
