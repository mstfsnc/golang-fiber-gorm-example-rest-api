package repositories

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sample-app/internal/models"
	"sample-app/pkg/metric"
	"time"
)

type UserRepository struct {
	db     *gorm.DB
	metric *metric.Metric
}

func NewUserRepository(db *gorm.DB, metric *metric.Metric) UserRepository {
	return UserRepository{
		db:     db,
		metric: metric,
	}
}

func (r UserRepository) All() ([]models.User, error) {
	defer func(begin time.Time) {
		r.metric.Observe("UserRepository_All", begin)
	}(time.Now())

	var users []models.User
	if err := r.db.Limit(10).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r UserRepository) Retrieve(id uint64) (models.User, error) {
	defer func(begin time.Time) {
		r.metric.Observe("UserRepository_Retrieve", begin)
	}(time.Now())

	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r UserRepository) Create(user *models.User) error {
	defer func(begin time.Time) {
		r.metric.Observe("UserRepository_Create", begin)
	}(time.Now())

	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r UserRepository) Update(user *models.User, fields models.User) error {
	defer func(begin time.Time) {
		r.metric.Observe("UserRepository_Update", begin)
	}(time.Now())

	if err := r.db.Model(&user).Clauses(clause.Returning{}).Updates(fields).Error; err != nil {
		return err
	}
	return nil
}

func (r UserRepository) Delete(id uint64) error {
	defer func(begin time.Time) {
		r.metric.Observe("UserRepository_Delete", begin)
	}(time.Now())

	exec := r.db.Where("id = ?", id).Delete(&models.User{})
	if exec.Error != nil {
		return exec.Error
	}
	if exec.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
