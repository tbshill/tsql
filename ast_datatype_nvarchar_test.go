package tsql

import (
	"testing"
)

func TestNVarcharDataTypeNode_Parse(t *testing.T) {
	node := &NVarcharDataTypeNode{}
	rem, err := node.Parse("NVARCHAR(10)")

	if err != nil {
		t.Errorf("NVarcharDataTypeNode.Parse returned an error with valid input: %v", err)
	}

	if rem != "" {
		t.Errorf("NVarcharDataTypeNode.Parse did not consume all characters")
	}

	if node.Capacity != 10 {
		t.Errorf("Node was not 10")
	}

	rem, err = node.Parse("NVARCHAR(MAX)")
	if err != nil {
		t.Errorf("NVarcharDataTypeNode.Parse returned an error with valid input: %v", err)
	}

	if rem != "" {
		t.Errorf("NVarcharDataTypeNode.Parse did not consume all characters")
	}

	if node.Capacity != MaxNVarcharCapacity {
		t.Errorf("Node was not MAX")
	}
}

func TestNVarcharDataTypeNode_String(t *testing.T) {
	node := NVarcharDataTypeNode{
		Capacity: 10,
	}

	if node.String() != "NVARCHAR(10)" {
		t.Errorf("NVarcharDataTypeNode does not return the correct string")
	}

	node = NVarcharDataTypeNode{
		Capacity: MaxNVarcharCapacity,
	}

	if node.String() != "NVARCHAR(MAX)" {
		t.Errorf("NVarcharDataTypeNode does not return the correct string")
	}
}
