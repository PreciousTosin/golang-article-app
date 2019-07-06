package main

import (
	"fmt"
	"github.com/PreciousTosin/golang-article-app/routes"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func SetUpRouter() *gin.Engine {
	// Set the router as the defaut one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Initialize routes
	routes.InitializeRoutes(router)
	fmt.Println("ROUTES ARE INITIALIZED AGAIN!!!")

	return router
}

func main() {
	// Start serving the application
	router := SetUpRouter()
	_ = router.Run()

}
