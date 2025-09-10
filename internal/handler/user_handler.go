package handlers

import (
	"net/http"
	"strconv"

	"gin-api888/model"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	userIDStr := c.Param("id")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "Invalid user ID",
		})
		return
	}

	user, err := model.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "Failed to get user information",
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, Response{
			Code:    404,
			Message: "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "Success",
		Data:    user,
	})
}

// UpdateUserPhone 更新用户电话
func UpdateUserPhone(c *gin.Context) {
	userIDStr := c.Param("id")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "Invalid user ID",
		})
		return
	}

	var request struct {
		Phone string `json:"phone" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "Invalid request data",
		})
		return
	}

	// 验证手机号格式（可选）
	// if !isValidPhone(request.Phone) {
	//     c.JSON(http.StatusBadRequest, Response{
	//         Code:    400,
	//         Message: "Invalid phone number format",
	//     })
	//     return
	// }

	err = model.UpdateUserPhone(userID, request.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "Failed to update phone number",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "Phone number updated successfully",
	})
}
