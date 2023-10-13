package entities

type CartItem struct {
	ID        uint `gorm:"primaryKey"`
	CartID    uint // FK of the owner (Cart)
	ProductID uint // FK of the owner (Product)
	Quantity  int
}
