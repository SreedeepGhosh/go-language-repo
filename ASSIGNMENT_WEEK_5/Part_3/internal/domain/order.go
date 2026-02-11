package domain
import(
	"github.com/google/uuid"
	"gorm.io/gorm"
)
type Order struct{
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	ProductID uuid.UUID `json:"product_id" gorm:"type:uuid;not null"`
	OrderDate string    `json:"order_date"`
}
func (o *Order) BeforeCreate(tx *gorm.DB) error{
	if o.ID == uuid.Nil {
		o.ID = uuid.New()
	}
	return nil
}