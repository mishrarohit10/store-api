package main

import (
	"LibManSys/api/controllers"
	"LibManSys/api/initializers"
	"LibManSys/api/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("views/*")
	r.Use(static.Serve("/", static.LocalFile("./public", true)))
	r.Static("/public","./public")
	r.Use(cors.Default())

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
 
	r.GET("/login", controllers.Premium)
	r.GET("/logout", controllers.Logout)
	r.GET("/home", controllers.Home)
	r.GET("/signup", controllers.SignUpGet)
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

	r.Use(middlewares.IsAuthorized())
	r.POST("/createLib", controllers.LibCreate)
	r.POST("/addBooks", controllers.AddBooks)
	r.DELETE("/deleteBook/:id", controllers.RemoveBooks)
	r.PUT("/updateBook/:id", controllers.UpdateBook)
	r.POST("/raiseIssue", controllers.RaiseIssue)
	r.GET("/getIssues", controllers.ListIssue)
	r.PUT("/resolveIssue/:id", controllers.ResolveIssue)
	r.GET("/searchByTitle/:title", controllers.SearchByTitle)
	r.GET("/searchByAuthor/:author", controllers.SearchByAuthor)
	r.GET("/searchByPublisher/:publisher", controllers.SearchByPublisher)

	r.Run()
}      


