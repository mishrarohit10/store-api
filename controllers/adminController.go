package controllers

import (
	"LibManSys/api/initializers"
	"LibManSys/api/models"
	"LibManSys/api/utils"
	"github.com/gin-gonic/gin"
	"log"
	"time"
	// "github.com/gin-gonic/gin/binding"
)

func Admin(c *gin.Context) {

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

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// c.JSON(200, gin.H{"success": "home page", "role": claims.Role})
	c.HTML(200, "admin.html", "admin")
	// c.JSON(200, gin.H{"success": "home page", "role": claims.Role})
}

func HTMLAddBooks(c *gin.Context) {
	c.HTML(200, "addBooks.html", "HTML")
}

func HTMLRemoveBooks(c *gin.Context) {
	c.HTML(200, "removeBook.html", "HTML")
}

func HTMLUpdateBooks(c *gin.Context) {
	c.HTML(200, "updateBooks.html", "HTML")
}

func AddBooks(c *gin.Context) {

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

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// c.JSON(200, gin.H{"success": "home page", "role": claims.Role})

	var addBooks struct {
		ISBN            int    `json:"ISBN" gorm:"primary_key"`
		LibID           int    `json:"libID"`
		Title           string `json:"title"`
		Authors         string `json:"authors"`
		Publisher       string `json:"publisher"`
		Version         int    `json:"version"`
		TotalCopies     int    `json:"totalCopies"`
		AvailableCopies int    `json:"availableCopies"`
		LibAdminEmail   string
	}

	// var addBooks models.BookInventory

	if err := c.ShouldBindJSON(&addBooks); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	log.Println(addBooks.AvailableCopies, "----------available copies")

	inventory := models.BookInventory{ISBN: addBooks.ISBN, LibID: addBooks.LibID, Title: addBooks.Title, Authors: addBooks.Authors, Publisher: addBooks.Publisher, Version: addBooks.Version, TotalCopies: addBooks.TotalCopies, AvailableCopies: addBooks.AvailableCopies}

	ID := addBooks.ISBN

	log.Println(inventory.TotalCopies, "---------- inventory total copies")
	log.Println(inventory.AvailableCopies, "---------- inventory available copies")

	var exists models.BookInventory
	result := initializers.DB.First(&exists, ID)

	initializers.DB.Select("total_copies").Find(&exists)

	log.Println(exists.AvailableCopies, "---------- exists available copies")

	if result.Error != nil {
		// create inventory
		log.Println("inside result")
		initializers.DB.Save(&inventory)
		c.JSON(200, inventory)
	} else {
		// update TotalCopies

		var updatedInventory models.BookInventory

		if err := initializers.DB.First(&updatedInventory, ID).Error; err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		updatedInventory.TotalCopies = inventory.TotalCopies + exists.TotalCopies
		updatedInventory.AvailableCopies = inventory.AvailableCopies + exists.AvailableCopies

		if err := initializers.DB.Save(&updatedInventory).Error; err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, updatedInventory)
	}
}

func RemoveBooks(c *gin.Context) {

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

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// c.JSON(200, gin.H{"success": "home page", "role": claims.Role})

	//  remove book by id
	ID := c.Param("id")

	var inventory models.BookInventory

	result := initializers.DB.Find(&inventory, ID)

	if result.Error != nil {
		c.JSON(400, gin.H{"message": "book not found"})
		return
	}

	log.Println(inventory.TotalCopies)

	if inventory.TotalCopies > 0 {
		inventory.TotalCopies = inventory.TotalCopies - 1

		if err := initializers.DB.Save(&inventory).Error; err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	} else {
		c.JSON(200, gin.H{"message": "available copies is zero"})
		return
	}

	c.JSON(200, gin.H{"message": "done"})
}

func UpdateBook(c *gin.Context) {

	log.Println("inside updatebook")

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

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// c.JSON(200, gin.H{"success": "home page", "role": claims.Role})

	ID := c.Param("id")

	log.Println(ID, "this is ID")

	var book models.BookInventory

	result := initializers.DB.First(&book, ID)

	log.Println(result, "this is result")

	if result.Error != nil {
		c.JSON(400, gin.H{"message": "record not found"})
		// return
	}

	var updatedBook models.BookInventory
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(204, gin.H{"error": err.Error()})
		// return
	}

	log.Println(updatedBook, "this is updated book")

	book.ISBN = updatedBook.ISBN
	book.LibID = updatedBook.LibID
	book.Title = updatedBook.Title
	book.Authors = updatedBook.Authors
	book.Publisher = updatedBook.Publisher
	book.Version = updatedBook.Version
	book.TotalCopies = updatedBook.TotalCopies
	book.AvailableCopies = updatedBook.AvailableCopies

	log.Println(book, "this is  book")

	error := initializers.DB.Save(&book)

	if error.Error != nil {
		c.JSON(400, gin.H{"message": "record not found"})
		// return
	}

	// if err := initializers.DB.Where("id=?", ID).Update(&book); err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }

	log.Println("saved")
	log.Println(book, "this is saved book")
	c.JSON(201, gin.H{"message": "done"})
}

func ListIssue(c *gin.Context) {

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

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// c.JSON(200, gin.H{"success": "home page", "role": claims.Role})

	var lists []models.RequestEvents
	result := initializers.DB.Find(&lists)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Record not found"})
		return
	}
	// c.HTML(200, "admin.html", lists)
	c.JSON(200, lists)
}

func ResolveIssue(c *gin.Context) {

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

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// c.JSON(200, gin.H{"success": "home page", "role": claims.Role})

	ID := c.Param("id")

	log.Println(ID, "this is id")
	// var issue models.IssueRegistry

	var issue models.RequestEvents

	result := initializers.DB.First(&issue, ID)

	log.Println(issue, "issue reg")

	if result.Error != nil {
		c.JSON(400, gin.H{"message": "record not found"})
		return
	}

	// var resolvedEvent struct {
	// 	ReqID        int `json:"id" gorm:"primary_key"`
	// 	BookID       int `json:"bookId"`
	// 	ReaderID     int	`json:"readerId"`
	// 	RequestDate  time.Time
	// 	ApprovalDate time.Time
	// 	ApprovalID   int
	// 	RequestType  string
	// }

	var updateIssueRegistry models.IssueRegistry

	log.Println(issue.BookID, "book id")

	updateIssueRegistry.ISBN = issue.BookID
	updateIssueRegistry.ReaderID = issue.ReaderID
	updateIssueRegistry.IssueStatus = "accepted"
	updateIssueRegistry.IssueDate = time.Now()
	updateIssueRegistry.ExpectedReturnDate = time.Now().AddDate(0, 0, 7)
	updateIssueRegistry.ReturnDate = time.Now().AddDate(0, 0, 7)

	log.Println(updateIssueRegistry.ISBN, "updated issue reg ISBN")

	res := initializers.DB.Save(&updateIssueRegistry)

	print(res.Error, "this is error")

	c.JSON(204, updateIssueRegistry)
}
