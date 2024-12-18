package repository

type UserRepository interface {

	CreateUser(*models.UserDB) (*models.UserDB, error)
	
	GetUser(email string) (*models.UserDB, error)
	
	UpdateUser(*models.UserDB) (*models.UserDB, error)
	
	DeleteUser(*models.UserDB) (*models.UserDB, error)
}