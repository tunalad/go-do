package task

const (
	TODO  Status = iota
	DOING Status = iota
	DONE  Status = iota
)

type Status int

type Task struct {
	Title    string
	Status   Status
	Priority bool
}

var StatusIcons = map[Status]string{
	TODO:  " ",
	DOING: "~",
	DONE:  "×",
}

// constructor \\

func NewTask(title string, status Status, priority bool) Task {
	return Task{
		title, status, priority,
	}
}

// setters \\

func (t *Task) SetTitle(newTitle string) {
	t.Title = newTitle
}

func (t *Task) SetStatus(newStatus Status) {
	t.Status = newStatus
}

func (t *Task) SetPriority(newPriority bool) {
	t.Priority = newPriority
}

// functions \\

func GetPriority(tasks []Task, noPriorityOnly bool) []Task {
	var filteredTasks []Task

	for _, task := range tasks {
		if (noPriorityOnly && !task.Priority) || (!noPriorityOnly && task.Priority) {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks
}
