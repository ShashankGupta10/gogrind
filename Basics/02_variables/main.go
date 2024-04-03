package main

import "fmt"

const CONSTANT_VALUE int = 3999

func main() {
	var name string = "shashank"
	fmt.Println(name)
	fmt.Printf("The type of the variable is: %T \n", name)

	var number int = 200
	fmt.Println(number)
	fmt.Printf("The type of the variable is: %T \n", number)

	var booleanVal bool = true
	fmt.Println(booleanVal)
	fmt.Printf("The type of the variable is: %T \n", booleanVal)

	fmt.Println(CONSTANT_VALUE)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("The type of the variable is: %T \n", smallVal)

	var floatVal float64 = 255.7846845768456489758
	fmt.Println(floatVal)
	fmt.Printf("The type of the variable is: %T \n", floatVal)
}
