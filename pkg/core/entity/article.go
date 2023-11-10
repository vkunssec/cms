package entity

type Article struct {
	Id      string `json:"_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
}
