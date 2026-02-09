package main

import (
	"flag"
	"fmt"
	"os"

	"GTD/internal"
)

func printHelp() {
	fmt.Println(`
GTD - simple CLI todo manager

Usage:
.\gtd add    --title "task title"
.\gtd list
.\gtd done   --id <task_id>
.\gtd delete --id <task_id>
.\gtd help

Examples:
.\gtd add --title "learn golang"
.\gtd done --id 1
`)
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {

	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		title := addCmd.String("title", "", "title of the task")
		addCmd.Parse(os.Args[2:])

		if *title == "" {
			fmt.Println("Error: --title is required")
			return
		}

		store, err := internal.LoadStore()
		if err != nil {
			fmt.Println("Error loading store:", err)
			return
		}

		todo := internal.Todo{
			ID:    store.NextID,
			Title: *title,
			Done:  false,
		}

		store.Todos = append(store.Todos, todo)
		store.NextID++

		internal.SaveStore(store)
		fmt.Println("Task added successfully.")

	case "list":
		store, _ := internal.LoadStore()

		if len(store.Todos) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		for _, t := range store.Todos {
			status := " "
			if t.Done {
				status = "x"
			}
			fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Title)
		}

	case "done":
		doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
		id := doneCmd.Int("id", -1, "id of the task")
		doneCmd.Parse(os.Args[2:])

		if *id == -1 {
			fmt.Println("Error: --id is required")
			return
		}

		store, _ := internal.LoadStore()
		found := false

		for i := range store.Todos {
			if store.Todos[i].ID == *id {
				store.Todos[i].Done = true
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Error: task not found")
			return
		}

		internal.SaveStore(store)
		fmt.Println("Task marked as done.")

	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deleteCmd.Int("id", -1, "id of the task to delete")
		deleteCmd.Parse(os.Args[2:])

		if *id == -1 {
			fmt.Println("Error: --id is required")
			return
		}

		store, _ := internal.LoadStore()
		index := -1

		for i, t := range store.Todos {
			if t.ID == *id {
				index = i
				break
			}
		}

		if index == -1 {
			fmt.Println("Error: task not found")
			return
		}

		store.Todos = append(store.Todos[:index], store.Todos[index+1:]...)
		internal.SaveStore(store)

		fmt.Println("Task deleted successfully.")

	case "help", "--help", "-h":
		printHelp()

	default:
		fmt.Println("Unknown command:", os.Args[1])
		printHelp()
	}
}
