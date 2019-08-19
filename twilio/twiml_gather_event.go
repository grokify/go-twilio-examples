package twilio

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
