package todo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/AshirwadPradhan/todo"
)

func TestAdd(t *testing.T) {
	todos := todo.TodoList{}

	taskName := "New Task Test"
	todos.Add(taskName)

	if todos[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, todos[0].Task)
	}
}

func TestMarkComplete(t *testing.T) {
	todos := todo.TodoList{}

	taskName := "New Task Test"
	todos.Add(taskName)

	if todos[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, todos[0].Task)
	}

	if todos[0].IsCompleted {
		t.Errorf("Task should not have been completed")
	}

	todos.MarkComplete(0)
	if !todos[0].IsCompleted {
		t.Errorf("Task should be completed")
	}
}

func TestDelete(t *testing.T) {
	todos := todo.TodoList{}

	todoItems := []string{
		"Test Todo 1",
		"Test Todo 2",
		"Test Todo 3",
		"Test Todo 4",
		"Test Todo 5",
	}

	for _, v := range todoItems {
		todos.Add(v)
	}

	if todos[0].Task != todoItems[0] {
		t.Errorf("Expected %q, got %q instead", todoItems[0], todos[0].Task)
	}

	todos.Delete(2)
	var rem int = 4
	if len(todos) != rem {
		t.Errorf("Expected length after deletion %d, got %d instead", rem, len(todos))
	}

	todos.Delete(1)
	rem = 3
	if len(todos) != rem {
		t.Errorf("Expected length after deletion %d, got %d instead", rem, len(todos))
	}

	todos.Delete(3)
	if len(todos) != rem {
		t.Errorf("Expected length after deletion %d, got %d instead", rem, len(todos))
	}
}

func TestSaveGet(t *testing.T) {
	t1 := todo.TodoList{}
	t2 := todo.TodoList{}

	taskName := "New Task Test"
	t1.Add(taskName)
	if t1[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, t1[0].Task)
	}

	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("error creating tempfile: %s", err)
	}
	defer os.Remove(tf.Name())

	if err := t1.Save(tf.Name()); err != nil {
		t.Fatalf("error saving todos in file %s", err)
	}

	if err := t2.Get(tf.Name()); err != nil {
		t.Fatalf("error reading todos from file %s", err)
	}

	if t1[0].Task != t2[0].Task {
		t.Errorf("Expected %q, got %q instead", t1[0].Task, t2[0].Task)
	}
}
