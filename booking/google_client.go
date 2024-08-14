package booking

import (
	"context"
	"encoding/json"

	"google.golang.org/api/calendar/v3"
)

func authRedirect(ctx context.Context, params CreateEventPageRequest) (freeBusyResponse *calendar.FreeBusyResponse, err error) {
	_ = json.Unmarshal([]byte(`{"kind":"calendar#freeBusy","timeMin":"2024-08-13T16:09:53.000Z","timeMax":"2024-08-15T16:09:53.000Z","calendars":{"rohitkeshwani07@gmail.com":{"busy":[{"start":"2024-08-13T17:30:00Z","end":"2024-08-14T02:00:00Z"},{"start":"2024-08-14T02:30:00Z","end":"2024-08-14T03:00:00Z"},{"start":"2024-08-14T03:30:00Z","end":"2024-08-14T08:00:00Z"},{"start":"2024-08-14T14:30:00Z","end":"2024-08-14T15:30:00Z"},{"start":"2024-08-14T17:30:00Z","end":"2024-08-15T02:30:00Z"},{"start":"2024-08-15T03:30:00Z","end":"2024-08-15T08:00:00Z"},{"start":"2024-08-15T14:30:00Z","end":"2024-08-15T15:30:00Z"}]}}}`), freeBusyResponse)
	return freeBusyResponse, nil
}
