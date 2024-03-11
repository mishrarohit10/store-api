// PATH: go-auth/controllers/auth.go

package controllers

import (
	"LibManSys/api/initializers"
	"LibManSys/api/models"
	"LibManSys/api/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var jwtKey = []byte("my_secret_key")

func Login(c *gin.Context) {

	var user models.User
	log.Println(user)
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("login ------__>")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User

	initializers.DB.Where("email = ?", user.Email).First(&existingUser)

	// log.Println(existingUser)

	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"error": "user does not exist"})
		return
	}

	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)

	if !errHash {
		c.JSON(400, gin.H{"error": "invalid password"})
		return
	}

	expirationTime := time.Now().Add(50 * time.Minute)

	claims := &models.Claims{
		Role: existingUser.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   existingUser.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}

	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
	c.JSON(200, gin.H{
		"success": "user logged in",
		"token":   tokenString,
		"role":    claims.Role,
		"user":    user.Email})
}

// PATH: go-auth/controllers/auth.go

func Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User

	initializers.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID != 0 {
		c.JSON(400, gin.H{"error": "user already exists"})
		return
	}

	var errHash error
	user.Password, errHash = utils.GenerateHashPassword(user.Password)

	if errHash != nil {
		c.JSON(500, gin.H{"error": "could not generate password hash"})
		return
	}

	if len(user.Role) == 0 {
		user.Role = "reader"
	}

	initializers.DB.Create(&user)

	c.JSON(200, gin.H{"role": user.Role})
}

// PATH: go-auth/controllers/auth.go

func Home(c *gin.Context) {

	// cookie, _ := c.Cookie("token")

	// // if err != nil {
	// // 	c.JSON(401, gin.H{"error": "unauthorized"})
	// // 	return
	// // }
	// if len(cookie) == 0 {
	// 	c.HTML(200, "index.html", gin.H{"role": ""})
	// 	return
	// }
	// claims, err := utils.ParseToken(cookie)

	// if err != nil {
	// 	c.JSON(401, gin.H{"error": "unauthorized"})
	// 	return
	// }

	// // if claims.Role != "admin" {
	// // 	c.JSON(401, gin.H{"error": "unauthorized"})
	// // 	return
	// // }

	// log.Println(cookie,"cookie")
	// log.Println(claims.Role,"role")

	// c.JSON(200, gin.H{"success": "home page", "role": claims.Role})
	log.Println("home")
	// c.HTML(200, "index.html", gin.H{"role": claims.Role})
	c.HTML(200, "index.html", "nice")
}

// PATH: go-auth/controllers/auth.go

func Premium(c *gin.Context) {

	log.Println("premium")
	email := c.PostForm("email")
	password := c.PostForm("password")

	log.Println(email, password)
	c.HTML(200, "login.html", gin.H{"message": "login Page"})
}

// PATH: go-auth/controllers/auth.go

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.HTML(200, "index.html",gin.H{"message": "user logged out"})
}

func SignUpGet(c *gin.Context) {
	c.HTML(200, "signup.html", gin.H{"page": "signup page"})
}
