package main

import "fmt"

type Task struct {
	ID       int
	Title    string
	Complete bool
}

var tasks []Task
var nextID int = 1

func addTask(title string) {
	task := Task{
		ID:       nextID,
		Title:    title,
		Complete: false,
	}
	tasks = append(tasks, task)
	nextID++
	fmt.Println("Added Task", title)
}

func listTasks() {
	fmt.Println("TO-DO List:")
	for _, task := range tasks {
		status := "Pending"
		if task.Complete {
			status = "Done"
		}
		fmt.Printf("ID: %d | %s | %s \n", task.ID, task.Title, status)
	}
}

func completeTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Complete = true
			fmt.Println("Marked as Complete:", task.Title)
			return
		}
	}
	fmt.Println("Task not found with ID:", id)
}

func main() {
	addTask("Learn Go Lang")
	addTask("Build To-Do app")
	listTasks()

	completeTask(1)
	listTasks()
}
