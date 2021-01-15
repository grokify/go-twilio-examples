package twilio

const (
	EventConnected = "connected"
	EventStart     = "start"
	EventMedia     = "media"
	EventStop      = "stop"
	EventMark      = "mark"
)

// StreamMessage is defined in
// https://www.twilio.com/docs/voice/twiml/stream
type StreamMessage struct {
	Event          string             `json:"event"`
	Mark           StreamMessageMark  `json:"mark,omitempty"`
	Media          StreamMessageMedia `json:"media,omitempty"`
	Protocol       string             `json:"protocol,omitempty"`
	SequenceNumber string             `json:"sequenceNumber"`
	Start          StreamMessageStart `json:"start,omitempty"`
	StreamSid      string             `json:"streamSid"`
	Stop           StreamMessageStop  `json:"stop,omitempty"`
	Version        string             `json:"version,omitempty"`
}

// StreamMessageStart is a property of a start message. This message contains
// important metadata about the stream and is sent immediately after
// the Connected message. It is only sent once at the start of the Stream.
type StreamMessageStart struct {
	AccountSid       string            `json:"accountSid"`
	CallSid          string            `json:"callSid"`
	StreamSid        string            `json:"streamSid"`
	Tracks           []string          `json:"tracks"`
	CustomParameters map[string]string `json:"customParameters"`
}

// StreamMessageMedia is a property of a media message. This message
// type encapsulates the raw audio data.
type StreamMessageMedia struct {
	Track     string `json:"track"`
	Chunk     string `json:"chunk"`
	Timestamp string `json:"timestamp"`
	Payload   string `json:"payload"`
}

// StreamMessageStop is a property of a stop message. A stop message
// will be sent when the Stream is either `<Stop>`ped or the Call has ended.
type StreamMessageStop struct {
	AccountSid string `json:"accountSid"`
	CallSid    string `json:"callSid"`
}

// StreamMessageMark is the property of a `mark` message. The `mark` event
// is sent only during bi-directional streaming by using the `<Connect>`
// verb. It is used to track, or label, when media has completed.
type StreamMessageMark struct {
	Name string `json:"name"`
}

/*

Example Start Message

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
*/
