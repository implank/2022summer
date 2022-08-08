package v1

import (
	"2022summer/model/database"
	"2022summer/model/response"
	"2022summer/service"
	"2022summer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateGroup
// @Tags Group
// @Accept json
// @Produce json
// @Param data body response.CreateGroupQ true "团队名称，团队介绍"
// @Success 200 {object} response.CreateGroupA
// @Router /group/create_group [post]
func CreateGroup(c *gin.Context) {
	var data response.CreateGroupQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.PARAMERROR)
		return
	}
	poster := c.MustGet("user").(database.User)
	group := database.Group{
		GroupName: data.GroupName,
		GroupInfo: data.GroupInfo,
		UserID:    poster.UserID,
	}
	if err := service.CreateGroup(&group); err != nil {
		c.JSON(http.StatusOK, response.DBERROR)
		return
	}
	c.JSON(http.StatusOK, response.CreateGroupA{
		CommonA: response.CommonA{
			Message: "创建成功",
			Success: true,
		},
		Group: group,
	})
}

// ModifyGroup
// @Tags Group
// @Accept json
// @Produce json
// @Param data body response.ModifyGroupQ true "团队id，团队名称，团队介绍"
// @Success 200 {object} response.ModifyGroupA
// @Router /group/modify_group [post]
func ModifyGroup(c *gin.Context) {
	var data response.ModifyGroupQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.PARAMERROR)
		return
	}
	poster := c.MustGet("user").(database.User)
	group, notFound := service.QueryGroupByGroupID(data.GroupID)
	if notFound {
		c.JSON(http.StatusOK, response.PARAMERROR)
		return
	}
	identity, notFound := service.QueryIdentity(poster.UserID, group.GroupID)
	if notFound || identity.Status != 3 {
		c.JSON(http.StatusOK, response.NOAUTH)
		return
	}
	group.GroupName = data.GroupName
	group.GroupInfo = data.GroupInfo
	if err := service.UpdateGroup(&group); err != nil {
		c.JSON(http.StatusOK, response.DBERROR)
		return
	}
	c.JSON(http.StatusOK, response.ModifyGroupA{
		CommonA: response.CommonA{
			Message: "修改成功",
			Success: true,
		},
		Group: group,
	})
}

// GetIdentity
// @Tags Group
// @Accept json
// @Produce json
// @Param data body response.GetIdentityQ true "团队id"
// @Success 200 {object} response.GetIdentityA
// @Router /group/get_my_identity [post]
func GetIdentity(c *gin.Context) {
	data := utils.BindJsonAndValid(c, &response.GetIdentityQ{}).(*response.GetIdentityQ)
	poster, _ := c.Get("user")
	identity, notFound := service.QueryIdentity(poster.(database.User).UserID, data.GroupID)
	if notFound {
		c.JSON(http.StatusOK, response.GetIdentityA{CommonA: response.USERNOTINGROUP})
		return
	}
	c.JSON(http.StatusOK, response.GetIdentityA{
		CommonA: response.CommonA{
			Message: "获取成功",
			Success: true,
		},
		Status: identity.Status,
	})
}

// GetMembers
// @Tags Group
// @Accept json
// @Produce json
// @Param data body response.GetMembersQ true "团队id"
// @Success 200 {object} response.GetMembersA
// @Router /group/get_group_members [post]
func GetMembers(c *gin.Context) {
	var data response.GetMembersQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.PARAMERROR)
		return
	}
	poster, _ := c.Get("user")
	print(data.GroupID)
	if _, notFound := service.QueryIdentity(poster.(database.User).UserID, data.GroupID); notFound {
		c.JSON(http.StatusOK, response.GetMembersA{CommonA: response.NOAUTH})
		return
	}
	members := service.GetGroupMembers(data.GroupID)
	c.JSON(http.StatusOK, response.GetMembersA{
		CommonA: response.CommonA{
			Message: "获取成功",
			Success: true,
		},
		Count:   len(members),
		Members: members,
	})
}

// InviteMember
// @Tags Group
// @Accept json
// @Produce json
// @Param data body response.InviteMemberQ true "团队id，用户id"
// @Success 200 {object} response.InviteMemberA
// @Router /group/invite_member [post]
func InviteMember(c *gin.Context) {
	var data response.InviteMemberQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.PARAMERROR)
		return
	}
	poster := c.MustGet("user").(database.User)
	identity, notFound := service.QueryIdentity(poster.UserID, data.GroupID)
	if notFound || identity.Status == 1 {
		c.JSON(http.StatusOK, response.NOAUTH)
		return
	}
	user, notFound := service.QueryUserByUsername(data.Username)
	if notFound {
		c.JSON(http.StatusOK, response.USERNOTEXSIT)
		return
	}
	group, notFound := service.QueryGroupByGroupID(data.GroupID)
	if notFound {
		c.JSON(http.StatusOK, response.InviteMemberA{
			CommonA: response.CommonA{
				Message: "团队不存在",
				Success: false,
			},
		})
	}
	if _, notFound = service.QueryIdentity(user.UserID, data.GroupID); !notFound {
		c.JSON(http.StatusOK, response.InviteMemberA{
			CommonA: response.CommonA{
				Message: "用户已在团队中",
				Success: false,
			},
		})
		return
	}
	message := database.Message{
		SenderID:   poster.UserID,
		ReceiverID: user.UserID,
		Content:    poster.Username + " 邀请您加入团队 " + group.GroupName,
		GroupID:    group.GroupID,
		Type:       1,
	}
	if err := service.CreateMessage(&message); err != nil {
		c.JSON(http.StatusOK, response.DBERROR)
		return
	}
	c.JSON(http.StatusOK, response.InviteMemberA{
		CommonA: response.CommonA{
			Message: "邀请发送成功",
			Success: true,
		},
	})
}

// RemoveMember
// @Tags Group
// @Accept json
// @Produce json
// @Param data body response.RemoveMemberQ true "用户id，团队id"
// @Success 200 {object} response.RemoveMemberA
// @Router /group/remove_member [post]
func RemoveMember(c *gin.Context) {
	data := utils.BindJsonAndValid(c, &response.RemoveMemberQ{}).(*response.RemoveMemberQ)
	poster, _ := c.Get("user")
	identity, notFound := service.QueryIdentity(poster.(database.User).UserID, data.GroupID)
	if notFound || identity.Status == 1 {
		c.JSON(http.StatusOK, response.RemoveMemberA{CommonA: response.NOAUTH})
		return
	}
	identity, notFound = service.QueryIdentity(data.UserID, data.GroupID)
	if notFound {
		c.JSON(http.StatusOK, response.RemoveMemberA{CommonA: response.USERNOTINGROUP})
		return
	}
	if identity.Status >= 2 {
		c.JSON(http.StatusOK, response.RemoveMemberA{CommonA: response.NOAUTH})
		return
	}
	if err := service.DeleteIdentity(&identity); err != nil {
		c.JSON(http.StatusOK, response.RemoveMemberA{CommonA: response.DBERROR})
		return
	}
	c.JSON(http.StatusOK, response.RemoveMemberA{
		CommonA: response.CommonA{
			Message: "移除成功",
			Success: true,
		},
	})
}

// SetMemberStatus
// @Tags Group
// @Accept json
// @Produce json
// @Param data body response.SetMemberStatusQ true "用户id，团队id，状态"
// @Success 200 {object} response.SetMemberStatusA
// @Router /group/set_member_status [post]
func SetMemberStatus(c *gin.Context) {
	var data response.SetMemberStatusQ
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusOK, response.PARAMERROR)
		return
	}
	poster, _ := c.Get("user")
	identity1, notFound := service.QueryIdentity(poster.(database.User).UserID, data.GroupID)
	if notFound || identity1.Status == 1 {
		c.JSON(http.StatusOK, response.SetMemberStatusA{CommonA: response.NOAUTH})
		return
	}
	identity2, notFound := service.QueryIdentity(data.UserID, data.GroupID)
	if notFound {
		c.JSON(http.StatusOK, response.SetMemberStatusA{CommonA: response.USERNOTINGROUP})
		return
	}
	switch data.Status {
	case 1:
		if identity2.Status == 1 {
			c.JSON(http.StatusOK, response.SetMemberStatusA{
				CommonA: response.CommonA{
					Message: "对方已是普通成员",
					Success: false,
				},
			})
		} else if identity2.Status == 2 && identity1.Status == 3 {
			identity2.Status = 1
			_ = service.UpdateIdentity(&identity2)
			c.JSON(http.StatusOK, response.SetMemberStatusA{
				CommonA: response.CommonA{
					Message: "设置普通成员成功",
					Success: true,
				},
			})
		} else {
			c.JSON(http.StatusOK, response.SetMemberStatusA{CommonA: response.NOAUTH})
		}
	case 2:
		if identity2.Status == 2 {
			c.JSON(http.StatusOK, response.SetMemberStatusA{
				CommonA: response.CommonA{
					Message: "对方已是管理员",
					Success: false,
				},
			})
		} else if identity2.Status == 1 && identity1.Status >= 2 {
			identity2.Status = 2
			_ = service.UpdateIdentity(&identity2)
			c.JSON(http.StatusOK, response.SetMemberStatusA{
				CommonA: response.CommonA{
					Message: "设置管理员成功",
					Success: true,
				},
			})
		} else {
			c.JSON(http.StatusOK, response.SetMemberStatusA{CommonA: response.NOAUTH})
		}
	}
}

// GetGroups
// @Tags Group
// @Accept json
// @Produce json
// @Success 200 {object} response.GetGroupsA
// @Router /group/get_groups [post]
func GetGroups(c *gin.Context) {
	poster := c.MustGet("user").(database.User)
	groups := service.GetUserHasGroups(poster.UserID)
	count := len(groups)
	if count == 0 {
		c.JSON(http.StatusOK, response.GetGroupsA{
			CommonA: response.CommonA{
				Message: "您还没有团队",
				Success: true,
			},
			Count: 0,
		})
		return
	}
	c.JSON(http.StatusOK, response.GetGroupsA{
		CommonA: response.CommonA{
			Message: "获取团队成功",
			Success: true,
		},
		Count:  count,
		Groups: groups,
	})
}
