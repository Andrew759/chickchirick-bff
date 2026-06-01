package request

type CreateUserRequest struct {
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
	Login    string  `json:"login"`
	Phone    string  `json:"phone"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}
