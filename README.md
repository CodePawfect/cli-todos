
# CLI To-Do Manager

I have a bad memory and I like CLI tools. This CLI tool is a simple To-Do manager to keep track of tasks with minimal effort.

## Features
- **Add tasks:** `-add "task description"`
- **List tasks:** `-list`
- **Toggle task status:** `-toggle [index]`
- **Delete tasks:** `-delete [index]`
- **Clear all tasks:** `-clear`

## How It Works
Tasks are saved to a `todos.json` file. You can view, add, toggle (mark as finished or in-progress), delete, and clear tasks as needed.

### Usage

```bash
# Add a new task
go run main.go -add "Buy groceries"

# List all tasks
go run main.go -list

# Toggle the status of task at index 0
go run main.go -toggle 0

# Delete the task at index 0
go run main.go -delete 0

# Clear all tasks
go run main.go -clear
