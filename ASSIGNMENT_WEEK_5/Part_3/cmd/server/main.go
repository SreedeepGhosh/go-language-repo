package main
import (
	"log"
	"github.com/gin-gonic/gin"
	"assignment_week_5/Part_3/internal/database"
	"assignment_week_5/Part_3/internal/delivery/http"
	"assignment_week_5/Part_3/internal/domain"
	"assignment_week_5/Part_3/internal/repository"
	"assignment_week_5/Part_3/internal/usecase"
)
func main() {
	db := database.ConnectDB()
	err := db.AutoMigrate(
		&domain.Product{},
		&domain.Order{},
		&domain.Category{},
	)
	if err != nil {
		log.Fatal(err)
	}
	productRepo := repository.NewProductRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	productUC := usecase.NewProductUsecase(productRepo)
	orderUC := usecase.NewOrderUsecase(orderRepo)
	handler := http.NewHandler(productUC, orderUC)
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/products", handler.GetProducts)
		api.POST("/products", handler.CreateProduct)
		api.PUT("/products/:id", handler.UpdateProduct)
		api.DELETE("/products/:id", handler.DeleteProduct)
		api.GET("/tasks", handler.GetTasks)
		api.POST("/tasks", handler.CreateTask)
	}
	log.Println("Server running on port 8080")
	r.Run(":8080")
}