package auth

import (
	"errors"
	"time"

	"github.com/Flake-chat/Flake-auth/model"
	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	config *Config
}

func New(config *Config) *Auth {
	return &Auth{
		config: config,
	}
}

type Claims struct {
	jwt.RegisteredClaims
	ID int
	Name string
	Exp int64
}

func (j *Auth) Createtoken(u *model.User) (string, error) {
	var key = []byte(j.config.Token)
	payload := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		ID:   u.ID,
		Name: u.Name,
		Exp:  time.Now().Add(time.Hour * 72).Unix(),
	})

	
	t, err := payload.SignedString(key)
	if err != nil {
		return "", err
	}
	return t, nil
}

func (j *Auth) Get_User(t string) (Claims, error)  {
	var Claims Claims
	var key = []byte(j.config.Token)
	token, err := jwt.ParseWithClaims(t, &Claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	} )
	if err != nil {
		return Claims, err
	}
	if !token.Valid {
		err := errors.New("token unvlaid")
		return Claims, err
	}
	return Claims, nil
}