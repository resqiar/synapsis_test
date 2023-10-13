package entities

type Category struct {
	ID       uint `gorm:"primaryKey"`
	Title    string
	Desc     string
	Products []Product `gorm:"many2many:product_categories;"`
}
