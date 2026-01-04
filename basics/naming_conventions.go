package basics

import "fmt"

type Employee struct {
	FirstName string
	LasName   string
	Age       string
}

func main() {
	// PascalCase
	//ex: CalaculateArea, UserInfo
	//Structs, Interfaces, Enums

	//snake_case
	//ex: user_id, first_name -> namefiles or variables

	//UPPERCASE
	// Use case is constans

	//mixcase
	// Ex: javaScript, htmlDocument -> namefiles or variables

	const MAXRETRIES = 5

	var employeeId = 1001

	fmt.Println("EmployeeId: ", employeeId)

}
