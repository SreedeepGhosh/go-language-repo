package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
	r.GET("/products", h.GetProducts)
	r.POST("/products", h.CreateProduct)
	r.PUT("/products/:id", h.UpdateProduct)
	r.DELETE("/products/:id", h.DeleteProduct)
	r.GET("/tasks", h.GetTasks)
}
