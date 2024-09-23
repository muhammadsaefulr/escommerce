package entity

type ShoppingCartItems struct {
	ID        int           `gorm:"primary_key"`
	CartId    string        `json:"cart_id" validate:"required"`
	ProductId string        `json:"product_id" validate:"required"`
	Product   *ProductItems `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Quantity  int
}
