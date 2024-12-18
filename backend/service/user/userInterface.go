package service

import (
	"movie-reservation-system/models"
	"github.com/google/uuid"
)

type UserService interface {

	CreateUser(*models.UserRequest) (*models.UserDB)
	
	GetUser(email string) (*models.UserDB)
	
	UpdateUser(*models.UserUpdateRequest) (*models.UserDB)
	
	DeleteUser(email string) (*models.UserDB)
}