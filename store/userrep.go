package store

import (
	"github.com/Flake-chat/Flake-auth/model"
)

type Userrep struct {
	store *Store
}

func (r *Userrep) Create(u *model.User) (*model.User, error) {

	if err := u.Valudate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}
	if err := r.store.db.QueryRow("INSERT INTO users (name, fullname, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id",
		u.Name, u.Fullname, u.Email, u.Password_hash).Scan(&u.ID); err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *Userrep) Find(u *model.User) error {

	if err := r.store.db.QueryRow("SELECT id,email,password_hash,name FROM users WHERE email = $1", u.Email).Scan(
		&u.ID,
		&u.Email,
		&u.Password_hash,
		&u.Name,
	); err != nil {
		return err
	}
	check := u.CheckPass()
	if check != nil {
		return check
	}

	return nil
}
