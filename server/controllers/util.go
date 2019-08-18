package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/BTBurke/twiml"
	hum "github.com/grokify/gotilla/net/httputilmore"
	log "github.com/sirupsen/logrus"
)

const (
	TwilioApiCallsJsonURLFormat = `https://api.twilio.com/2010-04-01/Accounts/%s/Calls.json`
)

func processResponse(w http.ResponseWriter, r *http.Request, res *twiml.Response) {
	// Validate and encode the response.  Validation is done
	// automatically before the response is encoded.
	b, err := res.Encode()
	if err != nil {
		log.Warn("C3_InProgress_Error_Encode")
		http.Error(w, http.StatusText(502), 502)
		return
	}

	// Write the XML response to the http.ReponseWriter
	if _, err := w.Write(b); err != nil {
		log.Warn("C3_InProgress_Error_Write")
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

func BuildTwilioCallURL(accountSid string) string {
	return fmt.Sprintf(TwilioApiCallsJsonURLFormat, accountSid)
}

type GatherEvent struct {
	msg           string
	AccountSid    string
	ApiVersion    string
	Called        string
	CalledCity    string
	CalledCountry string
	Caller        string
	CallerCity    string
	CallerCountry string
	CallerState   string
	CallerZip     string
	CallStatus    string
	Digits        string
	FinishedOnKey string
	From          string
	FromCity      string
	FromState     string
	FromZip       string
	ToCity        string
	ToCountry     string
	ToState       string
	ToZip         string
}

/*
msg=Gather+End&
Called=%2B17752204114&
Digits=1&
ToState=NV&CallerCountry=US&Direction=outbound-api&CallerState=WI&ToZip=89701&CallSid=CA98dfba576b41a4650469b78592e24e86&To=%2B17752204114&CallerZip=&ToCountry=US&FinishedOnKey=&ApiVersion=2010-04-01&CalledZip=89701&CalledCity=CARSON+CITY&CallStatus=in-progress&From=%2B14146221701&AccountSid=AC791ebe2a4510eb2254d59a9645afe474&CalledCountry=US&CallerCity=&Caller=%2B14146221701&FromCountry=US&ToCity=CARSON+CITY&FromCity=&CalledState=NV&FromZip=&FromState=WI
*/
