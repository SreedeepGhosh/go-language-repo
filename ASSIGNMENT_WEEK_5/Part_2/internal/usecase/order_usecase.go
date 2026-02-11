package usecase
import (
	"assignment_week_5/Part_2/internal/domain"
	"assignment_week_5/Part_2/internal/repository"
)
type OrderUsecase struct{
	repo *repository.OrderRepository
}
func NewOrderUsecase(r *repository.OrderRepository) *OrderUsecase{
	return &OrderUsecase{repo: r}
}
func (u *OrderUsecase) GetTasks() ([]domain.Order, error){
	return u.repo.GetTasks()
}