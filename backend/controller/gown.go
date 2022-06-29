package controller

import (
	"net/http"

	"github.com/Worrakarnp/suta/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// POST /gown
func CreateGown(c *gin.Context) {
	var branch entity.Branch
	var size entity.Size
	var blazer entity.Blazer
	var degree entity.Degree
	var gown entity.Gown

	// ผลลัพธ์ที่ได้จะถูก bind เข้าตัวแปร gown
	if err := c.ShouldBindJSON(&gown); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา branch ด้วย id
	if tx := entity.DB().Where("id = ?", gown.BranchID).First(&branch); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "branch not found"})
		return
	}

	// ค้นหา size ด้วย id
	if tx := entity.DB().Where("id = ?", gown.SizeID).First(&size); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "size not found"})
		return
	}

	// ค้นหา blazer ด้วย id
	if tx := entity.DB().Where("id = ?", gown.BlazerID).First(&blazer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "blazer not found"})
		return
	}

	// ค้นหา degree ด้วย id
	if tx := entity.DB().Where("id = ?", gown.DegreeID).First(&degree); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "degree not found"})
		return
	}

	// สร้าง gown
	g := entity.Gown{
		Name:    gown.Name,    // ตั้งค่าฟิลด์ Name
		Student: gown.Student, // ตั้งค่าฟิลด์ Student
		Status:  gown.Status,  // ตั้งค่าฟิลด์ Status
		Size:    size,         // โยงความสัมพันธ์กับ Entity Size
		Blazer:  blazer,       // โยงความสัมพันธ์กับ Entity Blazer
		Degree:  degree,       // โยงความสัมพันธ์กับ Entity Degree
		Branch:  branch,       // โยงความสัมพันธ์กับ Entity Branch
		Tel:     gown.Tel,     // ตั้งค่าฟิลด์ Tel
		Date:    gown.Date,    // ตั้งค่าฟิลด์ Date
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// บันทึก
	if err := entity.DB().Create(&g).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": g})
}

// GET /gown/:id
func GetGown(c *gin.Context) {
	var gown entity.Gown
	id := c.Param("id")
	if err := entity.DB().Preload("Branch").Preload("Size").Preload("Blazer").Preload("Degree").Raw("SELECT * FROM gown WHERE id = ?", id).Find(&gown).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gown})
}

// GET /gown
func ListGown(c *gin.Context) {
	var gown []entity.Gown
	if err := entity.DB().Preload("Branch").Preload("Size").Preload("Blazer").Preload("Degree").Raw("SELECT * FROM gowns").Find(&gown).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gown})
}

// DELETE /gown/:id
func DeleteGown(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM gown WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gown not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /gown
func UpdateGown(c *gin.Context) {
	var gown entity.Gown
	if err := c.ShouldBindJSON(&gown); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", gown.ID).First(&gown); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gown not found"})
		return
	}

	if err := entity.DB().Save(&gown).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gown})
}
