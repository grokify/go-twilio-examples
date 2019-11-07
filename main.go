package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/grokify/go-appointment-reminder-demo/controllers"
)

const DefaultPort string = "8081"

func main() {
	http.HandleFunc("/reminder_start", controllers.HandleReminderStart())
	http.HandleFunc("/reminder_process", controllers.HandleReminderProcess())

	port := os.Getenv("PORT")
	portStr := ":" + DefaultPort
	if len(port) > 0 {
		portStr = ":" + port
	}
	fmt.Printf("Running on [%v]\n", portStr)
	http.ListenAndServe(portStr, nil)
}
