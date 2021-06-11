package data

// Notification struct
type Notification struct {
	ID       string `json:"id"`
	Receiver *User  `json:"receiver"`
	Content  string `json:"content"`
}

// User struct
type User struct {
	Name   string `json:"name"`
	Number string `json:"number"`
}

var notifications []Notification

func GetNotifications() []Notification {
	return notifications
}

func CreateNotification(notification Notification) Notification {
	notifications = append(notifications, notification)
	return notification
}
