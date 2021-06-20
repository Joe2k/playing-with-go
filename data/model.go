package data

import "database/sql"

// Notification struct
type Notification struct {
	ID      int    `json:"id"`
	Number  int    `json:"number"`
	Message string `json:"message"`
}

func (n *Notification) GetNotification(db *sql.DB) error {
	return db.QueryRow("SELECT number, message FROM notifications WHERE id=$1", n.ID).Scan(&n.Number, &n.Message)
}

func (n *Notification) CreateNotification(db *sql.DB) error {
	return db.QueryRow("INSERT INTO notifications(number, message) VALUES($1, $2) RETURNING id", n.Number, n.Message).Scan(&n.ID)
}
