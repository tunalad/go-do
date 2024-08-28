package utils

import (
	"bufio"
	"fmt"
	"go-do/task"
	"os"
	"strconv"
	"strings"
	//"github.com/adrg/xdg"
)

func CsvToArray(path string) []task.Task {
	tasks := []task.Task{}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening the csv")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		if len(parts) >= 3 {
			title := parts[0]
			status, _ := strconv.Atoi(parts[1])
			priority, _ := strconv.ParseBool(parts[2])

			tasks = append(tasks, task.NewTask(title, task.Status(status), priority))
		}
	}

	return tasks
}

func ArrayToCsv(tasks []task.Task, path string) {
	// opening the file
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error opening the CSV:", err)
		return
	}
	defer file.Close()

	// creating a writer for the file
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// array to csv
	for _, t := range tasks {
		newRow := fmt.Sprintf("%s,%d,%t\n", t.Title, t.Status, t.Priority)

		_, err = writer.WriteString(newRow)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

func RemoveFromCsv(tasks []task.Task, itemIndex int) []task.Task {
	new := make([]task.Task, 0)
	new = append(new, tasks[:itemIndex]...)
	return append(new, tasks[itemIndex+1:]...)
}
