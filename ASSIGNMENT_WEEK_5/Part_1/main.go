package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New()
	router.Use(LoggerMiddleware())
	router.Use(CORSMiddleware())
	router.POST("/books", addBook)
	router.GET("/books", getAllBooks)
	router.GET("/books/:id", getBookByID)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)
	router.Run(":8080")
}