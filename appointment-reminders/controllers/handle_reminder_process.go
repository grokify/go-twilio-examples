package controllers

import (
	"net/http"
	"strconv"
	"strings"

	twilio "github.com/grokify/go-twilio-examples"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/twiml"
	"github.com/rs/zerolog/log"
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

		log.Info().
			Str("requestUri", r.URL.RequestURI()).
			Int("numTriesParsed", numTries).
			Msg("Handle_Reminder_Process")

		// Bind the request
		var evt twilio.GatherEvent
		if err := twiml.Bind(&evt, r); err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		err := fmtutil.PrintJSON(evt)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

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
				res.Add(
					&twiml.Say{
						Language: "en",
						Text:     "You have reached the maximum number of retries allowed. Please hang up and call our office if you have any questions on your appointment."},
					&twiml.Hangup{})
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
