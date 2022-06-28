package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sut64/team14/entity"
)

// POST /member
func CreateMember(c *gin.Context) {
	var branch entity.Branch
	var category entity.Category
	var subdistrict entity.SubDistrict
	var district entity.District
	var province entity.Province
	var member entity.Member

	// ผลลัพธ์ที่ได้จะถูก bind เข้าตัวแปร member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา branch ด้วย id
	if tx := entity.DB().Where("id = ?", member.BranchID).First(&branch); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "branch not found"})
		return
	}

	// ค้นหา category ด้วย id
	if tx := entity.DB().Where("id = ?", member.CategoryID).First(&category); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category not found"})
		return
	}

	// ค้นหา subdistrict ด้วย id
	if tx := entity.DB().Where("id = ?", member.SubDistrictID).First(&subdistrict); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subdistrict not found"})
		return
	}

	// ค้นหา district ด้วย id
	if tx := entity.DB().Where("id = ?", member.DistrictID).First(&district); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "district not found"})
		return
	}

	// ค้นหา province ด้วย id
	if tx := entity.DB().Where("id = ?", member.ProvinceID).First(&province); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "province not found"})
		return
	}

	// สร้าง member
	m := entity.Member{
		Name:        member.Name,    // ตั้งค่าฟิลด์ Name
		Num:         member.Num,     // ตั้งค่าฟิลด์ Num
		Branch:      branch,         // โยงความสัมพันธ์กับ Entity Branch
		Email:       member.Email,   // ตั้งค่าฟิลด์ Email
		Tel:         member.Tel,     // ตั้งค่าฟิลด์ Tel
		Address:     member.Address, // ตั้งค่าฟิลด์ Address
		SubDistrict: subdistrict,    // โยงความสัมพันธ์กับ Entity subdistrict
		District:    district,       // โยงความสัมพันธ์กับ Entity district
		Province:    province,       // โยงความสัมพันธ์กับ Entity Province
		Date:        member.Date,    // ตั้งค่าฟิลด์ Date
		Category:    category,       // โยงความสัมพันธ์กับ Entity category
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// บันทึก
	if err := entity.DB().Create(&m).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": m})
}

// GET /member/:id
func GetMember(c *gin.Context) {
	var member entity.Member
	id := c.Param("id")
	if err := entity.DB().Preload("Branch").Preload("Category").Preload("SubDistrict").Preload("District").Preload("Province").Raw("SELECT * FROM member WHERE id = ?", id).Find(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": member})
}

// GET /member
func ListMember(c *gin.Context) {
	var member []entity.Member
	if err := entity.DB().Preload("Branch").Preload("Category").Preload("SubDistrict").Preload("District").Preload("Province").Raw("SELECT * FROM members").Find(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": member})
}

// DELETE /member/:id
func DeleteMember(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM member WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "member not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /member
func UpdateMember(c *gin.Context) {
	var member entity.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", member.ID).First(&member); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "member not found"})
		return
	}

	if err := entity.DB().Save(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": member})
}
