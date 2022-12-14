package v1

import (
	"2022summer/model/database"
	"2022summer/model/response"
	"2022summer/service"
	"2022summer/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// GetProjPPages
// @Summary 获取某个项目的不在回收站的设计原型的页面
// @Tags 设计原型
// @Accept json
// @Produce json
// @Param data body response.GetProjPPagesQ true "项目ID"
// @Success 200 {object} response.GetProjPPagesA
// @Router /ppage/get_proj_ppages [post]
func GetProjPPages(c *gin.Context) {
	var data response.GetProjPPagesQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.GetProjPPagesA{Message: "输入数据不符合要求", Success: false})
		return
	}
	ppages := service.GetProjPPages(data.ProjID, 1)
	x := len(ppages)
	if x == 0 {
		c.JSON(http.StatusOK, response.GetProjPPagesA{
			Message: "没有找到捏",
			Success: false,
			Count:   uint64(x),
			PPages:  ppages})
		return
	}
	c.JSON(http.StatusOK, response.GetProjPPagesA{
		Message: "成功搜索到以下页面",
		Success: true,
		Count:   uint64(x),
		PPages:  ppages})
}

// GetPPageByID
// @Summary 获取设计原型的某个页面
// @Tags 设计原型
// @Accept json
// @Produce json
// @Param data body response.GetPPageByIDQ true "设计原型的页面ID"
// @Success 200 {object} response.GetPPageByIDA
// @Router /ppage/get_ppage_by_id [post]
func GetPPageByID(c *gin.Context) {
	var data response.GetPPageByIDQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.GetPPageByIDA{Message: "输入数据不符合要求", Success: false})
		return
	}
	ppage, notFound := service.QueryPPageByPPageID(data.PPageID)
	if notFound {
		c.JSON(http.StatusOK, response.GetPPageByIDA{Message: "页面不存在", Success: false, PPage: ppage})
		return
	}
	c.JSON(http.StatusOK, response.GetPPageByIDA{Message: "成功获取页面信息", Success: true, PPage: ppage})
}

// CreatePPage
// @Summary 创建设计原型的一个页面
// @Tags 设计原型
// @Accept json
// @Produce json
// @Param data body response.CreatePPageQ true "页面名称，页面数据（可选），页面所属项目ID"
// @Success 200 {object} response.CreatePPageA
// @Router /ppage/create_ppage [post]
func CreatePPage(c *gin.Context) {
	var data response.CreatePPageQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.CreatePPageA{Message: "输入数据不符合要求", Success: false})
		return
	}
	if _, notFound := service.QueryPPageByPPageName(data.PPageName, data.ProjID); !notFound {
		c.JSON(http.StatusOK, response.CreatePPageA{Message: "页面名已存在", Success: false})
		return
	}
	err := service.CreatePPage(&database.PPage{
		PPageName: data.PPageName,
		PPageData: data.PPageData,
		ProjID:    data.ProjID})
	if err != nil {
		c.JSON(http.StatusOK, response.CreatePPageA{Message: "创建页面失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.CreatePPageA{Message: "创建页面成功", Success: true})
}

// UpdatePPage
// @Summary 修改设计原型的某个页面的名称或数据
// @Tags 设计原型
// @Accept json
// @Produce json
// @Param data body response.UpdatePPageQ true "设计原型的页面ID，页面名称（可选，如果没填或者为空字符串，则不修改），页面数据（可选，如果没填或者为空字符串，则不修改）"
// @Success 200 {object} response.UpdatePPageA
// @Router /ppage/update_ppage [post]
func UpdatePPage(c *gin.Context) {
	var data response.UpdatePPageQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.UpdatePPageA{Message: "输入数据不符合要求", Success: false})
		return
	}
	ppage, notFound := service.QueryPPageByPPageID(data.PPageID)
	if notFound {
		c.JSON(http.StatusOK, response.UpdatePPageA{Message: "页面不存在", Success: false})
		return
	}
	PPageTmp, notFound := service.QueryPPageByPPageName(data.PPageName, ppage.ProjID)
	if !notFound && ppage.PPageID != PPageTmp.PPageID {
		c.JSON(http.StatusOK, response.UpdatePPageA{Message: "页面名已存在，同一设计原型中不能有同名页面", Success: false})
		return
	}
	if data.PPageName != "" {
		ppage.PPageName = data.PPageName
	}
	if data.PPageData != "" {
		ppage.PPageData = data.PPageData
	}
	err := service.UpdatePPage(&ppage)
	if err != nil {
		c.JSON(http.StatusOK, response.UpdatePPageA{Message: "修改页面失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.UpdatePPageA{Message: "修改页面成功", Success: true})
}

// MovePPageToBin
// @Summary 设计原型移入回收站
// @Tags 设计原型
// @Accept json
// @Produce json
// @Param data body response.MovePPageToBinQ true "设计原型的页面ID"
// @Success 200 {object} response.MovePPageToBinA
// @Router /ppage/move_ppage_to_bin [post]
func MovePPageToBin(c *gin.Context) {
	var data response.MovePPageToBinQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.MovePPageToBinA{Message: "输入数据不符合要求", Success: false})
		return
	}
	PPage, notFound := service.QueryPPageByPPageID(data.PPageID)
	if notFound {
		c.JSON(http.StatusOK, response.MovePPageToBinA{Message: "设计原型不存在", Success: false})
		return
	}
	PPage.Status = 2
	err := service.UpdatePPage(&PPage)
	if err != nil {
		c.JSON(http.StatusOK, response.MovePPageToBinA{Message: "移入回收站失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.MovePPageToBinA{Message: "移入回收站成功", Success: true})
}

// GenSharedPPageToken
// @Summary 获取设计原型的共享页面的token
// @Tags 设计原型
// @Accept json
// @Produce json
// @Param data body response.GenSharedPPageTokenQ true "设计原型的页面ID"
// @Success 200 {object} response.GenSharedPPageTokenA
// @Router /ppage/gen_shared_ppage_token [post]
func GenSharedPPageToken(c *gin.Context) {
	var data response.GenSharedPPageTokenQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.GenSharedPPageTokenA{
			Message: "输入数据不符合要求",
			Success: false,
		})
		return
	}
	pPage, notFound := service.QueryPPageByPPageID(data.PPageID)
	if notFound {
		c.JSON(http.StatusOK, response.GenSharedPPageTokenA{
			Message: "设计原型不存在",
			Success: false,
		})
		return
	}
	if pPage.Status == 2 {
		c.JSON(http.StatusOK, response.GenSharedPPageTokenA{
			Message: "设计原型已经移入回收站",
			Success: false,
		})
		return
	}
	raw := fmt.Sprintf("%d", data.PPageID) + time.Now().String()
	md5 := utils.GetMd5(raw)
	sharedPPage := database.SharedPPage{
		PPageID:   data.PPageID,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Token:     md5,
	}
	err := service.CreateSharedPPage(&sharedPPage)
	if err != nil {
		c.JSON(http.StatusOK, response.GenSharedPPageTokenA{
			Message: "生成token失败",
			Success: false,
		})
		return
	}
	c.JSON(http.StatusOK, response.GenSharedPPageTokenA{
		Message: "生成token成功",
		Success: true,
		Token:   md5,
	})
}

// GetSharedPPage
// @Summary 获取设计原型的共享页面
// @Tags 设计原型
// @Accept json
// @Produce json
// @Param data body response.GetSharedPPageQ true "设计原型的页面ID"
// @Success 200 {object} response.GetSharedPPageA
// @Router /get_shared_ppage [post]
func GetSharedPPage(c *gin.Context) {
	var data response.GetSharedPPageQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.GetSharedPPageA{
			Message: "输入数据不符合要求",
			Success: false,
		})
		return
	}
	sharedPPage, notFound := service.QuerySharedPPageByToken(data.Token)
	if notFound {
		c.JSON(http.StatusOK, response.GetSharedPPageA{
			Message: "共享页面不存在",
			Success: false,
		})
		return
	}
	if sharedPPage.ExpiresAt <= time.Now().Unix() {
		c.JSON(http.StatusOK, response.GetSharedPPageA{
			Message: "共享页面已过期",
			Success: false,
		})
		_ = service.DeleteSharedPPage(&sharedPPage)
		return
	}
	pPage, notFound := service.QueryPPageByPPageID(sharedPPage.PPageID)
	if notFound {
		c.JSON(http.StatusOK, response.GetSharedPPageA{
			Message: "设计原型不存在",
			Success: false,
		})
		return
	}
	c.JSON(http.StatusOK, response.GetSharedPPageA{
		Message: "获取共享页面成功",
		Success: true,
		PPage:   pPage,
	})
}

// CloseSharedPPage
// @Summary 关闭设计原型的共享页面
// @Tags 设计原型
// @Accept json
// @Produce json
// @Param data body response.CloseSharedPPageQ true "设计原型的页面ID"
// @Success 200 {object} response.CloseSharedPPageA
// @Router /ppage/close_shared_ppage [post]
func CloseSharedPPage(c *gin.Context) {
	var data response.CloseSharedPPageQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.CloseSharedPPageA{
			Message: "输入数据不符合要求",
			Success: false,
		})
		return
	}
	_, notFound := service.QueryPPageByPPageID(data.PPageID)
	if notFound {
		c.JSON(http.StatusOK, response.CloseSharedPPageA{
			Message: "设计原型不存在",
			Success: false,
		})
		return
	}
	sharedPPage, notFound := service.QuerySharedPPageByPPageID(data.PPageID)
	if notFound {
		c.JSON(http.StatusOK, response.CloseSharedPPageA{
			Message: "共享页面不存在",
			Success: false,
		})
		return
	}
	_ = service.DeleteSharedPPage(&sharedPPage)
	c.JSON(http.StatusOK, response.CloseSharedPPageA{
		Message: "关闭共享页面成功",
		Success: true,
	})
}
