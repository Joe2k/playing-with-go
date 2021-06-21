package app

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Joe2k/playing-with-go/data"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Main Handlers

func (a *App) getNotification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Notification ID")
		return
	}

	n := data.Notification{ID: id}
	if err := n.GetNotification(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Notification not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, n)
}

func (a *App) createNotification(w http.ResponseWriter, r *http.Request) {
	var n data.Notification
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&n); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}
	defer r.Body.Close()

	if err := n.CreateNotification(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, n)
}

func (a *App) getNotifications(w http.ResponseWriter, r *http.Request) {
	notifications, err := data.GetNotifications(a.DB)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, notifications)
}

// Helper Funcitions

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
