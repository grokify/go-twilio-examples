package main

import (
	"fmt"
	"net/http"

	"github.com/grokify/mogo/encoding/jsonutil"
	"github.com/grokify/mogo/net/httputilmore"
	"github.com/grokify/twiml"
	"github.com/rs/zerolog/log"

	"github.com/grokify/go-twilio-examples"
)

// HandleCall
func (svc *Service) HandleCall(w http.ResponseWriter, r *http.Request) {
	log.Info().Str("stage", "start").Msg("HandleCall")
	// Bind the request
	var cr twiml.VoiceRequest
	if err := twiml.Bind(&cr, r); err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	log.Info().
		Str("call_status", cr.CallStatus).
		Str("voice_request", string(jsonutil.MustMarshal(cr, true))).
		Str("stage", "parsed_request").Msg("HandleCall")

	// Create a new response container
	res := twiml.NewResponse()

	switch status := cr.CallStatus; status {

	// Call is already in progress, tell Twilio to continue
	case twiml.InProgress:
		w.WriteHeader(200)
	// Call is ringing but has not been connected yet, tell Twilio to continue
	case twiml.Ringing:
		callLength := 10
		log.Info().Str("twilio.status", status).Msg("C1_CallInProgress")
		res.Add(
			&twilio.Start{
				Children: []twiml.Markup{
					&twilio.Stream{
						URL: "wss://ringforce.ngrok.io/media"}}},
			&twiml.Say{
				Language: "en",
				Text:     fmt.Sprintf("I will stream the next %d seconds of audio through your websocket", callLength)},
			&twiml.Pause{
				Length: callLength},
		)
		twilio.TwimlResponseProcess(w, r, res)
	case twiml.Queued:
		w.WriteHeader(200)
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
		w.Header().Set(
			httputilmore.HeaderContentType,
			httputilmore.ContentTypeAppXMLUtf8)
		w.WriteHeader(200)
	}
}
