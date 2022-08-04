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
	poster := c.MustGet("user").(model.User)
	group := model.Group{
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

// GetIdentity
// @Tags Group
// @Accept json
// @Produce json
// @Param data body response.GetIdentityQ true "团队id"
// @Success 200 {object} response.GetIdentityA
// @Router /group/get_identity [post]
func GetIdentity(c *gin.Context) {
	data := utils.BindJsonAndValid(c, &response.GetIdentityQ{}).(*response.GetIdentityQ)
	poster, _ := c.Get("user")
	identity, notFound := service.QueryIdentity(poster.(database.User).UserID, data.GroupID)
	if notFound {
		c.JSON(http.StatusOK, response.GetIdentityA{CommonA: response.NOMENBER})
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
	data := utils.BindJsonAndValid(c, &response.GetMembersQ{}).(*response.GetMembersQ)
	poster, _ := c.Get("user")
	if _, notFound := service.QueryIdentity(poster.(database.User).UserID, data.GroupID); notFound {
		c.JSON(http.StatusOK, response.GetMembersA{CommonA: response.NOMENBER})
	}
	members := service.GetGroupMembers(data.GroupID)
	c.JSON(http.StatusOK, response.GetMembersA{
		CommonA: response.CommonA{
			Message: "获取成功",
			Success: true,
		},
		Members: members,
	})
}

// AddMember
// @Tags Group
// @Accept json
// @Produce json
// @Param data body response.AddMemberQ true "用户id，团队id"
// @Success 200 {object} response.AddMemberA
// @Router /group/add_member [post]
func AddMember(c *gin.Context) {
	data := utils.BindJsonAndValid(c, &response.AddMemberQ{}).(*response.AddMemberQ)
	poster, _ := c.Get("user")
	identity, notFound := service.QueryIdentity(poster.(database.User).UserID, data.GroupID)
	if notFound || identity.Status == 1 {
		c.JSON(http.StatusOK, response.AddMemberA{CommonA: response.NOAUTH})
		return
	}
	identity = database.Identity{
		UserID:  data.UserID,
		GroupID: data.GroupID,
		Status:  1,
	}
	if err := service.CreateIdentity(&identity); err != nil {
		c.JSON(http.StatusOK, response.AddMemberA{CommonA: response.DBERROR})
	}
	c.JSON(http.StatusOK, response.AddMemberA{
		CommonA: response.CommonA{
			Message: "添加成功",
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
		c.JSON(http.StatusOK, response.RemoveMemberA{CommonA: response.NOMENBER})
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
		c.JSON(http.StatusOK, response.SetMemberStatusA{CommonA: response.NOMENBER})
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
