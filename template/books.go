package template

type Book struct {
	Id  int64  `json:"id,omitempty"`
	Author string `json:"author,omitempty"`
	Language string `json:"language,omitempty"`
	Link string `json:"link,omitempty"`
	Pages int64 `json:"pages,omitempty"`
	Title string `json:"title,omitempty"`
}

type Books struct {
	Books []Book
}
