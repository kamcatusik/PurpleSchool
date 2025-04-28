package auth

import (
	"4-order-api/internal/user"
	"4-order-api/pkg/rand"
	"errors"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthRepository(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) Register(number string) (*user.User, error) {
	// идем в БД проверям на наличие номер

	//если нет то генерим sesionID и выдаем иначе генерим новый

	user := &user.User{
		Number:    number,
		SessionID: rand.RandSession(),
	}
	createdUser, err := service.UserRepository.CreateUser(user)
	if err != nil {
		return nil, errors.New("ошибка создания пользователя")
	}
	return createdUser, nil
}
func (service *AuthService) Update(user *user.User) (*user.User, error) {

	user.SessionID = rand.RandSession()

	_, err := service.UserRepository.UpdateSessionId(user)
	if err != nil {
		return nil, errors.New("ошибка обновления пользователя")
	}
	return user, nil
}
