package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type todoItem struct {
	Task        string
	IsCompleted bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type TodoList []todoItem

func (t *TodoList) Add(task string) {
	todoItem := todoItem{
		Task:        task,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todoItem)
}

func (t *TodoList) MarkComplete(pos int) error {
	// Here pos is 0-based index position
	if pos < 0 || pos >= len(*t) {
		return fmt.Errorf("the selected item at position %d does not exists", pos)
	}

	(*t)[pos].IsCompleted = true
	(*t)[pos].CompletedAt = time.Now()

	return nil
}

func (t *TodoList) Delete(pos int) error {
	// Here pos is 0-based index position
	if pos < 0 || pos >= len(*t) {
		return fmt.Errorf("the selected item at position %d does not exists", pos)
	}
	if pos == 0 {
		*t = (*t)[1:]
	}

	if pos == len(*t)-1 {
		*t = (*t)[:len(*t)-1]
	}
	*t = append((*t)[:pos], (*t)[pos+1:]...)
	return nil
}

func (t *TodoList) Save(filename string) error {
	js, err := json.Marshal(t)
	if err != nil {
		return fmt.Errorf("error while marshalling the todo items to json: %v", err)
	}
	err = ioutil.WriteFile(filename, js, 0644)
	if err != nil {
		return fmt.Errorf("error while saving the file to disk: %v", err)
	}
	return nil
}

func (t *TodoList) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return fmt.Errorf("error while reading the file from disk: %v", err)
	}
	if len(file) == 0 {
		return nil
	}
	if err := json.Unmarshal(file, t); err != nil {
		return fmt.Errorf("error while unmarshaling todo items from json: %v", err)
	}
	return nil
}

// Implements Stringer interface
func (t *TodoList) String() string {
	out := ""

	for i, item := range *t {
		prefix := " "
		if item.IsCompleted {
			prefix = "X"
		}
		out += fmt.Sprintf("%s %d: %s\n", prefix, i+1, item.Task)
	}
	return out
}

func (t *TodoList) VerbosePrint() {
	fmt.Printf("%4s %1s: %20s \t %25s \t %25s\n", "Mark", "#", "Task", "CreatedAt", "CompletedAt")
	for i, item := range *t {
		prefix := " "
		if item.IsCompleted {
			prefix = "X"
		}
		fmt.Printf("%4s %1d: %20s \t %15s \t %15s\n", prefix, i+1, item.Task, item.CreatedAt, item.CompletedAt)
	}
}
