package v1

import (
	"2022summer/model"
	"2022summer/service"
	"2022summer/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Register
func Register(c *gin.Context) {
	data := utils.BindJsonAndValid(c, &model.RegisterQ{}).(*model.RegisterQ)
	if _, notFound := service.QueryUserByUsername(data.Username); !notFound {
		c.JSON(http.StatusOK, gin.H{
			"message": "用户名已存在",
			"success": false,
		})
		return
	}
	HashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err := service.CreateUser(&model.User{
		Username: data.Username,
		Password: string(HashedPassword)}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "注册失败",
			"success": false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
		"success": true,
	})
}

// Login
func Login(c *gin.Context) {
	data := utils.BindJsonAndValid(c, &model.LoginQ{}).(*model.LoginQ)
	user, notFound := service.QueryUserByUsername(data.Username)
	if notFound {
		c.JSON(http.StatusOK, gin.H{
			"message": "用户名不存在",
			"success": false,
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "密码错误",
			"success": false,
		})
		return
	}
	token := utils.GenerateToken(user.UserID)
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   token,
		"success": true,
	})
}
