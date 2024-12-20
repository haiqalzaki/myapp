package proto

type Person struct {
	Name  string
	Age   uint
	Email string
}

func Proto_Struct(name string, age uint, email string) Person {

	object_person := Person{
		Name:  name,
		Age:   age,
		Email: email,
	}

	return object_person
}
