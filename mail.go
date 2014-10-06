package main

import (
	"github.com/sendgrid/sendgrid-go"
	"log"
	"os"
)

var (
	recipient    string
	sender       string
	sendgridUser string
	sendgridKey  string
	sg           *sendgrid.SGClient
)

func init() {
	if sendgridUser = os.Getenv("SENDGRID_USER"); sendgridUser == "" {
		log.Fatal("'SENDGRID_USER' must be set in the environment")
	}
	if sendgridKey = os.Getenv("SENDGRID_KEY"); sendgridKey == "" {
		log.Fatal("'SENDGRID_KEY' must be set in the environment")
	}
	if recipient = os.Getenv("RECIPIENT"); recipient == "" {
		log.Fatal("'RECIPIENT' must be set in the environment")
	}
	if sender = os.Getenv("SENDER"); sender == "" {
		log.Fatal("'SENDER' must be set in the environment")
	}

	sg = sendgrid.NewSendGridClient(sendgridUser, sendgridKey)
}

func notify(msg string) {
	message := sendgrid.NewMail()
	message.AddTo(recipient)
	message.SetSubject("UEM UPDATE")
	message.SetText(msg)
	message.SetFrom(sender)

	if err := sg.Send(message); err == nil {
		log.Printf("Email sent to %s from %s", recipient, sender)
	} else {
		log.Printf("ERROR (SENDGRID) => %s", err.Error())
	}
}
