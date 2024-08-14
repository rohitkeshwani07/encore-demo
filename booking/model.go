package booking

import (
	"time"

	"encore.dev/storage/sqldb"
	"github.com/google/uuid"
)

type Calender struct {
	Id                 uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();"`
	ExternalId         string
	ExternalType       string
	OwnerId            uuid.UUID
	ExternalPermission string
}

type User struct {
	Id    uuid.UUID
	Email string
}

type Event struct {
	Id           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	CalenderId   uuid.UUID
	OwnerId      uuid.UUID
	StartTime    time.Time
	EndTime      time.Time
	Seq          uint8
	EventPayload string
	ICalPayload  string
	ICalUID      string
	ExternalId   string
	ExternalType string
	Attendees    []*Attendees `gorm:"serializer:json"`
	Guest        []*Guest     `gorm:"serializer:json"`
}

type Attendees struct {
	Email      string
	CalenderId uuid.UUID
}

type Guest struct {
	Email string
}

type EventPage struct {
	Id            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Slug          string
	Title         string
	Description   string
	DurationInSec int
	Expiry        time.Time
	EventType     string
	Schedule      *Schedule `gorm:"serializer:json"`
	HostId        uuid.UUID
}

type Schedule struct {
	BeforBufferTime int
	AfterBufferTime int
	SpotStepInSec   int
	PeriodType      int
	MinNoticeInSec  int
	Availability    []*Availability `gorm:"serializer:json"`
}

type Availability struct {
	Date      time.Time
	StartTime time.Time
	EndTime   time.Time
}

type Booking struct {
	ID          uuid.UUID
	EventPageId uuid.UUID
	Start       time.Time
	End         time.Time
	Email       string
}

var (
	bookingDB = sqldb.NewDatabase("booking_system", sqldb.DatabaseConfig{
		Migrations: "./migrations",
	})
)
