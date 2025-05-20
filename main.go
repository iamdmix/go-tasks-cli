package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"text/tabwriter"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
)

type Task struct {
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
}

const taskFile = "tasks.json"

func loadTasks() []Task {
	var tasks []Task
	data, err := os.ReadFile(taskFile)
	if err == nil {
		json.Unmarshal(data, &tasks)
	}
	return tasks
}

func saveTasks(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	_ = os.WriteFile(taskFile, data, 0644)
}

func addTask(args []string) {
	if len(args) == 0 {
		color.Red("Task description missing")
		return
	}
	description := strings.Join(args, " ")
	tasks := loadTasks()
	tasks = append(tasks, Task{
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
	})
	saveTasks(tasks)
	color.Cyan("Added task: %s", description)
}

func listTasks() {
	tasks := loadTasks()
	if len(tasks) == 0 {
		color.Yellow("No tasks found.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Header
	fmt.Fprintln(w, "ID\tTask\tCreated\tDone")

	for i, task := range tasks {
		created := humanize.Time(task.CreatedAt)
		done := "false"
		if task.Done {
			done = "true"
		}
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", i+1, task.Description, created, done)
	}

	w.Flush()
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
		var remaining []Task
		for _, t := range tasks {
			if !t.Done {
				remaining = append(remaining, t)
			}
		}
		saveTasks(remaining)
		color.Cyan("Cleared all done tasks.")
	} else {
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
