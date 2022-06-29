package controller

import (
	"net/http"

	"github.com/Worrakarnp/suta/entity"
	"github.com/gin-gonic/gin"
)

// POST /degrees
func CreateDegree(c *gin.Context) {
	var degree entity.Degree
	if err := c.ShouldBindJSON(&degree); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&degree).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": degree})
}

// GET /degree/:id
func GetDegree(c *gin.Context) {
	var degree entity.Degree
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM degrees WHERE id = ?", id).Scan(&degree).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": degree})
}

// GET /degrees
func ListDegree(c *gin.Context) {
	var degrees []entity.Degree
	if err := entity.DB().Raw("SELECT * FROM degrees").Scan(&degrees).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": degrees})
}

// DELETE /degrees/:id
func DeleteDegree(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM degrees WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "degree not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /degrees
func UpdateDegree(c *gin.Context) {
	var degree entity.Degree
	if err := c.ShouldBindJSON(&degree); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", degree.ID).First(&degree); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "degree not found"})
		return
	}

	if err := entity.DB().Save(&degree).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": degree})
}
