package main

import (
	"github.com/gin-gonic/gin"
	"workshop/routes"
)

func main() {
	r := gin.Default()

	routes.Router(r)

	err := r.Run()
	if err != nil {
		return
	}
}
