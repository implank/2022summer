package middleware

import (
	"2022summer/model/response"
	"2022summer/service"
	"2022summer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		id, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, response.AuthResponse{
				Message: "认证失败",
				Success: false,
				Code:    401,
			})
			c.Abort()
			return
		}
		if user, notFound := service.QueryUserByUserID(id); notFound {
			c.JSON(http.StatusOK, response.AuthResponse{
				Message: "" + user.Username + "用户不存在",
				Success: false,
				Code:    402,
			})
			c.Abort()
		} else {
			c.Set("user", user)
		}
	}
}
