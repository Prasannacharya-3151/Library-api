package routes

import (
	"library-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		books := api.Group("/books")
		{
			books.POST("", handlers.CreateBook)
			books.GET("", handlers.GetAllBooks)
			books.GET("/:id", handlers.GetBookByID)
			books.PUT("/:id", handlers.UpdateBook)
			books.DELETE("/:id", handlers.DeleteBook)
		}
	}
}