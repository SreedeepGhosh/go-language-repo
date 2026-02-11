package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	httpDelivery "assignment_week_5/Part_2/internal/delivery/http"
	"assignment_week_5/Part_2/internal/repository"
	"assignment_week_5/Part_2/internal/usecase"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=strongpassword dbname=inventory_db port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	productRepo := repository.NewProductRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	productUC := usecase.NewProductUsecase(productRepo)
	orderUC := usecase.NewOrderUsecase(orderRepo)

	handler := httpDelivery.NewHandler(productUC, orderUC)

	r := gin.Default()
	httpDelivery.RegisterRoutes(r, handler)

	r.Run(":8081")
}

