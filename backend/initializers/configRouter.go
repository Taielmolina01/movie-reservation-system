package initializers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	authController "movie-reservation-system/controller"
	userController "movie-reservation-system/controller"
	authRepository "movie-reservation-system/repository/auth"
	userRepository "movie-reservation-system/repository/user"
	authService "movie-reservation-system/service/auth"
	userService "movie-reservation-system/service/user"
)

func Init(db *gorm.DB, config *Configuration) *gin.Engine {

	userRepo := userRepository.CreateRepositoryImpl(db)

	if db == nil {
		panic("db is nil")
	}

	if config == nil {
		panic("config is nil")
	}

	userController := setUpUserLayers(db, userRepo)
	authController := setUpAuthLayers(db, userRepo, config)

	router := gin.Default()

	addCorsConfiguration(router)

	setUpUserRoutes(router, userController)
	setUpAuthRoutes(router, authController)

	return router
}

func setUpUserLayers(db *gorm.DB, userRepo userRepository.UserRepository) *userController.UserController {
	userService := userService.NewUserServiceImpl(userRepo)

	userController := userController.NewUserController(userService)

	return userController
}

func setUpAuthLayers(db *gorm.DB, userRepo userRepository.UserRepository, config *Configuration) *authController.AuthController {
	authRepo := authRepository.NewAuthRepositoryImpl(db)

	authService := authService.NewAuthService(authRepo, userRepo, config.JwtAlgorithm, config.JwtSecret)

	authController := authController.NewAuthController(authService)

	return authController
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
		usersGroup.POST("", userController.CreateUser)
		usersGroup.GET("/:email", userController.GetUser)
		usersGroup.PUT("/:email", userController.UpdateUser)
		usersGroup.PUT("/:email/password", userController.UpdateUserPassword)
		usersGroup.DELETE("/:email", userController.DeleteUser)
	}
}

func setUpAuthRoutes(router *gin.Engine, authController *authController.AuthController) {
	router.POST("/login", authController.Login)
	router.POST("/logout/:email", authController.Logout)
}
