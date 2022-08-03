package v1

import (
	"2022summer/model"
	"2022summer/model/response"
	"2022summer/service"
	"2022summer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateProj
// @Summary 创建项目
// @Tags Proj1
// @Accept json
// @Produce json
// @Param data body response.CreateProjQ true "项目名称，项目详情（可选），所属团队ID"
// @Success 200 {object} response.CreateProjA
// @Failure 500 {object} response.CreateProjA
// @Router /create_proj [post]
func CreateProj(c *gin.Context) {
	poster, _ := c.Get("user")
	data := utils.BindJsonAndValid(c, &response.CreateProjQ{}).(*response.CreateProjQ)
	if _, notFound := service.QueryProjByProjName(data.ProjName); !notFound {
		c.JSON(http.StatusOK, response.CreateProjA{Message: "项目名已存在", Success: false})
		return
	}
	err := service.CreateProj(&model.Proj{
		ProjName: data.ProjName,
		ProjInfo: data.ProjInfo,
		GroupID:  data.GroupID,
		UserID:   poster.(model.User).UserID})
	if err != nil {
		c.JSON(http.StatusOK, response.CreateProjA{Message: "创建项目失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.CreateProjA{Message: "创建项目成功", Success: true})
}

// UpdateProj
// @Summary 修改项目名称、项目描述
// @Tags Proj1
// @Accept json
// @Produce json
// @Param data body response.UpdateProjQ true "项目ID，项目名称（必填，可以填原名，不能和其他项目同名），项目详情（可选）"
// @Success 200 {object} response.UpdateProjA
// @Failure 500 {object} response.UpdateProjA
// @Router /update_proj [post]
func UpdateProj(c *gin.Context) {
	data := utils.BindJsonAndValid(c, &response.UpdateProjQ{}).(*response.UpdateProjQ)
	proj, notFound := service.QueryProjByProjID(data.ProjID)
	if notFound {
		c.JSON(http.StatusOK, response.UpdateProjA{Message: "项目不存在", Success: false})
		return
	}
	projTmp, notFound := service.QueryProjByProjName(data.ProjName)
	if !notFound && proj.ProjID != projTmp.ProjID {
		c.JSON(http.StatusOK, response.UpdateProjA{Message: "项目名已存在", Success: false})
		return
	}
	proj.ProjName = data.ProjName
	proj.ProjInfo = data.ProjInfo
	err := service.UpdateProj(&proj)
	if err != nil {
		c.JSON(http.StatusOK, response.UpdateProjA{Message: "修改项目失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.UpdateProjA{Message: "修改项目成功", Success: true})
}

// MoveProjBin
// @Summary 移入或移出回收站
// @Tags Proj1
// @Accept json
// @Produce json
// @Param data body response.MoveProjBinQ true "项目ID"
// @Success 200 {object} response.MoveProjBinA
// @Failure 500 {object} response.MoveProjBinA
// @Router /move_proj_to_bin [post]
func MoveProjBin(c *gin.Context) {
	data := utils.BindJsonAndValid(c, &response.MoveProjBinQ{}).(*response.MoveProjBinQ)
	proj, notFound := service.QueryProjByProjID(data.ProjID)
	if notFound {
		c.JSON(http.StatusOK, response.MoveProjBinA{Message: "项目不存在", Success: false})
		return
	}
	if proj.Status == 1 {
		proj.Status = 2
	} else {
		proj.Status = 1
	}
	err := service.UpdateProjStatus(&proj)
	if err != nil {
		c.JSON(http.StatusOK, response.MoveProjBinA{Message: "移入或移出回收站失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.MoveProjBinA{Message: "移入或移出回收站成功", Success: true})
}

// DeleteProj
// @Summary 删除项目
// @Tags Proj1
// @Accept json
// @Produce json
// @Param data body response.DeleteProjQ true "项目ID"
// @Success 200 {object} response.DeleteProjA
// @Failure 500 {object} response.DeleteProjA
// @Router /delete_proj [post]
func DeleteProj(c *gin.Context) {
	data := utils.BindJsonAndValid(c, &response.DeleteProjQ{}).(*response.DeleteProjQ)
	proj, notFound := service.QueryProjByProjID(data.ProjID)
	if notFound {
		c.JSON(http.StatusOK, response.DeleteProjA{Message: "项目不存在", Success: false})
		return
	}
	err := service.DeleteProj(&proj)
	if err != nil {
		c.JSON(http.StatusOK, response.DeleteProjA{Message: "删除项目失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.DeleteProjA{Message: "删除项目成功", Success: true})
}

// GetProjAll
// @Summary 全部项目
// @Tags Proj1
// @Accept json
// @Produce json
// @Param data body response.GetProjAllQ true "无"
// @Success 200 {object} response.GetProjAllA
// @Failure 500 {object} response.GetProjAllA
// @Router /get_proj_all [post]
func GetProjAll(c *gin.Context) {
	poster, _ := c.Get("user")
	// data := utils.BindJsonAndValid(c, &response.GetProjAllQ{}).(*response.GetProjAllQ)
	projs := service.GetUserProjs(poster.(model.User).UserID, 1, 3)
	// projs := service.GetUserProjsInGroup(poster.(model.User).UserID, data.GroupID, 1, 3)
	x := len(projs)
	if x == 0 {
		c.JSON(http.StatusOK, response.GetProjAllA{
			Message: "您还没有项目捏",
			Success: false,
			Count:   uint64(x),
			Projs:   projs})
		return
	}
	c.JSON(http.StatusOK, response.GetProjAllA{
		Message: "成功获取全部项目",
		Success: true,
		Count:   uint64(x),
		Projs:   projs})
}

// GetProjCreate
// @Summary 我创建的
// @Tags Proj1
// @Accept json
// @Produce json
// @Param data body response.GetProjCreateQ true "无"
// @Success 200 {object} response.GetProjCreateA
// @Failure 500 {object} response.GetProjCreateA
// @Router /get_proj_create [post]
func GetProjCreate(c *gin.Context) {
	poster, _ := c.Get("user")
	projs := service.GetUserProjs(poster.(model.User).UserID, 1, 1)
	x := len(projs)
	if x == 0 {
		c.JSON(http.StatusOK, response.GetProjCreateA{
			Message: "您还没有项目捏",
			Success: false,
			Count:   uint64(x),
			Projs:   projs})
		return
	}
	c.JSON(http.StatusOK, response.GetProjCreateA{
		Message: "成功获取您创建的项目",
		Success: true,
		Count:   uint64(x),
		Projs:   projs})
}

// GetProjJoin
// @Summary 我参与的
// @Tags Proj1
// @Accept json
// @Produce json
// @Param data body response.GetProjJoinQ true "无"
// @Success 200 {object} response.GetProjJoinA
// @Failure 500 {object} response.GetProjJoinA
// @Router /get_proj_join [post]
func GetProjJoin(c *gin.Context) {
	poster, _ := c.Get("user")
	projs := service.GetUserProjs(poster.(model.User).UserID, 1, 2)
	x := len(projs)
	if x == 0 {
		c.JSON(http.StatusOK, response.GetProjJoinA{
			Message: "您还没有项目捏",
			Success: false,
			Count:   uint64(x),
			Projs:   projs})
		return
	}
	c.JSON(http.StatusOK, response.GetProjJoinA{
		Message: "成功获取您参与的项目",
		Success: true,
		Count:   uint64(x),
		Projs:   projs})
}

// GetProjByName
// @Summary 搜索框
// @Tags Proj1
// @Accept json
// @Produce json
// @Param data body response.GetProjByNameQ true "项目名称（不一定是全名，子串搜索）"
// @Success 200 {object} response.GetProjByNameA
// @Failure 500 {object} response.GetProjByNameA
// @Router /get_proj_by_name [post]
func GetProjByName(c *gin.Context) {
	data := utils.BindJsonAndValid(c, &response.GetProjByNameQ{}).(*response.GetProjByNameQ)
	projs := service.GetProjsByProjNameBur(data.ProjName)
	x := len(projs)
	if x == 0 {
		c.JSON(http.StatusOK, response.GetProjByNameA{
			Message: "没有找到捏",
			Success: false,
			Count:   uint64(x),
			Projs:   projs})
		return
	}
	c.JSON(http.StatusOK, response.GetProjByNameA{
		Message: "成功搜索到以下项目",
		Success: true,
		Count:   uint64(x),
		Projs:   projs})
}

/* * * * * * * * * * * */

func GetProjByID(c *gin.Context) { // 获取项目信息

}

func GetProjPrototypes(c *gin.Context) { // 获取项目的设计原型

}

func GetProjUmls(c *gin.Context) { // 获取项目的 Uml 图

}

func GetProjDocuments(c *gin.Context) { // 获取项目的文档

}

func CreatePrototype(c *gin.Context) { // 创建设计原型

}

func CreateUml(c *gin.Context) { // 创建 Uml

}

func CreateDocument(c *gin.Context) { // 创建文档

}

func GetSthByName(c *gin.Context) { // 搜索框

}
