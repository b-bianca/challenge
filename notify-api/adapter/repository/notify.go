package repository

import (
	"context"
	"time"

	"github.com/b-bianca/melichallenge/notify-api/internal/domain/entity"
	"github.com/google/uuid"
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

	if result := dbFn.Table("notify").Where("created_at >= ? AND ack = ?", twoMinutesAgo, false).Find(&nt); result.Error != nil {
		return nil, result.Error
	}
	for _, n := range nt {
		dbFn.Table("notify").Where("id = ?", n.ID).Updates(map[string]interface{}{"ack": true})
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

	var m []*entity.Message

	if result := dbFn.Table("message").Find(&m); result.Error != nil {
		return nil, result.Error
	}

	return &entity.MessageList{
		Result: m,
	}, nil
}

func (n *Notify) FetchUser(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	dbFn := n.db.WithContext(ctx)

	var u *entity.User

	if result := dbFn.Table("users").Where("id = ?", userID).Find(&u); result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}
