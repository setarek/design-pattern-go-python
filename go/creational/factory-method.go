package main

import "fmt"

// Product interface
type Notification interface {
	send(message string)
}

// Concerete products
type SMSNotification struct{}

func (n *SMSNotification) send(message string) {
	fmt.Println("Send SMS: ", message)
}

type EmailNotification struct{}

func (n *EmailNotification) send(message string) {
	fmt.Println("Send Email: ", message)
}

// Creator Interface
type NotificationFactory interface {
	CreateNotification() Notification
}

// Concerete Creators

type EmailNotificationFactory struct{}

func (n *EmailNotificationFactory) CreateNotification() Notification {
	return &EmailNotification{}
}

type SMSNotificationFactory struct{}

func (n *SMSNotificationFactory) CreateNotification() Notification {
	return &SMSNotification{}
}

func notify(factory NotificationFactory, message string) {
	notificationFactory := factory.CreateNotification()
	notificationFactory.send(message)
}

func main() {

	var factory NotificationFactory
	notificationType := "email"

	switch notificationType {
	case "email":
		factory = &EmailNotificationFactory{}
	case "sms":
		factory = &SMSNotificationFactory{}
	default:
		fmt.Println("Invalid notification type")
		return
	}

	notify(factory, "Hello..")
}
