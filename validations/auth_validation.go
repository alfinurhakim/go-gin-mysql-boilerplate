package validations

type LoginValidation struct {
	UserName string `json:"user_name" validate:"min=6,max=45" binding:"required"`
	Password string `json:"password" validate:"min=6,max=225" binding:"required"`
}

type RegisterValidation struct {
	Name     string `json:"name" validate:"min=6,max=45" binding:"required"`
	UserName string `json:"user_name" validate:"min=6,max=225" binding:"required"`
	Password string `json:"password" validate:"min=6,max=225" binding:"required"`
}
