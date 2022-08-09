package v1

import (
	"2022summer/model/database"
	"2022summer/model/response"
	"2022summer/service"
	"2022summer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetProjDocuments
// @Summary 获取某个项目的不在回收站的文档
// @Tags 共享文档
// @Accept json
// @Produce json
// @Param data body response.GetProjDocumentsQ true "项目ID"
// @Success 200 {object} response.GetProjDocumentsA
// @Router /doc/get_proj_documents [post]
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
		Message:   "成功搜索到以下文档",
		Success:   true,
		Count:     uint64(x),
		Documents: documents})
}

// EnterDocument
// @Summary 进入文档
// @Tags 共享文档
// @Accept json
// @Produce json
// @Param data body response.EnterDocumentQ true "文档ID"
// @Success 200 {object} response.EnterDocumentA
// @Router /doc/enter_document [post]
func EnterDocument(c *gin.Context) {
	var data response.EnterDocumentQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.PARAMERROR)
		return
	}
	// todo there should be an auth check here
	doc, notFound := service.QueryDocumentByDocumentID(data.DocumentID)
	if notFound {
		c.JSON(http.StatusOK, response.EnterDocumentA{
			CommonA: response.CommonA{
				Message: "文档不存在",
				Success: false,
			},
		})
		return
	}
	doc.Count += 1
	service.UpdateDocument(&doc)
	c.JSON(http.StatusOK, response.EnterDocumentA{
		CommonA: response.CommonA{
			Message: "成功",
			Success: true,
		},
		Document: doc,
		Rank:     doc.Count,
	})
}

// QuitDocument
// @Summary 退出文档
// @Tags 共享文档
// @Accept json
// @Produce json
// @Param data body response.QuitDocumentQ true "文档ID"
// @Success 200 {object} response.QuitDocumentA
// @Router /doc/quit_document [post]
func QuitDocument(c *gin.Context) {
	var data response.QuitDocumentQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.PARAMERROR)
		return
	}
	// todo there should be an auth check here
	doc, notFound := service.QueryDocumentByDocumentID(data.DocumentID)
	if notFound {
		c.JSON(http.StatusOK, response.QuitDocumentA{
			CommonA: response.CommonA{
				Message: "文档不存在",
				Success: false,
			},
		})
		return
	}
	doc.Count -= 1
	service.UpdateDocument(&doc)
	c.JSON(http.StatusOK, response.QuitDocumentA{
		CommonA: response.CommonA{
			Message: "成功",
			Success: true,
		},
		Document: doc,
		Rank:     doc.Count,
	})
}

// CreateDocument
// @Summary 创建文档
// @Tags 共享文档
// @Accept json
// @Produce json
// @Param data body response.CreateDocumentQ true "文档名称，文档所属项目ID"
// @Success 200 {object} response.CreateDocumentA
// @Router /doc/create_document [post]
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
	proj, _ := service.QueryProjByProjID(data.ProjID)
	err := service.CreateDocument(&database.Document{
		DocumentName: data.DocumentName,
		Status:       1,
		ProjID:       data.ProjID,
		DirID:        proj.DocumentID,
	})
	if err != nil {
		c.JSON(http.StatusOK, response.CreateDocumentA{Message: "创建文档失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.CreateDocumentA{Message: "创建文档成功", Success: true})
}

// UploadDocument
// @Summary 上传文档
// @Tags 共享文档
// @Accept json
// @Produce json
// @Param data body response.UploadDocumentQ true "文档ID，文档内容"
// @Success 200 {object} response.UploadDocumentA
// @Router /doc/upload_document [post]
func UploadDocument(c *gin.Context) {
	var data response.UploadDocumentQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.PARAMERROR)
		return
	}
	document, notFound := service.QueryDocumentByDocumentID(data.DocumentID)
	if notFound {
		c.JSON(http.StatusOK, response.UploadDocumentA{
			CommonA: response.CommonA{
				Message: "文档不存在",
				Success: false,
			},
		})
		return
	}
	document.Content = data.Content
	service.UpdateDocument(&document)
	//filename := strings.Split(document.DocumentURL, "/")[len(strings.Split(document.DocumentURL, "/"))-1]
	//saveDir := "./media/documents/"
	//savePath := path.Join(saveDir, filename)
	//file, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE, 0666)
	//if err != nil {
	//	c.JSON(http.StatusOK, response.UploadDocumentA{
	//		CommonA: response.CommonA{
	//			Message: "文档不存在",
	//			Success: false,
	//		},
	//	})
	//	return
	//}
	//file.Write([]byte(data.Context))
	//file.Close()
	c.JSON(http.StatusOK, response.UploadDocumentA{
		CommonA: response.CommonA{
			Message: "上传文件成功",
			Success: true,
		},
		Document: document,
	})
}

// UpdateDocument
// @Summary 修改文档名称
// @Tags 共享文档
// @Accept json
// @Produce json
// @Param data body response.UpdateDocumentQ true "文档ID，文档的新名字（必填，可以填原名，不能和其他项目同名）"
// @Success 200 {object} response.UpdateDocumentA
// @Router /doc/update_document [post]
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

// MoveDocumentToBin
// @Summary 文档移入回收站
// @Tags 共享文档
// @Accept json
// @Produce json
// @Param data body response.MoveDocumentToBinQ true "文档 ID"
// @Success 200 {object} response.MoveDocumentToBinA
// @Router /doc/move_document_to_bin [post]
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

// CreateDocFile
// @Summary 创建文档文件
// @Tags 共享文档
// @Accept json
// @Produce json
// @Param data body response.CreateDocFileQ true "父目录ID，文档名称，创建的是目录还是文件夹"
// @Success 200 {object} response.CreateDocFileA
// @Router /doc/create_doc_file [post]
func CreateDocFile(c *gin.Context) {
	var data response.CreateDocFileQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.PARAMERROR)
		return
	}
	Dir, notFound := service.QueryDocumentByDocumentID(data.DirID)
	if notFound {
		c.JSON(http.StatusOK, response.CreateDocFileA{
			Message: "文件夹不存在",
			Success: false,
		})
		return
	}
	if Dir.IsDir == 0 {
		c.JSON(http.StatusOK, response.CreateDocFileA{
			Message: "不是文件夹",
			Success: false,
		})
		return
	}
	if Dir.IsFixed == 1 || (Dir.ProjID != 0 && data.IsDir == 1) {
		c.JSON(http.StatusOK, response.CreateDocFileA{
			Message: "文件夹内不能添加文件夹",
			Success: false,
		})
		return
	}
	if Dir.Status == 2 {
		c.JSON(http.StatusOK, response.CreateDocFileA{
			Message: "文件夹已经被移入回收站",
			Success: false,
		})
		return
	}
	files := service.GetDocumentsInDir(Dir.DocumentID)
	for _, file := range files {
		if file.DocumentName == data.Filename {
			c.JSON(http.StatusOK, response.CreateDocFileA{
				Message: "文件已存在",
				Success: false,
			})
			return
		}
	}
	doc := database.Document{
		DocumentName: data.Filename,
		Status:       1,
		ProjID:       Dir.ProjID,
		DirID:        Dir.DirID,
		IsDir:        data.IsDir,
	}
	err := service.CreateDocument(&doc)
	if err != nil {
		c.JSON(http.StatusOK, response.DBERROR)
		return
	}
	c.JSON(http.StatusOK, response.CreateDocFileA{
		Message: "创建文件成功",
		Success: true,
	})
}

// GetDocFiles
// @Summary 获取团队文件
// @Tags 共享文档
// @Accept json
// @Produce json
// @Param data body response.GetDocFilesQ true "团队ID"
// @Success 200 {object} response.GetDocFilesA
// @Router /doc/get_doc_files [post]
func GetDocFiles(c *gin.Context) {
	var data response.GetDocFilesQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, response.PARAMERROR)
		return
	}
	group, notFound := service.QueryGroupByGroupID(data.GroupID)
	if notFound {
		c.JSON(http.StatusOK, response.GetDocFilesA{
			Message: "团队不存在",
			Success: false,
		})
		return
	}
	files := Tree(group.DocumentID)
	c.JSON(http.StatusOK, response.GetDocFilesA{
		Message: "获取文件列表成功",
		Success: true,
		Files:   files,
	})
}

func Tree(DocumentID uint64) (files []database.File) {
	documents := service.GetDocumentsInDir(DocumentID)
	for _, document := range documents {
		if document.IsDir == 1 {
			files = append(files, database.File{
				FileID:         document.DocumentID,
				Filename:       document.DocumentName,
				IsDir:          1,
				ContainedFiles: Tree(document.DocumentID),
			})
		} else {
			files = append(files, database.File{
				FileID:   document.DocumentID,
				Filename: document.DocumentName,
				IsDir:    0,
			})
		}
	}
	return files
}
