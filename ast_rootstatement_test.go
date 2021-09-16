package tsql

import (
	"testing"
	"encoding/json"
	"fmt"
)

func TestStatementNode_String(t *testing.T){
	t.SkipNow()
}

func TestStatementNode_Parse(t *testing.T){
	n := new(StatementNode)

	in := `CREATE TABLE Patients (
	FirstName NVARCHAR(100),
	LastName NVARCHAR(100),
	DateOfBirth DATE
);`

	rem, err := n.Parse(in)
	if err != nil {
		t.Errorf("Statement returned an error with valid input: %v", err)
	}

	if rem != ""{
		t.Errorf("Statement did not consume all of the input")
	}

	bytes, err := json.MarshalIndent(n,"","    ")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(bytes))

}
