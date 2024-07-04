package request

type UpdateUserRequest struct {
	Id       int    `validate:"required"`
	Username string `validate:"required, max=200 min=3" json:"username"`
	Password string `validate:"required, min=6" json:"password"`
}
