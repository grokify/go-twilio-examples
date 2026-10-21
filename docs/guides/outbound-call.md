# Outbound Call

A CLI that places a call through the Twilio Calls API and points Twilio at a TwiML callback URL, such as the [Appointment Reminders](appointment-reminders.md) server.

## Configure

Copy the sample file and fill in your values. `.env` is gitignored; never commit it.

```bash
cd outbound-call
cp .env.sample .env
```

| Variable | Description |
|----------|-------------|
| `TWILIO_DEMO_SID` | Twilio Account SID |
| `TWILIO_DEMO_TOKEN` | Twilio auth token |
| `TWILIO_DEMO_TO` | Number to call |
| `TWILIO_DEMO_FROM` | Your Twilio number |
| `TWILIO_DEMO_LOG` | Number for logging |
| `TWILIO_DEMO_CALLBACKURL` | TwiML URL, e.g. `https://12345678.ngrok.io/reminder_start` |

An alternate `.env` location can be given with `ENVPATH`.

## Run

```bash
cd outbound-call
go run .
```

The program prints the loaded options as JSON, then the API status, for example `Called with status [201]`. The `TWILIO_DEMO_TO` number receives the call.
