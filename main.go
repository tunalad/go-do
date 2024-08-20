package main

import (
	"bufio"
	"flag"
	"fmt"
	"go-do/task"
	"os"
	"strconv"
	"strings"
	//"github.com/adrg/xdg"
	//"github.com/cheynewallace/tabby"
)

func printTasks(tasks []task.Task) {
	fmt.Println(strings.Repeat("-", 55))
	fmt.Println("| ID |", strings.Repeat(" ", 16), "Task", strings.Repeat(" ", 16), "| State |")
	fmt.Println(strings.Repeat("-", 55))

	if len(tasks) < 1 {
		fmt.Println("|    |", strings.Repeat(" ", 16), "None", strings.Repeat(" ", 16), "|       |")
	} else {
		for i := 0; i < len(tasks); i++ {
			printTask(i, tasks[i])
		}
	}

	// tasksPriority := task.GetPriority(tasks, false)
	// tasksLessPriority := task.GetPriority(tasks, true)

	//for i := 0; i < len(tasksPriority); i++ {
	//	printTask(i, tasksPriority[i])
	//}

	// fmt.Println(strings.Repeat("-", 50))

	//for i := 0; i < len(tasksLessPriority); i++ {
	//	printTask(i, tasksLessPriority[i])
	//}

	fmt.Println(strings.Repeat("-", 55))
}

func printTask(id int, t task.Task) {
	icon := task.StatusIcons[t.Status]
	stringLength := len(t.Title)

	if stringLength < 37 {
		unusedSpace := 37 - stringLength
		padLeft := unusedSpace / 2
		padRight := padLeft

		indexSpace := 4
		indexChars := len(strconv.FormatInt(int64(id), 10))
		indexSpaceLeft := indexSpace - indexChars

		if padLeft+padRight+len(t.Title) > 36 {
			padRight = padRight - 1
		}

		fmt.Println("|"+strconv.FormatInt(int64(id), 10)+strings.Repeat(" ", indexSpaceLeft)+"|", strings.Repeat(" ", padLeft), t.Title, strings.Repeat(" ", padRight), "|  ", icon, "  |")
	} else {
		printLongTask(id, t)
	}
}

func printLongTask(id int, t task.Task) {
	maxLen := 36
	icon := task.StatusIcons[t.Status]

	for len(t.Title) > maxLen {
		fmt.Println("|", id, "|", t.Title[:maxLen], "|  ", icon, "  |")
		t.Title = t.Title[maxLen:]
	}

	unusedSpace := maxLen - len(t.Title)
	fmt.Println("|    |", t.Title, strings.Repeat(" ", unusedSpace), "|  ", icon, "  |")
}

func csvToArray(path string) []task.Task {
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

func addToCsv(path string, newItem string, status task.Status, priority bool) {
	// old ooga booga monkey implementation
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening the csv:", err)
		return
	}
	defer file.Close()

	// Construct the new row as a comma-separated string
	newRow := fmt.Sprintf("%s,%d,%t\n", newItem, status, priority)

	// Create a new writer
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Write the new row to the CSV file
	_, err = writer.WriteString(newRow)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Item appended successfully.")
}

func removeFromCsv(tasks []task.Task, itemIndex int) []task.Task {
	new := make([]task.Task, 0)
	new = append(new, tasks[:itemIndex]...)
	return append(new, tasks[itemIndex+1:]...)
}

func arrayToCsv(tasks []task.Task, path string) {
	// opening the file
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error opening the csv:", err)
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

func main() {
	var newTask string
	var itemToRemove int
	var markDone int
	var markTodo int
	var markDoing int
	var editTask string

	flag.StringVar(&newTask, "add", "", "Add new task")
	flag.StringVar(&editTask, "edit", "", "Edit existing task")
	flag.IntVar(&itemToRemove, "remove", -1, "Remove item from the list")

	flag.IntVar(&markTodo, "todo", -1, "Mark task as 'TODO'")
	flag.IntVar(&markDoing, "doing", -1, "Mark task as 'DOING'")
	flag.IntVar(&markDone, "done", -1, "Mark task as 'DONE'")

	tasks := []task.Task{
		// task.NewTask("wake up", task.DONE, true),
		// task.NewTask("do dishes", task.TODO, false),
		// task.NewTask("make music", task.TODO, false),
		// task.NewTask("code a bit", task.DOING, false),
		// task.NewTask("procrastinate", task.DOING, false),
		// task.NewTask("eat the breakfast", task.DONE, false),
		// task.NewTask("ksdasdasdajsdlk jaslkd sdsdjalksjdls", task.DONE, false),
	}

	tasks = csvToArray("./tasks.csv")

	flag.Parse()
	if len(newTask) > 0 {
		// addToCsv("./tasks.csv", "{{.newTask}},{{.task.TODO}},false")
		// addToCsv("./tasks.csv", newTask, task.TODO, false)
		// tasks = csvToArray("./tasks.csv")
		tasks = append(tasks, task.NewTask(newTask, task.TODO, false))
		printTasks(tasks)
		arrayToCsv(tasks, "./tasks.csv")
	} else if itemToRemove >= 0 {
		tasks = removeFromCsv(tasks, itemToRemove)
		printTasks(tasks)
		arrayToCsv(tasks, "./tasks.csv")
	} else if markTodo >= 0 {
		tasks[markTodo].SetStatus(task.TODO)
		printTasks(tasks)
		arrayToCsv(tasks, "./tasks.csv")
	} else if markDoing >= 0 {
		tasks[markDoing].SetStatus(task.DOING)
		printTasks(tasks)
		arrayToCsv(tasks, "./tasks.csv")
	} else if markDone >= 0 {
		tasks[markDone].SetStatus(task.DONE)
		printTasks(tasks)
		arrayToCsv(tasks, "./tasks.csv")
	} else if editTask != "" {
		index, _ := strconv.Atoi(editTask)
		args := flag.Args()

		fmt.Println(args)

		if len(args) != 1 {
			fmt.Println("Error: Invalid format for -edit argument. Usage: -edit <index> '<new value>'")
			return
		}

		newValue := args[0]
		tasks[index].SetTitle(newValue)
		printTasks(tasks)
		arrayToCsv(tasks, "./tasks.csv")
	} else {
		printTasks(tasks)
	}
}
