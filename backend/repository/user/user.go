package repository

import (
	"movie-reservation-system/models"
	"movie-reservation-system/errors"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func CreateRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (ur *UserRepositoryImpl) CreateUser(user *models.UserDB) (*models.UserDB, error) {
	result := ur.db.Create(user)
	if userResponse.Error != nil {
		return nil, result.Error
	}
	return result, nil
}
	
func (ur *UserRepositoryImpl) GetUser(email string) (user *models.UserDB, error) {
	result := db.First(&models.UserDB, "email = ?", email)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.ErrorUserNotExist{Email: email}.Error()
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (ur *UserRepositoryImpl) UpdateUser(*models.UserDB) (user *models.UserDB, error) {
	result := ur.db.Save(user)

    if result.Error != nil {
        return nil, result.Error
    }

    return result, nil
}

func (ur *UserRepositoryImpl) DeleteUser(*models.UserDB) (user *models.UserDB, error) {
	result := ur.db.Delete(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}