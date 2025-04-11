package requests

type AddTaskRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
}