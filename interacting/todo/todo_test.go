package todo_test

import (
	"os"
	"testing"

	"github.com/t-ober/go-clis/interacting/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %s, got %s", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %s, got %s", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("Task should not be done")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("Task should be done")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	for _, task := range tasks {
		l.Add(task)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q instead", tasks[0], l[0].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"

	l1.Add(taskName)
	if l1[0].Task != taskName {
		t.Errorf("Expected %s, got %s", taskName, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Errorf("Error creating temp file")
	}

	l1.Save(tf.Name())
	l2.Get(tf.Name())

	if l1[0].Task != l2[0].Task {
		t.Errorf("Task mismatch after reading from file")
	}
}
