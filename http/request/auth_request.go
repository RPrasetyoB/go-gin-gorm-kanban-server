package request

type CreateRequest struct {
	Username string `validate:"required,min=3,max=200" json:"username"`
	Password string `validate:"required,min=6" json:"password"`
}

type LoginRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}
