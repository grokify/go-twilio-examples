package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/grokify/mogo/net/http/httputilmore"
	"github.com/grokify/mogo/time/month"
	"github.com/grokify/mogo/time/timeutil"
	"github.com/grokify/twiml"
	"github.com/rs/zerolog/log"
)

// https://www.twilio.com/docs/voice/twiml/gather
// https://github.com/BTBurke/twiml/blob/17fee1f07bf2c41d244d235c53db21cf610aa8a1/vocabulary.go

// HandleReminderStart will return XML to connect to the forwarding number
func HandleReminderStart() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Bind the request
		var cr twiml.VoiceRequest
		if err := twiml.Bind(&cr, r); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// Create a new response container
		res := twiml.NewResponse()

		switch status := cr.CallStatus; status {
		// Call is already in progress, tell Twilio to continue
		case twiml.InProgress:
			log.Info().Msg("C1_InProgress")
			nextWed := timeutil.NewTimeMore(time.Now(), 0).WeekdayNext(time.Wednesday)
			res.Add(&twiml.Say{
				Language: "en",
				Text: fmt.Sprintf("Hello, you have an appointment coming up on Wednesday %s %s",
					nextWed.Month(), month.DayofmonthToEnglish(uint16(nextWed.Day()))),
			})
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
				http.Error(w, http.StatusText(http.StatusBadGateway), http.StatusBadGateway)
				return
			}
			if _, err := w.Write(b); err != nil {
				http.Error(w, http.StatusText(http.StatusBadGateway), http.StatusBadGateway)
				return
			}
			w.Header().Set(
				httputilmore.HeaderContentType,
				httputilmore.ContentTypeAppXMLUtf8)
			w.WriteHeader(200)
			return
		}
	}
}
