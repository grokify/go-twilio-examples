package controllers

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/BTBurke/twiml"
	"github.com/grokify/gotilla/fmt/fmtutil"
	log "github.com/sirupsen/logrus"
)

// https://www.twilio.com/docs/voice/twiml/gather
// https://www.twilio.com/blog/2014/10/making-and-receiving-phone-calls-with-golang.html
// CallRequest will return XML to connect to the forwarding number
func Menu1() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Bind the request
		var evt GatherEvent
		if err := twiml.Bind(&evt, r); err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		log.Infof("I_GATHER_DIGITS [%v]", evt.Digits)
		fmtutil.PrintJSON(evt)

		// Create a new response container
		res := twiml.NewResponse()

		timestamp := getTimestampNow()
		log.Infof("I_MENU1_S1 [%v]", timestamp)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, http.StatusText(502), 502)
			return
		}
		log.Infof("GATHER_BODY_RAW: %s\n", string(body))

		qry, err := url.ParseQuery(string(body))
		if err != nil {
			http.Error(w, http.StatusText(502), 502)
			return
		}
		keys := []string{}
		for k := range qry {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		log.Infof("GATHER_BODY_KEYS: %s\n", strings.Join(keys, ","))

		evt.Digits = strings.TrimSpace(evt.Digits)
		switch evt.Digits {
		case "1":
			c := twiml.Say{
				Language: "en",
				Text:     "Thank you for confirming. We will see you then. Good bye.",
			}
			res.Add(&c)
			processResponse(w, r, res)
			return
		case "2":
			c := twiml.Say{
				Language: "en",
				Text:     "Our office will call you back to reschedule. Good bye.",
			}
			res.Add(&c)
			processResponse(w, r, res)
			return
		case "3":
			c := twiml.Say{
				Language: "en",
				Text:     "We will note the cancellation. Please call us again when are you can book another appointment. Good bye.",
			}
			res.Add(&c)
			processResponse(w, r, res)
			return
		}
		c := twiml.Say{
			Language: "en",
			Text:     "I did not understand that.",
		}
		res.Add(&c)

		addMainMenu(res)

		log.Info("GATHER_S1")
		log.Info(string(body))

		processResponse(w, r, res)
		return
	}
}
