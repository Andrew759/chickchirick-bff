package request

type CreateUserRequest struct {
	Name     string
	Surname  string
	Login    string
	Phone    *string
	Email    *string
	Password *string
}
