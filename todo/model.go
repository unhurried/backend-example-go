package todo

type Todo struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Content  string `json:"content"`
}
