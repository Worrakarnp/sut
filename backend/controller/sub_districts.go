package controller

import (
	"net/http"

	"github.com/Worrakarnp/suta/entity"
	"github.com/gin-gonic/gin"
)

// POST /sub_districts
func CreateSubDistrict(c *gin.Context) {
	var sub_district entity.SubDistrict
	if err := c.ShouldBindJSON(&sub_district); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&sub_district).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sub_district})
}

// GET /sub_district/:id
func GetSubDistrict(c *gin.Context) {
	var sub_district entity.SubDistrict
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM sub_districts WHERE id = ?", id).Scan(&sub_district).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sub_district})
}

// GET /sub_districts
func ListSubDistrict(c *gin.Context) {
	var sub_districts []entity.SubDistrict
	if err := entity.DB().Raw("SELECT * FROM sub_districts").Scan(&sub_districts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sub_districts})
}

// DELETE /sub_districts/:id
func DeleteSubDistrict(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM sub_districts WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sub_district not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /sub_districts
func UpdateSubDistrict(c *gin.Context) {
	var sub_district entity.SubDistrict
	if err := c.ShouldBindJSON(&sub_district); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", sub_district.ID).First(&sub_district); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sub_district not found"})
		return
	}

	if err := entity.DB().Save(&sub_district).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sub_district})
}
