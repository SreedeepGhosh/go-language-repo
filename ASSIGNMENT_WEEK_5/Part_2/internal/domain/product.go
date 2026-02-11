package domain
import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)
type Product struct{
	ID  uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	ProductName string    `json:"product_name"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CategoryID 	uuid.UUID `gorm:"type:uuid" json:"category_id"`
}
func (p *Product) BeforeCreate(tx *gorm.DB)error{
	if p.ID==uuid.Nil{
		p.ID=uuid.New()
	}
	return nil
}