package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	secretKey = []byte("secret")
)

type JwtClaim struct {
	UserId int `json:"userid"`
	jwt.StandardClaims
}

func createJWTtoken(userID int) (jwtToken string, err error) {
	expTime := time.Now().Add(1 * time.Hour)
	jc := JwtClaim{
		UserId: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jc)
	jwtToken, err = token.SignedString(secretKey)
	if err != nil {
		return
	}
	return
}

func validateJwtToken(token string) (int, error) {
	jc := JwtClaim{}
	var ok bool
	jwtToken, err := jwt.ParseWithClaims(token, &jc, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return 0, err
	}
	jwtValidToken := &JwtClaim{}

	if jwtValidToken, ok = jwtToken.Claims.(*JwtClaim); !ok {
		return 0, errors.New("parsing error")
	}

	if jwtValidToken.ExpiresAt < time.Now().Unix() {
		return 0, errors.New("token expired")
	}

	return jwtValidToken.UserId, nil
}
