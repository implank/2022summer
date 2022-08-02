package v1

import (
	"2022summer/model"
	"2022summer/model/response"
	"2022summer/service"
	"2022summer/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Register
// @Tags User
// @Accept json
// @Produce json
// @Param data body response.RegisterQ true "用户名，密码"
// @Success 200 {object} response.RegisterA
// @Router /register [post]
func Register(c *gin.Context) {
	data := utils.BindJsonAndValid(c, &response.RegisterQ{}).(*response.RegisterQ)
	if _, notFound := service.QueryUserByUsername(data.Username); !notFound {
		c.JSON(http.StatusOK, response.RegisterA{
			Message: "用户名已存在",
			Success: false,
		})
		return
	}
	HashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err := service.CreateUser(&model.User{
		Username: data.Username,
		Password: string(HashedPassword)}); err != nil {
		c.JSON(http.StatusOK, response.RegisterA{
			Message: "注册失败",
			Success: false,
		})
		return
	}
	c.JSON(http.StatusOK, response.RegisterA{
		Message: "注册成功",
		Success: true,
	})
}

// Login
// @Tags User
// @Accept json
// @Produce json
// @Param data body response.LoginQ true "用户名，密码"
// @Success 200 {object} response.LoginA
// @Router /login [post]
func Login(c *gin.Context) {
	data := utils.BindJsonAndValid(c, &response.LoginQ{}).(*response.LoginQ)
	user, notFound := service.QueryUserByUsername(data.Username)
	if notFound {
		c.JSON(http.StatusOK, response.LoginA{
			Message: "用户不存在",
			Success: false,
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		c.JSON(http.StatusOK, response.LoginA{
			Message: "密码错误",
			Success: false,
		})
		return
	}
	token := utils.GenerateToken(user.UserID)
	c.JSON(http.StatusOK, response.LoginA{
		Message: "登录成功",
		Success: true,
		Token:   token,
	})
}

// GetUserInfo
// @Tags User
// @Accept json
// @Produce json
// @Param data body response.GetUserInfoQ true "用户名，密码"
// @Success 200 {object} response.GetUserInfoA
// @Router /info [post]
func GetUserInfo(c *gin.Context) {
	poster, _ := c.Get("user")
	data := utils.BindJsonAndValid(c, &response.GetUserInfoQ{}).(*response.GetUserInfoQ)
	user, notFound := service.QueryUserByUserID(data.UserID)
	print(data.UserID)
	if notFound {
		c.JSON(http.StatusOK, response.GetUserInfoA{
			Message: "用户不存在",
			Success: false,
			Poster:  poster.(model.User),
		})
		return
	}
	c.JSON(http.StatusOK, response.GetUserInfoA{
		Message: "获取用户信息成功",
		Success: true,
		Poster:  poster.(model.User),
		User:    user,
	})
}
