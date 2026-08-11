package api

import "github.com/gin-gonic/gin"

//thanks swizzle!
func CookieToAuthHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("Auth")
		if err == nil && cookie != "" {
			c.Request.Header.Set("Authorization", "Bearer "+cookie)
		}

		c.Next()
	}
}
