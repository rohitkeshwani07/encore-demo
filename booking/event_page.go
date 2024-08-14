package booking

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const DefaultBookingDuration = 1 * time.Hour
const DefaultBeforBufferTime = 0
const DefaultAfterBufferTime = 0

type CreateEventPageRequest struct {
	Title       string
	Description string
	EventType   string
	Schedule    *Schedule
}

type CreateEventPageResponse struct {
	Id  uuid.UUID
	Url string
}

//encore:api public method=POST path=/event_page/create
func CreateEventPage(ctx context.Context, params CreateEventPageRequest) (*CreateEventPageResponse, error) {

	return &CreateEventPageResponse{
		Url: params.Title,
	}, nil
}
