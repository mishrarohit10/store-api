package main

import (
	"LibManSys/api/controllers"
	"LibManSys/api/initializers"
	"LibManSys/api/middlewares"

	// "time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// config.AllowOrigins = []string{"http://localhost:5173/"}
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"*"}
	config.AllowCredentials = true

	r.Use(cors.New(config))

	r.LoadHTMLGlob("views/*")
	r.Use(static.Serve("/", static.LocalFile("./public", true)))
	r.Static("/public", "./public")

	r.POST("/api/signup", controllers.Signup)
	r.POST("/api/login", controllers.Login)
	r.GET("/signup", controllers.SignUpGet)
	r.GET("/login", controllers.Premium)
	r.GET("/logout", controllers.Logout)
	r.GET("/home", controllers.Home)
	r.GET("/", controllers.HTMLUnauthorized)

	r.Use(middlewares.IsAuthorized())

	r.GET("/admin", controllers.Admin)
	r.GET("/reader", controllers.Reader)
	r.GET("/owner", controllers.Owner)
	r.GET("/addBooks", controllers.HTMLAddBooks)
	r.GET("/removeBooks", controllers.HTMLRemoveBooks)
	r.GET("/updateBooks", controllers.HTMLUpdateBooks)
	r.GET("/createLib", controllers.HTMLCreateLib)
	r.GET("/getTitle", controllers.HTMLtitle)
	r.GET("/getAuthor", controllers.HTMLauthor)
	r.GET("/getPublisher", controllers.HTMLpublisher)
	r.GET("/raiseIssue", controllers.HTMLRaiseIssue)
	r.GET("/success", controllers.HTMLSuccess)
	r.GET("/resolveIssue", controllers.HTMLResolveIssue)

	r.POST("/api/createLib", controllers.LibCreate)
	r.POST("/api/addBooks", controllers.AddBooks)
	r.DELETE("/api/deleteBook/:id", controllers.RemoveBooks)
	r.PUT("/api/updateBook/:id", controllers.UpdateBook)
	r.POST("/api/raiseIssue", controllers.RaiseIssue)
	r.GET("/api/getIssues", controllers.ListIssue)
	r.PUT("/api/resolveIssue/:id", controllers.ResolveIssue)
	r.GET("/api/searchByTitle/:title", controllers.SearchByTitle)
	r.GET("/api/searchByAuthor/:author", controllers.SearchByAuthor)
	r.GET("/api/searchByPublisher/:publisher", controllers.SearchByPublisher)
	r.GET("/api/getAllBooks", controllers.GetAllBooks)

	r.Run()
}
