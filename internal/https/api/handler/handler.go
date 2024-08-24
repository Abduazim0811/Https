package handler

import (
	"Items/internal/entity/items"
	"Items/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemsHandler struct {
	service *service.ItemsService
}

func NewService(service *service.ItemsService) *ItemsHandler {
	return &ItemsHandler{service: service}
}

func (i *ItemsHandler) CreateItems(c *gin.Context) {
	var req items.CreateItems
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := i.service.Create(req)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Itesm created"})
}

func (i *ItemsHandler) GetbyIDItems(c *gin.Context){
	id := c.Param("id")
	items, err := i.service.GetById(id)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

func (i *ItemsHandler) GetAllItems(c *gin.Context){
	items, err := i.service.Getall()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

func (i *ItemsHandler) UpdateItems(c *gin.Context){
	id := c.Param("id")
	var req items.Items
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	req.Id = id
	err := i.service.Updateitems(id, req)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Updated successfully"})
}

func (i *ItemsHandler) DeleteItems(c *gin.Context){
	id := c.Param("id")
	
	err := i.service.Deleteitems(id)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Items deleted"})
}
