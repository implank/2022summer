package v1

import (
	"2022summer/global"
	"2022summer/model/response"
	"2022summer/service"
	"2022summer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"path"
	"strings"
	"time"
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
	ppages, umls, documents := service.GetFilesByNameBur(data.Name, 1, data.GroupID)
	c.JSON(http.StatusOK, response.GetFilesByNameA{
		Message:        "成功搜索到以下内容",
		Success:        true,
		CountPPage:     uint64(len(ppages)),
		PPage:          ppages,
		CountUmls:      uint64(len(umls)),
		Umls:           umls,
		CountDocuments: uint64(len(documents)),
		Documents:      documents})
}

// UploadImage
// @Summary 上传图片 前缀为/media/images/xxxx
// @Tags 项目管理的第二页
// @Param 			 avatar  formData  file true "avatar"
// @Produce json
// @Success 200 {object} response.UploadImageA
// @Router /upload_image [post]
func UploadImage(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusOK, response.PARAMERROR)
		return
	}
	raw := time.Now().String() + file.Filename
	md5 := utils.GetMd5(raw)
	suffix := strings.Split(file.Filename, ".")[1]
	saveDir := global.VP.Get("image_dir").(string)
	saveName := md5 + "." + suffix
	savePath := path.Join(saveDir, saveName)
	_ = c.SaveUploadedFile(file, savePath)
	url := saveName
	c.JSON(http.StatusOK, response.UploadImageA{
		Message: "上传成功",
		Success: true,
		Url:     url,
	})
}

// ConvertHtmlToPdf
// @Summary 上传html文件，转换为pdf文件 url前缀为 api/v1/temp/xxxx.pdf
// @Tags 项目管理的第二页
// @Param 			 htmlfile  formData  file true "htmlfile"
// @Produce json
// @Success 200 {object} response.ConvertHtmlToPdfA
// @Router /convert_html_to_pdf [post]
func ConvertHtmlToPdf(c *gin.Context) {
	file, _ := c.FormFile("htmlfile")
	saveDir := global.VP.Get("temp_dir").(string)
	savePath := path.Join(saveDir, file.Filename)
	_ = c.SaveUploadedFile(file, savePath)
	raw := time.Now().String() + file.Filename
	md5 := utils.GetMd5(raw)
	outFileName := md5 + ".pdf"
	err := exec.Command("pandoc", savePath, "-f", "html", "-t", "pdf", "-s", "-o", path.Join(saveDir, outFileName)).Run()
	if err != nil {
		panic(err)
	}
	_ = exec.Command("rm", savePath).Run()
	c.JSON(http.StatusOK, response.ConvertHtmlToPdfA{
		Message: "上传成功",
		Success: true,
		Url:     outFileName,
	})
}

// ConvertHtmlToDocx
// @Summary 上传html文件，转换为docx文件 url前缀为 api/v1/temp/xxxx.docx
// @Tags 项目管理的第二页
// @Param 			 htmlfile  formData  file true "htmlfile"
// @Produce json
// @Success 200 {object} response.ConvertHtmlToPdfA
// @Router /convert_html_to_docx [post]
func ConvertHtmlToDocx(c *gin.Context) {
	file, _ := c.FormFile("htmlfile")
	saveDir := global.VP.Get("temp_dir").(string)
	savePath := path.Join(saveDir, file.Filename)
	_ = c.SaveUploadedFile(file, savePath)
	raw := time.Now().String() + file.Filename
	md5 := utils.GetMd5(raw)
	outFileName := md5 + ".docx"
	err := exec.Command("pandoc", savePath, "-f", "html", "-t", "docx", "-s", "-o", path.Join(saveDir, outFileName)).Run()
	if err != nil {
		panic(err)
	}
	_ = exec.Command("rm", savePath).Run()
	c.JSON(http.StatusOK, response.ConvertHtmlToPdfA{
		Message: "上传成功",
		Success: true,
		Url:     outFileName,
	})
}
