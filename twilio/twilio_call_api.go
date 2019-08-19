package twilio

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	hum "github.com/grokify/gotilla/net/httputilmore"
)

const (
	TwilioApiCallsJsonURLFormat = `https://api.twilio.com/2010-04-01/Accounts/%s/Calls.json`
)

func BuildTwilioCallURL(accountSid string) string {
	return fmt.Sprintf(TwilioApiCallsJsonURLFormat, accountSid)
}

type TwilioCallsOpts struct {
	To          string
	From        string
	CallbackURL string
}

func (opts *TwilioCallsOpts) String() string {
	v := url.Values{}
	v.Set("To", opts.To)
	v.Set("From", opts.From)
	v.Set("Url", opts.CallbackURL)
	return v.Encode()
}

func (opts *TwilioCallsOpts) StringsReader() *strings.Reader {
	return strings.NewReader(opts.String())
}

func MakeCall(client *http.Client, apiUrl string, opts TwilioCallsOpts) (*http.Response, error) {
	rb := opts.StringsReader()
	fmt.Println(opts.String())
	req, err := http.NewRequest(http.MethodPost, apiUrl, rb)
	if err != nil {
		return nil, err
	}
	req.Header.Add(hum.HeaderContentType, hum.ContentTypeAppFormUrlEncoded)
	return client.Do(req)
}
