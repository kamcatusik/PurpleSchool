package user

import (
	"4-order-api/pkg/db"
)

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}
func (repo *UserRepository) CreateUser(user *User) (*User, error) {
	result := repo.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
func (repo *UserRepository) FindUserByNum(number string) (*User, error) {
	var user User
	result := repo.Database.DB.First(&user, "number= ? ", number)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (repo *UserRepository) UpdateSessionId(user *User) (*User, error) {
	//fmt.Printf("Имя пользователя: %v Пароль: %v\n", user.Number, user.SessionID)
	result := repo.Database.DB.Model(&User{}).Where("id = ?", user.ID).Update("session_id", user.SessionID)
	if result.Error != nil {
		return nil, result.Error
	}
	//fmt.Printf("Имя пользователя: %v Пароль: %v\n", user.Number, user.SessionID)
	return user, nil

}
func (repo *UserRepository) FindUserBySession(session string) (*User, error) {
	var user User
	result := repo.Database.DB.First(&user, "session_id = ? ", session)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (repo *UserRepository) UpdateCode(user *User, code string) (*User, error) {

	result := repo.Database.DB.Model(&User{}).Where("id = ?", user.ID).Update("code", code)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil

}
