package middlewares

import (
	"LibManSys/api/utils"
	"log"
	"github.com/gin-gonic/gin"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")

		if err != nil {
			log.Println("1")
			c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(cookie)

		log.Println(claims)

		if err != nil {
			log.Println(err.Error())
			log.Println("1")
			c.HTML(401, "login.html", gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
    
		c.Set("role", claims.Role)
		c.Next()
	}
}
pe9cdzh4bbzvt9p