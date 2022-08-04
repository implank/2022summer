package initialize

import (
	v1 "2022summer/api/v1"
	"2022summer/docs"
	"2022summer/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func SetupRouter(r *gin.Engine) {
	r.Use(middleware.Cors())         // 跨域
	r.Use(middleware.LoggerToFile()) // 日志

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/test", testGin)

	baseGroup := r.Group("/api/v1")
	{
		baseGroup.POST("/register", v1.Register)
		baseGroup.POST("/login", v1.Login)
	}

	userGroup := baseGroup.Group("/user", middleware.AuthRequired())
	{
		userGroup.POST("/info", v1.GetUserInfo)
		userGroup.POST("/modify_password", v1.ModifyPassword)
		userGroup.POST("/modify_info", v1.ModifyInfo)
	}

	groupGroup := userGroup.Group("/group", middleware.AuthRequired())
	{
		groupGroup.POST("/create_group", v1.CreateGroup)
		groupGroup.POST("/get_identity", v1.GetIdentity)
		groupGroup.POST("/get_group_members", v1.GetMembers)
		groupGroup.POST("/add_member", v1.AddMember)
		groupGroup.POST("/remove_member", v1.RemoveMember)
		groupGroup.POST("/set_member_status", v1.SetMemberStatus)
	}

	projGroup := r.Group("/api/v1/proj", middleware.AuthRequired())
	{
		projGroup.POST("/create_proj", v1.CreateProj)         // 创建项目
		projGroup.POST("/update_proj", v1.UpdateProj)         // 修改项目名称、项目描述
		projGroup.POST("/move_proj_to_bin", v1.MoveProjBin)   // 移入或移出回收站
		projGroup.POST("/get_proj_all", v1.GetProjAll)        // 全部项目
		projGroup.POST("/get_proj_create", v1.GetProjCreate)  // 我创建的
		projGroup.POST("/get_proj_join", v1.GetProjJoin)      // 我参与的
		projGroup.POST("/get_proj_by_name", v1.GetProjByName) // 搜索框
	}

	fileGroup := r.Group("/api/v1/file", middleware.AuthRequired())
	{
		fileGroup.POST("/get_proj_by_id", v1.GetProjByID)               // 获取项目信息
		fileGroup.POST("/get_proj_prototypes", v1.GetProjPrototypes)    // 获取项目的设计原型
		fileGroup.POST("/get_proj_umls", v1.GetProjUmls)                // 获取项目的 Uml 图
		fileGroup.POST("/get_proj_documents", v1.GetProjDocuments)      // 获取项目的文档
		fileGroup.POST("/create_prototype", v1.CreatePrototype)         // 创建设计原型
		fileGroup.POST("/create_uml", v1.CreateUml)                     // 创建 Uml
		fileGroup.POST("/create_document", v1.CreateDocument)           // 创建文档
		fileGroup.POST("/update_prototype", v1.UpdatePrototype)         // 修改设计原型名称
		fileGroup.POST("/update_uml", v1.UpdateUml)                     // 修改 Uml 名称
		fileGroup.POST("/update_document", v1.UpdateDocument)           // 修改文档名称
		fileGroup.POST("/move_prototype_to_bin", v1.MovePrototypeToBin) // 设计原型移入回收站
		fileGroup.POST("/move_uml_to_bin", v1.MoveUmlToBin)             // Uml 移入回收站
		fileGroup.POST("/move_document_to_bin", v1.MoveDocumentToBin)   // 文档移入回收站
		fileGroup.POST("/get_files_by_name", v1.GetFilesByName)         // 搜索框
	}

	pPageGroup := r.Group("/api/v1/ppage", middleware.AuthRequired())
	{
		pPageGroup.POST("/get_ppages", v1.GetPPages)         // 获取某个设计原型的所有页面
		pPageGroup.POST("/get_ppage_by_id", v1.GetPPageByID) // 获取设计原型的某个页面
		pPageGroup.POST("/create_ppage", v1.CreatePPage)     // 创建设计原型的一个页面
		pPageGroup.POST("/update_ppage", v1.UpdatePPage)     // 修改设计原型的某个页面的名称或数据
		pPageGroup.POST("/delete_ppage", v1.DeletePPage)     // 删除设计原型的某个页面
	}

	binGroup := r.Group("/api/v1/bin", middleware.AuthRequired())
	{
		binGroup.POST("/delete_proj", v1.DeleteProj)                       // 删除项目
		binGroup.POST("/move_prototype_from_bin", v1.MovePrototypeFromBin) // 设计原型移出回收站
		binGroup.POST("/move_uml_from_bin", v1.MoveUmlFromBin)             // Uml 移出回收站
		binGroup.POST("/move_document_from_bin", v1.MoveDocumentFromBin)   // 文档移出回收站
		binGroup.POST("/delete_prototype", v1.DeletePrototype)             // 删除设计原型
		binGroup.POST("/delete_uml", v1.DeleteUml)                         // 删除 Uml
		binGroup.POST("/delete_document", v1.DeleteDocument)               // 删除文档
	}
}

func testGin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"success": true,
	})
}
