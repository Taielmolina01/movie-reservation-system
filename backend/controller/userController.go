package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	ownErrors "movie-reservation-system/errors"
	"movie-reservation-system/models"
	"movie-reservation-system/service/user"
	"net/http"
)

type UserController struct {
	UserService user.UserService
}

func NewUserController(userService user.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var request models.UserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	user, err := uc.UserService.CreateUser(&request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": user,
	})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	email := ctx.Param("email")

	user, err := uc.UserService.GetUser(email)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	email := ctx.Param("email")
	var request models.UserUpdateRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	user, err := uc.UserService.UpdateUser(email, &request)

	if err != nil {
		if errors.Is(err, ownErrors.ErrorUserNotExist{}) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func (uc *UserController) UpdateUserPassword(ctx *gin.Context) {
	var request models.UserUpdatePasswordRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	user, err := uc.UserService.UpdateUserPassword(&request)

	if err != nil {
		if errors.Is(err, ownErrors.ErrorUserNotExist{}) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	email := ctx.Param("email")

	user, err := uc.UserService.DeleteUser(email)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
