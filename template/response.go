package template

type Response struct {
	Books []Book `json:"books"`
	Error string `json:"message"`
}
