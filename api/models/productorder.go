package models

import "errors"

type ProductStatus uint8

const (
	ProductStatus_Unavailable = 0
	ProductStatus_Available   = 1
)

// type User struct {
// 	gorm.Model
// 	CreditCards []CreditCard `gorm:"foreignKey:UserRefer"`
//   }
  
//   type CreditCard struct {
// 	gorm.Model
// 	Number    string
// 	UserRefer uint
//   }

type Productorder struct {
	Model
	Detail  string        `gorm:"size:512;not null;unique" json:"detail"`
	Sumcost float64       `gorm:"type:decimal(10,2);not null;default:0.0" json:"sumcost"`
	Quantity uint16	`gorm:"default:0;unsigned" json:"quantity"`
	Status  ProductStatus `gorm:"char(1);default:0" json:"status"`
	// ProductID uint64 `foreignkey:"ProductID" json:"product_id"`
	UserID uint64 `gorm:"UserID" json:"user_id"`
	
}

var (
	ErrProductorderEmptyName = errors.New("Productorder.Detail can't be empty")
)

func (p *Productorder) Validate() error {
	if p.Detail == "" {
		return ErrProductorderEmptyName
	}
	return nil
}

func (p *Productorder) CheckStatus() {
	p.Status = ProductStatus_Unavailable
	if p.Quantity > 0 {
		p.Status = ProductStatus_Available
	}
}