package request

type BookCreateRequest struct {
	Title  string `validate:"required" json:"title"`
	Author string `validate:"required,min=3" json:"author"`
	Cover  string `json:"cover"`
}
