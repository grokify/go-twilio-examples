# Twilio Examples in Go

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Docs][docs-mkdoc-svg]][docs-mkdoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/grokify/go-twilio-examples/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/grokify/go-twilio-examples/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/grokify/go-twilio-examples/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/grokify/go-twilio-examples/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/grokify/go-twilio-examples/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/grokify/go-twilio-examples/actions/workflows/go-sast-codeql.yaml
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/go-twilio-examples
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/go-twilio-examples
 [docs-mkdoc-svg]: https://img.shields.io/badge/Go-dev%20guide-blue.svg
 [docs-mkdoc-url]: https://grokify.github.io/go-twilio-examples
 [viz-svg]: https://img.shields.io/badge/visualizaton-Go-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=grokify%2Fgo-twilio-examples
 [loc-svg]: https://tokei.rs/b1/github/grokify/go-twilio-examples
 [repo-url]: https://github.com/grokify/go-twilio-examples
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/go-twilio-examples/blob/main/LICENSE

Example code for building Twilio voice applications in Go: TwiML-driven call flows, outbound calls via the REST API, and real-time audio via Media Streams.

## Examples

Full guides are in [`docs/`](docs/index.md) (MkDocs site).

| Directory | Description |
|-----------|-------------|
| [`appointment-reminders`](docs/guides/appointment-reminders.md) | TwiML server that speaks a reminder and handles keypad input. |
| [`outbound-call`](docs/guides/outbound-call.md) | CLI that places a call via the Twilio Calls API. |
| [`media-streams`](docs/guides/media-streams.md) | Receives live call audio over a WebSocket and saves `.ulaw`/`.wav`. |

The root package holds shared helpers (Calls API, TwiML verbs, stream message types); see the [docs home](docs/index.md).

## Quick Start

Your server must be reachable by Twilio, so use a tunnel such as [ngrok](https://ngrok.com/):

```bash
# Terminal 1: run the appointment reminder server
go run ./appointment-reminders

# Terminal 2: expose it
ngrok http 8081

# Terminal 3: place a call (set TWILIO_DEMO_CALLBACKURL to https://<ngrok-host>/reminder_start)
cd outbound-call && cp .env.sample .env  # then edit .env
go run .
```

## Notes

- `media-streams` requires `BASE_URL_WSS` (e.g. `wss://<your-host>.ngrok.io`).
- Requires Go 1.26 or later.
- Never commit your `.env` file; it holds your Twilio credentials.

## License

[MIT](LICENSE)
