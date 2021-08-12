package todo

type Todo struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
}
