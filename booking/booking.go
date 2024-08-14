package booking

import (
	"context"
	"time"
)

type BookParams struct {
	Start time.Time
	Email string
}

//encore:api public method=POST path=/booking/:slug
func Book(ctx context.Context, slug string, p *BookParams) error {

	return nil
}
