# Go Tasks CLI

A simple command-line task manager written in Go.

## Features

- Add, list, and remove tasks
- Mark tasks as done and undo them
- Clear all tasks or only done tasks
- Color-coded output for better readability:
  - Done tasks and their checkmarks appear in green
  - Pending tasks appear in white
  - Success messages in cyan
  - Errors and warnings in red

## Installation

Make sure you have Go installed. Then, clone this repository and install dependencies:

```bash
git clone https://github.com/iamdmix/go-tasks-cli.git
cd go-tasks-cli
go get github.com/fatih/color
```

## Usage

Run commands using:

```bash
go run main.go [command] [arguments]
```

### Commands

| Command                  | Description                          |
|--------------------------|--------------------------------------|
| `add <task description>` | Add a new task                       |
| `list`                   | List all tasks                       |
| `remove <task number>`   | Remove a task                        |
| `done <task number>`     | Mark a task as done                  |
| `undo <task number>`     | Unmark a task as done                |
| `clear [done]`           | Clear all tasks or only done tasks   |
| `help`                   | Show this help message               |

### Examples

Add a task:

```bash
go run main.go add "Buy groceries"
```

List tasks:

```bash
go run main.go list
```

Mark task #1 as done:

```bash
go run main.go done 1
```

Undo done on task #1:

```bash
go run main.go undo 1
```

Remove task #1:

```bash
go run main.go remove 1
```

Clear all done tasks:

```bash
go run main.go clear done
```

Clear all tasks:

```bash
go run main.go clear
```

Show help:

```bash
go run main.go help
```

## File Storage

Tasks are stored locally in a JSON file named `tasks.json` in the same directory as the program.