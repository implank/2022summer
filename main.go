package main

import (
	"2022summer/initialize"
)

func main() {
	err := initialize.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer initialize.Close()

	r := initialize.SetupRouter()

	if err := r.Run("localhost:8000"); err != nil {
		panic(err)
	}
}
