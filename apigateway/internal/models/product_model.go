package models

type ProductWithoutID struct {
	Name        string  `json:"name"  validate:"required"`
	Description string  `json:"description"  validate:"required"`
	Price       float64 `json:"price"  validate:"required,gt=0,numeric"`
	Quantity    int32   `json:"quantity" validate:"required,gt=0,numeric"`
}

type ProductComplete struct {
	Name        string  `json:"name"  validate:"required"`
	Description string  `json:"description"  validate:"required"`
	Price       float64 `json:"price"  validate:"required,gt=0,numeric"`
	Quantity    int32   `json:"quantity" validate:"required,numeric"`
	IsActive    int32   `json:"is_active" validate:"required,gt=0,lt=3,numeric"`
}

type ProductStock struct {
	Quantity int32 `json:"quantity" validate:"gte=0,numeric"`
	Action   int32 `json:"action" validate:"required,gt=0,lt=4,numeric"`
}
