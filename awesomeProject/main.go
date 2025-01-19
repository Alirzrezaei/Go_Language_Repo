package main

import (
	"fmt"
	"net/http"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
var task1 = "Watch Go crash course"
var task2 = "Watch Go full course"
var task3 = "Reward myself with a long trip"

var tasks = []string{task1, task2, task3}

func main() {

	fmt.Println("###### Welcome to our Todolist App! ######")

	http.HandleFunc("/", helloUser)

	http.HandleFunc("/show-tasks", showtasks)

	http.ListenAndServe(":8080", nil)

	fmt.Println("List of my Todos")

	//for index, element := range tasks {
	//	fmt.Printf("%d: %s\n", index+1, element)
	//}
	//printTasks(tasks)

	//tasks = addTask(tasks, "Come home after Trip")

	fmt.Println("********************************")
	fmt.Println("New List of my Todos")
	//printTasks(tasks)

}

func showtasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range tasks {
		fmt.Fprintln(writer, task)

	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello User, Welcome to our Todolist App!"
	fmt.Fprintf(writer, greeting)
}
func printTasks(tasks []string) {
	for index, element := range tasks {
		fmt.Printf("%d: %s\n", index+1, element)
	}
}

func addTask(tasks []string, newTask string) []string {
	return append(tasks, newTask)

}
