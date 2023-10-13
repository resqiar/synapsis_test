package entities

type Cart struct {
	ID        uint       `gorm:"primaryKey"`
	UserID    uint       // FK of the owner (user)
	CartItems []CartItem // one-to-many to cart items
}
