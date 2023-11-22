package tg

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

const (
	tgURL   = "https://api.telegram.org/bot%s/sendMessage"
	reqBody = "chat_id=%s&text=%s"
	encode  = "application/x-www-form-urlencoded"
)

type Message struct {
	message  string
	groupsID []string
	token    string
}

func New(_message, _token string, _groupsID []string) *Message {
	return &Message{
		message:  _message,
		groupsID: _groupsID,
		token:    _token,
	}
}

func (m *Message) Send() error {
	wg := sync.WaitGroup{}
	errChan := make(chan error, len(m.groupsID))

	for _, chatID := range m.groupsID {
		wg.Add(1)
		go m.sendAsync(strings.TrimSpace(chatID), &wg, errChan)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Message) sendAsync(chatID string, wg *sync.WaitGroup, errChan chan error) {
	defer wg.Done()
	url := fmt.Sprintf(tgURL, m.token)
	body := fmt.Sprintf(reqBody, chatID, m.message)
	resp, err := http.Post(url, encode, bytes.NewBufferString(body))
	if err != nil {
		errChan <- err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		errChan <- fmt.Errorf("Message in chat with id:%s not send", chatID)
	}
}
