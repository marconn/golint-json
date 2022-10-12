package main

import (
	"testing"
)

func TestLint(t *testing.T) {
	type TestTable struct {
		filePath   string
		shouldFail bool
	}

	tests := []TestTable{
		{
			filePath:   "./testData/ok.json",
			shouldFail: false,
		},
		{
			filePath:   "./testData/error_invalid_format.json",
			shouldFail: true,
		},
		{
			filePath:   "./testData/error_missing_value.json",
			shouldFail: true,
		},
		{
			filePath:   "./testData/error_missing_closing_bracket.json",
			shouldFail: true,
		},
	}

	for _, test := range tests {
		err := lint(test.filePath)

		if test.shouldFail && err == nil {
			t.Errorf("Test should fail for \"%s\"", test.filePath)
		}
	}
}
