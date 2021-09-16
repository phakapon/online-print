package models

import "errors"

type Location struct {
	Model
	LocationDetail string  `gorm:"size:256;not null;unique" json:"location_detail"`

}

var (
	ErrLocationEmptyName = errors.New("Location.LocationDetail can't be empty")
)

func (p *Location) Validate() error {
	if p.LocationDetail == "" {
		return ErrLocationEmptyName
	}
	return nil
}