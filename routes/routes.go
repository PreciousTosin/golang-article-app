package routes

import "github.com/PreciousTosin/golang-article-app/controllers"

func initializeRoutes() {

	// Handle the index route
	router.GET("/", controllers.ShowIndexPage)

	// Handle GET requests at /article/view/some_article_id
	router.GET("/article/view/:article_id", controllers.GetArticle)
}
