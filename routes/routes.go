package routes

import "github.com/PreciousTosin/golang-article-app/controllers"

func initializeRoutes() {

	// Handle the index route
	router.GET("/", controllers.ShowIndexPage)

	// Group article related routes together
	articleRoutes := router.Group("/article")
	{
		// Handle GET requests at /article/view/some_article_id
		articleRoutes.GET("/view/:article_id", controllers.GetArticle)

		// Handle the GET requests at /article/create
		// Show the article creation page
		// Ensure that the user is logged in by using the middleware
		articleRoutes.GET("/create", controllers.ShowArticleCreationPage)

		// Handle POST requests at /article/create
		// Ensure that the user is logged in by using the middleware
		articleRoutes.POST("/create", controllers.CreateArticle)
	}
}
