package twilio

const (
	EventConnected = "connected"
	EventStart     = "start"
	EventMedia     = "media"
	EventStop      = "stop"
)

// StreamMessage is defined in
// https://www.twilio.com/docs/voice/twiml/stream
type StreamMessage struct {
	Event          string             `json:"event"`
	Media          Media              `json:"media"`
	Protocol       string             `json:"protocol"`
	SequenceNumber string             `json:"sequenceNumber"`
	Start          StreamMessageStart `json:"start"`
	StreamSid      string             `json:"streamSid"`
	Stop           interface{}        `json:"stop"`
	Version        string             `json:"version"`
}

// Media event
type Media struct {
	Track     string `json:"track"`
	Chunk     string `json:"chunk"`
	Timestamp string `json:"timestamp"`
	Payload   string `json:"payload"`
}

type StreamMessageStart struct {
	AccountSid       string            `json:"accountSid"`
	CallSid          string            `json:"callSid"`
	StreamSid        string            `json:"streamSid"`
	Tracks           []string          `json:"tracks"`
	CustomParameters map[string]string `json:"customParameters"`
}

/*
{
 "event": "start",
 "sequenceNumber": "2",
 "start": {
   "streamSid": "MZ18ad3ab5a668481ce02b83e7395059f0",
   "accountSid": "AC123",
   "callSid": "CA123",
   "tracks": [
     "inbound",
     "outbound"
   ],
   "customParameters": {
     "FirstName": "Jane",
     "LastName": "Doe",
     "RemoteParty": "Bob",
   },
   "mediaFormat": {
     "encoding": "audio/x-mulaw",
     "sampleRate": 8000,
     "channels": 1
   }
 },
"streamSid": "MZ18ad3ab5a668481ce02b83e7395059f0"
}
------
{
	"event":"media",
	"sequenceNumber":"424",
	"media":{
		"track":"inbound",
		"chunk":"423",
		"timestamp":"8585",
		"payload":"deadbeef1==="
	},
	"streamSid":"MZ92fd0a8252929adb108b724f970eaebb"
}",
*/
