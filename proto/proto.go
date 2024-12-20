package proto

type Person struct {
	Name  string
	Age   uint
	Email string
}

// Golang function capitalization determines visibility
// Uppercase method name = exported / accessible outside
// Lowercase method name = local / inaccessible outside

func Person_Create(name string, age uint, email string) Person {

	return Person{
		Name:  name,
		Age:   age,
		Email: email,
	}
}

func (p Person) Walk() string {
	return p.Name + " is walking..." + p.eating()
}

func (p Person) eating() string {
	return " whilst eating ice cream!"
}

func Person_Get(person *Person) *Person {
	return person
}
