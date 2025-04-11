package responses

type UpdateTaskResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
}