package initializers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	userController "movie-reservation-system/controller"
	userRepository "movie-reservation-system/repository/user"
	userService "movie-reservation-system/service/user"
)

func Init() *gin.Engine {

	LoadEnvVariables()

	db, err := ConnectDB()

	if err != nil {
		log.Fatal()
	}

	userRepo := userRepository.CreateRepositoryImpl(db)

	userService := userService.NewUserServiceImpl(userRepo)

	userController := userController.NewUserController(userService)

	router := gin.Default()

	setUpUserRoutes(router, userController)

	return router
}

func addCorsConfiguration(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	router.Use(cors.New(config))
}

func setUpUserRoutes(router *gin.Engine, userController *userController.UserController) {
	usersGroup := router.Group("/users")
	{
		usersGroup.POST("/", userController.CreateUser)
		usersGroup.GET("/:email", userController.GetUser)
		usersGroup.PUT("/:email", userController.UpdateUser)
		usersGroup.DELETE("/:email", userController.DeleteUser)
	}
}
