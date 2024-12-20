package main

import (
	"fmt"
	"myapp/proto"
	"myapp/utils"
	helper "mymodule"
)

func main() {
	fmt.Println("Hello, Go!")
	utils.HelperFunction()

	fmt.Println("This is a helper function from another module.")
	helper.HelperFunction()

	var x uint = 287
	var y uint = 599

	var totalSum uint = utils.AddNumbers(x, y)

	fmt.Printf("This is the sum of two variables = %d\n", totalSum)

	var userName, userEmail string
	var userAge uint

	// get input from user
	fmt.Println("Enter name:") // output
	fmt.Scanln(&userName)      // input
	fmt.Println("Enter age:")
	fmt.Scanln(&userAge)
	fmt.Println("Enter email:")
	fmt.Scanln(&userEmail)

	// create new Person struct instance using Person_Create function
	person := proto.Person_Create(userName, userAge, userEmail)

	// display object data
	fmt.Println("This is the user name:", person.Name)
	fmt.Println("This is the user age:", person.Age)
	fmt.Println("This is the user email:", person.Email)

	// call Walk method (function associated with Person struct)
	fmt.Println(person.Walk())
}
