package repository

import (
	"online-print/api/models"
	"time"

	"github.com/jinzhu/gorm"
)

type LocationusersRepository interface {
	Save(*models.Locationuser) (*models.Locationuser, error)
	Find(uint64) (*models.Locationuser, error)
	FindAll() ([]*models.Locationuser, error)
	Update(*models.Locationuser) error
	Delete(uint64) error
	Count() (int64, error)
	Paginate(*Metadata) (*Pagination, error)
	Search(string) ([]*models.Locationuser, error)
}

type locationusersRepositoryImpl struct {
	db *gorm.DB
}

func NewLocationusersRepository(db *gorm.DB) *locationusersRepositoryImpl {
	return &locationusersRepositoryImpl{db}
}

func (r *locationusersRepositoryImpl) Save(locationuser *models.Locationuser) (*models.Locationuser, error) {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Locationuser{}).Create(locationuser).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return locationuser, tx.Commit().Error
}

func (r *locationusersRepositoryImpl) Find(locationuser_id uint64) (*models.Locationuser, error) {
	locationuser := &models.Locationuser{}
	err := r.db.Debug().Model(&models.Locationuser{}).Where("id = ?", locationuser_id).Find(locationuser).Error
	return locationuser, err
}

func (r *locationusersRepositoryImpl) FindAll() ([]*models.Locationuser, error) {
	locationusers := []*models.Locationuser{}
	err := r.db.Debug().Model(&models.Locationuser{}).Find(&locationusers).Error
	return locationusers, err
}

func (r *locationusersRepositoryImpl) Update(locationuser *models.Locationuser) error {
	tx := r.db.Begin()

	columns := map[string]interface{}{
		"location_id":        locationuser.LocationID,
		"user_id":       locationuser.UserID,
		// "sumcost":    locationuser.Sumcost,
		//locationid
		//paymentid
		// "quantity":    locationuser.Quantity,
		// "product_id": locationuser.ProductID,//ทำเสร็จค่อยว่ากัน
		"updated_at":  time.Now(),
	}

	err := tx.Debug().Model(&models.Locationuser{}).Where("id = ?", locationuser.ID).UpdateColumns(columns).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *locationusersRepositoryImpl) Delete(locationuser_id uint64) error {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Locationuser{}).Where("id = ?", locationuser_id).Delete(&models.Locationuser{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *locationusersRepositoryImpl) Count() (int64, error) {
	var c int64
	err := r.db.Debug().Model(&models.Locationuser{}).Count(&c).Error
	return c, err
}

func (r *locationusersRepositoryImpl) Paginate(meta *Metadata) (*Pagination, error) {
	locationusers := []*models.Locationuser{}

	err := r.db.Debug().
		Model(&models.Locationuser{}).
		Offset(meta.Offset).
		Limit(meta.Limit).
		Find(&locationusers).Error

	return &Pagination{
		Elements: locationusers,
		Metadata: meta,
	}, err
}

func (r *locationusersRepositoryImpl) Search(search string) ([]*models.Locationuser, error) {
	locationusers := []*models.Locationuser{}

	err := r.db.Debug().
		Model(&models.Locationuser{}).
		Where("name like ?", "%"+search+"%").
		Find(&locationusers).Error

	return locationusers, err
}
