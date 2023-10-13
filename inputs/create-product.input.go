package inputs

type CreateProductInput struct {
	Title       string `validate:"required,min=3,max=100"`
	Description string `validate:"omitempty,max=100"`
	ImageURL    string `validate:"omitempty,url"`
}
