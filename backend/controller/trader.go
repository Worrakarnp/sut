package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sut64/team14/entity"
)

// POST /trader
func CreateTrader(c *gin.Context) {
	var trader entity.Trader

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร trader
	if err := c.ShouldBindJSON(&trader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// สร้าง trader
	t := entity.Trader{
		Name:    trader.Name,    // ตั้งค่าฟิลด์ Name
		Tel:     trader.Tel,     // ตั้งค่าฟิลด์ Tel
		Address: trader.Address, // ตั้งค่าฟิลด์ Address
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// บันทึก
	if err := entity.DB().Create(&t).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": t})
}

// GET /trader/:id
func GetTrader(c *gin.Context) {
	var trader entity.Trader
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM trader WHERE id = ?", id).Find(&trader).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": trader})
}

// GET /trader
func ListTrader(c *gin.Context) {
	var trader []entity.Trader
	if err := entity.DB().Raw("SELECT * FROM traders").Find(&trader).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": trader})
}

// DELETE /trader/:id
func DeleteTrader(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM trader WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "trader not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /trader
func UpdateTrader(c *gin.Context) {
	var trader entity.Trader
	if err := c.ShouldBindJSON(&trader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", trader.ID).First(&trader); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "trader not found"})
		return
	}

	if err := entity.DB().Save(&trader).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": trader})
}
