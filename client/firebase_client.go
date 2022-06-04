package client

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type FirebaseClient struct {
	app    *firebase.App
	client *messaging.Client
}

func NewFirebaseClient() *FirebaseClient {
	opt := option.WithCredentialsFile("firebase.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Printf("Error [NewFirebaseClient]: %v", err)
		return nil
	}

	client, err := app.Messaging(context.Background())

	if err != nil {
		log.Printf("Error [NewFirebaseClient]: %v", err)
		return nil
	}

	return &FirebaseClient{
		app:    app,
		client: client,
	}
}

func (fc *FirebaseClient) SendTopic(message string) error {
	notification := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Backend Notification",
			Body:  message,
		},
		Topic: "live_score",
	}

	_, err := fc.client.Send(context.Background(), notification)

	return err
}
