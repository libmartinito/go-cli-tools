package todo_test

import (
	"os"
	"testing"

	"github.com/libmartinito/go-cli-tools/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l.Items[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l.Items[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l.Items[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l.Items[0].Task)
	}

	if l.Items[0].Done {
		t.Errorf("New task should not be completed.")
	}

	l.Complete(1)

	if !l.Items[0].Done {
		t.Errorf("New task should be completed.")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	for _, v := range tasks {
		l.Add(v)
	}

	if l.Items[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q instead.", tasks[0], l.Items[0].Task)
	}

	l.Delete(2)

	if len(l.Items) != 2 {
		t.Errorf("Expected list length %d, got %d instead.", 2, len(l.Items))
	}

	if l.Items[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead.", tasks[2], l.Items[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)

	if l1.Items[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l1.Items[0].Task)
	}

	tf, err := os.CreateTemp("", "todo")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	if l1.Items[0].Task != l2.Items[0].Task {
		t.Errorf("Task %q should match %q task.", l1.Items[0].Task, l2.Items[0].Task)
	}
}
