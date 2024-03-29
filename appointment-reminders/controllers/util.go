package controllers

import (
	"net/http"
	"strconv"

	"github.com/grokify/mogo/net/http/httputilmore"
	"github.com/grokify/twiml"
	"github.com/rs/zerolog/log"
)

const AppRetryLimit int = 2

func addMainMenu(res *twiml.Response, numTries uint16) {
	res.Add(
		&twiml.Gather{
			Action:      "/reminder_process?numTries=" + strconv.Itoa(int(numTries)),
			FinishOnKey: "#",
			Input:       "speech dtmf",
			Timeout:     3,
			NumDigits:   1,
			Children: []twiml.Markup{
				&twiml.Say{
					Language: "en",
					Text:     "Please press 1 to confirm. 2 to reschedule. or 3 to cancel.",
				},
			},
		},
		&twiml.Say{
			Language: "en",
			Text:     "I didn't receive an answer.",
		},
		&twiml.Redirect{
			Method: http.MethodPost,
			URL:    "/reminder_menu?numTries=" + strconv.Itoa(int(numTries)+AppRetryLimit),
		},
	)
}

func processResponse(w http.ResponseWriter, r *http.Request, res *twiml.Response) {
	// Validate and encode the response.  Validation is done
	// automatically before the response is encoded.
	b, err := res.Encode()
	if err != nil {
		log.Warn().Msg("C3_InProgress_Error_Encode [502]")
		http.Error(w, http.StatusText(http.StatusBadGateway), http.StatusBadGateway)
		return
	}

	// Write the XML response to the http.ReponseWriter
	if _, err := w.Write(b); err != nil {
		log.Warn().Msg("C3_InProgress_Error_Write [502]")
		http.Error(w, http.StatusText(http.StatusBadGateway), http.StatusBadGateway)
		return
	}
	w.Header().Set(
		httputilmore.HeaderContentType,
		httputilmore.ContentTypeAppXMLUtf8)
	w.WriteHeader(200)
	log.Info().Msg("C3_Success")
}
