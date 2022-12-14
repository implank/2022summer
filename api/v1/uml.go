package v1

import (
	"2022summer/model/database"
	"2022summer/model/response"
	"2022summer/service"
	"2022summer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetProjUmls
// @Summary 获取某个项目的不在回收站的 Uml
// @Tags Uml
// @Accept json
// @Produce json
// @Param data body response.GetProjUmlsQ true "项目ID"
// @Success 200 {object} response.GetProjUmlsA
// @Router /uml/get_proj_umls [post]
func GetProjUmls(c *gin.Context) {
	var data response.GetProjUmlsQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.GetProjUmlsA{Message: "输入数据不符合要求", Success: false})
		return
	}
	umls := service.GetProjUmls(data.ProjID, 1)
	x := len(umls)
	if x == 0 {
		c.JSON(http.StatusOK, response.GetProjUmlsA{
			Message: "没有找到捏",
			Success: false,
			Count:   uint64(x),
			Umls:    umls})
		return
	}
	c.JSON(http.StatusOK, response.GetProjUmlsA{
		Message: "成功搜索到以下Uml",
		Success: true,
		Count:   uint64(x),
		Umls:    umls})
}

// GetUmlByID
// @Summary 获取某个设计原型
// @Tags Uml
// @Accept json
// @Produce json
// @Param data body response.GetUmlByIDQ true "Uml ID"
// @Success 200 {object} response.GetUmlByIDA
// @Router /uml/get_uml_by_id [post]
func GetUmlByID(c *gin.Context) {
	var data response.GetUmlByIDQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.GetUmlByIDA{Message: "输入数据不符合要求", Success: false})
		return
	}
	uml, notFound := service.QueryUmlByUmlID(data.UmlID)
	if notFound {
		c.JSON(http.StatusOK, response.GetUmlByIDA{Message: "页面不存在", Success: false, Uml: uml})
		return
	}
	c.JSON(http.StatusOK, response.GetUmlByIDA{Message: "成功获取页面信息", Success: true, Uml: uml})
}

// CreateUml
// @Summary 创建 Uml
// @Tags Uml
// @Accept json
// @Produce json
// @Param data body response.CreateUmlQ true "Uml名称，Uml所属项目ID"
// @Success 200 {object} response.CreateUmlA
// @Router /uml/create_uml [post]
func CreateUml(c *gin.Context) {
	var data response.CreateUmlQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.CreateUmlA{Message: "输入数据不符合要求", Success: false})
		return
	}
	if _, notFound := service.QueryUmlByUmlName(data.UmlName, data.ProjID); !notFound {
		c.JSON(http.StatusOK, response.CreateUmlA{Message: "Uml名已存在", Success: false})
		return
	}
	/*raw := fmt.Sprintf("%d", data.ProjID) + time.Now().String() + data.UmlName
	md5 := utils.GetMd5(raw)
	dir := "./media/umls"
	name := md5 + ".meow"
	filePath := path.Join(dir, name)
	file, err := os.Create(filePath)
	defer utils.CloseFile(file)*/
	err := service.CreateUml(&database.Uml{
		UmlName: data.UmlName,
		// UmlURL:  "http://" + global.VP.GetString("host") + "/media/umls/" + name,
		ProjID: data.ProjID})
	if err != nil {
		c.JSON(http.StatusOK, response.CreateUmlA{Message: "创建Uml失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.CreateUmlA{Message: "创建Uml成功", Success: true})
}

// UpdateUml
// @Summary 修改 Uml 名称
// @Tags Uml
// @Accept json
// @Produce json
// @Param data body response.UpdateUmlQ true "Uml ID，Uml的新名字（必填，可以填原名，不能和其他项目同名）"
// @Success 200 {object} response.UpdateUmlA
// @Router /uml/update_uml [post]
func UpdateUml(c *gin.Context) {
	var data response.UpdateUmlQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.UpdateUmlA{Message: "输入数据不符合要求", Success: false})
		return
	}
	uml, notFound := service.QueryUmlByUmlID(data.UmlID)
	if notFound {
		c.JSON(http.StatusOK, response.UpdateUmlA{Message: "Uml不存在", Success: false})
		return
	}
	umlTmp, notFound := service.QueryUmlByUmlName(data.UmlName, uml.ProjID)
	if !notFound && uml.UmlID != umlTmp.UmlID {
		c.JSON(http.StatusOK, response.UpdateUmlA{Message: "Uml名已存在，同一项目中不能有同名Uml", Success: false})
		return
	}
	uml.UmlName = data.UmlName
	err := service.UpdateUml(&uml)
	if err != nil {
		c.JSON(http.StatusOK, response.UpdateUmlA{Message: "修改Uml失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.UpdateUmlA{Message: "修改Uml成功", Success: true})
}

// UploadUml
// @Summary Uml 修改 Uml 内容
// @Tags Uml
// @Accept json
// @Produce json
// @Param data body response.UploadUmlQ true "Uml ID，Uml 数据"
// @Success 200 {object} response.UploadUmlA
// @Router /uml/upload_uml [post]
func UploadUml(c *gin.Context) {
	var data response.UploadUmlQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.UploadUmlA{Message: "输入数据不符合要求", Success: false})
		return
	}
	uml, notFound := service.QueryUmlByUmlID(data.UmlID)
	if notFound {
		c.JSON(http.StatusOK, response.UploadUmlA{Message: "Uml不存在", Success: false})
		return
	}
	uml.UmlData = data.UmlData
	err := service.UpdateUml(&uml)
	if err != nil {
		c.JSON(http.StatusOK, response.UploadUmlA{Message: "修改Uml失败", Success: false})
		return
	}
}

// MoveUmlToBin
// @Summary Uml 移入回收站
// @Tags Uml
// @Accept json
// @Produce json
// @Param data body response.MoveUmlToBinQ true "Uml ID"
// @Success 200 {object} response.MoveUmlToBinA
// @Router /uml/move_uml_to_bin [post]
func MoveUmlToBin(c *gin.Context) {
	var data response.MoveUmlToBinQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.MoveUmlToBinA{Message: "输入数据不符合要求", Success: false})
		return
	}
	uml, notFound := service.QueryUmlByUmlID(data.UmlID)
	if notFound {
		c.JSON(http.StatusOK, response.MoveUmlToBinA{Message: "Uml不存在", Success: false})
		return
	}
	uml.Status = 2
	err := service.UpdateUml(&uml)
	if err != nil {
		c.JSON(http.StatusOK, response.MoveUmlToBinA{Message: "移入回收站失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.MoveUmlToBinA{Message: "移入回收站成功", Success: true})
}
