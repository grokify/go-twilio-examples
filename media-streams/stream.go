package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/grokify/simplego/audio/ulaw"
	"github.com/grokify/simplego/time/timeutil"
	"github.com/rs/zerolog/log"

	"github.com/grokify/go-twilio-examples/media-streams/utility"
	"github.com/grokify/go-twilio-examples/twilio"
)

// HandleMediaStream will upgrade connection to websocket and save the audio to file
func HandleMediaStream(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Begin_Handle_Media_Stream")
	upgrader := websocket.Upgrader{}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warn().Err(err).Msg("upgrade_to_wss")
		return
	}

	defer utility.SafeClose(c)

	inboundBytes := []byte{}

	streamSid := ""

	loop := true
	for loop == true {
		_, messageBytes, err := c.ReadMessage()
		utility.PanicIfErr(err)
		log.Info().Str("body", string(messageBytes)).
			Msg("inbound-wss-message")

		msg := twilio.StreamMessage{}
		err = json.Unmarshal(messageBytes, &msg)
		utility.PanicIfErr(err)

		switch msg.Event {
		case twilio.EventConnected:
			log.Debug().Str("protocol", msg.Protocol).Str("version", msg.Version).
				Msg("stream.event.connect received")
		case twilio.EventStart:
			streamSid = msg.StreamSid
			log.Debug().Str("message", fmt.Sprintf("%#v", msg.Start)).
				Msg("stream.event.start received")
		case twilio.EventMedia:
			if msg.Media.Track == "inbound" {
				chunk, err := base64.StdEncoding.DecodeString(msg.Media.Payload)
				utility.PanicIfErr(err)
				inboundBytes = append(inboundBytes, chunk...)
			}
		case twilio.EventStop:
			utility.DebugLogf("Ending audio stream: %#v", msg.Stop)
			loop = false
		default:
			utility.LogWarningf("Unrecognized event type: %s", msg.Event)
			loop = false
		}
	}
	WriteFiles(streamSid, inboundBytes)
}

func WriteFiles(streamSid string, inboundBytes []byte) {
	if len(streamSid) > 0 && len(inboundBytes) > 0 {
		filebase := "media_" + strconv.Itoa(int(timeutil.Dt14Now())) + "_" + streamSid
		err := ioutil.WriteFile(filebase+".ulaw", inboundBytes, 0644)
		utility.PanicIfErr(err)
		err = ulaw.WriteFileWavFromUlaw(filebase+".wav", inboundBytes)
		utility.PanicIfErr(err)
		log.Info().
			Str("wav", filebase+".wav").
			Str("ulaw", filebase+".ulaw").
			Msg("wrote_files")
	}
}
