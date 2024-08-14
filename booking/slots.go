package booking

import (
	"context"
	"time"
)

type Slot struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type BookableSlot struct {
	Date  time.Time `json:"date"`
	Slots []*BookableSlot
}

type SlotsResponse struct {
	Slots []*BookableSlot
}

//encore:api public method=GET path=/slots/:slug
func GetBookableSlots(ctx context.Context, slug string) (*SlotsResponse, error) {
	return &SlotsResponse{}, nil
}
