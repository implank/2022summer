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

// GetProjByID
// @Summary 获取项目信息
// @Tags 项目管理的第二页
// @Accept json
// @Produce json
// @Param data body response.GetProjByIDQ true "这个接口还没测"
// @Success 200 {object} response.GetProjByIDA
// @Router /file/get_proj_by_id [post]
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

// GetProjPrototypes
// @Summary 获取项目的设计原型
// @Tags 项目管理的第二页
// @Accept json
// @Produce json
// @Param data body response.GetProjPrototypesQ true "这个接口还没测"
// @Success 200 {object} response.GetProjPrototypesA
// @Router /file/get_proj_prototypes [post]
func GetProjPrototypes(c *gin.Context) {
	var data response.GetProjPrototypesQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.GetProjPrototypesA{Message: "输入数据不符合要求", Success: false})
		return
	}
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

// GetProjUmls
// @Summary 获取项目的 Uml 图
// @Tags 项目管理的第二页
// @Accept json
// @Produce json
// @Param data body response.GetProjUmlsQ true "这个接口还没测"
// @Success 200 {object} response.GetProjUmlsA
// @Router /file/get_proj_umls [post]
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
		Message: "成功搜索到以下设计原型",
		Success: true,
		Count:   uint64(x),
		Umls:    umls})
}

// GetProjDocuments
// @Summary 获取项目的文档
// @Tags 项目管理的第二页
// @Accept json
// @Produce json
// @Param data body response.GetProjDocumentsQ true "这个接口还没测"
// @Success 200 {object} response.GetProjDocumentsA
// @Router /file/get_proj_documents [post]
func GetProjDocuments(c *gin.Context) {
	var data response.GetProjDocumentsQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.GetProjDocumentsA{Message: "输入数据不符合要求", Success: false})
		return
	}
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

// CreatePrototype
// @Summary 创建设计原型
// @Tags 项目管理的第二页
// @Accept json
// @Produce json
// @Param data body response.CreatePrototypeQ true "这个接口还没测"
// @Success 200 {object} response.CreatePrototypeA
// @Router /file/create_prototype [post]
func CreatePrototype(c *gin.Context) {
	var data response.CreatePrototypeQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.CreatePrototypeA{Message: "输入数据不符合要求", Success: false})
		return
	}
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

// CreateUml
// @Summary 创建 Uml
// @Tags 项目管理的第二页
// @Accept json
// @Produce json
// @Param data body response.CreateUmlQ true "这个接口还没测"
// @Success 200 {object} response.CreateUmlA
// @Router /file/create_uml [post]
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

// CreateDocument
// @Summary 创建文档
// @Tags 项目管理的第二页
// @Accept json
// @Produce json
// @Param data body response.CreateDocumentQ true "这个接口还没测"
// @Success 200 {object} response.CreateDocumentA
// @Router /file/create_document [post]
func CreateDocument(c *gin.Context) {
	var data response.CreateDocumentQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.CreateDocumentA{Message: "输入数据不符合要求", Success: false})
		return
	}
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

// UpdatePrototype
// @Summary 修改设计原型名称
// @Tags 项目管理的第二页
// @Accept json
// @Produce json
// @Param data body response.UpdatePrototypeQ true "这个接口还没测"
// @Success 200 {object} response.UpdatePrototypeA
// @Router /file/update_prototype [post]
func UpdatePrototype(c *gin.Context) {
	var data response.UpdatePrototypeQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.UpdatePrototypeA{Message: "输入数据不符合要求", Success: false})
		return
	}
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

// UpdateUml
// @Summary 修改 Uml 名称
// @Tags 项目管理的第二页
// @Accept json
// @Produce json
// @Param data body response.UpdateUmlQ true "这个接口还没测"
// @Success 200 {object} response.UpdateUmlA
// @Router /file/update_uml [post]
func UpdateUml(c *gin.Context) {
	var data response.UpdateUmlQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.UpdateUmlA{Message: "输入数据不符合要求", Success: false})
		return
	}
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

// UpdateDocument
// @Summary 修改文档名称
// @Tags 项目管理的第二页
// @Accept json
// @Produce json
// @Param data body response.UpdateDocumentQ true "这个接口还没测"
// @Success 200 {object} response.UpdateDocumentA
// @Router /file/update_document [post]
func UpdateDocument(c *gin.Context) {
	var data response.UpdateDocumentQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.UpdateDocumentA{Message: "输入数据不符合要求", Success: false})
		return
	}
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

// MovePrototypeToBin
// @Summary 设计原型移入回收站
// @Tags 项目管理的第二页
// @Accept json
// @Produce json
// @Param data body response.MovePrototypeToBinQ true "这个接口还没测"
// @Success 200 {object} response.MovePrototypeToBinA
// @Router /file/move_prototype_to_bin [post]
func MovePrototypeToBin(c *gin.Context) {
	var data response.MovePrototypeToBinQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.MovePrototypeToBinA{Message: "输入数据不符合要求", Success: false})
		return
	}
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

// MoveUmlToBin
// @Summary Uml 移入回收站
// @Tags 项目管理的第二页
// @Accept json
// @Produce json
// @Param data body response.MoveUmlToBinQ true "这个接口还没测"
// @Success 200 {object} response.MoveUmlToBinA
// @Router /file/move_uml_to_bin [post]
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

// MoveDocumentToBin
// @Summary 文档移入回收站
// @Tags 项目管理的第二页
// @Accept json
// @Produce json
// @Param data body response.MoveDocumentToBinQ true "这个接口还没测"
// @Success 200 {object} response.MoveDocumentToBinA
// @Router /file/move_document_to_bin [post]
func MoveDocumentToBin(c *gin.Context) {
	var data response.MoveDocumentToBinQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.MoveDocumentToBinA{Message: "输入数据不符合要求", Success: false})
		return
	}
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

// GetFilesByName
// @Summary 搜索框
// @Tags 项目管理的第二页
// @Accept json
// @Produce json
// @Param data body response.GetFilesByNameQ true "这个接口还没测"
// @Success 200 {object} response.GetFilesByNameA
// @Router /file/get_files_by_name [post]
func GetFilesByName(c *gin.Context) {
	var data response.GetFilesByNameQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.GetFilesByNameA{Message: "输入数据不符合要求", Success: false})
		return
	}
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
