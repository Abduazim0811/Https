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

// CreateItems godoc
// @Summary Create a new item
// @Description Create a new item in the items collection
// @Tags items
// @Accept  json
// @Produce  json
// @Param item body items.CreateItems true "Item to create"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /items [post]
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

// GetByIDItems godoc
// @Summary Get item by ID
// @Description Get a specific item by its ID
// @Tags items
// @Accept  json
// @Produce  json
// @Param id path string true "Item ID"
// @Success 200 {object} items.Items
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /items/{id} [get]
func (i *ItemsHandler) GetbyIDItems(c *gin.Context){
	id := c.Param("id")
	items, err := i.service.GetById(id)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// GetAllItems godoc
// @Summary Get all items
// @Description Retrieve all items from the items collection
// @Tags items
// @Accept  json
// @Produce  json
// @Success 200 {array} items.Items
// @Failure 500 {object} string
// @Router /items [get]
func (i *ItemsHandler) GetAllItems(c *gin.Context){
	items, err := i.service.Getall()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// UpdateItems godoc
// @Summary Update an item
// @Description Update an existing item by its ID
// @Tags items
// @Accept  json
// @Produce  json
// @Param id path string true "Item ID"
// @Param item body items.Items true "Updated item data"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /items/{id} [put]
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

// DeleteItems godoc
// @Summary Delete an item
// @Description Delete an item by its ID
// @Tags items
// @Accept  json
// @Produce  json
// @Param id path string true "Item ID"
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /items/{id} [delete]
func (i *ItemsHandler) DeleteItems(c *gin.Context){
	id := c.Param("id")
	
	err := i.service.Deleteitems(id)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Items deleted"})
}
