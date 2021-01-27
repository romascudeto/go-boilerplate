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
		grpRestExternal.GET("/example1", controllers.GetExternal1)
		grpRestExternal.GET("/example2", controllers.GetExternal2)
	}
	grpPractice := r.Group("/practice")
	{
		grpPractice.GET("/channeling", controllers.Channeling)
		grpPractice.GET("/channeling2", controllers.Channeling2)
	}
	return r
}
