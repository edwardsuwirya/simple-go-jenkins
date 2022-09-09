package models

type ProductCategory struct {
	ID           string `gorm:"primaryKey"`
	CategoryName string `gorm:"column:category_name;size:100;not null"`
	Products     []Product
}

func (p *ProductCategory) TableName() string {
	return "m_product_category"
}
