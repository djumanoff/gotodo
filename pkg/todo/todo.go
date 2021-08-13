package todo

type Todo struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Status  string `json:"status"`
	Body    string `json:"body"`
	OwnerID string `json:"owner_id"`
}

type TodoUpdate struct {
	ID     int64   `json:"-"`
	Title  *string `json:"title"`
	Status *string `json:"status"`
	Body   *string `json:"body"`
}

type TodoQuery struct {
	OwnerID string `bson:"owner_id"`
	Status  string `bson:"status"`
}
