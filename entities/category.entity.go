package entities

type Category struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string `gorm:"unique"`
	Desc     string
	Products []Product `gorm:"many2many:product_categories;"`
}
