package domain

type Task struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
}

func (t *Task) GetId() uint {
	return t.ID
}

func (t *Task) GetTitle() string {
	return t.Title
}

func (t *Task) GetDescription() string {
	return t.Description
}

func (t *Task) GetCompleted() bool {
	return t.Completed
}