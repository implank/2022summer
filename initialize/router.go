package initialize

import (
	"2022summer/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default()) // 跨域

	Group := r.Group("api/v1")
	{
		router.InitRouter(Group)
	}
	return r
}
