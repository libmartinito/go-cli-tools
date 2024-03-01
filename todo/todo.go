package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List struct {
	Items         []item
	Verbose       bool
	HideCompleted bool
}

func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	l.Items = append(l.Items, t)
}

func (l *List) Complete(i int) error {
	if i <= 0 || i > len(l.Items) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	l.Items[i-1].Done = true
	l.Items[i-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Delete(i int) error {
	if i <= 0 || i > len(l.Items) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	l.Items = append(l.Items[:i-1], l.Items[i:]...)

	return nil
}

func (l *List) Save(filename string) error {
	js, err := json.Marshal(l.Items)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, &l.Items)
}

func (l *List) String() string {
	formatted := ""

	for k, t := range l.Items {
		prefix := "  "

		if l.HideCompleted && t.Done {
			continue
		}

		if t.Done {
			prefix = "X "
		}

		formatted += fmt.Sprintf("%s%d: %s", prefix, k+1, t.Task)

		if l.Verbose && t.Done {
			formatted += fmt.Sprintf("  Completed: %v\n", t.CompletedAt.Format("2006-01-02 15:04:05"))
		}

		if l.Verbose {
			formatted += fmt.Sprintf("  Created: %v\n", t.CreatedAt.Format("2006-01-02 15:04:05"))
		}

		if !l.Verbose {
			formatted += "\n"
		}
	}

	return formatted
}
