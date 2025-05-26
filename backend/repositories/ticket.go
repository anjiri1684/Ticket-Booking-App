package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/anjiri1684/ticket-booking-project-v1/models"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}
func (r *TicketRepository)GetMany(ctx context.Context, userId uint) ([]*models.Ticket, error) {
	tickets := []*models.Ticket{}

	res := r.db.Model(&models.Ticket{}).Where("user_id = ?", userId).Preload("Event").Order("updated_at desc").Find(&tickets)

	if res.Error != nil {
		return nil, res.Error
	}

	return tickets, nil
}

func (r *TicketRepository)GetOne(ctx context.Context,userId uint, ticketId uint) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	res := r.db.Model(ticket).Where("id = ?", ticketId).Where("user_id = ?", userId).Preload("Event").First(ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil
}

func (r *TicketRepository)CreateOne(ctx context.Context,userId uint, ticket *models.Ticket) (*models.Ticket, error) {

	ticket.UserID = userId

	res := r.db.Model(ticket).Create(ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return r.GetOne(ctx,userId, ticket.ID)
}

func (r *TicketRepository) UpdateOne(ctx context.Context, userId uint, ticketId uint, updateData map[string]interface{}) (*models.Ticket, error) {
	if ticketId == 0 || userId == 0 {
		return nil, errors.New("invalid ticketId or userId")
	}

	ticket := &models.Ticket{}

	// Apply the update with strict user scoping
	updateRes := r.db.WithContext(ctx).
		Model(ticket).
		Where("id = ? AND user_id = ?", ticketId, userId).
		Updates(updateData)

	// Handle DB errors
	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	// Handle no rows updated
	if updateRes.RowsAffected == 0 {
		return nil, fmt.Errorf("no ticket found for id=%d and user_id=%d", ticketId, userId)
	}

	// Fetch and return updated ticket
	return r.GetOne(ctx, userId, ticketId)
}


func NewTicketRepository(db *gorm.DB) models.TicketRepository {
	return &TicketRepository{
		db: db,
	}
}