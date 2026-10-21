package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/grokify/go-twilio-examples/appointment-reminders/controllers"
)

const DefaultPort string = "8081"

func main() {
	http.HandleFunc("/reminder_start", controllers.HandleReminderStart())
	http.HandleFunc("/reminder_menu", controllers.HandleReminderMenu())
	http.HandleFunc("/reminder_process", controllers.HandleReminderProcess())

	port := os.Getenv("PORT")
	portStr := ":" + DefaultPort
	if len(port) > 0 {
		portStr = ":" + port
	}
	fmt.Printf("Running on [%v]\n", portStr)
	srv := &http.Server{
		Addr:              portStr,
		ReadHeaderTimeout: 10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
