package main
import (
	"log"
	"time"
)
func processBookInBackground(bookID int){
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("Background task completed for book:", bookID)
	}()
}