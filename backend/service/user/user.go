package user

import (
	"golang.org/x/crypto/bcrypt"
	"movie-reservation-system/errors"
	"movie-reservation-system/models"
	"movie-reservation-system/repository/user"
	"movie-reservation-system/service"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func mapUserRequestToUserDB(req *models.UserRequest) *models.UserDB {
	return &models.UserDB{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
		Role:     req.Role,
	}
}

func (us *UserServiceImpl) CreateUser(req *models.UserRequest) (*models.UserDB, error) {
	// Validate fields of request
	if err := service.ValidateUserFields(req); err != nil {
		return nil, err
	}

	// Call to the db to validate that the user doesnt already exist
	_, userError := us.GetUser(req.Email)

	if userError != nil {
		return nil, errors.ErrorUserAlreadyExist{}
	}

	// Must hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	if err != nil {
		return nil, errors.ErrorEncriptyngPassword{}
	}

	req.Password = string(hashedPassword)

	newUser := mapUserRequestToUserDB(req)

	// Save user in the db
	return us.UserRepository.CreateUser(newUser)
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

	if err := service.ValidateAndUpdateUser(req, user); err != nil {
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
