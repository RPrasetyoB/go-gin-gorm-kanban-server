package request

type ItemRequest struct {
	Todo_id             int    `validate:"required" json:"todo_id"`
	Name                string `validate:"required" json:"name"`
	Progress_percentage int    `validate:"required" json:"progress_percentage"`
}
