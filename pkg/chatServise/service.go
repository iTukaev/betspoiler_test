package chatServise

import (
	"bytes"
	"fmt"
)

type chat struct {
	buf bytes.Buffer
}

type Chat interface {
	AddMessage(name string, message string) error
	GetAllMessage() *bytes.Buffer
}

func NewChat() Chat {
	return new(chat)
}

func (c *chat) AddMessage(name string, message string) error {

	if _, err := c.buf.WriteString(fmt.Sprintf("%s : %s\n", name, message)); err != nil {
		return err
	}
	return nil
}

func (c *chat) GetAllMessage() *bytes.Buffer {
	return &c.buf
}