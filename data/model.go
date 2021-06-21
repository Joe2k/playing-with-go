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

func (n *Notification) UpdateNotification(db *sql.DB) error {
	_, err := db.Exec("UPDATE notifications SET number=$1, message=$2 WHERE id=$3", n.Number, n.Message, n.ID)

	return err
}

func (n *Notification) DeleteNotification(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM notifications WHERE id=$1", n.ID)

	return err
}

func GetNotifications(db *sql.DB) ([]Notification, error) {
	rows, err := db.Query("SELECT id, number, message FROM notifications")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	notifications := []Notification{}

	for rows.Next() {
		var n Notification
		if err := rows.Scan(&n.ID, &n.Number, &n.Message); err != nil {
			return nil, err
		}
		notifications = append(notifications, n)
	}

	return notifications, nil
}
