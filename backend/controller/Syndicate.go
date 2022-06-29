package controller

import (
	"net/http"

	"github.com/Worrakarnp/suta/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// POST /syndicate
func CreateSyndicate(c *gin.Context) {
	var branch entity.Branch
	var syndicate entity.Syndicate

	// ผลลัพธ์ที่ได้จะถูก bind เข้าตัวแปร syndicate
	if err := c.ShouldBindJSON(&syndicate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา branch ด้วย id
	if tx := entity.DB().Where("id = ?", syndicate.BranchID).First(&branch); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "branch not found"})
		return
	}

	// สร้าง syndicate
	s := entity.Syndicate{
		Name:   syndicate.Name, // ตั้งค่าฟิลด์ Name
		Rank:   syndicate.Rank, // ตั้งค่าฟิลด์ Rank
		Branch: branch,         // โยงความสัมพันธ์กับ Entity Branch
		Tel:    syndicate.Tel,  // ตั้งค่าฟิลด์ Tel
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// บันทึก
	if err := entity.DB().Create(&s).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": s})
}

// GET /syndicate/:id
func GetSyndicate(c *gin.Context) {
	var syndicate entity.Syndicate
	id := c.Param("id")
	if err := entity.DB().Preload("Branch").Raw("SELECT * FROM syndicate WHERE id = ?", id).Find(&syndicate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": syndicate})
}

// GET /syndicate
func ListSyndicate(c *gin.Context) {
	var syndicate []entity.Syndicate
	if err := entity.DB().Preload("Branch").Raw("SELECT * FROM syndicates").Find(&syndicate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": syndicate})
}

// DELETE /syndicate/:id
func DeleteSyndicate(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM syndicate WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "syndicate not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /syndicate
func UpdateSyndicate(c *gin.Context) {
	var syndicate entity.Syndicate
	if err := c.ShouldBindJSON(&syndicate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", syndicate.ID).First(&syndicate); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "syndicate not found"})
		return
	}

	if err := entity.DB().Save(&syndicate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": syndicate})
}
