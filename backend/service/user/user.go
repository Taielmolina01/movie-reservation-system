package service

import (
	"movie-reservation-system/repository"
	"movie-reservation-system/models"
	"strings"
	"errors"
)



type UserServiceImpl struct {
	userRepository repository.User
}

func NewUserServiceImpl(userRepository repository.User) UserService {
	return &UserServiceImpl{userRepository: userRepository}
}

func (us *UserServiceImpl) CreateUser(req *models.UserRequest) (*models.UserDB, error) {
	// Validate fields of request
	TrimStructFields(req)
	if req.Email == ''{
		return nil, errors.ErrorUserMustHaveEmail{}.Error()
	} else if req.Name == '' {
		return nil, errors.ErrorUserMustHaveName{}.Error()
	} else if len(req.Password) < 8 {
		return nil, errors.ErrorPasswordMustHaveLenght8{}.Error()
	} else if !Contains(models.GetRoles(), req.Role) {
		return nil, errors.ErrorUserRoleInvalid{req.Role}.Error() 
	}

	// Call to the db to validate that the user doesnt already exist

	// Save user in the db
}

func (us *UserServiceImpl) GetUser(email string) (*models.UserDB, error) {
	// Get user from the db
}

func (us *UserServiceImpl) UpdateUser(req *models.UserUpdateRequest) (*models.UserDB, error) {
	// Get user from the db
	user, err := us.GetUser(email) // Should read the email from the endpoint path

	if err != nil {
		return nil, errors.ErrorUserNotExist{email}.Error()
	}

	// Updating fields
	if req.Email != nil {
		user.Email = req.Email
	} else if req.Name != nil {
		user.Name = req.Name
	} else if req.Role != nil {
		user.Role = req.Role
	}

	// Save updated user in the db
}	

func (us *UserServiceImpl) UpdatePassword(req *models.UserUpdatePasswordRequest) (*models.UserDB, error) {
	// Get user from the db
	user, err := us.GetUser(email) // Should read the email from the endpoint path

	if err != nil {
		return nil, errors.ErrorUserNotExist{email}.Error()
	}

	// Validate password fields
	if !ValidatePassword(user.Password, req.OldPassword) {
		return nil, errors.ErrorWrongOldPassword{}.Error()
	}

	if len(req.NewPassword) < 8 {
		return nil, errors.ErrorPasswordMustHaveLenght8{}.Error()
	}

	// Update password
	user.Password = req.NewPassword

	// Save updated user in the db
}

func (us *UserServiceImpl) DeleteUser(email string) (*models.UserDB, error) {
	// Get user from the db
	user, err := us.GetUser(email) // Should read the email from the endpoint path

	if err != nil {
		return nil, errors.ErrorUserNotExist{email}.Error()
	}

	// Delete user from the db
}
	


