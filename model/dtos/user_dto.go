package dtos

type UserRegisterRequest struct {
	Email    string `validate:"required,email,max=100" json:"email"`
	Password string `validate:"required,min=4,max=100" json:"password"`
	Name     string `validate:"required,min=1,max=100" json:"name"`
}

type UserLoginRequest struct {
	Email    string `validate:"required,email,max=100" json:"email"`
	Password string `validate:"required,min=8,max=100" json:"password"`
}

type UserResponse struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
