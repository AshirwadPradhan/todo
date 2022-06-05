package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AshirwadPradhan/todo"
)

const filename = "/tmp/todo.json"

func main() {

	task := flag.String("task", "", "Task to be included in the todo list")
	list := flag.Bool("list", false, "List All Tasks")
	complete := flag.Int("complete", 0, "Items to be completed")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Starter Todo App\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2022\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	flag.Parse()
	todos := &todo.TodoList{}

	if err := todos.Get(filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		for _, t := range *todos {
			if !t.IsCompleted {
				fmt.Println(t.Task)
			}
		}
	case *complete > 0:
		if err := todos.MarkComplete(*complete - 1); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := todos.Save(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		todos.Add(*task)
		if err := todos.Save(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

}
