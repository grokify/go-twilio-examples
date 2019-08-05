package demo

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

type Options struct {
	Action      string `short:"a" long:"action" description:"A site" required:"true"`
	Username    string `short:"u" long:"username" description:"A site" required:"false"`
	Password    string `short:"p" long:"password" description:"A token" required:"false"`
	To          string `short:"t" long:"to" description:"An object" required:"false"`
	From        string `short:"f" long:"from" description:"An object" required:"false"`
	Log         string `short:"l" long:"log" description:"An action (create|update|delete)" required:"false"`
	CallbackURL string `short:"c" long:"callbackurl" description:"An action (create|update|delete)" required:"false"`
}

/*func call(w http.ResponseWriter, r *http.Request) {
	// Let's set some initial default variables
	accountSid := "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	authToken := "YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Calls.json"

	// Build out the data for our message
	v := url.Values{}
	v.Set("To", "+155555555555")
	v.Set("From", "+15555551234")
	v.Set("Url", "[CHANGE_TO_YOUR_NGROK_URL]/twiml")
	rb := *strings.NewReader(v.Encode())

	// Create Client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

}
*/
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
	//	req.Header.Add("Accept", hum.ContentTypeAppJsonUtf8)
	req.Header.Add(hum.HeaderContentType, hum.ContentTypeAppFormUrlEncoded)
	return client.Do(req)
}
