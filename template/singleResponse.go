package template

type SingleResponse struct {
	Book Book `json:"book"`
	Error string `json:"message"`
}