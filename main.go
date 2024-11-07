package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type Todo struct {
	CreatedAt  time.Time
	FinishedAt time.Time
	Name       string
	Status     string
}

var (
	InProgress = "X"
	Finished   = "\u2705" // ✅
	Todos      = []Todo{}
)

func main() {
	add := flag.String("add", "", "Add a new task to the todo list")
	list := flag.Bool("list", false, "List all Todos")
	deleteIndex := flag.Int("delete", -1, "Delete a task by its index in the list")
	toggle := flag.Int("toggle", -1, "Toggle the status of a task by its index in the list")
	clear := flag.Bool("clear", false, "Clear all tasks from the todo list")

	flag.Parse()

	readTodosFromFile("todos.json", &Todos)

	if *toggle != -1 && *toggle < len(Todos) {
		todo := &Todos[*toggle]
		if todo.Status == InProgress {
			todo.Status = Finished
			todo.FinishedAt = time.Now()
		} else {
			todo.Status = InProgress
			todo.FinishedAt = time.Time{}
		}
	}

	if *add != "" {
		Todos = append(Todos, Todo{
			CreatedAt:  time.Now(),
			FinishedAt: time.Time{},
			Name:       *add,
			Status:     InProgress,
		})
	}

	if *deleteIndex != -1 && *deleteIndex < len(Todos) {
		Todos = append(Todos[:*deleteIndex], Todos[*deleteIndex+1:]...)
	}

	if *list {
		t := createTable()
		fmt.Println(t)
		return
	}

	if *clear {
		err := clearFile("todos.json")
		if err != nil {
			fmt.Println("Error clearing file:", err)
		} else {
			fmt.Println("All tasks cleared from todos.json")
		}
		return
	}

	t := createTable()
	fmt.Println(t)

	saveTodosToFile("todos.json", Todos)
}

// createTable creates a table with todos with predifined styles
func createTable() *table.Table {
	t := table.New().
		Border(lipgloss.NormalBorder()).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == table.HeaderRow {
				return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Align(lipgloss.Center).Bold(true)
			}
			return lipgloss.NewStyle().Foreground(lipgloss.Color("#dddddd")).Align(lipgloss.Center)
		}).
		Headers("INDEX", "TASK", "STATUS", "CREATED AT", "FINISHED AT")

	i := 0
	for _, v := range Todos {
		iStr := strconv.Itoa(i) // Convert i to string
		if v.FinishedAt.IsZero() {
			t.Row(iStr, v.Name, v.Status, v.CreatedAt.Format(time.RFC822), "")
		} else {
			t.Row(iStr, v.Name, v.Status, v.CreatedAt.Format(time.RFC822), v.FinishedAt.Format(time.RFC822))
		}
		i++
	}
	return t
}

// clearFile clears all content of a file
func clearFile(filename string) error {
	return os.Truncate(filename, 0)
}

// saveTodosToFile saves the Todos slice to a JSON file.
func saveTodosToFile(filename string, todos []Todo) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Make JSON pretty-printed
	err = encoder.Encode(todos)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println("Todos saved to", filename)
}

// readTodosFromFile reads the Todos from a JSON file and init the Todo slice with them
func readTodosFromFile(filename string, todos *[]Todo) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(data, todos)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}
