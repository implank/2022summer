package v1

import (
	"2022summer/model/response"
	"2022summer/service"
	"2022summer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetFilesByName
// @Summary 搜索框，只搜不在回收站的
// @Tags 项目管理的第二页
// @Accept json
// @Produce json
// @Param data body response.GetFilesByNameQ true "文件名称（不一定是全名，子串搜索，为空时返回数据库中全部不在回收站的文件，包括设计原型、Uml、文档）"
// @Success 200 {object} response.GetFilesByNameA
// @Router /file/get_files_by_name [post]
func GetFilesByName(c *gin.Context) {
	var data response.GetFilesByNameQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.GetFilesByNameA{Message: "输入数据不符合要求", Success: false})
		return
	}
	ppages, umls, documents := service.GetFilesByNameBur(data.Name, 1)
	c.JSON(http.StatusOK, response.GetFilesByNameA{
		Message:        "成功搜索到以下项目",
		Success:        true,
		CountPPage:     uint64(len(ppages)),
		PPage:          ppages,
		CountUmls:      uint64(len(umls)),
		Umls:           umls,
		CountDocuments: uint64(len(documents)),
		Documents:      documents})
}
