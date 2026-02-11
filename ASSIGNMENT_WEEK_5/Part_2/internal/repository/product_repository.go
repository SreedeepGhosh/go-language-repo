package repository
import(
	"assignment_week_5/Part_2/internal/domain"
	"gorm.io/gorm"
)
type ProductRepository struct{
	db *gorm.DB
}
func NewProductRepository(db *gorm.DB) *ProductRepository{
	return &ProductRepository{db: db}
}
func (r *ProductRepository) GetAll() ([]domain.Product, error){
	var products []domain.Product
	result:=r.db.Find(&products)
	if result.RowsAffected==0{
		return nil, gorm.ErrRecordNotFound
	}
	return products, result.Error
}
func (r *ProductRepository) Create(p *domain.Product) error{
	return r.db.Create(p).Error
}
func (r *ProductRepository) Update(p *domain.Product) error{
	result:=r.db.Save(p)
	if result.RowsAffected==0{
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
func (r *ProductRepository) Delete(id string) error{
	result:=r.db.Delete(&domain.Product{}, "id = ?", id)
	if result.RowsAffected==0{
		return gorm.ErrRecordNotFound
	}
	return result.Error
}