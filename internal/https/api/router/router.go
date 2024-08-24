package router

import (
	"Items/internal/storage"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() {
	handler := storage.Handler()

	r := gin.Default()

	r.POST("/items", handler.CreateItems)
	r.GET("/items/:id", handler.GetbyIDItems)
	r.GET("/items", handler.GetAllItems)
	r.PUT("/items/:id", handler.UpdateItems)
	r.DELETE("/items/:id", handler.DeleteItems)
	r.GET("/ping", handler.Greet)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.RunTLS(":9000", "./internal/tls/items.pem", "./internal/tls/items-key.pem")
	if err != nil {
		log.Fatal("Failed to run HTTPS server:", err)
	}
}
