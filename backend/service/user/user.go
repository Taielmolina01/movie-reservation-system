package user

import (
	"movie-reservation-system/repository/user"
	"movie-reservation-system/models"
	"movie-reservation-system/errors"
	"movie-reservation-system/service"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (us *UserServiceImpl) CreateUser(req *models.UserRequest) (*models.UserDB, error) {
	// Validate fields of request
	if err := service.ValidateUserFields(req); err != nil {
		return nil, err
	}

	// Call to the db to validate that the user doesnt already exist
	user, userError := us.GetUser(req.Email)

	if userError != nil {
		return nil, errors.ErrorUserAlreadyExist{}
	}

	// Must hash the password


	// Save user in the db
	return us.UserRepository.CreateUser(user)
}

func (us *UserServiceImpl) GetUser(email string) (*models.UserDB, error) {
	// Get user from the db
	return us.UserRepository.GetUser(email)
}

func (us *UserServiceImpl) UpdateUser(email string, req *models.UserUpdateRequest) (*models.UserDB, error) {
	// Get user from the db
	user, err := us.GetUser(email) 

	if err != nil {
		return nil, errors.ErrorUserNotExist{email}
	}

	if err := service.ValidateUserUpdateFields(req); err != nil {
		return nil, err
	}

	// Save updated user in the db
	return us.UserRepository.UpdateUser(user)
}	

func (us *UserServiceImpl) UpdateUserPassword(req *models.UserUpdatePasswordRequest) (*models.UserDB, error) {
	// Get user from the db
	user, err := us.GetUser(req.Email)

	if err != nil {
		return nil, errors.ErrorUserNotExist{req.Email}
	}

	// Validate password fields
	if !service.ValidatePassword(user.Password, req.OldPassword) {
		return nil, errors.ErrorWrongOldPassword{}
	}

	if len(req.NewPassword) < 8 {
		return nil, errors.ErrorPasswordMustHaveLenght8{}
	}

	// Update password
	user.Password = req.NewPassword // Must hash the password here

	// Save updated user in the db
	return us.UserRepository.UpdateUser(user)
}

func (us *UserServiceImpl) DeleteUser(email string) (*models.UserDB, error) {
	// Get user from the db
	user, err := us.GetUser(email) 

	if err != nil {
		return nil, errors.ErrorUserNotExist{email}
	}

	// Delete user from the db
	return us.UserRepository.DeleteUser(user)
}
	


