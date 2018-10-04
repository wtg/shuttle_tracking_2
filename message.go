package shuttletracker

import (
	"errors"
	"time"
)

// Message represents a message displayed to users.
type Message struct {
	Message string    `json:"message"`
	Enabled bool      `json:"enabled"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// MessageService is an interface for interacting with Messages.
type MessageService interface {
	Message() (*Message, error)
	SetMessage(message *Message) error
}

var (
	// ErrMessageNotFound indicates that a Message is not in the database.
	ErrMessageNotFound = errors.New("message not found")
)
