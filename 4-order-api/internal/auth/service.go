package auth

import (
	"crypto/rand"
	"errors"
	"math/big"
	"ps-go-adv/4-order-api/internal/user"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Login(phone string) (string, error) {
	currentUser, _ := service.UserRepository.FindByPhone(phone)
	if currentUser == nil {
		currentUser = &user.User{
			Phone: phone,
		}
		_, err := service.UserRepository.Create(currentUser)
		if err != nil {
			return "", err
		}
	}
	for {
		sessionID, err := service.genSessionID()
		if err != nil {
			return "", err
		}
		existingUser, _ := service.UserRepository.FindBySessionID(sessionID)
		if existingUser == nil {
			currentUser.SessionID = sessionID
			break
		}
	}
	currentUser.VerificationCode = "123456"
	_, err := service.UserRepository.Update(currentUser)
	if err != nil {
		return "", err
	}
	return currentUser.SessionID, nil
}

func (service *AuthService) Verify(sessionID, code string) (string, error) {
	existedUser, _ := service.UserRepository.FindBySessionID(sessionID)
	if existedUser == nil {
		return "", errors.New(SessionNotFound)
	}
	if existedUser.VerificationCode != code {
		return "", errors.New(ErrWrongCode)
	}
	return existedUser.Phone, nil
}

func (service *AuthService) genSessionID() (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    out := make([]byte, 16)
    for i := range out {
        num, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
        if err != nil {
            return "", err
        }
        out[i] = chars[num.Int64()]
    }
    return string(out), nil
}