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

	var sum uint = utils.AddNumbers(x, y)

	fmt.Printf("This is the sum of two variables = %d\n", sum)

	person := proto.Proto_Struct("Mathias", 27, "mathias@kol.com")

	fmt.Println("This is the person name:", person.Name)
	fmt.Println("This is the person age:", person.Age)
	fmt.Println("This is the person email:", person.Email)
}
