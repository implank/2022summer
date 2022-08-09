package v1

import (
	"2022summer/model/database"
	"2022summer/model/response"
	"2022summer/service"
	"2022summer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateProj
// @Summary 创建项目
// @Tags 项目管理
// @Accept json
// @Produce json
// @Param data body response.CreateProjQ true "项目名称，项目详情（可选），所属团队ID，是否置顶（可选，1 不置顶、2 置顶）"
// @Success 200 {object} response.CreateProjA
// @Router /proj/create_proj [post]
func CreateProj(c *gin.Context) {
	poster, _ := c.Get("user")
	var data response.CreateProjQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.CreateProjA{Message: "输入数据不符合要求", Success: false})
		return
	}
	if _, notFound := service.QueryProjByProjName(data.ProjName, data.GroupID); !notFound {
		c.JSON(http.StatusOK, response.CreateProjA{Message: "项目名已存在", Success: false})
		return
	}
	err := service.CreateProj(&database.Proj{
		ProjName: data.ProjName,
		ProjInfo: data.ProjInfo,
		GroupID:  data.GroupID,
		UserID:   poster.(database.User).UserID,
		Top:      data.Top})
	if err != nil {
		c.JSON(http.StatusOK, response.CreateProjA{Message: "创建项目失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.CreateProjA{Message: "创建项目成功", Success: true})
}

// UpdateProj
// @Summary 修改项目名称、项目描述、是否置顶
// @Tags 项目管理
// @Accept json
// @Produce json
// @Param data body response.UpdateProjQ true "项目ID，项目名称（必填，可以填原名，不能和其他项目同名），项目详情（可选），是否置顶（可选，1 不置顶、2 置顶）"
// @Success 200 {object} response.UpdateProjA
// @Router /proj/update_proj [post]
func UpdateProj(c *gin.Context) {
	var data response.UpdateProjQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.UpdateProjA{Message: "输入数据不符合要求", Success: false})
		return
	}
	proj, notFound := service.QueryProjByProjID(data.ProjID)
	if notFound {
		c.JSON(http.StatusOK, response.UpdateProjA{Message: "项目不存在", Success: false})
		return
	}
	projTmp, notFound := service.QueryProjByProjName(data.ProjName, proj.GroupID)
	if !notFound && proj.ProjID != projTmp.ProjID {
		c.JSON(http.StatusOK, response.UpdateProjA{Message: "项目名已存在", Success: false})
		return
	}
	proj.ProjName = data.ProjName
	proj.ProjInfo = data.ProjInfo
	if data.Top == 1 || data.Top == 2 {
		proj.Top = data.Top
	}
	err := service.UpdateProj(&proj)
	if err != nil {
		c.JSON(http.StatusOK, response.UpdateProjA{Message: "修改项目失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.UpdateProjA{Message: "修改项目成功", Success: true})
}

// MoveProjBin
// @Summary 移入或移出回收站
// @Tags 项目管理
// @Accept json
// @Produce json
// @Param data body response.MoveProjBinQ true "项目ID"
// @Success 200 {object} response.MoveProjBinA
// @Router /proj/move_proj_to_bin [post]
func MoveProjBin(c *gin.Context) {
	var data response.MoveProjBinQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.MoveProjBinA{Message: "输入数据不符合要求", Success: false})
		return
	}
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

// GetProjAll
// @Summary 不在回收站的某组"全部项目"
// @Tags 项目管理
// @Accept json
// @Produce json
// @Param data body response.GetProjAllQ true "团队ID，排序方式（1 按创建时间排序、2 按修改时间排序、3 按编辑次数），是否降序"
// @Success 200 {object} response.GetProjAllA
// @Router /proj/get_proj_all [post]
func GetProjAll(c *gin.Context) {
	poster, _ := c.Get("user")
	var data response.GetProjAllQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.GetProjAllA{Message: "输入数据不符合要求", Success: false})
		return
	}
	projs := service.GetUserProjsInGroup(poster.(database.User).UserID, data.GroupID, 1, 3, data.OrderBy, data.IsDesc)
	// projs := service.GetUserProjs(poster.(database.User).UserID, 1, 3)
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
// @Summary 不在回收站的某组"我创建的"
// @Tags 项目管理
// @Accept json
// @Produce json
// @Param data body response.GetProjCreateQ true "团队ID，排序方式（1 按创建时间排序、2 按修改时间排序、3 按编辑次数），是否降序"
// @Success 200 {object} response.GetProjCreateA
// @Router /proj/get_proj_create [post]
func GetProjCreate(c *gin.Context) {
	poster, _ := c.Get("user")
	var data response.GetProjCreateQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.GetProjCreateA{Message: "输入数据不符合要求", Success: false})
		return
	}
	projs := service.GetUserProjsInGroup(poster.(database.User).UserID, data.GroupID, 1, 1, data.OrderBy, data.IsDesc)
	// projs := service.GetUserProjs(poster.(database.User).UserID, 1, 1)
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
// @Summary 不在回收站的某组"我参与的"
// @Tags 项目管理
// @Accept json
// @Produce json
// @Param data body response.GetProjJoinQ true "团队ID，排序方式（1 按创建时间排序、2 按修改时间排序、3 按编辑次数），是否降序"
// @Success 200 {object} response.GetProjJoinA
// @Router /proj/get_proj_join [post]
func GetProjJoin(c *gin.Context) {
	poster, _ := c.Get("user")
	var data response.GetProjJoinQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.GetProjJoinA{Message: "输入数据不符合要求", Success: false})
		return
	}
	projs := service.GetUserProjsInGroup(poster.(database.User).UserID, data.GroupID, 1, 2, data.OrderBy, data.IsDesc)
	// projs := service.GetUserProjs(poster.(database.User).UserID, 1, 2)
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
// @Summary 搜索框，只搜不在回收站的
// @Tags 项目管理
// @Accept json
// @Produce json
// @Param data body response.GetProjByNameQ true "项目名称（不一定是全名，子串搜索，为空时返回数据库中全部不在回收站的项目）"
// @Success 200 {object} response.GetProjByNameA
// @Router /proj/get_proj_by_name [post]
func GetProjByName(c *gin.Context) {
	var data response.GetProjByNameQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.GetProjByNameA{Message: "输入数据不符合要求", Success: false})
		return
	}
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

// GetProjByID
// @Summary 获取项目信息
// @Tags 项目管理
// @Accept json
// @Produce json
// @Param data body response.GetProjByIDQ true "项目ID"
// @Success 200 {object} response.GetProjByIDA
// @Router /proj/get_proj_by_id [post]
func GetProjByID(c *gin.Context) {
	var data response.GetProjByIDQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.GetProjByIDA{Message: "输入数据不符合要求", Success: false})
		return
	}
	proj, notFound := service.QueryProjByProjID(data.ProjID)
	if notFound {
		c.JSON(http.StatusOK, response.GetProjByIDA{Message: "项目不存在", Success: false, Proj: proj})
		return
	}
	c.JSON(http.StatusOK, response.GetProjByIDA{Message: "成功获取项目信息", Success: true, Proj: proj})
}

// CopyProj
// @Summary 复制项目
// @Tags 项目管理
// @Accept json
// @Produce json
// @Param data body response.CopyProjQ true "项目ID"
// @Success 200 {object} response.CopyProjA
// @Router /proj/copy_proj [post]
func CopyProj(c *gin.Context) {
	poster, _ := c.Get("user")
	var data response.CopyProjQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.CopyProjA{Message: "输入数据不符合要求", Success: false})
		return
	}
	proj, err := service.CopyProj(data.ProjID, poster.(database.User).UserID)
	if err != nil {
		c.JSON(http.StatusOK, response.CopyProjA{Message: "复制项目失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.CopyProjA{Message: "复制项目成功", Success: false, Proj: proj})
}
