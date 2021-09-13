package tsql

import (
	"testing"
)

func TestMultiPartName_String(t *testing.T) {
	tests := []struct {
		multiPartName  *MultiPartNameNode
		expectedString string
	}{
		{
			multiPartName: &MultiPartNameNode{
				Server:   "",
				Database: "",
				Schema:   "",
				Object:   "MyTable",
			},
			expectedString: "[MyTable]",
		},
		{
			multiPartName: &MultiPartNameNode{
				Server:   "",
				Database: "BI_STAGING",
				Schema:   "dbo",
				Object:   "MyTable",
			},
			expectedString: "[BI_STAGING].[dbo].[MyTable]",
		},
		{
			multiPartName: &MultiPartNameNode{
				Server:   "",
				Database: "BI_STAGING",
				Schema:   "",
				Object:   "MyTable",
			},
			expectedString: "[BI_STAGING]..[MyTable]",
		},
	}

	for _, test := range tests {
		t.Run(test.expectedString, func(t *testing.T) {
			if got := test.multiPartName.String(); got != test.expectedString {
				t.Errorf("Expected: '%s', Got: '%s'", test.expectedString, got)
			}
		})
	}
}

func TestMultiPartNameNode_Parse(t *testing.T) {
	tests := []struct {
		in     string
		expect string
	}{
		{
			in:     "MyTable",
			expect: "[MyTable]",
		},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			m := &MultiPartNameNode{}
			_, err := m.Parse(test.in)
			if err != nil {
				t.Errorf("Failed to parse the multipart name: %v", err)
				return
			}
			if m.String() != test.expect {
				t.Errorf("MultiPart Name was not parsed correctly.")
			}
		})
	}
}
