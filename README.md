# Go Tasks CLI

A lightweight and intuitive command-line task manager written in Go.  
Easily manage your daily tasks locally with simple commands to add, list, remove, and mark tasks as done.

---

## ✨ Features

- **Add tasks** with custom descriptions
- **List** all tasks with their status (pending/done)
- **Remove** tasks by their number
- **Mark** tasks as done

---

## 🚀 Requirements

- Go **1.18** or higher installed

---

## ⚡ Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/go-tasks-cli.git
   cd go-tasks-cli
   ```

2. **Build the executable (optional):**
   ```bash
   go build -o go-tasks
   ```

3. **Or run directly without building:**
   ```bash
   go run main.go <command> [arguments]
   ```

---

## 🛠️ Usage

```bash
go run main.go add "Buy groceries"
go run main.go list
go run main.go done 1
go run main.go remove 1
```

### Available Commands

| Command | Description                     |
|---------|---------------------------------|
| `add`   | Add a new task                  |
| `list`  | List all tasks                  |
| `done`  | Mark a task as done by its number |
| `remove`| Remove a task by its number     |

---

## 💾 Data Storage

Tasks are stored locally in a JSON file named `tasks.json`.  
This file is automatically created and updated as you manage your tasks.

---


Enjoy productive task management from the comfort of your terminal!