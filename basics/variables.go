package main

import "fmt"

var middleName = "Cane"

func main() {
	var age int
	var name string = "Jhon"
	var name1 = "Jane"

	//ERROR ❌ TO USE "var count := 10"
	//CORRECT ✅
	count := 10
	lastName := "Smith"

	fmt.Print(middleName)
	// dEFAULT VALUES
	// Numeric Tyoe: 0
	// Boolean Tyoes: false
	// String Types: ""
	// Pointers, slices, maps, functions and structurs: nill

	// ---- SCOPE ----
	printname(firstName)

}

func printname(){
	firstName := "Neo"
	fmt.Println(firstName)
}
