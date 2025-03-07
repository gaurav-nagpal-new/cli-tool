package main

import (
	"cli-tool/usecase"
	"fmt"
)

// A simple CLI ToDo list - support CRUD operations
// store all todos in a file - only name and ID should be stored

type TodoApp interface {
	Create()
	Update()
	Read()
	Delete()
}

func main() {
	var t TodoApp = usecase.TodoA{}

	for {
		fmt.Println("-----------------")
		fmt.Println("1. Create a todo")
		fmt.Println("2. Update a todo")
		fmt.Println("3. Read all")
		fmt.Println("4. Delete a todo")
		fmt.Println("Press any other key to exit")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("was not able to read input ", err.Error())
			return
		}

		switch choice {
		case 1:
			t.Create()
		case 2:
			t.Update()
		case 3:
			t.Read()
		case 4:
			t.Delete()
		default:
			fmt.Println("Bye Bye !!")
			return
		}
	}

}
