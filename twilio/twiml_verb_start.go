package twilio

import (
	"encoding/xml"

	"github.com/grokify/twiml"
)

// Start TwiML
type Start struct {
	XMLName  xml.Name       `xml:"Start"`
	Children []twiml.Markup `xml:",omitempty"`
}

// Validate returns an error if the TwiML is constructed improperly
func (s *Start) Validate() error { return nil }

// Type returns the XML name of the verb
func (s *Start) Type() string { return "Start" }

// Stream TwiML
type Stream struct {
	XMLName xml.Name `xml:"Stream"`
	URL     string   `xml:"url,attr,omitempty"`
}

// Validate returns an error if the TwiML is constructed improperly
func (s *Stream) Validate() error { return nil }

// Type returns the XML name of the verb
func (s *Stream) Type() string { return "Stream" }
