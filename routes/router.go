package routes

import (
	"go-project/controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter(r *gin.Engine) *gin.Engine {
	grpEmployee := r.Group("/employee")
	{
		grpEmployee.GET("", controllers.GetEmployee)
		grpEmployee.GET("/:id", controllers.GetEmployeeByID)
	}
	grpRestExternal := r.Group("/external")
	{
		grpRestExternal.GET("", controllers.GetExternal)
	}
	return r
}
