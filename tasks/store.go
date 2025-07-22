package tasks

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"
)

const dataFile = "tasks.csv"

func loadFile(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}
	return f, nil
}

func CloseFile(f *os.File) error {
	_ = syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}

func AddTask(newTask Task) error {
	file, err := loadFile(dataFile)
	if err != nil {
		return err
	}
	defer CloseFile(file)

	tasks, err := LoadTasksFromFile(file)
	if err != nil {
		return err
	}

	// Assign ID
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	newTask.ID = maxID + 1

	tasks = append(tasks, newTask)

	return SaveTasksToFile(file, tasks)
}

func LoadTasksFromFile(file *os.File) ([]Task, error) {
	_, _ = file.Seek(0, 0)
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return []Task{}, nil
	}

	tasks := []Task{}
	for i, r := range records {
		if i == 0 {
			continue
		}
		id, _ := strconv.Atoi(r[0])
		createdAt, _ := time.Parse(time.RFC3339, r[2])
		isComplete, _ := strconv.ParseBool(r[3])
		tasks = append(tasks, Task{
			ID:          id,
			Description: r[1],
			CreatedAt:   createdAt,
			IsComplete:  isComplete,
		})
	}
	return tasks, nil
}

func SaveTasksToFile(file *os.File, tasks []Task) error {
	_ = file.Truncate(0)
	_, _ = file.Seek(0, 0)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Description", "CreatedAt", "IsComplete"})

	for _, t := range tasks {
		record := []string{
			strconv.Itoa(t.ID),
			t.Description,
			t.CreatedAt.Format(time.RFC3339),
			strconv.FormatBool(t.IsComplete),
		}
		writer.Write(record)
	}

	return writer.Error()
}

func OpenFileForRead() (*os.File, error) {
	return os.OpenFile(dataFile, os.O_RDONLY|os.O_CREATE, os.ModePerm)
}

func OpenFileForReadWrite() (*os.File, error) {
	f, err := os.OpenFile(dataFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}
	return f, nil
}
