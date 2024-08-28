package main

import (
	"flag"
	"fmt"
	"go-do/task"
	"go-do/utils"
	"strconv"
	"strings"
	//"github.com/adrg/xdg"
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

	tasks := utils.CsvToArray("./tasks.csv")

	flag.Parse()
	if len(newTask) > 0 {
		tasks = append(tasks, task.NewTask(newTask, task.TODO, false))
		printTasks(tasks)
		utils.ArrayToCsv(tasks, "./tasks.csv")
	} else if itemToRemove >= 0 {
		tasks = utils.RemoveFromCsv(tasks, itemToRemove)
		printTasks(tasks)
		utils.ArrayToCsv(tasks, "./tasks.csv")
	} else if markTodo >= 0 {
		tasks[markTodo].SetStatus(task.TODO)
		printTasks(tasks)
		utils.ArrayToCsv(tasks, "./tasks.csv")
	} else if markDoing >= 0 {
		tasks[markDoing].SetStatus(task.DOING)
		printTasks(tasks)
		utils.ArrayToCsv(tasks, "./tasks.csv")
	} else if markDone >= 0 {
		tasks[markDone].SetStatus(task.DONE)
		printTasks(tasks)
		utils.ArrayToCsv(tasks, "./tasks.csv")
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
		utils.ArrayToCsv(tasks, "./tasks.csv")
	} else {
		printTasks(tasks)
	}
}
