package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/Joe2k/playing-with-go/app"
	"github.com/joho/godotenv"
)

var a app.App

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)

	code := m.Run()
	clearTable()
	os.Exit(code)
}

// All Testing Methods

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/notifications", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestCreateProduct(t *testing.T) {
	clearTable()

	var jsonStr = []byte(`{"number":1234567, "message": "Test Message"}`)

	req, _ := http.NewRequest("POST", "/notification", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["number"] != 1234567.0 {
		t.Errorf("Expected product name to be ‘1234567’. Got ‘%v’", m["number"])
	}
	if m["message"] != "Test Message" {
		t.Errorf("Expected product price to be ‘Test Message’. Got ‘%v’", m["message"])
	}
	if m["id"] != 1.0 {
		t.Errorf("Expected product ID to be ‘1’. Got ‘%v’", m["id"])
	}
}

func TestGetProduct(t *testing.T) {
	clearTable()
	addProducts(1)

	req, _ := http.NewRequest("GET", "/notification/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateProduct(t *testing.T) {
	clearTable()
	addProducts(1)

	req, _ := http.NewRequest("GET", "/notification/1", nil)
	response := executeRequest(req)

	var originalNotification map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalNotification)

	var jsonStr = []byte(`{"number":123456, "message":"New Updated Message"}`)
	req, _ = http.NewRequest("PUT", "/notification/1", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var n map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &n)

	if n["id"] != originalNotification["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalNotification["id"], n["id"])
	}

	if n["number"] == originalNotification["number"] {
		t.Errorf("Expected the number to change from '%v' to '%v'. Got '%v'", originalNotification["number"], n["number"], n["number"])
	}

	if n["message"] == originalNotification["message"] {
		t.Errorf("Expected the message to change from '%v' to '%v'. Got '%v'", originalNotification["message"], n["message"], n["message"])
	}
}

func TestGetNonExistentProduct(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/notification/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Notification not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Notification not found'. Got '%s'", m["error"])
	}
}

// Helper Methods for Testing

func clearTable() {
	a.DB.Exec("DELETE FROM notifications")
	a.DB.Exec("ALTER SEQUENCE notifications_id_seq RESTART WITH 1")
}

func addProducts(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO notifications(number, message) VALUES($1, $2)", (i+1)*1000000, "Notification "+strconv.Itoa(i))
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
