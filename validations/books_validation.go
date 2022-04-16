package validations

type BooksValidation struct {
	BookNumber      string `json:"book_number" validate:"min=1" binding:"required"`
	BookTitle       string `json:"book_title" validate:"min=1" binding:"required"`
	Author          string `json:"author" validate:"min=1" binding:"required"`
	PublicationYear int64  `json:"publication_year" validate:"min=1" binding:"required"`
	Publisher       string `json:"publisher" validate:"min=1" binding:"required"`
}
