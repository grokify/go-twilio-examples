# Twilio Examples in Go

Example code for building Twilio voice applications in Go: TwiML-driven call flows, outbound calls via the REST API, and real-time audio via Media Streams.

| Guide | Description |
|-------|-------------|
| [Appointment Reminders](guides/appointment-reminders.md) | TwiML server that speaks a reminder and handles keypad input with `<Gather>`. |
| [Outbound Call](guides/outbound-call.md) | CLI that places a call through the Twilio Calls API. |
| [Media Streams](guides/media-streams.md) | Server that receives live call audio over a WebSocket and saves it as `.ulaw`/`.wav`. |

## Shared Package

The root package (`github.com/grokify/go-twilio-examples`) provides helpers used by the examples:

- `api_call.go`: Calls API URL builder and `MakeCall`
- `twiml_response.go`: `TwimlResponseProcess` to encode and write a TwiML response
- `twiml_verb_start.go`: `<Start>` and `<Stream>` TwiML verbs, built on [`grokify/twiml`](https://github.com/grokify/twiml)
- `twiml_event_gather.go`: `GatherEvent` webhook parameters
- `stream_message.go`: `StreamMessage` types for Media Streams WebSocket messages

## Prerequisites

- Go 1.26 or later
- A Twilio account (Account SID, auth token, and a voice-capable number)
- A public tunnel such as [ngrok](https://ngrok.com/), since Twilio must reach your server
