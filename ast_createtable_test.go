package tsql

import (
	"fmt"
	"testing"
)

func TestCreateTable(t *testing.T) {
	in := `CREATE TABLE MyTable (
	mycol int,
	my2ndCol varchar(100),
	supperCool TEXT
);`

	var createNode CreateObjectNode
	if _, err := createNode.Parse(in); err != nil {
		t.Errorf("Failed to parse a create table: %v", err)
	}
	fmt.Println(createNode.String())
}
