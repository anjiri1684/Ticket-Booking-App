package models

import (
	"context"
	"time"
)

type Event struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}



type EventRepository interface {
	Getmany(ctx context.Context) ([]*Event, error)
	GetOne(ctx context.Context, eventId uint) (*Event, error)
	CreateOne(ctx context.Context, event *Event)(*Event, error)
	UpdateOne(ctx context.Context, eventId uint, updatedata map[string]interface{}) (*Event, error)
	DeleteOne(ctx context.Context, eventId uint)  error
}
