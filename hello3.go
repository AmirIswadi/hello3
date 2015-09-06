package main 

import "fmt"

//this is to add comment to this document. 

func main() {
	var x string = "Hello World!!"
	var name string = "Amir"
	var age string = "34"

  	fmt.Println(x)
	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Printf("\nWhat is yours?\n")
	fmt.Println("1 + 1 =", 1 + 1)
	fmt.Println(len("Hello World"))
  	fmt.Println("Hello World"[1])
  	fmt.Println("Hello " + "World")
}
