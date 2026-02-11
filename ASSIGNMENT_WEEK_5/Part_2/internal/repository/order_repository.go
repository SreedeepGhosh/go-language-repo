package repository
import(
	"assignment_week_5/Part_2/internal/domain"
	"gorm.io/gorm"
)
type OrderRepository struct{
	db *gorm.DB
}
func NewOrderRepository(db *gorm.DB) *OrderRepository{
	return &OrderRepository{db: db}
}
func (r *OrderRepository) GetTasks() ([]domain.Order, error){
	var orders []domain.Order
	result := r.db.Find(&orders)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return orders, result.Error
}