package domain
import "github.com/google/uuid"
type Order struct{
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	ProductID uuid.UUID `json:"product_id"`
	OrderDate string    `json:"order_date"`
}
