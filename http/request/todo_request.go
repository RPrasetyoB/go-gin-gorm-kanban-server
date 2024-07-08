package request

type TodoRequest struct {
	Title       string `validate:"required" json:"title"`
	Description string `validate:"required" json:"description"`
}
