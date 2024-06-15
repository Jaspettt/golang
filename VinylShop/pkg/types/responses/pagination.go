package responses

type Pagination struct {
	Content interface{} `json:"content"`
	Next    string      `json:"next"`
	Prev    string      `json:"prev"`
}
