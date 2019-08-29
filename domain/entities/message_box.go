package entities

import (
	"log"
	"sync"
)

// MessageBox aggregator
type MessageBox struct {
	Messages    []*Message
	MaxMessages int

	messagesMutex sync.Mutex
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
	box.messagesMutex.Lock()
	defer box.messagesMutex.Unlock()

	copy := box.Messages
	box.Messages = []*Message{}

	return copy
}

// Append a new message
func (box *MessageBox) Append(message *Message) {
	box.messagesMutex.Lock()
	defer box.messagesMutex.Unlock()

	boxMessagesLen := len(box.Messages)
	if boxMessagesLen >= box.MaxMessages {
		box.Messages = box.Messages[:boxMessagesLen-1]
		log.Printf("WARNING -> Removing messages because list is too long")
	}

	box.Messages = append(box.Messages, message)
}
