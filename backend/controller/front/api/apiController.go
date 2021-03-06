package controller

import (
	// "fmt"
	"handmade_mask_shop/domain"
	"handmade_mask_shop/repository"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)


var item domain.Item
var itemImage domain.ItemImage
var category domain.Category


func GetItem(c *gin.Context) {
	u64, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(u64)
	item, err := repository.GetItemByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "adminUser not found"})
	}
	c.JSON(http.StatusOK, item)
}


func GetDisplayItem(c *gin.Context) {
	u64, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(u64)
	p := repository.Item{}
  item, err := p.GetDisplayItem(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, item)
}


func GetDisplayItems(c *gin.Context) {
  items, err := repository.GetDisplayItems()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "items not found"})
	}
	c.JSON(http.StatusOK, items)
}


func GetItemImages(c * gin.Context) {
	itemImages, err := repository.GetAllItemImages()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, itemImages)
}


func GetDisplayItemsCategoryId(c *gin.Context) {
	u64, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(u64)
  items, err := repository.GetDisplayItemsCategoryId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "items not found"})
	}
	c.JSON(http.StatusOK, items)
}


func GetCategories(c *gin.Context) {
	categories, err := repository.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, categories)
}