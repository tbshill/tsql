package tsql

import (
	"fmt"
	"testing"
)

func TestColumnDefinitionNode_Parse(t *testing.T) {
	in := "MyCol varchar(10)"
	col := new(ColumnDefinitionNode)
	rem, err := col.Parse(in)
	if err != nil {
		t.Errorf("Failed on valid input: %v", err)
	}
	if rem != "" {
		t.Errorf("Did not consume all of the string:%s", rem)
	}
	fmt.Println(col.String())
}