package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/slack-go/slack"
)

type file struct {
	Token    string    `json:"bot_token"`
	Messages []message `json:"channels"`
}

type message struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func (f file) validate() error {
	var errMessages []string

	if f.Token == "" {
		errMessages = append(errMessages, "token can't be empty")
	}

	if len(f.Messages) == 0 {
		errMessages = append(errMessages, "messages can't be empty")
	}

	for i := range f.Messages {
		if f.Messages[i].Channel == "" {
			errMessages = append(errMessages, fmt.Sprintf("message #%d has empty channel", i))
		}

		if f.Messages[i].Text == "" {
			errMessages = append(errMessages, fmt.Sprintf("message #%d has empty text", i))
		}
	}

	if len(errMessages) > 0 {
		return errors.New(strings.Join(errMessages, ", "))
	}

	return nil
}

func run() error {
	f := flag.String("f", "message.json", "path to message.json")
	flag.Parse()

	if f == nil || *f == "" {
		return fmt.Errorf("path to file can't be empty")
	}

	byteValue, err := ioutil.ReadFile(*f)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	var file file
	if err := json.Unmarshal(byteValue, &file); err != nil {
		return fmt.Errorf("unmarshalling: %w", err)
	}

	if err := file.validate(); err != nil {
		return fmt.Errorf("validate file: %w", err)
	}

	api := slack.New(file.Token)
	if _, err := api.AuthTest(); err != nil {
		return fmt.Errorf("slack client auth check: %w", err)
	}

	for i, msg := range file.Messages {
		channelID, _, err := api.PostMessage(msg.Channel, slack.MsgOptionText(msg.Text, false))
		if err != nil {
			return fmt.Errorf("sending message #%d to channel %q: %w", i, channelID, err)
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
