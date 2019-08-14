package entities

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MessageBoxSuite struct {
	suite.Suite
}

func (suite *MessageBoxSuite) TestNewMessageBox() {
	box := NewMessageBox()
	suite.NotNil(box)
	suite.Equal(100, box.MaxMessages)
}

func (suite *MessageBoxSuite) TestAppend() {
	box := NewMessageBox()
	suite.NotNil(box)
	suite.Equal(100, box.MaxMessages)

	message := NewMessage()
	message.Body = "Testing"
	wg := sync.WaitGroup{}

	for i := 0; i < 300; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			box.Append(message)
		}()
	}

	wg.Wait()

	suite.Equal(100, len(box.Messages))
}

func (suite *MessageBoxSuite) TestRead() {
	box := NewMessageBox()
	suite.NotNil(box)
	suite.Equal(100, box.MaxMessages)

	message := NewMessage()
	message.Body = "Testing"
	wg := sync.WaitGroup{}

	for i := 0; i < 300; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			box.Append(message)
		}()
	}

	wg.Wait()

	suite.Equal(100, len(box.Messages))

	messages := box.Read()

	suite.Zero(len(box.Messages))

	for _, readedMessage := range messages {
		suite.Equal(readedMessage, message)
	}

}

func TestMessageBoxSuite(t *testing.T) {
	suite.Run(t, new(MessageBoxSuite))
}
