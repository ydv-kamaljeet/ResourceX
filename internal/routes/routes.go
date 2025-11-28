package routes

import (
	"book.com/internal/handlers"
	"book.com/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	book := r.Group("/books")
	book.Use(
		middleware.Logger(),
		middleware.ResponseMiddleware(),
	)
	{
		book.GET("/search", handlers.SearchBook)
		book.POST("/upload", handlers.UploadBook)

	}

	r.Static("/css", "./frontend/css")
	r.Static("/js", "./frontend/js")

	// HTML pages
	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/index.html")
	})
	r.GET("/upload", func(c *gin.Context) {
		c.File("./frontend/upload.html")
	})
	r.GET("/search", func(c *gin.Context) {
		c.File("./frontend/search.html")
	})
}
