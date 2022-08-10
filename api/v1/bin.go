package v1

import (
	"2022summer/model/database"
	"2022summer/model/response"
	"2022summer/service"
	"2022summer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteProj
// @Summary 删除项目
// @Tags 回收站
// @Accept json
// @Produce json
// @Param data body response.DeleteProjQ true "项目ID"
// @Success 200 {object} response.DeleteProjA
// @Router /bin/delete_proj [post]
func DeleteProj(c *gin.Context) {
	var data response.DeleteProjQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.DeleteProjA{Message: "输入数据不符合要求", Success: false})
		return
	}
	proj, notFound := service.QueryProjByProjID(data.ProjID)
	if notFound {
		c.JSON(http.StatusOK, response.DeleteProjA{Message: "项目不存在", Success: false})
		return
	}
	/*documents := service.GetProjDocuments(data.ProjID, 2)
	for _, value := range documents {
		filename := strings.Split(value.DocumentURL, "/")[len(strings.Split(value.DocumentURL, "/"))-1]
		saveDir := "./media/documents/"
		savePath := path.Join(saveDir, filename)
		if err := os.Remove(savePath); err != nil {
			c.JSON(http.StatusOK, response.DeleteDocumentA{Message: "删除项目失败", Success: false})
			return
		}
	}
	documents = service.GetProjDocuments(data.ProjID, 1)
	for _, value := range documents {
		filename := strings.Split(value.DocumentURL, "/")[len(strings.Split(value.DocumentURL, "/"))-1]
		saveDir := "./media/documents/"
		savePath := path.Join(saveDir, filename)
		if err := os.Remove(savePath); err != nil {
			c.JSON(http.StatusOK, response.DeleteDocumentA{Message: "删除项目失败", Success: false})
			return
		}
	}*/
	err := service.DeleteProj(&proj)
	if err != nil {
		c.JSON(http.StatusOK, response.DeleteProjA{Message: "删除项目失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.DeleteProjA{Message: "删除项目成功", Success: true})
}

// MovePPageFromBin
// @Summary 设计原型的某个页面移出回收站
// @Tags 回收站
// @Accept json
// @Produce json
// @Param data body response.MovePPageFromBinQ true "设计原型的页面ID"
// @Success 200 {object} response.MovePPageFromBinA
// @Router /bin/move_PPage_from_bin [post]
func MovePPageFromBin(c *gin.Context) {
	var data response.MovePPageFromBinQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.MovePPageToBinA{Message: "输入数据不符合要求", Success: false})
		return
	}
	ppage, notFound := service.QueryPPageByPPageID(data.PPageID)
	if notFound {
		c.JSON(http.StatusOK, response.MovePPageToBinA{Message: "设计原型不存在", Success: false})
		return
	}
	proj, _ := service.QueryProjByProjID(ppage.ProjID)
	if proj.Status == 2 {
		c.JSON(http.StatusOK, response.MovePPageToBinA{Message: "请先将该设计原型所属项目移出回收站", Success: false})
		return
	}
	ppage.Status = 1
	err := service.UpdatePPage(&ppage)
	if err != nil {
		c.JSON(http.StatusOK, response.MovePPageToBinA{Message: "移出回收站失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.MovePPageToBinA{Message: "移出回收站成功", Success: true})
}

// MoveUmlFromBin
// @Summary Uml 移出回收站
// @Tags 回收站
// @Accept json
// @Produce json
// @Param data body response.MoveUmlFromBinQ true "UmlID"
// @Success 200 {object} response.MoveUmlFromBinA
// @Router /bin/move_uml_from_bin [post]
func MoveUmlFromBin(c *gin.Context) {
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
	proj, _ := service.QueryProjByProjID(uml.ProjID)
	if proj.Status == 2 {
		c.JSON(http.StatusOK, response.MoveUmlToBinA{Message: "请先将该Uml所属项目移出回收站", Success: false})
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

// MoveDocumentFromBin
// @Summary 文档移出回收站
// @Tags 回收站
// @Accept json
// @Produce json
// @Param data body response.MoveDocumentFromBinQ true "文档ID"
// @Success 200 {object} response.MoveDocumentFromBinA
// @Router /bin/move_document_from_bin [post]
func MoveDocumentFromBin(c *gin.Context) {
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
	proj, _ := service.QueryProjByProjID(document.ProjID)
	if proj.Status == 2 {
		c.JSON(http.StatusOK, response.MoveDocumentToBinA{Message: "请先将该文档所属项目移出回收站", Success: false})
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

// DeletePPage
// @Summary 删除设计原型的某个页面
// @Tags 回收站
// @Accept json
// @Produce json
// @Param data body response.DeletePPageQ true "设计原型的页面ID"
// @Success 200 {object} response.DeletePPageA
// @Router /bin/delete_ppage [post]
func DeletePPage(c *gin.Context) {
	var data response.DeletePPageQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.DeletePPageA{Message: "输入数据不符合要求", Success: false})
		return
	}
	ppage, notFound := service.QueryPPageByPPageID(data.PPageID)
	if notFound {
		c.JSON(http.StatusOK, response.DeletePPageA{Message: "页面不存在", Success: false})
		return
	}
	err := service.DeletePPage(&ppage)
	if err != nil {
		c.JSON(http.StatusOK, response.DeletePPageA{Message: "删除页面失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.DeletePPageA{Message: "删除页面成功", Success: true})
}

// DeleteUml
// @Summary 删除 Uml
// @Tags 回收站
// @Accept json
// @Produce json
// @Param data body response.DeleteUmlQ true "Uml ID"
// @Success 200 {object} response.DeleteUmlA
// @Router /bin/delete_uml [post]
func DeleteUml(c *gin.Context) {
	var data response.DeleteUmlQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.DeleteUmlA{Message: "输入数据不符合要求", Success: false})
		return
	}
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

// DeleteDocument
// @Summary 删除文档
// @Tags 回收站
// @Accept json
// @Produce json
// @Param data body response.DeleteDocumentQ true "文档ID"
// @Success 200 {object} response.DeleteDocumentA
// @Router /bin/delete_document [post]
func DeleteDocument(c *gin.Context) {
	var data response.DeleteDocumentQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.DeleteDocumentA{Message: "输入数据不符合要求", Success: false})
		return
	}
	document, notFound := service.QueryDocumentByDocumentID(data.DocumentID)
	if notFound {
		c.JSON(http.StatusOK, response.DeleteDocumentA{Message: "文档不存在", Success: false})
		return
	}
	/*filename := strings.Split(document.DocumentURL, "/")[len(strings.Split(document.DocumentURL, "/"))-1]
	saveDir := "./media/documents/"
	savePath := path.Join(saveDir, filename)
	if err := os.Remove(savePath); err != nil {
		c.JSON(http.StatusOK, response.DeleteDocumentA{Message: "删除文档失败", Success: false})
		return
	}*/
	err := service.DeleteDocument(&document)
	if err != nil {
		c.JSON(http.StatusOK, response.DeleteDocumentA{Message: "删除文档失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.DeleteDocumentA{Message: "删除文档成功", Success: true})
}

// GetProjInBin
// @Summary 回收站中的所有项目
// @Tags 回收站
// @Accept json
// @Produce json
// @Param data body response.GetProjInBinQ true "团队ID"
// @Success 200 {object} response.GetProjInBinA
// @Router /bin/get_projs_in_bin [post]
func GetProjInBin(c *gin.Context) {
	poster, _ := c.Get("user")
	var data response.GetProjInBinQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.GetProjInBinA{Message: "输入数据不符合要求", Success: false})
		return
	}
	projs := service.GetUserProjsInGroup(poster.(database.User).UserID, data.GroupID, 2, 3, 1, true)
	x := len(projs)
	if x == 0 {
		c.JSON(http.StatusOK, response.GetProjJoinA{
			Message: "回收站没有项目捏",
			Success: false,
			Count:   uint64(x),
			Projs:   projs})
		return
	}
	c.JSON(http.StatusOK, response.GetProjJoinA{
		Message: "成功获取回收站中的项目",
		Success: true,
		Count:   uint64(x),
		Projs:   projs})
}

// GetFilesInBin
// @Summary 回收站中的设计原型页面 / Uml / 文档，这些设计原型 / Uml / 文档所在的项目并没有被移到回收站
// @Tags 回收站
// @Accept json
// @Produce json
// @Param data body response.GetFilesInBinQ true "团队ID"
// @Success 200 {object} response.GetFilesInBinA
// @Router /bin/get_files_in_bin [post]
func GetFilesInBin(c *gin.Context) {
	var data response.GetFilesInBinQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.GetFilesInBinA{Message: "输入数据不符合要求", Success: false})
		return
	}
	PPages := service.GetProjPPagesInBin(data.GroupID)
	umls := service.GetProjUmlsInBin(data.GroupID)
	documents := service.GetProjDocumentsInBin(data.GroupID)
	c.JSON(http.StatusOK, response.GetFilesInBinA{
		Message:        "回收站中有以下设计原型 / Uml / 文档",
		Success:        true,
		CountPPages:    uint64(len(PPages)),
		PPages:         PPages,
		CountUmls:      uint64(len(umls)),
		Umls:           umls,
		CountDocuments: uint64(len(documents)),
		Documents:      documents})
}
