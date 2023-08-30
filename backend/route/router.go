package route

import (
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-66/controllers"
	"github.com/rg-km/final-project-engineering-66/middleware"
	"github.com/rg-km/final-project-engineering-66/repository"
)

func NewRouter(userRepository repository.UserRepository, controller *controllers.UserController) *gin.Engine {
	router := gin.Default()
	middleware.SetupCorsMiddleware(router)

	api := router.Group("/api")
	{
		api.POST("/register", controller.Register)
		api.POST("/login", controller.Login)
		api.GET("/user", middleware.DeserializeUser(userRepository), controller.FindAll)
		api.GET("/user/:id", middleware.DeserializeUser(userRepository), controller.FindByID)
		api.PUT("/user/:id", middleware.DeserializeUser(userRepository), controller.UpdateUser)
		api.DELETE("/user/:id", middleware.DeserializeUser(userRepository), controller.DeleteUser)
	}
	return router
}
