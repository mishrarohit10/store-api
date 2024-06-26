package controllers

import (
	"LibManSys/api/initializers"
	"LibManSys/api/models"
	"LibManSys/api/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
)

func LibCreate(c *gin.Context) {

	log.Println("inside libcreate fn")

	// cookie, err := c.Cookie("token")

	// if err != nil {
	// 	c.JSON(401, gin.H{"error": "unauthorized"})
	// 	return
	// }

	// claims, err := utils.ParseToken(cookie)

	// if err != nil {
	// 	c.JSON(401, gin.H{"error": "unauthorized"})
	// 	return
	// }

	// if claims.Role != "owner" {
	// 	c.JSON(401, gin.H{"error": "unauthorized"})
	// 	return
	// }

	// c.JSON(200, gin.H{"success": "home page", "role": claims.Role})

	var library struct {
		LibName  string `json:"name"`
		UserName string `json:"username"`
		Role     string `json:"role"`
	}

	log.Println(library)
	if err := c.ShouldBindBodyWith(&library, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Println(library)
	log.Println(library.LibName, "name og teh lib")
	var existingLib models.Library

	if err := initializers.DB.Where("name = ?", library.LibName).First(&existingLib).Error; err != nil {
		// c.JSON(400, gin.H{"error": err.Error()})
		log.Println("1st fn")
		// return
	}

	log.Println(existingLib.Name, "-------------------------------LIB NAME")
	if len(existingLib.Name) != 0 {
		log.Println("err2231")
		c.JSON(400, gin.H{"error": "library already exists try with new name"})
		return
	}

	// log.Println(library.ID, library.LibName, "libIDs")

	log.Println(library.LibName)

	lib := models.Library{Name: library.LibName}

	if err := initializers.DB.Create(&lib).Error; err != nil {
		log.Println("err2233")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Println(lib)

	var user models.User

	if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		log.Println("err2234")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var newUser models.User

	newUser.Name = user.Name
	newUser.Role = user.Role

	if len(newUser.Role) == 0 {
		newUser.Role = "owner"
	}

	result2 := initializers.DB.Create(&newUser)

	if result2.Error != nil {
		log.Println("err2235")
		c.JSON(400, gin.H{"errror": result2.Error})
		return
	}
	c.JSON(200, gin.H{"message": "done"})
}

func Owner(c *gin.Context) {

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

	if claims.Role != "owner" {
		c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
		return
	}

	// c.JSON(200, gin.H{"success": "home page", "role": claims.Role})

	c.HTML(200, "owner.html", "owner")
}

func HTMLCreateLib(c *gin.Context) {
	c.HTML(200, "createLib.html", "owner")
}
