package service

import (
	"movie-reservation-system/models"
)

type TokenService interface {

	Login(*models.UserLoginRequest)

	Logout()
}