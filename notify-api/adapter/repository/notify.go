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
	fiveMinutesAgo := time.Now().Add(-5 * time.Minute)

	if result := dbFn.Table("notify").Where("created_at >= ?", fiveMinutesAgo).Find(&nt); result.Error != nil {
		return nil, result.Error
	}
	return &entity.NotificationList{
		Result: nt,
	}, nil

}

func (n *Notify) SendMessage(ctx context.Context, m *entity.Message) (*entity.Message, error) {
	dbFn := n.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true})

	if result := dbFn.Table("message").Create(m); result.Error != nil {
		return nil, result.Error
	}

	return m, nil
}

func (n *Notify) FetchMessage(ctx context.Context) (*entity.MessageList, error) {
	dbFn := n.db.WithContext(ctx)

	var m *entity.MessageList

	if result := dbFn.Table("message").Find(&m); result.Error != nil {
		return nil, result.Error
	}

	return m, nil
}
