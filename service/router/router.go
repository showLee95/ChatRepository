package router

import (
	"chatim/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	func StuRouter() *gin.Engine {
//		router := gin.Default()
//		apiGroup := router.Group("/api")
//		routes.SetApiGroupRoutes(apiGroup)
//		return router
//		}
func StartRoute() *gin.Engine {
	c := gin.Default()
	v1 := c.Group("/api/v1")
	v1.POST("/upfile", controller.UpFile())
	c.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return c
}
