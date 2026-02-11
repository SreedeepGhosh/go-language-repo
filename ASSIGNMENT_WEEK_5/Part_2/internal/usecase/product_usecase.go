package usecase
import(
	"assignment_week_5/Part_2/internal/domain"
	"assignment_week_5/Part_2/internal/repository"
)
type ProductUsecase struct{
	repo *repository.ProductRepository
}
func NewProductUsecase(r *repository.ProductRepository) *ProductUsecase{
	return &ProductUsecase{repo: r}
}
func (u *ProductUsecase) GetAll() ([]domain.Product, error){
	return u.repo.GetAll()
}
func (u *ProductUsecase) Create(p *domain.Product) error{
	return u.repo.Create(p)
}
func (u *ProductUsecase) Update(p *domain.Product) error{
	return u.repo.Update(p)
}
func (u *ProductUsecase) Delete(id string) error{
	return u.repo.Delete(id)
}