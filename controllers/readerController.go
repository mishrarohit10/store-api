package controllers

import (
	"LibManSys/api/initializers"
	"LibManSys/api/models"
	"LibManSys/api/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func SearchByTitle(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "reader" {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
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
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "reader" {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
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

	log.Println(book, "books-----------------------------------------------------")

	type response struct {
		Title         string `json:"title"`
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
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "reader" {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	// c.JSON(200, gin.H{"success": "home page", "role": claims.Role})

	Publisher := c.Param("publisher")

	log.Println(Publisher, "------------Publisher-----------")

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

}

func RaiseIssue(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "reader" {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	// c.JSON(200, gin.H{"success": "home page", "role": claims.Role})

	var raise struct {
		BookID int `json:"bookID"`
		// Email  string `json:"email"`
	}

	if err := c.ShouldBindJSON(&raise); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var book models.BookInventory

	log.Println(raise.BookID, "this is book id")

	if err := initializers.DB.Where("isbn =?", raise.BookID).Find(&book).Error; err != nil {
		log.Println("invalid id")
		c.JSON(400, gin.H{"error": "invalid book id"})
		return
	}

	if book.ISBN == 0 {
		log.Println(book.ISBN, "this is not good invalid id")
		c.JSON(400, gin.H{"error": "invalid book id"})
		return
	}

	// if Error != nil {
	// 	log.Println("invalid id")
	// 	c.JSON(400, gin.H{"error": "invalid book id"})
	// 	return
	// }``

	log.Println(book, "this is book")

	log.Println(raise.BookID, "this is issue")

	var issue models.RequestEvents

	issue.BookID = raise.BookID

	if err := initializers.DB.Save(&issue).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Issue raised"})
}

func GetAllBooks(c *gin.Context) {

	// cookie, err := c.Cookie("token")

	// if err != nil {
	// 	c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
	// 	return
	// }

	// claims, err := utils.ParseToken(cookie)

	// if err != nil {
	// 	c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
	// 	return
	// }

	// if claims.Role != "reader"  {
	// 	c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
	// 	return
	// }

	// c.JSON(200, gin.H{"success": "home page", "role": claims.Role})

	var books []models.BookInventory

	result := initializers.DB.Find(&books)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(200, books)
}

func Reader(c *gin.Context) {
	cookie, err := c.Cookie("token")

	if err != nil {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "reader" {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	// c.JSON(200, gin.H{"success": "home page", "role": claims.Role})
	c.HTML(200, "reader.html", "reader")
}

func HTMLtitle(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "reader" {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	c.HTML(200, "searchTitle.html", "reader")
}

func HTMLauthor(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "reader" {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	c.HTML(200, "searchAuthor.html", "reader")
}

func HTMLpublisher(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "reader" {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	c.HTML(200, "searchPublisher.html", "reader")
}

func HTMLRaiseIssue(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "reader" {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	c.HTML(200, "raiseIssue.html", "reader")
}

func HTMLSuccess(c *gin.Context) {
	c.HTML(200, "success.html", "success")
}

func HTMLUnauthorized(c *gin.Context) {
	c.HTML(200, "error.html", "error")
}

func HTMLResolveIssue(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "reader" {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	c.HTML(200, "resolveIssue.html", "reader")
}
