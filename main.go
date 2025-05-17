package main

import(
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"strings"
	"strconv"
)
type Task struct{
	Description string `json:"description"`
	Done bool `json:"done"`
}
const taskFile = "tasks.json"

func loadTasks() []Task {
	var tasks []Task

	data, err := ioutil.ReadFile(taskFile)
	if err == nil {
		json.Unmarshal(data,&tasks)
	}
	return tasks
}

func saveTasks(tasks []Task){
	data, _ := json.MarshalIndent(tasks,""," ")
	_ = ioutil.WriteFile(taskFile, data, 0644) 
}

func addTask(args []string){
	if len(args) ==0 {
		fmt.Println("Task description missing")
		return 
	}
	description := strings.Join(args," ")
	tasks := loadTasks()
	tasks = append(tasks, Task{Description: description, Done:false})
	saveTasks(tasks)
	fmt.Println("Added task:",description)
}

func listTasks(){
	tasks:= loadTasks()

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for i, task := range tasks {
		status := " "
		if task.Done {
			status="x"
		}
		fmt.Printf("%d. [%s] %s\n",i+1, status, task.Description)
	}
}

func removeTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide task number to remove.")
		return
	}
	index, err := strconv.Atoi(args[0])
	if err != nil || index < 1 {
		fmt.Println("Invalid task number")
		return
	}

	tasks := loadTasks()

	if index > len(tasks) {
		fmt.Println("Task number out of range. ")
		return
	}

	removed := tasks[index-1].Description
	tasks = append(tasks[:index-1], tasks[index:]...)
	saveTasks(tasks)
	fmt.Println("Removed task:",removed)
}

func main(){
	if len(os.Args) < 2{
		fmt.Println("Usage: go run main.go [add|list|remove] <tasks>")
		return
	}

	command := os.Args[1]

	switch command{
	case "add":
		addTask(os.Args[2:])
	case "list":
		listTasks()
	case "remove":
		removeTask(os.Args[2:])
	default:
		fmt.Println("Unknown Command: ", command)
	}
}