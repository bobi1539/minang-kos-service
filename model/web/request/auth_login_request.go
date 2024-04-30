package request

type AuthLoginRequest struct {
	Email    string `validate:"required,max=255,min=1" json:"email"`
	Password string `validate:"required,max=255,min=1" json:"password"`
}
