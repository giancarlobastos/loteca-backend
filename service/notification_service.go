package service

import (
	"github.com/giancarlobastos/loteca-backend/client"
)

type NotificationService struct {
	firebaseClient *client.FirebaseClient
}

func NewNotificationService(
	firebaseClient *client.FirebaseClient) *NotificationService {
	return &NotificationService{
		firebaseClient: firebaseClient,
	}
}

func (ns *NotificationService) SendNotification(message string) error {
	return ns.firebaseClient.SendTopic(message)
}
