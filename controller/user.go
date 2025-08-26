package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wys1976/gin-api888/model"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func (uc *UserController) GetUser(c *gin.Context) {
	var req struct {
		UserId string `json:"userId"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var user model.User
	if err := uc.DB.Where("id = ?", req.UserId).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"phone":    user.Phone,
		"age":      user.Age,
	})
}

func (uc *UserController) UpdatePhone(c *gin.Context) {
	var req struct {
		UserId string `json:"userId"`
		Phone  string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := uc.DB.Model(&model.User{}).
		Where("id = ?", req.UserId).
		Update("phone", req.Phone).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "phone updated successfully"})
}
