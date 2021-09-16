package models

import "errors"

type Locationuser struct {
	Model
	LocationID string  `gorm:"foreignkey:LocationID" json:"location_id"`
	UserID string  `gorm:"foreignkey:UserID" json:"user_id"`
	
}

var (
	ErrLocationuserEmptyName = errors.New("Locationuser.LocationID can't be empty")
)

func (p *Locationuser) Validate() error {
	if p.LocationID == "" {
		return ErrLocationuserEmptyName
	}
	return nil
}