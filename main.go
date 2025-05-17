package main

import(
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"strings"
	"strconv"

	"github.com/fatih/color"
)

type Task struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

const taskFile = "tasks.json"

func loadTasks() []Task {
	var tasks []Task

	data, err := ioutil.ReadFile(taskFile)
	if err == nil {
		json.Unmarshal(data, &tasks)
	}
	return tasks
}

func saveTasks(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	_ = ioutil.WriteFile(taskFile, data, 0644)
}

func addTask(args []string) {
	if len(args) == 0 {
		color.Red("Task description missing")
		return
	}
	description := strings.Join(args, " ")
	tasks := loadTasks()
	tasks = append(tasks, Task{Description: description, Done: false})
	saveTasks(tasks)
	color.Cyan("Added task: %s", description)
}

func listTasks() {
	tasks := loadTasks()

	if len(tasks) == 0 {
		color.Yellow("No tasks found.")
		return
	}

	green := color.New(color.FgGreen).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()

	for i, task := range tasks {
		status := " "
		if task.Done {
			status = "x"
			fmt.Printf("%d. [%s] %s\n", i+1, green(status), green(task.Description))
		} else {
			fmt.Printf("%d. [%s] %s\n", i+1, white(status), white(task.Description))
		}
	}
}

func removeTask(args []string) {
	if len(args) == 0 {
		color.Red("Please provide task number to remove.")
		return
	}
	index, err := strconv.Atoi(args[0])
	if err != nil || index < 1 {
		color.Red("Invalid task number")
		return
	}

	tasks := loadTasks()

	if index > len(tasks) {
		color.Red("Task number out of range.")
		return
	}

	removed := tasks[index-1].Description
	tasks = append(tasks[:index-1], tasks[index:]...)
	saveTasks(tasks)
	color.Cyan("Removed task: %s", removed)
}

func markDone(args []string) {
	if len(args) == 0 {
		color.Red("Please provide a task number to mark as done.")
		return
	}
	index, err := strconv.Atoi(args[0])
	if err != nil || index < 1 {
		color.Red("Invalid task number.")
		return
	}

	tasks := loadTasks()

	if index > len(tasks) {
		color.Red("Task number out of range.")
		return
	}

	if tasks[index-1].Done {
		color.Yellow("Task is already marked as done.")
		return
	}

	tasks[index-1].Done = true
	saveTasks(tasks)
	color.Cyan("Marked task as done: %s", tasks[index-1].Description)
}

func undoDone(args []string) {
	if len(args) == 0 {
		color.Red("Please provide task number to unmark as done.")
		return
	}
	index, err := strconv.Atoi(args[0])
	if err != nil || index < 1 {
		color.Red("Invalid task number")
		return
	}

	tasks := loadTasks()

	if index > len(tasks) {
		color.Red("Task number out of range.")
		return
	}

	tasks[index-1].Done = false
	saveTasks(tasks)
	color.Cyan("Unmarked task as done: %s", tasks[index-1].Description)
}

func clearTasks(args []string) {
	tasks := loadTasks()
	if len(tasks) == 0 {
		color.Yellow("No tasks to clear.")
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
		color.Cyan("Cleared all done tasks.")
	} else {
		// Clear all tasks
		saveTasks([]Task{})
		color.Cyan("Cleared all tasks.")
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

func main() {
	if len(os.Args) < 2 {
		color.Red("Usage: go run main.go [add|list|remove|done|undo|clear|help] <tasks>")
		return
	}

	command := os.Args[1]

	switch command {
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
		color.Red("Unknown Command: %s", command)
		printHelp()
	}
}
