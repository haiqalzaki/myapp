package proto

import (
	"testing"
)

func TestProto_Struct(t *testing.T) {
	tests := []struct {
		Name     string
		Age      uint
		Email    string
		Expected Person
	}{
		{
			Name:  "Bob",
			Age:   40,
			Email: "bob@carlo.com",
			Expected: Person{
				Name:  "Bob",
				Age:   40,
				Email: "bob@carlo.com",
			},
		},
		{
			Name:  "Ayu",
			Age:   22,
			Email: "ayu@hr.com",
			Expected: Person{
				Name:  "Ayu",
				Age:   22,
				Email: "ayu@hr.com",
			},
		},
		{
			Name:  "Garina",
			Age:   35,
			Email: "garina@marconi.com",
			Expected: Person{
				Name:  "Garina",
				Age:   35,
				Email: "garina@marconi.com",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := Person_Create(test.Name, test.Age, test.Email)

			if result != test.Expected {
				t.Errorf("Person_Create(%s, %d, %s) = %+v; want %+v", test.Name, test.Age, test.Email, result, test.Expected)
			}
		})
	}
}

func TestWalk(t *testing.T) {
	person := Person{
		Name:  "John",
		Age:   30,
		Email: "john@example.com",
	}

	result := person.Walk()

	expected := "John is walking... whilst eating ice cream!"

	if result != expected {
		t.Errorf("Walk() = %v; want %v", result, expected)
	}
}

func TestEating(t *testing.T) {
	// Create a Person instance
	person := Person{
		Name:  "John",
		Age:   30,
		Email: "john@example.com",
	}

	// Call eating() method
	result := person.eating()

	// Expected result
	expected := " whilst eating ice cream!"

	// Check if the result matches the expected output
	if result != expected {
		t.Errorf("eating() = %v; want %v", result, expected)
	}
}
