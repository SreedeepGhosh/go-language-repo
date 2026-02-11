package main
import (
	"log"
	"time"
	"github.com/gin-gonic/gin"
)
func LoggerMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		start:=time.Now()
		c.Next()
		log.Printf("%s %s took %v", c.Request.Method, c.Request.URL.Path, time.Since(start))
	}
}
func CORSMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}