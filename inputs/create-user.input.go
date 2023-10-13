package inputs

type CreateUserInput struct {
	Username string `validate:"required,min=3,max=100"`
	Password string `validate:"required,min=8,max=100"`
}
