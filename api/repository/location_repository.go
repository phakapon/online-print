package repository

import (
	"online-print/api/models"
	"time"

	"github.com/jinzhu/gorm"
)

type LocationsRepository interface {
	Save(*models.Location) (*models.Location, error)
	Find(uint64) (*models.Location, error)
	FindAll() ([]*models.Location, error)
	Update(*models.Location) error
	Delete(uint64) error
	Count() (int64, error)
	Paginate(*Metadata) (*Pagination, error)
	Search(string) ([]*models.Location, error)
}

type locationsRepositoryImpl struct {
	db *gorm.DB
}

func NewLocationsRepository(db *gorm.DB) *locationsRepositoryImpl {
	return &locationsRepositoryImpl{db}
}

func (r *locationsRepositoryImpl) Save(location *models.Location) (*models.Location, error) {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Location{}).Create(location).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return location, tx.Commit().Error
}

func (r *locationsRepositoryImpl) Find(location_id uint64) (*models.Location, error) {
	location := &models.Location{}
	err := r.db.Debug().Model(&models.Location{}).Where("id = ?", location_id).Find(location).Error
	return location, err
}

func (r *locationsRepositoryImpl) FindAll() ([]*models.Location, error) {
	locations := []*models.Location{}
	err := r.db.Debug().Model(&models.Location{}).Find(&locations).Error
	return locations, err
}

func (r *locationsRepositoryImpl) Update(location *models.Location) error {
	tx := r.db.Begin()

	columns := map[string]interface{}{
		"location_detail":        location.LocationDetail,
		// "status":       location.Status,
		// "sumcost":    location.Sumcost,
		//locationid
		//paymentid
		// "quantity":    location.Quantity,
		// "product_id": location.ProductID,//ทำเสร็จค่อยว่ากัน
		"updated_at":  time.Now(),
	}

	err := tx.Debug().Model(&models.Location{}).Where("id = ?", location.ID).UpdateColumns(columns).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *locationsRepositoryImpl) Delete(location_id uint64) error {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Location{}).Where("id = ?", location_id).Delete(&models.Location{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *locationsRepositoryImpl) Count() (int64, error) {
	var c int64
	err := r.db.Debug().Model(&models.Location{}).Count(&c).Error
	return c, err
}

func (r *locationsRepositoryImpl) Paginate(meta *Metadata) (*Pagination, error) {
	locations := []*models.Location{}

	err := r.db.Debug().
		Model(&models.Location{}).
		Offset(meta.Offset).
		Limit(meta.Limit).
		Find(&locations).Error

	return &Pagination{
		Elements: locations,
		Metadata: meta,
	}, err
}

func (r *locationsRepositoryImpl) Search(search string) ([]*models.Location, error) {
	locations := []*models.Location{}

	err := r.db.Debug().
		Model(&models.Location{}).
		Where("name like ?", "%"+search+"%").
		Find(&locations).Error

	return locations, err
}
