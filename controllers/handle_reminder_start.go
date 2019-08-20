package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/BTBurke/twiml"
	hum "github.com/grokify/gotilla/net/httputilmore"
	tu "github.com/grokify/gotilla/time/timeutil"
	log "github.com/sirupsen/logrus"
)

// https://www.twilio.com/docs/voice/twiml/gather
// https://github.com/BTBurke/twiml/blob/17fee1f07bf2c41d244d235c53db21cf610aa8a1/vocabulary.go

// HandleReminderStart will return XML to connect to the forwarding number
func HandleReminderStart() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Bind the request
		var cr twiml.VoiceRequest
		if err := twiml.Bind(&cr, r); err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		// Create a new response container
		res := twiml.NewResponse()

		switch status := cr.CallStatus; status {

		// Call is already in progress, tell Twilio to continue
		case twiml.InProgress:
			log.Info("C1_InProgress")
			nextWed := tu.NextWeekday(time.Wednesday)
			d := twiml.Say{
				Language: "en",
				Text: fmt.Sprintf("Hello, you have an appointment coming up on Wednesday %s %s",
					nextWed.Month(), tu.DayofmonthToEnglish(uint16(nextWed.Day()))),
			}
			res.Add(&d)
			addMainMenu(res, uint16(0))
			processResponse(w, r, res)
		// Call is ringing but has not been connected yet, tell Twilio to continue
		case twiml.Ringing:
			w.WriteHeader(200)
			return
		case twiml.Queued:
			w.WriteHeader(200)
			return

		// Call is over, hang up
		default:
			res.Add(&twiml.Hangup{})
			b, err := res.Encode()
			if err != nil {
				http.Error(w, http.StatusText(502), 502)
				return
			}
			if _, err := w.Write(b); err != nil {
				http.Error(w, http.StatusText(502), 502)
				return
			}
			w.Header().Set(hum.HeaderContentType, hum.ContentTypeAppXmlUtf8)
			w.WriteHeader(200)
			return
		}
	}
}
