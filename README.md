# Appointment Reminder Demo

This is a simple remidner demo app using Twiml along with a call initiation demo example.

# Installation & Usage

You can simply set up an run the server. No configuration is required to run this server.

```
$ go get github.com/grokify/twilio-appointment-reminder-demo
$ cd twilio-appointment-reminder-demo
$ go run main.go
```

This will use a default server port, 8081. Optionally set one in the environment:

```
$ PORT=8080 go run main.go
```

## Ngrok

Your server must be available online. An easy way to set this up is to use ngrok tunneling. For example:

```
$ ngrok http 8081
```

# Demo Phone Call

To set up a demo a phone call, create an `examples/call/.env` file using the [`examples/call/.env.sample`](examples/call/.env.sample) file and set all the parameters for your demo.

Of note, configure the `TWILIO_DEMO_CALLBACKURL` to be your hostname plus the call start endpoint `reminder_start`. For example:

```
TWILIO_DEMO_CALLBACKURL=https://12345678.ngrok.io/reminder_start
```