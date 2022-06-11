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
	// listOfMessages := make([]Message, length)
	return &nopMessageTracker{
		capacity: length,
		message:  nil,
		size:     0,
	}
}

type nopMessageTracker struct {
	capacity int
	message  []*Message
	size     int
}

// Add will add a message to the tracker, deleting the oldest message if necessary.
func (dmt *nopMessageTracker) Add(message *Message) (err error) {
	//check for the duplicate message
	_, err = getMessageIndexByID(message.ID, dmt.message)
	if err == nil {
		return
	}
	err = nil

	//check if the size of the list of message is full
	//if yes, then delete the oldest
	if dmt.capacity == dmt.size {
		dmt.message = append(dmt.message[:0], dmt.message[1:]...)
	}

	//add the incoming message to the last of the list of messages
	dmt.message = append(dmt.message, message)

	//keep the size of the list of messages updated
	dmt.size = len(dmt.message)
	return
}

// Delete will delete a message from tracker
func (dmt *nopMessageTracker) Delete(id string) (err error) {
	//check for the index of the message
	key, err := getMessageIndexByID(id, dmt.message)
	if err != nil {
		return err
	}
	//delete the messages by keeping the list of messages ordered
	dmt.message = append(dmt.message[:key], dmt.message[key+1:]...)

	//keep the size of the list of messages updated
	dmt.size = len(dmt.message)
	return
}

// Get returns a message for a given ID.  Message is retained in tracker
func (dmt *nopMessageTracker) Message(id string) (message *Message, err error) {
	//check for the index of the message
	key, err := getMessageIndexByID(id, dmt.message)
	if err != nil {
		return nil, err
	}

	return dmt.message[key], nil
}

// All returns messages in the order in which they were received
func (dmt *nopMessageTracker) Messages() (messages []*Message) {
	return dmt.message
}
func getMessageIndexByID(id string, mt []*Message) (int, error) {
	for key, value := range mt {
		if value.ID == id {
			return key, nil
		}
	}
	return -1, ErrMessageNotFound
}
