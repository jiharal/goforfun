package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type NotificationData struct {
	Message        string `json:"message"`
	CountdownValue int    `json:"countdown_value"`
}

var notificationDataChannel chan NotificationData = make(chan NotificationData)

func main() {
	go Notify()
	r := mux.NewRouter()
	r.HandleFunc("/notifications", Notification)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func CreateTimer(countdownValue int, msg string) {
	timer := time.NewTimer(time.Duration(countdownValue) * time.Second)
	go func() {
		<-timer.C
		SendNotification(msg)
	}()
}

func Notify() {
	for {
		select {
		case data := <-notificationDataChannel:
			CreateTimer(data.CountdownValue, data.Message)
		}
	}
}

func SendNotification(message string) {
	fmt.Println("Notification:", message)
}

func Notification(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var data NotificationData

	err = json.Unmarshal(body, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	notificationDataChannel <- data
	w.Write([]byte("OK"))
}
