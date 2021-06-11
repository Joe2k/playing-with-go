package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/Joe2k/playing-with-go/data"
)

func GetNotifications(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(data.GetNotifications())
}

func CreateNotification(w http.ResponseWriter, r *http.Request) {
	var notification data.Notification
	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		fmt.Println("error getting post: %w", err)
		return
	}
	notification.ID = strconv.Itoa(rand.Intn(100000000))
	json.NewEncoder(w).Encode(data.CreateNotification(notification))
}
