package log

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Client struct {
	application string
	instance    string
	address     string
	Channel     chan *LogPayload
}

func NewClient(application string, instance string, address string) *Client {
	client := &Client{application: application, instance: instance, address: address}
	client.Channel = make(chan *LogPayload)
	go func() {
		for true {
			message := <-client.Channel
			err := client.SendLog(message)
			if err != nil {
				logrus.Error(err)
			}
		}
	}()
	return client
}
func (c *Client) SendLog(payload *LogPayload) error {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(*payload)
	if err != nil {
		return err
	}
	go func() {
		_, err = http.Post(c.address, "application/json", b)
		if err != nil {
			logrus.Error(err)
		}
	}()
	return err
}
