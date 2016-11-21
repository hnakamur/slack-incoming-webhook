package main

import (
	"flag"
	"log"
	"os"

	slack "github.com/hnakamur/slack-incoming-webhook"
)

func main() {
	var url string
	var payload slack.Payload
	flag.StringVar(&url, "url", "", "slack incoming webhook URL")
	flag.StringVar(&payload.Text, "text", "", "text message to send")
	flag.StringVar(&payload.Username, "username", "", "sender's username")
	flag.StringVar(&payload.IconEmoji, "icon-emoji", "", "sender's icon emoji")
	flag.StringVar(&payload.IconURL, "icon-url", "", "sender's icon URL")
	flag.StringVar(&payload.Channel, "channel", "", "slack channel to send message")
	flag.Parse()

	err := slack.Send(url, payload)
	if err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}
}
