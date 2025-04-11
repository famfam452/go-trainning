package requests

type UpdateTaskRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
}