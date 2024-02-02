package request

type UserCreateRequest struct {
	Name     string `validate:"required,min=3" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=5" json:"password"`
	Address  string `validate:"min=3" json:"address"`
	Phone    string `validate:"max=12,min=12" json:"phone"`
}

type UserUpdateRequest struct {
	Name    string `validate:"required,min=3" json:"name"`
	Address string `validate:"min=3" json:"address"`
	Phone   string `validate:"max=12,min=12" json:"phone"`
}

type UserEmailRequest struct {
	Email string `validate:"required,email" json:"email"`
}
