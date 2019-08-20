package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/BTBurke/twiml"
	"github.com/grokify/go-appointment-reminder-demo/twilio"
	"github.com/grokify/gotilla/fmt/fmtutil"
	log "github.com/sirupsen/logrus"
)

// https://www.twilio.com/docs/voice/twiml/gather
// https://www.twilio.com/blog/2014/10/making-and-receiving-phone-calls-with-golang.html
// CallRequest will return XML to connect to the forwarding number
func HandleReminderProcess() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		qry := r.URL.Query()
		numTries := 0
		numTriesRaw := strings.TrimSpace(qry.Get("numTries"))
		if len(numTriesRaw) > 0 {
			numTriesTry, err := strconv.Atoi(numTriesRaw)
			if err == nil {
				numTries = numTriesTry
			}
		}
		log.WithFields(log.Fields{
			"requestUri":     r.URL.RequestURI(),
			"numTriesParsed": strconv.Itoa(numTries),
		}).Info("Handle_Reminder_Process")

		// Bind the request
		var evt twilio.GatherEvent
		if err := twiml.Bind(&evt, r); err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		fmtutil.PrintJSON(evt)

		// Create a new response container
		res := twiml.NewResponse()

		switch strings.TrimSpace(evt.Digits) {
		case "1":
			res.Add(&twiml.Say{
				Language: "en",
				Text:     "Thank you for confirming. We will see you then. Good bye."})
			processResponse(w, r, res)
		case "2":
			res.Add(&twiml.Say{
				Language: "en",
				Text:     "Our office will call you back to reschedule. Good bye."})
			processResponse(w, r, res)
		case "3":
			res.Add(&twiml.Say{
				Language: "en",
				Text:     "We will note the cancellation. Please call us again when are you can book another appointment. Good bye."})
			processResponse(w, r, res)
		default:
			if numTries > AppRetryLimit {
				res.Add(&twiml.Say{
					Language: "en",
					Text:     "You have reached the maximum number of retries allowed. Please hang up and call our office if you have any questions on your appointment."})
				res.Add(&twiml.Hangup{})
				processResponse(w, r, res)
			} else {
				res.Add(&twiml.Say{
					Language: "en",
					Text:     "I did not understand that. Please try again."})
				addMainMenu(res, uint16(numTries+1))
				processResponse(w, r, res)
			}
		}
	}
}
