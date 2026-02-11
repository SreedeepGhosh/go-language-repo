package http
import(
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"assignment_week_5/Part_3/internal/domain"
	"assignment_week_5/Part_3/internal/usecase"
)
type Handler struct{
	productUC *usecase.ProductUsecase
	orderUC   *usecase.OrderUsecase
}
func NewHandler(p *usecase.ProductUsecase, o *usecase.OrderUsecase) *Handler{
	return &Handler{
		productUC: p,
		orderUC:   o,
	}
}
func (h *Handler) GetProducts(c *gin.Context){
	products, err := h.productUC.GetAll()
	if err != nil{
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "No products found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
func (h *Handler) CreateProduct(c *gin.Context){
	var p domain.Product
	if err := c.ShouldBindJSON(&p); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.productUC.Create(&p); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, p)
}
func (h *Handler) UpdateProduct(c *gin.Context){
	id := c.Param("id")
	var p domain.Product
	if err := c.ShouldBindJSON(&p); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parsedID, err := uuid.Parse(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	p.ID = parsedID
	if err := h.productUC.Update(&p); err != nil{
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, p)
}
func (h *Handler) DeleteProduct(c *gin.Context){
	id := c.Param("id")
	if err := h.productUC.Delete(id); err!=nil{
		if err == gorm.ErrRecordNotFound{
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
func (h *Handler) GetTasks(c *gin.Context){
	orders, err := h.orderUC.GetTasks()
	if err!=nil{
		if err == gorm.ErrRecordNotFound{
			c.JSON(http.StatusNotFound, gin.H{"error": "No tasks found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}
func (h *Handler) CreateTask(c *gin.Context){
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.orderUC.CreateTask(&order); err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	go sendNotificationEmail(order.ID.String())
	c.JSON(http.StatusCreated, order)
}
func sendNotificationEmail(orderID string){
	time.Sleep(2*time.Second)
	log.Println("Notification email sent for order:", orderID)
}