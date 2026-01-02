package main

import "fmt"

type Task struct {
	ID    int
	Title string
	Done  bool
}

func listTasks() {
	tasks := []Task{{1, "Hoc Git", false}, {2, "Code Golang", true}}
	for _, t := range tasks {
		status := " "
		if t.Done {
			status = "✓"
		}
		fmt.Printf("[%s] %d. %s\n", status, t.ID, t.Title)
	}
}

// TODO: Implement storage
func main() {
	fmt.Println("CLI To-Do Manager v0.1")
	fmt.Println("Current Tasks:")
	listTasks()
}
