package inputs

type CreateCategoryInput struct {
	Title       string `validate:"required,min=3,max=100"`
	Description string `validate:"omitempty,max=100"`
}
