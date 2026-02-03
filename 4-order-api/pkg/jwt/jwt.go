package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(phone string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone": phone,
	})
	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func (j *JWT) Validate(tokenString string) (jwt.MapClaims, error) {
    if tokenString == "" {
        return nil, errors.New(ErrEmptyToken)
    }
    token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
        if t.Method != jwt.SigningMethodHS256 {
            return nil, errors.New(ErrUnexpectedSigningMethod)
        }
        return []byte(j.Secret), nil
    })
    if err != nil {
        return nil, err
    }
    if !token.Valid {
        return nil, errors.New(ErrInvalidToken)
    }
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return nil, errors.New(ErrInvalidClaims)
    }
    return claims, nil
}