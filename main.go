package main

import (
	"2022summer/initialize"
	"github.com/gin-gonic/gin"
)

func main() {
	initialize.InitViper()

	initialize.InitMySQL()
	defer initialize.Close()

	r := gin.Default()
	initialize.SetupRouter(r)
	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}
