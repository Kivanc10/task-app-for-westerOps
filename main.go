package main

import (
	"github.com/Kivanc10/task-app-for-westerOps/route"
)

func main() {
	// fmt.Println("hello dude")
	// db = mydb.Connect()
	// if todo, err := mydb.CreateTodo("First todo", "user 1", db); err != nil {
	// 	fmt.Println("error ", err)
	// } else {
	// 	fmt.Println("created todo --> ", todo)
	// }
	route.HandleRequest()
}
