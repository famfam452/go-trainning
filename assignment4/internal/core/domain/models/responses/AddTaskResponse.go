package responses

type AddTaskResponse struct {
	ID uint `json:"id"`
	Completed bool `json:"completed"`
}