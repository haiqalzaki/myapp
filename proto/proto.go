package proto

type Person struct {
	Name  string
	Age   uint
	Email string
}

func Person_Create(name string, age uint, email string) Person {

	object_person := Person{
		Name:  name,
		Age:   age,
		Email: email,
	}

	return object_person
}

// Golang function capitalization determines visibility
// Uppercase method name = exported / accessible outside
// Lowercase method name = local / inaccessible outside

func (p Person) Walk() string {
	return p.Name + " is walking..." + p.eating()
}

func (p Person) eating() string {
	return " whilst eating ice cream!"
}
