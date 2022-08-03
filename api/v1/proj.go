package v1

import (
	"2022summer/model"
	"2022summer/model/response"
	"2022summer/service"
	"2022summer/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"time"
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
	data := utils.BindJsonAndValid(c, &response.GetProjByIDQ{}).(*response.GetProjByIDQ)
	proj, notFound := service.QueryProjByProjID(data.ProjID)
	if notFound {
		c.JSON(http.StatusOK, response.GetProjByIDA{Message: "项目不存在", Success: false, Proj: proj})
		return
	}
	c.JSON(http.StatusOK, response.GetProjByIDA{Message: "成功获取项目信息", Success: true, Proj: proj})
}

func GetProjPrototypes(c *gin.Context) { // 获取项目的设计原型
	data := utils.BindJsonAndValid(c, &response.GetProjPrototypesQ{}).(*response.GetProjPrototypesQ)
	prototypes := service.GetProjPrototypes(data.ProjID, 1)
	x := len(prototypes)
	if x == 0 {
		c.JSON(http.StatusOK, response.GetProjPrototypesA{
			Message:    "没有找到捏",
			Success:    false,
			Count:      uint64(x),
			Prototypes: prototypes})
		return
	}
	c.JSON(http.StatusOK, response.GetProjPrototypesA{
		Message:    "成功搜索到以下设计原型",
		Success:    true,
		Count:      uint64(x),
		Prototypes: prototypes})
}

func GetProjUmls(c *gin.Context) { // 获取项目的 Uml 图
	data := utils.BindJsonAndValid(c, &response.GetProjUmlsQ{}).(*response.GetProjUmlsQ)
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
		Message: "成功搜索到以下设计原型",
		Success: true,
		Count:   uint64(x),
		Umls:    umls})
}

func GetProjDocuments(c *gin.Context) { // 获取项目的文档
	data := utils.BindJsonAndValid(c, &response.GetProjDocumentsQ{}).(*response.GetProjDocumentsQ)
	documents := service.GetProjDocuments(data.ProjID, 1)
	x := len(documents)
	if x == 0 {
		c.JSON(http.StatusOK, response.GetProjDocumentsA{
			Message:   "没有找到捏",
			Success:   false,
			Count:     uint64(x),
			Documents: documents})
		return
	}
	c.JSON(http.StatusOK, response.GetProjDocumentsA{
		Message:   "成功搜索到以下设计原型",
		Success:   true,
		Count:     uint64(x),
		Documents: documents})
}

func CreatePrototype(c *gin.Context) { // 创建设计原型
	data := utils.BindJsonAndValid(c, &response.CreatePrototypeQ{}).(*response.CreatePrototypeQ)
	if _, notFound := service.QueryPrototypeByPrototypeName(data.PrototypeName, data.ProjID); !notFound {
		c.JSON(http.StatusOK, response.CreatePrototypeA{Message: "设计原型名已存在", Success: false})
		return
	}
	raw := fmt.Sprintf("%d", data.ProjID) + time.Now().String() + data.PrototypeName
	md5 := utils.GetMd5(raw)
	dir := "./media/prototypes"
	name := md5 + ".meow" // TODO 随便起的后缀名
	filePath := path.Join(dir, name)
	file, err := os.Create(filePath)
	defer utils.CloseFile(file)
	if err != nil {
		c.JSON(http.StatusOK, response.CreatePrototypeA{Message: "创建设计原型失败", Success: false})
		return
	}
	err = service.CreatePrototype(&model.Prototype{
		PrototypeName: data.PrototypeName,
		PrototypeURL:  "http://43.138.77.133:81/media/prototypes/" + name,
		ProjID:        data.ProjID})
	if err != nil {
		c.JSON(http.StatusOK, response.CreatePrototypeA{Message: "创建设计原型失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.CreatePrototypeA{Message: "创建设计原型成功", Success: true})
}

func CreateUml(c *gin.Context) { // 创建 Uml
	data := utils.BindJsonAndValid(c, &response.CreateUmlQ{}).(*response.CreateUmlQ)
	if _, notFound := service.QueryUmlByUmlName(data.UmlName, data.ProjID); !notFound {
		c.JSON(http.StatusOK, response.CreateUmlA{Message: "Uml名已存在", Success: false})
		return
	}
	raw := fmt.Sprintf("%d", data.ProjID) + time.Now().String() + data.UmlName
	md5 := utils.GetMd5(raw)
	dir := "./media/umls"
	name := md5 + ".meow" // TODO 随便起的后缀名
	filePath := path.Join(dir, name)
	file, err := os.Create(filePath)
	defer utils.CloseFile(file)
	err = service.CreateUml(&model.Uml{
		UmlName: data.UmlName,
		UmlURL:  "http://43.138.77.133:81/media/umls/" + name,
		ProjID:  data.ProjID})
	if err != nil {
		c.JSON(http.StatusOK, response.CreateUmlA{Message: "创建Uml失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.CreateUmlA{Message: "创建Uml成功", Success: true})
}

func CreateDocument(c *gin.Context) { // 创建文档
	data := utils.BindJsonAndValid(c, &response.CreateDocumentQ{}).(*response.CreateDocumentQ)
	if _, notFound := service.QueryDocumentByDocumentName(data.DocumentName, data.ProjID); !notFound {
		c.JSON(http.StatusOK, response.CreateDocumentA{Message: "文档名已存在", Success: false})
		return
	}
	raw := fmt.Sprintf("%d", data.ProjID) + time.Now().String() + data.DocumentName
	md5 := utils.GetMd5(raw)
	dir := "./media/documents"
	name := md5 + ".md"
	filePath := path.Join(dir, name)
	file, err := os.Create(filePath)
	defer utils.CloseFile(file)
	err = service.CreateDocument(&model.Document{
		DocumentName: data.DocumentName,
		DocumentURL:  "http://43.138.77.133:81/media/documents/" + name,
		ProjID:       data.ProjID})
	if err != nil {
		c.JSON(http.StatusOK, response.CreateDocumentA{Message: "创建文档失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.CreateDocumentA{Message: "创建文档成功", Success: true})
}

func UpdatePrototype(c *gin.Context) { // 修改设计原型名称
	data := utils.BindJsonAndValid(c, &response.UpdatePrototypeQ{}).(*response.UpdatePrototypeQ)
	prototype, notFound := service.QueryPrototypeByPrototypeID(data.PrototypeID)
	if notFound {
		c.JSON(http.StatusOK, response.UpdatePrototypeA{Message: "设计原型不存在", Success: false})
		return
	}
	prototypeTmp, notFound := service.QueryPrototypeByPrototypeName(data.PrototypeName, prototype.ProjID)
	if !notFound && prototype.PrototypeID != prototypeTmp.PrototypeID {
		c.JSON(http.StatusOK, response.UpdatePrototypeA{Message: "设计原型名已存在，同一项目中不能有同名设计原型", Success: false})
		return
	}
	prototype.PrototypeName = data.PrototypeName
	err := service.UpdatePrototype(&prototype)
	if err != nil {
		c.JSON(http.StatusOK, response.UpdatePrototypeA{Message: "修改设计原型失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.UpdatePrototypeA{Message: "修改设计原型成功", Success: true})
}

func UpdateUml(c *gin.Context) { // 修改 Uml 名称
	data := utils.BindJsonAndValid(c, &response.UpdateUmlQ{}).(*response.UpdateUmlQ)
	uml, notFound := service.QueryUmlByUmlID(data.UmlID)
	if notFound {
		c.JSON(http.StatusOK, response.UpdateUmlA{Message: "Unl不存在", Success: false})
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

func UpdateDocument(c *gin.Context) { // 修改文档名称
	data := utils.BindJsonAndValid(c, &response.UpdateDocumentQ{}).(*response.UpdateDocumentQ)
	document, notFound := service.QueryDocumentByDocumentID(data.DocumentID)
	if notFound {
		c.JSON(http.StatusOK, response.UpdateDocumentA{Message: "文档不存在", Success: false})
		return
	}
	documentTmp, notFound := service.QueryDocumentByDocumentName(data.DocumentName, document.ProjID)
	if !notFound && document.DocumentID != documentTmp.DocumentID {
		c.JSON(http.StatusOK, response.UpdateDocumentA{Message: "文档已存在，同一项目中不能有同名文档", Success: false})
		return
	}
	document.DocumentName = data.DocumentName
	err := service.UpdateDocument(&document)
	if err != nil {
		c.JSON(http.StatusOK, response.UpdateDocumentA{Message: "修改文档失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.UpdateDocumentA{Message: "修改文档成功", Success: true})
}

func MovePrototypeToBin(c *gin.Context) { // 设计原型移入回收站
	data := utils.BindJsonAndValid(c, &response.MovePrototypeToBinQ{}).(*response.MovePrototypeToBinQ)
	prototype, notFound := service.QueryPrototypeByPrototypeID(data.PrototypeID)
	if notFound {
		c.JSON(http.StatusOK, response.MovePrototypeToBinA{Message: "设计原型不存在", Success: false})
		return
	}
	prototype.Status = 2
	err := service.UpdatePrototype(&prototype)
	if err != nil {
		c.JSON(http.StatusOK, response.MovePrototypeToBinA{Message: "移入回收站失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.MovePrototypeToBinA{Message: "移入回收站成功", Success: true})
}

func MoveUmlToBin(c *gin.Context) { // Uml 移入回收站
	data := utils.BindJsonAndValid(c, &response.MoveUmlToBinQ{}).(*response.MoveUmlToBinQ)
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

func MoveDocumentToBin(c *gin.Context) { // 文档移入回收站
	data := utils.BindJsonAndValid(c, &response.MoveDocumentToBinQ{}).(*response.MoveDocumentToBinQ)
	document, notFound := service.QueryDocumentByDocumentID(data.DocumentID)
	if notFound {
		c.JSON(http.StatusOK, response.MoveDocumentToBinA{Message: "文档不存在", Success: false})
		return
	}
	document.Status = 2
	err := service.UpdateDocument(&document)
	if err != nil {
		c.JSON(http.StatusOK, response.MoveDocumentToBinA{Message: "移入回收站失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.MoveDocumentToBinA{Message: "移入回收站成功", Success: true})
}

func DeletePrototype(c *gin.Context) { // 删除设计原型
	data := utils.BindJsonAndValid(c, &response.DeletePrototypeQ{}).(*response.DeletePrototypeQ)
	prototype, notFound := service.QueryPrototypeByPrototypeID(data.PrototypeID)
	if notFound {
		c.JSON(http.StatusOK, response.DeletePrototypeA{Message: "设计原型不存在", Success: false})
		return
	}
	err := service.DeletePrototype(&prototype)
	if err != nil {
		c.JSON(http.StatusOK, response.DeletePrototypeA{Message: "删除设计原型失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.DeletePrototypeA{Message: "删除设计原型成功", Success: true})
}

func DeleteUml(c *gin.Context) { // 删除 Uml
	data := utils.BindJsonAndValid(c, &response.DeleteUmlQ{}).(*response.DeleteUmlQ)
	uml, notFound := service.QueryUmlByUmlID(data.UmlID)
	if notFound {
		c.JSON(http.StatusOK, response.DeleteUmlA{Message: "Uml不存在", Success: false})
		return
	}
	err := service.DeleteUml(&uml)
	if err != nil {
		c.JSON(http.StatusOK, response.DeleteUmlA{Message: "删除Uml失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.DeleteUmlA{Message: "删除Uml成功", Success: true})
}

func DeleteDocument(c *gin.Context) { // 删除文档
	data := utils.BindJsonAndValid(c, &response.DeleteDocumentQ{}).(*response.DeleteDocumentQ)
	document, notFound := service.QueryDocumentByDocumentID(data.DocumentID)
	if notFound {
		c.JSON(http.StatusOK, response.DeleteDocumentA{Message: "文档不存在", Success: false})
		return
	}
	err := service.DeleteDocument(&document)
	if err != nil {
		c.JSON(http.StatusOK, response.DeleteDocumentA{Message: "删除文档失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.DeleteDocumentA{Message: "删除文档成功", Success: true})
}

func GetFilesByName(c *gin.Context) { // 搜索框
	data := utils.BindJsonAndValid(c, &response.GetFilesByNameQ{}).(*response.GetFilesByNameQ)
	prototypes, umls, documents := service.GetFilesByNameBur(data.Name)
	c.JSON(http.StatusOK, response.GetFilesByNameA{
		Message:         "成功搜索到以下项目",
		Success:         true,
		CountPrototypes: uint64(len(prototypes)),
		Prototypes:      prototypes,
		CountUmls:       uint64(len(umls)),
		Umls:            umls,
		CountDocuments:  uint64(len(documents)),
		Documents:       documents})
}

/* * * * * * * * * * * */

func MovePrototypeFromBin(c *gin.Context) { // 设计原型移出回收站
	data := utils.BindJsonAndValid(c, &response.MovePrototypeFromBinQ{}).(*response.MovePrototypeFromBinQ)
	prototype, notFound := service.QueryPrototypeByPrototypeID(data.PrototypeID)
	if notFound {
		c.JSON(http.StatusOK, response.MovePrototypeToBinA{Message: "设计原型不存在", Success: false})
		return
	}
	proj, _ := service.QueryProjByProjID(prototype.ProjID)
	if proj.Status == 2 {
		c.JSON(http.StatusOK, response.MovePrototypeToBinA{Message: "请先将该设计原型所属项目移出回收站", Success: false})
		return
	}
	prototype.Status = 1
	err := service.UpdatePrototype(&prototype)
	if err != nil {
		c.JSON(http.StatusOK, response.MovePrototypeToBinA{Message: "移出回收站失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.MovePrototypeToBinA{Message: "移出回收站成功", Success: true})
}

func MoveUmlFromBin(c *gin.Context) { // Uml 移出回收站
	data := utils.BindJsonAndValid(c, &response.MoveUmlToBinQ{}).(*response.MoveUmlToBinQ)
	uml, notFound := service.QueryUmlByUmlID(data.UmlID)
	if notFound {
		c.JSON(http.StatusOK, response.MoveUmlToBinA{Message: "Uml不存在", Success: false})
		return
	}
	proj, _ := service.QueryProjByProjID(uml.ProjID)
	if proj.Status == 2 {
		c.JSON(http.StatusOK, response.MovePrototypeToBinA{Message: "请先将该Uml所属项目移出回收站", Success: false})
		return
	}
	uml.Status = 1
	err := service.UpdateUml(&uml)
	if err != nil {
		c.JSON(http.StatusOK, response.MoveUmlToBinA{Message: "移出回收站失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.MoveUmlToBinA{Message: "移出回收站成功", Success: true})
}

func MoveDocumentFromBin(c *gin.Context) { // 文档移出回收站
	data := utils.BindJsonAndValid(c, &response.MoveDocumentToBinQ{}).(*response.MoveDocumentToBinQ)
	document, notFound := service.QueryDocumentByDocumentID(data.DocumentID)
	if notFound {
		c.JSON(http.StatusOK, response.MoveDocumentToBinA{Message: "文档不存在", Success: false})
		return
	}
	proj, _ := service.QueryProjByProjID(document.ProjID)
	if proj.Status == 2 {
		c.JSON(http.StatusOK, response.MovePrototypeToBinA{Message: "请先将该文档所属项目移出回收站", Success: false})
		return
	}
	document.Status = 1
	err := service.UpdateDocument(&document)
	if err != nil {
		c.JSON(http.StatusOK, response.MoveDocumentToBinA{Message: "移出回收站失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.MoveDocumentToBinA{Message: "移出回收站成功", Success: true})
}
