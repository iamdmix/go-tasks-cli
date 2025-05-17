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

func markDone(args []string){
	if len(args) == 0{
		fmt.Println("Please provide a task number to mark as done. ")
		return
	}
	index, err := strconv.Atoi(args[0])
	if err != nil || index < 1 {
		fmt.Println("Invalid task number.")
		return
	}

	tasks := loadTasks()

	if index> len(tasks){
		fmt.Println("Task number out of range.")
		return
	}

	if tasks[index-1].Done {
		fmt.Println("Task is already marked as done.")
		return
	}

	tasks[index-1].Done = true
	saveTasks(tasks)
	fmt.Println("Marked task as done:",tasks[index-1].Description)
}

func undoDone(args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide task number to unmark as done.")
		return
	}
	index, err := strconv.Atoi(args[0])
	if err != nil || index < 1 {
		fmt.Println("Invalid task number")
		return
	}

	tasks := loadTasks()

	if index > len(tasks) {
		fmt.Println("Task number out of range.")
		return
	}

	tasks[index-1].Done = false
	saveTasks(tasks)
	fmt.Println("Unmarked task as done:", tasks[index-1].Description)
}


func clearTasks(args []string) {
	tasks := loadTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks to clear.")
		return
	}

	if len(args) > 0 && args[0] == "done" {
		// Clear only done tasks
		var activeTasks []Task
		for _, t := range tasks {
			if !t.Done {
				activeTasks = append(activeTasks, t)
			}
		}
		saveTasks(activeTasks)
		fmt.Println("Cleared all done tasks.")
	} else {
		// Clear all tasks
		saveTasks([]Task{})
		fmt.Println("Cleared all tasks.")
	}
}


func printHelp() {
	fmt.Println(`Usage: go run main.go [command] [arguments]

Commands:
  add <task description>    Add a new task
  list                      List all tasks
  remove <task number>      Remove a task
  done <task number>        Mark a task as done
  undo <task number>        Unmark a task as done
  clear [done]              Clear all tasks or only done tasks
  help                      Show this help message
`)
}

func main(){
	if len(os.Args) < 2{
		fmt.Println("Usage: go run main.go [add|list|remove|done] <tasks>")
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
	case "done":
		markDone(os.Args[2:])
	case "undo":
		undoDone(os.Args[2:])
	case "clear":
		clearTasks(os.Args[2:])
	case "help":
		printHelp()
	default:
		fmt.Println("Unknown Command: ", command)
		printHelp()
	}
}