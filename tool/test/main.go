package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//make a router
	router := gin.Default()
	//router
	admin := router.Group("/admin")
	{
		//admin.POST("/*uri", post_admin)
		admin.GET("/*uri", get_admin)
		//admin.PUT("/*uri", put)
		//admin.DELETE("/*uri", delete_admin)
	}
	router.Run("0.0.0.0:8080")
}

func get_admin(c *gin.Context) {
	uri := c.Param("uri")
	q := c.Query("q")
	s := c.Query("s")
	c.JSON(500, gin.H{
		"success": true,
		"uri": uri,
		"q": q,
		"s": s,
	})
	return
}
