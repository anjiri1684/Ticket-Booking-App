package models

import (
	"context"
	"time"
)

type Ticket struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	EventID   uint      `json:"eventId" gorm:"default:1"`
	Event     Event     `json:"event" gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Entered   bool      `json:"entered" gorm:"default:false"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}


type TicketRepository interface {
	GetMany(ctx context.Context)([]*Ticket, error)
	GetOne(ctx context.Context, ticketId uint) (*Ticket, error)
	CreateOne(ctx context.Context, ticket *Ticket) (*Ticket, error)
	UpdateOne(ctx context.Context, ticketId uint, updateData map[string]interface{}) (*Ticket, error)
}

type ValidateTicket struct {
	TicketId uint `json:"ticketId"`
}