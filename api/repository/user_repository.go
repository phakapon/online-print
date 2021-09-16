package repository

import (
	"online-print/api/models"
	"time"

	"github.com/jinzhu/gorm"
)

type UsersRepository interface {
	Save(*models.User) (*models.User, error)
	Find(uint64) (*models.User, error)
	FindAll() ([]*models.User, error)
	Update(*models.User) error
	Delete(uint64) error
	Count() (int64, error)
	Paginate(*Metadata) (*Pagination, error)
	Search(string) ([]*models.User, error)
}

type usersRepositoryImpl struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *usersRepositoryImpl {
	return &usersRepositoryImpl{db}
}

func (r *usersRepositoryImpl) Save(user *models.User) (*models.User, error) {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.User{}).Create(user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return user, tx.Commit().Error
}

func (r *usersRepositoryImpl) Find(user_id uint64) (*models.User, error) {
	user := &models.User{}
	err := r.db.Debug().Model(&models.User{}).Where("id = ?", user_id).Find(user).Error
	return user, err
}

func (r *usersRepositoryImpl) FindAll() ([]*models.User, error) {
	users := []*models.User{}
	err := r.db.Debug().Model(&models.User{}).Find(&users).Error
	return users, err
}

func (r *usersRepositoryImpl) Update(user *models.User) error {
	tx := r.db.Begin()

	columns := map[string]interface{}{
		"fullname": user.FullName,
		"phone":    user.Phone,
		// "sumcost":    user.Sumcost,
		//locationid
		//paymentid
		// "quantity":    user.Quantity,
		// "product_id": user.ProductID,//ทำเสร็จค่อยว่ากัน
		"updated_at": time.Now(),
	}

	err := tx.Debug().Model(&models.User{}).Where("id = ?", user.ID).UpdateColumns(columns).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *usersRepositoryImpl) Delete(user_id uint64) error {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.User{}).Where("id = ?", user_id).Delete(&models.User{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *usersRepositoryImpl) Count() (int64, error) {
	var c int64
	err := r.db.Debug().Model(&models.User{}).Count(&c).Error
	return c, err
}

func (r *usersRepositoryImpl) Paginate(meta *Metadata) (*Pagination, error) {
	users := []*models.User{}

	err := r.db.Debug().
		Model(&models.User{}).
		Offset(meta.Offset).
		Limit(meta.Limit).
		Find(&users).Error

	return &Pagination{
		Elements: users,
		Metadata: meta,
	}, err
}

func (r *usersRepositoryImpl) Search(search string) ([]*models.User, error) {
	users := []*models.User{}

	err := r.db.Debug().
		Model(&models.User{}).
		Where("name like ?", "%"+search+"%").
		Find(&users).Error

	return users, err
}
