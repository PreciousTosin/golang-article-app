package main

import (
	"github.com/PreciousTosin/golang-article-app/routes"
)

// var router *gin.Engine

func main() {
	// Start serving the application
	router := routes.SetUpRouter()
	_ = router.Run()

}
