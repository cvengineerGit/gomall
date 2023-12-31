package jwt

import (
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kalougata/gomall/pkg/config"
)

type JWT struct {
	key []byte
}

func New(conf *config.Config) *JWT {
	return &JWT{key: []byte(conf.JWT.Key)}
}

type MyCustomClaims struct {
	UserId    string
	LoginName string
	UserRule  string

	jwt.RegisteredClaims
}

func (j *JWT) BuildToken(claims MyCustomClaims, expiresAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		UserId:    claims.UserId,
		UserRule:  claims.UserRule,
		LoginName: claims.LoginName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "",
			Subject:   "",
			ID:        claims.UserId,
			Audience:  []string{},
		},
	})

	// Sign and get the complete encoded token as a string using the key
	tokenString, err := token.SignedString(j.key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWT) ParseToken(tokenString string) (*MyCustomClaims, error) {
	re := regexp.MustCompile(`(?i)Bearer `)
	tokenString = re.ReplaceAllString(tokenString, "")
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
