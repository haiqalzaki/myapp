package main

import (
	"dirman"
	"fmt"
	"myapp/proto"
	"myapp/utils"
	helper "mymodule"
	"os"
)

func main() {

	// ========================================================== //
	// =======================FUNC PROGRAM======================= //
	// ========================================================== //

	fmt.Println("Hello, Go!")
	utils.HelperFunction()

	fmt.Println("This is a helper function from another module.")
	helper.HelperFunction()

	var x uint = 287
	var y uint = 599

	var totalSum uint = utils.AddNumbers(x, y)

	fmt.Printf("This is the sum of two variables = %d\n", totalSum)

	// ========================================================== //
	// =======================STRUCT PROGRAM===================== //
	// ========================================================== //

	// declare variable to store input
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

	// passed Person struct object pointer to function, returned same pointer stored in personPtr, then display data
	var personPtr = proto.Person_Get(&person)
	fmt.Println(personPtr.Age)

	// Pause program exit
	fmt.Println("Press the Enter Key to terminate the console screen!")
	fmt.Scanln()

	// Iterate each item in current directory and display
	dirman.See_Dir()

	// ========================================================== //
	// =======================QUIT PROGRAM======================= //
	// ========================================================== //

	// declare var to store user exit status input
	var exitStatus string

	// get user input and store at the address of var
	fmt.Printf("Press y repeat program, n to exit >> ")
	fmt.Scanln(&exitStatus)

	// check user input for programming exit/continuity
	if exitStatus == "Y" || exitStatus == "y" {
		main()
	} else {
		os.Exit(0)
	}
}
