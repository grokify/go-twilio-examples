# Appointment Reminders

A TwiML demo server. When Twilio requests `/reminder_start`, the server speaks an appointment reminder and offers a keypad menu (`<Gather>`); the chosen digit is handled at `/reminder_process`.

## Endpoints

| Path | Purpose |
|------|---------|
| `/reminder_start` | Answers the call and speaks the reminder |
| `/reminder_menu` | Plays the keypad menu |
| `/reminder_process` | Handles the digit pressed (1, 2 or 3) |

## Run

No configuration is required. From the repository root:

```bash
go run ./appointment-reminders
```

The server listens on port `8081`. Override with `PORT`:

```bash
PORT=8080 go run ./appointment-reminders
```

## Expose the server

```bash
ngrok http 8081
```

## Place a demo call

Use the [Outbound Call](outbound-call.md) example, setting `TWILIO_DEMO_CALLBACKURL` to `https://<ngrok-host>/reminder_start`.
