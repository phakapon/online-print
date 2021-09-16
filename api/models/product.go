package models

import "errors"

type Product struct {
	Model
	File string  `gorm:"size:256;not null;unique" json:"file"`
	Cost float64 `gorm:"type:decimal(10,2);not null;default:0.0" json:"cost"`
	ProductorderID    uint64 `gorm:"foreignkey:ProductorderID" json:"productorder_id"`
}

var (
	ErrProductEmptyName = errors.New("Product.File can't be empty")
)

func (p *Product) Validate() error {
	if p.File == "" {
		return ErrProductEmptyName
	}
	return nil
}