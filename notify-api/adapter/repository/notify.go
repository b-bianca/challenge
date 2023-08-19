package repository

import (
	"context"
	"time"

	"github.com/b-bianca/melichallenge/notify-api/internal/domain/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Notify struct {
	db *gorm.DB
}

func NewNotifyRepository(db *gorm.DB) *Notify {
	return &Notify{db}
}

func (n *Notify) CreateNotify(ctx context.Context, nt *entity.Notification) (*entity.Notification, error) {
	dbFn := n.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true})

	if result := dbFn.Table("notify").Create(nt); result.Error != nil {
		return nil, result.Error
	}

	return nt, nil
}

func (n *Notify) FetchNotify(ctx context.Context) (*entity.NotificationList, error) {
	dbFn := n.db.WithContext(ctx)

	var nt []*entity.Notification
	twoMinutesAgo := time.Now().Add(-2 * time.Minute)

	if result := dbFn.Table("notify").Where("created_at >= ?", twoMinutesAgo).Find(&nt); result.Error != nil {
		return nil, result.Error
	}
	return &entity.NotificationList{
		Result: nt,
	}, nil

}
