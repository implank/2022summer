package main

import (
	"2022summer/initialize"
	"github.com/gin-gonic/gin"
)

func main() {
	initialize.InitViper()
	err := initialize.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer initialize.Close()

	r := gin.Default()
	initialize.SetupRouter(r)
	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}
