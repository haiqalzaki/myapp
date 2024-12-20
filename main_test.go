package main

import (
	"myapp/utils"
	helper "mymodule"
	"testing"
)

func TestUtilsHelperFunction(t *testing.T) {
	expected := "Helper function executed!"
	actual := utils.HelperFunction()
	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func TestMymoduleHelperFunction(t *testing.T) {
	expected := "Helper function from a separate module!"
	actual := helper.HelperFunction()
	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func TestUtilsAddNumbers(t *testing.T) {
	tests := []struct {
		num1     uint
		num2     uint
		expected uint
	}{
		{num1: 3, num2: 5, expected: 8},
		{num1: 0, num2: 0, expected: 0},
		{num1: 10, num2: 20, expected: 30},
		{num1: 100, num2: 200, expected: 300},
	}

	// for loop (Go)
	// used _ to ignore index (iterator)
	for _, test := range tests {
		result := utils.AddNumbers(test.num1, test.num2)
		if result != test.expected {
			t.Errorf("AddNumbers(%d, %d) = %d; want %d", test.num1, test.num2, result, test.expected)
		}
	}
}
