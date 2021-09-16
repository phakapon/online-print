package repository

import (
	"online-print/api/models"
	"time"

	"github.com/jinzhu/gorm"
)

type ProductsorderRepository interface {
	Save(*models.Productorder) (*models.Productorder, error)
	Find(uint64) (*models.Productorder, error)
	FindAll() ([]*models.Productorder, error)
	Update(*models.Productorder) error
	Delete(uint64) error
	Count() (int64, error)
	Paginate(*Metadata) (*Pagination, error)
	Search(string) ([]*models.Productorder, error)
}

type productsorderRepositoryImpl struct {
	db *gorm.DB
}

func NewProductsorderRepository(db *gorm.DB) *productsorderRepositoryImpl {
	return &productsorderRepositoryImpl{db}
}

func (r *productsorderRepositoryImpl) Save(productorder *models.Productorder) (*models.Productorder, error) {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Productorder{}).Create(productorder).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return productorder, tx.Commit().Error
}

func (r *productsorderRepositoryImpl) Find(productorder_id uint64) (*models.Productorder, error) {
	productorder := &models.Productorder{}
	err := r.db.Debug().Model(&models.Productorder{}).Where("id = ?", productorder_id).Find(productorder).Error
	return productorder, err
}

func (r *productsorderRepositoryImpl) FindAll() ([]*models.Productorder, error) {
	productsorder := []*models.Productorder{}
	err := r.db.Debug().Model(&models.Productorder{}).Find(&productsorder).Error
	return productsorder, err
}

func (r *productsorderRepositoryImpl) Update(productorder *models.Productorder) error {
	tx := r.db.Begin()

	columns := map[string]interface{}{
		"detail":        productorder.Detail,
		"status":       productorder.Status,
		"sumcost":    productorder.Sumcost,
		"quantity":    productorder.Quantity,
		// "product_id": productorder.ProductID,
		"user_id": productorder.UserID,
		"updated_at":  time.Now(),
	}

	err := tx.Debug().Model(&models.Productorder{}).Where("id = ?", productorder.ID).UpdateColumns(columns).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *productsorderRepositoryImpl) Delete(productorder_id uint64) error {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Productorder{}).Where("id = ?", productorder_id).Delete(&models.Productorder{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *productsorderRepositoryImpl) Count() (int64, error) {
	var c int64
	err := r.db.Debug().Model(&models.Productorder{}).Count(&c).Error
	return c, err
}

func (r *productsorderRepositoryImpl) Paginate(meta *Metadata) (*Pagination, error) {
	productsorder := []*models.Productorder{}

	err := r.db.Debug().
		Model(&models.Productorder{}).
		Offset(meta.Offset).
		Limit(meta.Limit).
		Find(&productsorder).Error

	return &Pagination{
		Elements: productsorder,
		Metadata: meta,
	}, err
}

func (r *productsorderRepositoryImpl) Search(search string) ([]*models.Productorder, error) {
	productsorder := []*models.Productorder{}

	err := r.db.Debug().
		Model(&models.Productorder{}).
		Where("name like ?", "%"+search+"%").
		Find(&productsorder).Error

	return productsorder, err
}
