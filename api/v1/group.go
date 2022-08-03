package v1

import (
	"2022summer/model"
	"2022summer/model/response"
	"2022summer/service"
	"2022summer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
	identity, notFound := service.QueryIdentity(poster.(model.User).UserID, data.GroupID)
	if notFound || identity.Status == 1 {
		c.JSON(http.StatusOK, response.RemoveMemberA{
			CommonA: response.CommonA{
				Message: "没有权限",
				Success: false,
			},
		})
		return
	}
	identity, notFound = service.QueryIdentity(data.UserID, data.GroupID)
	if notFound {
		c.JSON(http.StatusOK, response.RemoveMemberA{
			CommonA: response.CommonA{
				Message: "所操作用户不在团队中",
				Success: false,
			},
		})
		return
	}
	if identity.Status >= 2 {
		c.JSON(http.StatusOK, response.RemoveMemberA{
			CommonA: response.CommonA{
				Message: "对方是管理员，不能移除",
				Success: false,
			},
		})
		return
	}
	if err := service.DeleteIdentity(&identity); err != nil {
		c.JSON(http.StatusOK, response.RemoveMemberA{
			CommonA: response.CommonA{
				Message: "移除失败",
				Success: false,
			},
		})
		return
	}
	c.JSON(http.StatusOK, response.RemoveMemberA{
		CommonA: response.CommonA{
			Message: "移除成功",
			Success: true,
		},
	})
	return
}
