package entities

import (
	"log"
	"sync"
)

// MessageBox aggregator
type MessageBox struct {
	Messages      []*Message
	MessagesMutex sync.Mutex
	MaxMessages   int
}

// NewMessageBox constructor
func NewMessageBox() *MessageBox {
	return &MessageBox{
		MaxMessages: 100,
		Messages:    []*Message{},
	}
}

// Read drop older messages
func (box *MessageBox) Read() []*Message {
	box.MessagesMutex.Lock()
	defer box.MessagesMutex.Unlock()

	copy := box.Messages
	box.Messages = []*Message{}

	return copy
}

// Append a new message
func (box *MessageBox) Append(message *Message) {
	box.MessagesMutex.Lock()
	defer box.MessagesMutex.Unlock()

	boxMessagesLen := len(box.Messages)
	if boxMessagesLen >= box.MaxMessages {
		box.Messages = box.Messages[:boxMessagesLen-1]
		log.Printf("WARNING -> Removing messages because list is too long")
	}

	box.Messages = append(box.Messages, message)
}
