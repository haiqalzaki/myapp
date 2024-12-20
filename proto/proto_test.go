package proto

import "testing"

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
			result := Proto_Struct(test.Name, test.Age, test.Email)

			if result != test.Expected {
				t.Errorf("Proto_Struct(%s, %d, %s) = %+v; want %+v", test.Name, test.Age, test.Email, result, test.Expected)
			}
		})
	}
}
