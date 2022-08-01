package initialize

import (
	v1 "2022summer/api/v1"
	"2022summer/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func SetupRouter(r *gin.Engine) {
	r.Use(middleware.Cors())   // 跨域
	r.Use(middleware.Logger()) // 日志

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/test", testGin)

	userGroup := r.Group("/api/v1")
	{
		userGroup.POST("/register", v1.Register)
		userGroup.POST("/login", v1.Login)
		//userGroup.POST("/info", middleware.AuthRequired(), api.Info)
	}
}

func testGin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"success": true,
	})
}
