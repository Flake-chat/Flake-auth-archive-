package model

import (
	"errors"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            int    `json="id"`
	Name          string `json="name" validate:"required,min=2"`
	Fullname      string `json="fullname"validate:"required,min=2"`
	Email         string `json="email,omitempty" validate:"required,min=2"`
	Password      string `json="password,omitempty" validate:"required,min=2"`
	Password_hash string `json="-" validate:"required,max=0"`
}

func (u *User) Valudate() error {
	v := validator.New()
	err := v.Struct(u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) CheckPass() error {
	c := chechPass(u.Password_hash, u.Password)
	if c == false {
		e := errors.New("wrong password")
		return e
	}
	return nil
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := hashPass(u.Password)
		if err != nil {
			return err
		}
		u.Password_hash = enc
	}
	return nil
}

func hashPass(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func chechPass(h, p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	return err == nil
}
