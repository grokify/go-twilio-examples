# Media Streams

A server that answers a call with TwiML `<Start><Stream>` and receives the call audio over a WebSocket. Inbound audio is saved to `media_<timestamp>_<streamSid>.ulaw` and `.wav` files in the working directory.

## Configure

| Variable | Required | Description |
|----------|----------|-------------|
| `BASE_URL_WSS` | yes | Public WebSocket base URL, e.g. `wss://<your-host>.ngrok.io` |
| `PORT` | no | Listen port, default `8080` |
| `HTTP_ENGINE` | no | HTTP engine, default `nethttp` |

## Run

```bash
ngrok http 8080
BASE_URL_WSS=wss://<your-host>.ngrok.io go run ./media-streams
```

Point your Twilio number's voice webhook at `https://<your-host>.ngrok.io/voice`.

## Audio format

Twilio streams 8-bit mu-law mono at 8000 Hz (`audio/x-mulaw`). See the [`<Stream>` reference](https://www.twilio.com/docs/voice/twiml/stream).

## References

- [Twilio `<Stream>` TwiML](https://www.twilio.com/docs/voice/twiml/stream)
- [Best Practices for Audio Recordings](https://support.twilio.com/hc/en-us/articles/223180588-Best-Practices-for-Audio-Recordings)
- [Live Transcribing Phone Calls using Twilio Media Streams and Google Speech-to-Text](https://www.twilio.com/blog/live-transcribing-phone-calls-using-twilio-media-streams-and-google-speech-text)
- [go-audio/wav](https://github.com/go-audio/wav)
- [twilio/media-streams](https://github.com/twilio/media-streams)
