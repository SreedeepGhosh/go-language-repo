package main
import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)
var booksDB []Book
func addBook(c *gin.Context){
	var book Book
	if err:=c.ShouldBindJSON(&book); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	if err:=book.Validate(); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	for _, b:= range booksDB{
		if b.ID==book.ID{
			c.JSON(http.StatusBadRequest, gin.H{"error":"Book already exists"})
			return
		}
	}
	booksDB=append(booksDB, book)
	processBookInBackground(book.ID)
	c.JSON(http.StatusCreated, book)
}
func getAllBooks(c *gin.Context){
	c.JSON(http.StatusOK, booksDB)
}
func getBookByID(c *gin.Context){
	id, _:=strconv.Atoi(c.Param("id"))
	for _, book:= range booksDB{
		if book.ID==id{
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error":"Book not found"})
}
func updateBook(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedBook Book
	c.ShouldBindJSON(&updatedBook)
	for i, book := range booksDB{
		if book.ID==id{
			booksDB[i]=updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
func deleteBook(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	for i, book := range booksDB {
		if book.ID == id{
			booksDB = append(booksDB[:i], booksDB[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
