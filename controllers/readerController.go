package controllers

import (
	"LibManSys/api/initializers"
	"LibManSys/api/models"
	"LibManSys/api/utils"
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	"log"
	// "github.com/gin-gonic/gin/binding"
)

func SearchByTitle(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "reader" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// c.JSON(200, gin.H{"success": "home page", "role": claims.Role})

	Title := c.Param("title")

	log.Println(Title, "------------title-----------")

	var book models.BookInventory

	result := initializers.DB.Where("title =?", Title).First(&book)

	if result.Error != nil {
		log.Println(result.Error)
	}

	log.Println(book.Title, "book title")
	log.Println(book.AvailableCopies, "avail copies")

	var response struct {
		Status        string
		AvailableDate string
	}

	if book.AvailableCopies <= 0 {
		response.Status = "not available"
		var issueReg models.IssueRegistry

		initializers.DB.Where("isbn = ?", book.ISBN).First(&issueReg)
		response.AvailableDate = issueReg.ExpectedReturnDate.String()
	} else {
		response.Status = "Available"
		// response.AvailableDate = "Available now"
	}
	// log.Println(response)
	c.JSON(200, response)
}

func SearchByAuthor(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "reader" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// c.JSON(200, gin.H{"success": "home page", "role": claims.Role})

	Authors := c.Param("author")

	log.Println(Authors, "------------Authors-----------")


	var book []models.BookInventory

	result := initializers.DB.Where("authors =?", Authors).Find(&book)

	if result.Error != nil {
		log.Println("error db")
	}

	type response struct {
		Title         string
		Status        string
		AvailableDate string
	}

	var finalRes []response

	for i := 0; i < len(book); i++ {
		log.Println(book[i].Authors)
		log.Println(book[i].AvailableCopies)
		log.Println(book[i].Title)

		if book[i].AvailableCopies <= 0 {

			var issueReg models.IssueRegistry

			initializers.DB.Where("isbn = ?", book[i].ISBN).First(&issueReg)
			finalRes = append(finalRes, response{book[i].Title, "Not available", issueReg.ExpectedReturnDate.String()})
		} else {
			finalRes = append(finalRes, response{book[i].Title, "Available Now", "Available Now"})
		}
	}
	log.Println(finalRes)
	c.JSON(200, finalRes)
}

func SearchByPublisher(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "reader" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// c.JSON(200, gin.H{"success": "home page", "role": claims.Role})

	Publisher := c.Param("publisher")

	log.Println(Publisher, "------------Publisher-----------")

	// type EmailRequestBody struct {
	// 	Role  string `json:"role"`
	// 	Email string `json:"email"`
	// }

	// // log.Println("error is test")
	// var requestBody EmailRequestBody

	// if err := c.ShouldBindBodyWith(&requestBody, binding.JSON); err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// }

	// log.Println(requestBody.Email)

	// if requestBody.Role != "Reader" {
	// 	c.JSON(400, gin.H{"error": "unauthorized"})
	// 	return
	// }

	var book []models.BookInventory

	result := initializers.DB.Where("publisher =?", Publisher).Find(&book)

	if result.Error != nil {
		log.Println("error db")
	}

	type response struct {
		Title         string
		Status        string
		AvailableDate string
	}

	var finalRes []response

	for i := 0; i < len(book); i++ {
		log.Println(book[i].Publisher)
		log.Println(book[i].AvailableCopies)
		log.Println(book[i].Title)

		if book[i].AvailableCopies <= 0 {

			var issueReg models.IssueRegistry

			initializers.DB.Where("isbn = ?", book[i].ISBN).First(&issueReg)
			finalRes = append(finalRes, response{book[i].Title, "Not available", issueReg.ExpectedReturnDate.String()})
		} else {
			finalRes = append(finalRes, response{book[i].Title, "Available available", "Available now"})
		}
	}
	log.Println(finalRes)
	c.JSON(200, finalRes)

	// log.Println(book)
	// log.Println(book.Publisher)
	// log.Println(book.AvailableCopies)

	// if book.AvailableCopies <= 0 {
	// 	response.Status = "not available"
	// 	var issueReg models.IssueRegistry

	// 	initializers.DB.Where("isbn = ?",book.ISBN).First(&issueReg)
	// 	response.AvailableDate = issueReg.ExpectedReturnDate.String()
	// } else {
	// 	response.Status = "Available"
	// 	response.AvailableDate = "Available now"
	// }
	// log.Println(response)
	// c.JSON(200,response)
}

func RaiseIssue(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "reader" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// c.JSON(200, gin.H{"success": "home page", "role": claims.Role})

	var raise struct {
		BookID int    `json:"bookID"`
		Email  string `json:"email"`
	}

	if err := c.ShouldBindJSON(&raise); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	log.Println(raise.BookID, "this is issue")

	var issue models.RequestEvents

	issue.BookID = raise.BookID

	if err := initializers.DB.Save(&issue).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Issue raised"})
}

func Reader(c *gin.Context) {
	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "reader" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// c.JSON(200, gin.H{"success": "home page", "role": claims.Role})
	c.HTML(200, "reader.html", "reader")
}

func HTMLtitle(c *gin.Context) {
	c.HTML(200, "searchTitle.html", "reader")
}

func HTMLauthor(c *gin.Context) {
	c.HTML(200, "searchAuthor.html", "reader")
}

func HTMLpublisher(c *gin.Context) {
	c.HTML(200, "searchPublisher.html", "reader")
}

func HTMLRaiseIssue(c *gin.Context) {
	c.HTML(200, "raiseIssue.html", "reader")
}

func HTMLSuccess(c *gin.Context) {
	c.HTML(200, "success.html", "success")
}

func HTMLUnauthorized(c *gin.Context) {
	c.HTML(200, "error.html", "error")
}

func HTMLResolveIssue(c *gin.Context) {
	c.HTML(200, "resolveIssue.html", "reader")
}
