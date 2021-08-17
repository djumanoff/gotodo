package todo

type Todo struct {
	ID      int64  `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Status  string `json:"status" db:"status"`
	Body    string `json:"body" db:"body"`
	OwnerID string `json:"owner_id" db:"owner_id"`
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
