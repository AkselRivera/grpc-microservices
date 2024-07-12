package models

type selectedProduct struct {
	Id       string `json:"id" validate:"required"`
	Quantity int32  `json:"quantity" validate:"required,gt=0,numeric"`
}

type Order struct {
	Products []selectedProduct `json:"products" validate:"required,min=1"`
}

type ChangeOrderStatus struct {
	Status int32 `json:"status" validate:"required,gt=0,lt=3,numeric"`
}
