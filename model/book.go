package model

type Book struct {
	UUID        string   `json:"UUID"`
	Name        string   `json:"name"`
	AuthorList  []string `json:"authorList"`
	PublishDate string   `json:"publishDate"`
	ISBN        string   `json:"ISBN"`
}
