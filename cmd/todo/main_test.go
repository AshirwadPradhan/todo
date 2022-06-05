package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName  = "todo"
	filename = "/tmp/todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building Tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running Tests...")
	result := m.Run()

	fmt.Println("Cleaning Up...")
	os.Remove(binName)
	os.Remove(filename)
	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "Test Task 1"
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	binPath := filepath.Join(dir, binName)
	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(binPath, "--task", task)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(binPath, "--list")
		out, err := cmd.Output()
		if err != nil {
			t.Fatal(err)
		}

		expected := task + "\n"
		if expected != string(out) {
			t.Errorf("Expected output %s, got %s instead", expected, string(out))
		}
	})
}
