package service

import (
	"movie-reservation-system/models"
	"github.com/google/uuid"
)

type UserService interface {

	CreateUser(*models.UserRequest) (*models.UserDB, error)
	
	GetUser(email string) (*models.UserDB, error)
	
	UpdateUser(*models.UserUpdateRequest) (*models.UserDB, error)
	
	DeleteUser(email string) (*models.UserDB, error)
}