package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team14/entity"
)

// POST /sizes
func CreateSize(c *gin.Context) {
	var size entity.Size
	if err := c.ShouldBindJSON(&size); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&size).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": size})
}

// GET /size/:id
func GetSize(c *gin.Context) {
	var size entity.Size
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM sizes WHERE id = ?", id).Scan(&size).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": size})
}

// GET /sizes
func ListSize(c *gin.Context) {
	var sizes []entity.Size
	if err := entity.DB().Raw("SELECT * FROM sizes").Scan(&sizes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sizes})
}

// DELETE /sizes/:id
func DeleteSize(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM sizes WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "size not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /sizes
func UpdateSize(c *gin.Context) {
	var size entity.Size
	if err := c.ShouldBindJSON(&size); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", size.ID).First(&size); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "size not found"})
		return
	}

	if err := entity.DB().Save(&size).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": size})
}
