package models

type Product struct {
	ID                string          `gorm:"primaryKey"`
	ProductCategoryId string          `gorm:"column:category_id;size:36;not null"`
	ProductCategory   ProductCategory `gorm:"foreignKey:ProductCategoryId"`
	ProductName       string          `gorm:"column:product_name;size:100;not null"`
}

func (p *Product) TableName() string {
	return "m_product"
}
