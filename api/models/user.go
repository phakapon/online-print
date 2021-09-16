package models

import "errors"

type User struct {
	Model
	FullName  string        `gorm:"size:512;not null;unique" json:"full_name"`
	Phone  string        `gorm:"size:512;not null;unique" json:"phone"`
}

var (
	ErrUserEmptyName = errors.New("User.FullName can't be empty")
)

func (p *User) Validate() error {
	if p.FullName == "" {
		return ErrUserEmptyName
	}
	return nil
}