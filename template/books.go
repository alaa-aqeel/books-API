package template

type Book struct {
	Id  int64  `json:"id"`
	Author string `json:"author"`
	Language string `json:"language"`
	Link string `json:"link"`
	Pages int64 `json:"pages"`
	Title string `json:"title"`
}
type Books struct {
	Books []Book
}
