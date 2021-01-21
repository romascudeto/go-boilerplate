package routes

import (
	"go-project/controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter(r *gin.Engine) *gin.Engine {
	grp1 := r.Group("/employee")
	{
		grp1.GET("", controllers.GetEmployee)
		grp1.GET("/:id", controllers.GetEmployeeByID)
	}
	return r
}
