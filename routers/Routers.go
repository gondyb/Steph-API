package routers

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouters creates the gin routes
func SetupRouters() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/")
	{
		v1.GET("users", controllers.ListUsers)
	}

	return r
}
