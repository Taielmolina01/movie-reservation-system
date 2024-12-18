package controller

import (
	"movie-reservation-system/service"
)

type UserController struct {
	userService service.user.UserService
}