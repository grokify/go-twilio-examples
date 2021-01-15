package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/grokify/twiml"
	"github.com/rs/zerolog/log"
)

// https://www.twilio.com/docs/voice/twiml/gather
// https://www.twilio.com/blog/2014/10/making-and-receiving-phone-calls-with-golang.html
// CallRequest will return XML to connect to the forwarding number
func HandleReminderMenu() func(http.ResponseWriter, *http.Request) {
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

		// Create a new response container
		res := twiml.NewResponse()

		if numTries > AppRetryLimit {
			res.Add(
				&twiml.Say{
					Language: "en",
					Text:     "If you would like to change your appointment, please call our offices. Good bye."},
				&twiml.Hangup{})
			processResponse(w, r, res)
		} else {
			addMainMenu(res, uint16(numTries+1))
			processResponse(w, r, res)
		}
	}
}
