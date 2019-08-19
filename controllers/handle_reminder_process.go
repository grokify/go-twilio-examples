package controllers

import (
	"net/http"
	"strings"

	"github.com/BTBurke/twiml"
	"github.com/grokify/go-appointment-reminder-demo/twilio"
	"github.com/grokify/gotilla/fmt/fmtutil"
)

// https://www.twilio.com/docs/voice/twiml/gather
// https://www.twilio.com/blog/2014/10/making-and-receiving-phone-calls-with-golang.html
// CallRequest will return XML to connect to the forwarding number
func HandleReminderProcess() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Bind the request
		var evt twilio.GatherEvent
		if err := twiml.Bind(&evt, r); err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		fmtutil.PrintJSON(evt)

		// Create a new response container
		res := twiml.NewResponse()

		evt.Digits = strings.TrimSpace(evt.Digits)
		switch evt.Digits {
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
			res.Add(&twiml.Say{
				Language: "en",
				Text:     "I did not understand that."})
			addMainMenu(res)
			processResponse(w, r, res)
		}
	}
}
