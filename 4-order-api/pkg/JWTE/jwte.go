package jwte

import "github.com/golang-jwt/jwt/v5"

type JWTE struct {
	Secret string
}

func NewJWT(secret string) *JWTE {
	return &JWTE{
		Secret: secret,
	}
}
func (j *JWTE) Create(sessionId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sessionId": sessionId,
	})
	str, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return str, nil
}
