package repositories

import (
	"context"
	"fmt"

	"github.com/anjiri1684/ticket-booking-project-v1/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (r *EventRepository) Getmany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	res := r.db.Model(&models.Event{}).Order("updated_at desc").Find(&events)

	if res.Error != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	return events, nil
}

func (r *EventRepository)GetOne(ctx context.Context, eventId uint)(*models.Event, error) {
	event := &models.Event{}

	res := r.db.Model(event).Where("id = ?", eventId).First(event)

	if res.Error != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	return event, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event *models.Event) (*models.Event, error) {
	res := r.db.WithContext(ctx).Create(event)

	if res.Error != nil {
		return nil, fmt.Errorf("db error: %w", res.Error)
	}

	return event, nil
}


func (r *EventRepository)UpdateOne(ctx context.Context, eventId uint, updatedata map[string]interface{}) (*models.Event, error){
	event := &models.Event{}

	updateRes := r.db.Model(event).Where("id = ?", eventId).Updates(updatedata)

	if updateRes.Error != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	getRes := r.db.Where("id = ?", eventId).First(event)

	if getRes.Error != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	return event, nil

}

func (r *EventRepository)DeleteOne(ctx context.Context, eventId uint) error {
	res := r.db.Delete(&models.Event{}, eventId)

	return res.Error
}

func NewEventRepository(db *gorm.DB)models.EventRepository {
	return &EventRepository{
		db: db,
	}
}