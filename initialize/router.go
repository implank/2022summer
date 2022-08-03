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

	userGroup := r.Group("/api/v1")
	{
		userGroup.POST("/register", v1.Register)
		userGroup.POST("/login", v1.Login)
	}

	userGroup.POST("/info", middleware.AuthRequired(), v1.GetUserInfo)

	groupGroup := userGroup.Group("/group", middleware.AuthRequired())
	{
		groupGroup.POST("/remove_member", v1.RemoveMember)
	}

	projGroup := r.Group("/api/v1", middleware.AuthRequired())
	{
		projGroup.POST("/create_proj", v1.CreateProj)         // 创建项目
		projGroup.POST("/update_proj", v1.UpdateProj)         // 修改项目名称、项目描述
		projGroup.POST("/move_proj_to_bin", v1.MoveProjBin)   // 移入或移出回收站
		projGroup.POST("/delete_proj", v1.DeleteProj)         // 删除项目
		projGroup.POST("/get_proj_all", v1.GetProjAll)        // 全部项目
		projGroup.POST("/get_proj_create", v1.GetProjCreate)  // 我创建的
		projGroup.POST("/get_proj_join", v1.GetProjJoin)      // 我参与的
		projGroup.POST("/get_proj_by_name", v1.GetProjByName) // 搜索框

		projGroup.POST("/get_proj_by_id", v1.GetProjByID)               // 获取项目信息
		projGroup.POST("/get_proj_prototypes", v1.GetProjPrototypes)    // 获取项目的设计原型
		projGroup.POST("/get_proj_umls", v1.GetProjUmls)                // 获取项目的 Uml 图
		projGroup.POST("/get_proj_documents", v1.GetProjDocuments)      // 获取项目的文档
		projGroup.POST("/create_prototype", v1.CreatePrototype)         // 创建设计原型
		projGroup.POST("/create_uml", v1.CreateUml)                     // 创建 Uml
		projGroup.POST("/create_document", v1.CreateDocument)           // 创建文档
		projGroup.POST("/update_prototype", v1.UpdatePrototype)         // 修改设计原型名称
		projGroup.POST("/update_uml", v1.UpdateUml)                     // 修改 Uml 名称
		projGroup.POST("/update_document", v1.UpdateDocument)           // 修改文档名称
		projGroup.POST("/move_prototype_to_bin", v1.MovePrototypeToBin) // 设计原型移入回收站
		projGroup.POST("/move_uml_to_bin", v1.MoveUmlToBin)             // Uml 移入回收站
		projGroup.POST("/move_document_to_bin", v1.MoveDocumentToBin)   // 文档移入回收站
		projGroup.POST("/delete_prototype", v1.DeletePrototype)         // 删除设计原型
		projGroup.POST("/delete_uml", v1.DeleteUml)                     // 删除 Uml
		projGroup.POST("/delete_document", v1.DeleteDocument)           // 删除文档
		projGroup.POST("/get_sth_by_name", v1.GetFilesByName)           // 搜索框

		projGroup.POST("/move_prototype_from_bin", v1.MovePrototypeFromBin) // 设计原型移出回收站
		projGroup.POST("/move_uml_from_bin", v1.MoveUmlFromBin)             // Uml 移出回收站
		projGroup.POST("/move_document_from_bin", v1.MoveDocumentFromBin)   // 文档移出回收站
	}
}

func testGin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"success": true,
	})
}
