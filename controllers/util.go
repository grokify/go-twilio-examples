package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/BTBurke/twiml"
	hum "github.com/grokify/gotilla/net/httputilmore"
	log "github.com/sirupsen/logrus"
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
		log.Warn("C3_InProgress_Error_Encode [502]")
		http.Error(w, http.StatusText(502), 502)
		return
	}

	// Write the XML response to the http.ReponseWriter
	if _, err := w.Write(b); err != nil {
		log.Warn("C3_InProgress_Error_Write [502]")
		http.Error(w, http.StatusText(502), 502)
		return
	}
	w.Header().Set(hum.HeaderContentType, hum.ContentTypeAppXmlUtf8)
	w.WriteHeader(200)
	log.Info("C3_Success")
	return
}

func getTimestampNow() string {
	dt := time.Now()
	return dt.Format(time.RFC3339)
}
