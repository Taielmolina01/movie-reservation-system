package initializers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	userController "movie-reservation-system/controller"
	userRepository "movie-reservation-system/repository/user"
	userService "movie-reservation-system/service/user"
)

func Init(db *gorm.DB) *userController.UserController {

	userRepo := userRepository.CreateRepositoryImpl(db)

	userService := userService.NewUserServiceImpl(userRepo)

	userController := userController.NewUserController(userService)

	return userController
}

func SetUpRoutes(router *gin.Engine, userController *userController.UserController) {
	router.POST("/user", userController.CreateUser)
	router.GET("/user/:email", userController.GetUser)
	router.PUT("/user/:email", userController.UpdateUser)
	router.DELETE("/user/:email", userController.DeleteUser)
}
