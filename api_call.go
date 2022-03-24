package twilio

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/grokify/mogo/net/httputilmore"
)

const (
	TwilioApiCallsJsonURLFormat = `https://api.twilio.com/2010-04-01/Accounts/%s/Calls.json`
)

func BuildTwilioCallURL(accountSid string) string {
	return fmt.Sprintf(TwilioApiCallsJsonURLFormat, accountSid)
}

type TwilioCallsOpts struct {
	To          string `url:"To"`
	From        string `url:"From"`
	CallbackURL string `url:"Url"`
}

func (opts *TwilioCallsOpts) MustString() string {
	v, err := query.Values(opts)
	if err != nil {
		panic(err)
	}
	return v.Encode()
}

func (opts *TwilioCallsOpts) StringsReader() *strings.Reader {
	return strings.NewReader(opts.MustString())
}

func MakeCall(client *http.Client, apiUrl string, opts TwilioCallsOpts) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, apiUrl, opts.StringsReader())
	if err != nil {
		return nil, err
	}
	req.Header.Add(
		httputilmore.HeaderContentType,
		httputilmore.ContentTypeAppFormURLEncoded)
	return client.Do(req)
}
