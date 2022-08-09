package v1

import (
	"2022summer/model/database"
	"2022summer/model/response"
	"2022summer/service"
	"2022summer/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
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
	matched, _ := regexp.Match("^[A-Za-z\\d]{8,40}$", []byte(data.Password))
	if !matched {
		c.JSON(http.StatusOK, response.ModifyInfoA{
			CommonA: response.CommonA{
				Message: "密码格式不正确",
				Success: false,
			},
		})
		return
	}
	HashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	user := database.User{
		Username: data.Username,
		Password: string(HashedPassword),
	}
	if err := service.CreateUser(&user); err != nil {
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
		User:  user,
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
			Poster: poster.(database.User),
		})
		return
	}
	c.JSON(http.StatusOK, response.GetUserInfoA{
		CommonA: response.CommonA{
			Message: "获取成功",
			Success: true,
			Code:    200,
		},
		Poster: poster.(database.User),
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
	poster := c.MustGet("user").(database.User)
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
	poster := c.MustGet("user").(database.User)
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
			CommonA: response.CommonA{
				Message: "修改成功",
				Success: true,
			},
		})
	}
}

// GetMessages
// @Summary 获取用户消息，按时间排序
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param data body response.GetMessagesQ true "空json"
// @Success 200 {object} response.GetMessagesA
// @Router /user/get_messages [post]
func GetMessages(c *gin.Context) {
	var data response.GetMessagesQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.PARAMERROR)
		return
	}
	poster := c.MustGet("user").(database.User)
	messages := service.GetMessages(poster.UserID)
	count := len(messages)
	if count == 0 {
		c.JSON(http.StatusOK, response.GetMessagesA{
			CommonA: response.CommonA{
				Message: "您还没有消息",
				Success: true,
			},
			Count: 0,
		})
		return
	}
	c.JSON(http.StatusOK, response.GetMessagesA{
		CommonA: response.CommonA{
			Message: "获取消息成功",
			Success: true,
		},
		Count:    count,
		Messages: messages,
	})
}

// DeclineInvitation
// @Summary 拒绝加入团队邀请
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param data body response.DeclineInvitationQ true "消息ID"
// @Success 200 {object} response.DeclineInvitationA
// @Router /user/decline_invitation [post]
func DeclineInvitation(c *gin.Context) {
	var data response.DeclineInvitationQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.PARAMERROR)
		return
	}
	poster := c.MustGet("user").(database.User)
	message, notFound := service.QueryMessageByMessageID(data.MessageID)
	if notFound || message.Type != 1 || message.ReceiverID != poster.UserID {
		c.JSON(http.StatusOK, response.DeclineInvitationA{
			Message: "没有找到该消息 或 消息类型错误 或 消息接收者错误",
			Success: false,
		})
		return
	}
	message.Type = 2
	_ = service.UpdateMessage(&message)
	newMessage := database.Message{
		ReceiverID: message.SenderID,
		SenderID:   0,
		Content:    poster.Username + " 拒绝了您的邀请",
		Type:       4,
	}
	_ = service.CreateMessage(&newMessage)
	c.JSON(http.StatusOK, response.DeclineInvitationA{
		Message: "拒绝邀请成功",
		Success: true,
	})
}

// AcceptInvitation
// @Summary 接受加入团队邀请
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param data body response.AcceptInvitationQ true "消息ID"
// @Success 200 {object} response.AcceptInvitationA
// @Router /user/accept_invitation [post]
func AcceptInvitation(c *gin.Context) {
	var data response.AcceptInvitationQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.PARAMERROR)
		return
	}
	poster := c.MustGet("user").(database.User)
	message, notFound := service.QueryMessageByMessageID(data.MessageID)
	if notFound || message.Type != 1 || message.ReceiverID != poster.UserID {
		c.JSON(http.StatusOK, response.DeclineInvitationA{
			Message: "没有找到该消息 或 消息类型错误 或 消息接收者错误",
			Success: false,
		})
	}
	message.Type = 3
	_ = service.UpdateMessage(&message)
	identity := database.Identity{
		UserID:  poster.UserID,
		GroupID: message.GroupID,
		Status:  1,
	}
	_ = service.CreateIdentity(&identity)
	members := service.GetGroupMembers(message.GroupID)
	for _, member := range members {
		message = database.Message{
			ReceiverID: member.UserID,
			SenderID:   0,
			Content:    "新成员 " + poster.Username + " 加入了团队",
			Type:       5,
		}
		_ = service.CreateMessage(&message)
	}
	c.JSON(http.StatusOK, response.AcceptInvitationA{
		Message: "接受邀请成功",
		Success: true,
	})
}

// ReadAllMessages
// @Summary 标记所有消息为已读
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param data body response.ReadAllMessagesQ true "空json"
// @Success 200 {object} response.ReadAllMessagesA
// @Router /user/read_all_messages [post]
func ReadAllMessages(c *gin.Context) {
	post := c.MustGet("user").(database.User)
	messages := service.GetMessages(post.UserID)
	for _, message := range messages {
		message.Status = 1
		if err := service.UpdateMessage(&message); err != nil {
			c.JSON(http.StatusOK, response.DBERROR)
			return
		}
	}
	c.JSON(http.StatusOK, response.ReadAllMessagesA{
		CommonA: response.CommonA{
			Message: "操作成功",
			Success: true,
		},
	})
}
