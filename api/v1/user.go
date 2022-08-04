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
// @Tags 基本模块
// @Accept json
// @Produce json
// @Param data body response.RegisterQ true "用户名，密码"
// @Success 200 {object} response.RegisterA
// @Router /register [post]
func Register(c *gin.Context) {
	data := utils.BindJsonAndValid(c, &response.RegisterQ{}).(*response.RegisterQ)
	if _, notFound := service.QueryUserByUsername(data.Username); !notFound {
		c.JSON(http.StatusOK, response.RegisterA{
			CommonA: response.CommonA{
				Message: "用户名已存在",
				Success: false,
			},
		})
		return
	}
	HashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err := service.CreateUser(&model.User{
		Username: data.Username,
		Password: string(HashedPassword)}); err != nil {
		c.JSON(http.StatusOK, response.RegisterA{
			CommonA: response.CommonA{
				Message: "注册失败",
				Success: false,
			},
		})
		return
	}
	c.JSON(http.StatusOK, response.RegisterA{
		CommonA: response.CommonA{
			Message: "注册成功",
			Success: true,
		},
	})
}

// Login
// @Tags 基本模块
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
			CommonA: response.CommonA{
				Message: "用户不存在",
				Success: false,
			},
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		c.JSON(http.StatusOK, response.LoginA{
			CommonA: response.CommonA{
				Message: "密码错误",
				Success: false,
			},
		})
		return
	}
	token := utils.GenerateToken(user.UserID)
	c.JSON(http.StatusOK, response.LoginA{
		CommonA: response.CommonA{
			Message: "登录成功",
			Success: true,
		},
		Token: token,
	})
}

// GetUserInfo
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param data body response.GetUserInfoQ true "用户名，密码"
// @Success 200 {object} response.GetUserInfoA
// @Router /user/info [post]
func GetUserInfo(c *gin.Context) {
	poster, _ := c.Get("user")
	data := utils.BindJsonAndValid(c, &response.GetUserInfoQ{}).(*response.GetUserInfoQ)
	user, notFound := service.QueryUserByUserID(data.UserID)
	print(data.UserID)
	if notFound {
		c.JSON(http.StatusOK, response.GetUserInfoA{
			CommonA: response.CommonA{
				Message: "用户不存在",
				Success: false,
				Code:    200,
			},
			Poster: poster.(model.User),
		})
		return
	}
	c.JSON(http.StatusOK, response.GetUserInfoA{
		CommonA: response.CommonA{
			Message: "获取成功",
			Success: true,
			Code:    200,
		},
		Poster: poster.(model.User),
		User:   user,
	})
}

// ModifyPassword
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param data body response.ModifyPasswordQ true "新密码"
// @Success 200 {object} response.ModifyPasswordA
// @Router /user/modify_password [post]
func ModifyPassword(c *gin.Context) {
	var data response.ModifyPasswordQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.PARAMERROR)
		return
	}
	poster := c.MustGet("user").(model.User)
	poster.Password = data.Password
	if err := service.UpdateUser(&poster); err != nil {
		c.JSON(http.StatusOK, response.DBERROR)
		return
	}
	c.JSON(http.StatusOK, response.ModifyPasswordA{
		CommonA: response.CommonA{
			Message: "修改成功",
			Success: true,
		},
	})
}

// ModifyInfo
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param data body response.ModifyInfoQ true "用户信息"
// @Success 200 {object} response.ModifyInfoA
// @Router /user/modify_info [post]
func ModifyInfo(c *gin.Context) {
	var data response.ModifyInfoQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.PARAMERROR)
		return
	}
	poster := c.MustGet("user").(model.User)
	if _, notFound := service.QueryUserByUsername(data.Username); !notFound {
		c.JSON(http.StatusOK, response.ModifyInfoA{
			CommonA: response.CommonA{
				Message: "用户名已存在",
				Success: false,
			},
		})
		return
	}
	if _, notFound := service.QueryUserByEmail(data.Email); !notFound {
		c.JSON(http.StatusOK, response.ModifyInfoA{
			CommonA: response.CommonA{
				Message: "邮箱已存在",
				Success: false,
			},
		})
		return
	}
	poster.Username = data.Username
	poster.Email = data.Email
	poster.Age = data.Age
	poster.Sex = data.Sex
	if err := service.UpdateUser(&poster); err != nil {
		c.JSON(http.StatusOK, response.ModifyInfoA{
			response.CommonA{
				Message: "修改成功",
				Success: true,
			},
		})
	}
}
