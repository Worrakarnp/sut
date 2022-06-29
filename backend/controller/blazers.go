package controller

import (
	"net/http"

	"github.com/Worrakarnp/suta/entity"
	"github.com/gin-gonic/gin"
)

// POST /blazers
func CreateBlazer(c *gin.Context) {
	var blazer entity.Blazer
	if err := c.ShouldBindJSON(&blazer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&blazer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": blazer})
}

// GET /blazer/:id
func GetBlazer(c *gin.Context) {
	var blazer entity.Blazer
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM blazers WHERE id = ?", id).Scan(&blazer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": blazer})
}

// GET /blazers
func ListBlazer(c *gin.Context) {
	var blazers []entity.Blazer
	if err := entity.DB().Raw("SELECT * FROM blazers").Scan(&blazers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": blazers})
}

// DELETE /blazers/:id
func DeleteBlazer(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM blazers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "blazer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /blazers
func UpdateBlazer(c *gin.Context) {
	var blazer entity.Blazer
	if err := c.ShouldBindJSON(&blazer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", blazer.ID).First(&blazer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "blazer not found"})
		return
	}

	if err := entity.DB().Save(&blazer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": blazer})
}
