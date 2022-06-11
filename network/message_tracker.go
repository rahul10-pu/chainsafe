package network

import (
	"errors"
)

// MessageTracker tracks a configurable fixed amount of messages.
// Messages are stored first-in-first-out.  Duplicate messages should not be stored in the queue.
type MessageTracker interface {
	// Add will add a message to the tracker, deleting the oldest message if necessary
	Add(message *Message) (err error)
	// Delete will delete message from tracker
	Delete(id string) (err error)
	// Get returns a message for a given ID.  Message is retained in tracker
	Message(id string) (message *Message, err error)
	// Messages returns messages in FIFO order
	Messages() (messages []*Message)

	// FOLLOW UP: remove from code before sending to applicant
	// PeerCounts returns all the peers and their counts per message id accumulated by the tracker
	// PeerCounts(id string) (peerCounts map[string]int, err error)
}

// ErrMessageNotFound is an error returned by MessageTracker when a message with specified id is not found
var ErrMessageNotFound = errors.New("message not found")

func NewMessageTracker(length int) MessageTracker {
	return newNoopMessageTracker(length)
}

func newNoopMessageTracker(length int) *nopMessageTracker {
	return &nopMessageTracker{}
}

type nopMessageTracker struct{}

// Add will add a message to the tracker, deleting the oldest message if necessary.
func (dmt *nopMessageTracker) Add(message *Message) (err error) {
	return
}

// Delete will delete a message from tracker
func (dmt *nopMessageTracker) Delete(id string) (err error) {
	return
}

// Get returns a message for a given ID.  Message is retained in tracker
func (dmt *nopMessageTracker) Message(id string) (message *Message, err error) {
	return
}

// All returns messages in the order in which they were received
func (dmt *nopMessageTracker) Messages() (messages []*Message) {
	return
}
