package v1

import (
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
	err := service.DeleteProj(&proj)
	if err != nil {
		c.JSON(http.StatusOK, response.DeleteProjA{Message: "删除项目失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.DeleteProjA{Message: "删除项目成功", Success: true})
}

func MovePrototypeFromBin(c *gin.Context) { // 设计原型移出回收站
	var data response.MovePrototypeFromBinQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.MovePrototypeToBinA{Message: "输入数据不符合要求", Success: false})
		return
	}
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

func MoveDocumentFromBin(c *gin.Context) { // 文档移出回收站
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

func DeletePrototype(c *gin.Context) { // 删除设计原型
	var data response.DeletePrototypeQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.DeletePrototypeA{Message: "输入数据不符合要求", Success: false})
		return
	}
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

func DeleteDocument(c *gin.Context) { // 删除文档
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
	err := service.DeleteDocument(&document)
	if err != nil {
		c.JSON(http.StatusOK, response.DeleteDocumentA{Message: "删除文档失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.DeleteDocumentA{Message: "删除文档成功", Success: true})
}
