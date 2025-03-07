package usecase

import (
	"fmt"
	"slices"
)

type TodoA struct {
	ID   int
	Name string
}

var data []TodoA

var ID int // global ID for ToDos

func (t TodoA) Create() {
	fmt.Println("Enter todo")
	_, err := fmt.Scan(&t.Name)
	if err != nil {
		fmt.Println("error taking input ", err.Error())
		return
	}

	t.ID = len(data) + 1
	data = append(data, t)
}

func (t TodoA) Update() {
	if len(data) == 0 {
		fmt.Println("nothing to update")
		return
	}
	fmt.Println("Enter the todo ID")
	fmt.Scan(&t.ID)

	for i := 0; i < len(data); i++ {
		if data[i].ID == t.ID {
			fmt.Println("Enter the todo value")
			fmt.Scan(&data[i].Name)
			fmt.Println("updated")
			return
		}
	}
	fmt.Println("no todo found")
}

func (t TodoA) Read() {

	fmt.Println("Following are the ToDos")
	for _, t := range data {
		fmt.Printf("ID: %d Todo: %s\n", t.ID, t.Name)
	}
}

func (t TodoA) Delete() {
	if len(data) == 0 {
		fmt.Println("nothing to delete")
		return
	}
	fmt.Println("Enter the todo ID")
	fmt.Scan(&t.ID)

	for i := 0; i < len(data); i++ {
		if data[i].ID == t.ID {
			// data = append(data[:i], data[i+1:]...)
			data = slices.Delete(data, ID, ID+1)
			fmt.Println("deleted")
			return
		}
	}
	fmt.Println("no element deleted")
}
