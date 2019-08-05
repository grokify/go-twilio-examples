package demo

import (
	"fmt"
	"net/http"
	"time"

	"github.com/BTBurke/twiml"
	hum "github.com/grokify/gotilla/net/httputilmore"
	tu "github.com/grokify/gotilla/time/timeutil"
	log "github.com/sirupsen/logrus"
)

// CallRequest will return XML to connect to the forwarding number
func CallRequest() func(http.ResponseWriter, *http.Request) {
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
			// Add the verb to the response
			res.Add(&d)

			c := twiml.Say{
				Language: "en",
				Text:     "Please press 1 or say confirm to confirm. Press 2 or say cancel to cancel. Press 3 or say reschedule to reschedule.",
			}
			res.Add(&c)

			// https://www.twilio.com/docs/voice/twiml/gather
			// https://github.com/BTBurke/twiml/blob/17fee1f07bf2c41d244d235c53db21cf610aa8a1/vocabulary.go
			g := twiml.Gather{
				Input:     "speech dtmf",
				Timeout:   3,
				NumDigits: 1}
			res.Add(&g)

			// Validate and encode the response.  Validation is done
			// automatically before the response is encoded.
			b, err := res.Encode()
			if err != nil {
				log.Info("C3_InProgress_Error_Encode")
				http.Error(w, http.StatusText(502), 502)
				return
			}

			// Write the XML response to the http.ReponseWriter
			if _, err := w.Write(b); err != nil {
				log.Info("C3_InProgress_Error_Write")
				http.Error(w, http.StatusText(502), 502)
				return
			}
			w.Header().Set(hum.HeaderContentType, hum.ContentTypeAppXmlUtf8)
			w.WriteHeader(200)
			return

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
