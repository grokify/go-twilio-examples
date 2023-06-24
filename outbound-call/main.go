package main

import (
	"fmt"
	"os"

	twilio "github.com/grokify/go-twilio-examples"
	"github.com/grokify/goauth/authutil"
	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/log/logutil"
	"github.com/kelseyhightower/envconfig"
)

type CallOptions struct {
	Sid         string `short:"u" long:"username" description:"A site" required:"false"`
	Token       string `short:"p" long:"password" description:"A token" required:"false"`
	To          string `short:"t" long:"to" description:"A phone number" required:"false"`
	From        string `short:"f" long:"from" description:"A phone number" required:"false"`
	Log         string `short:"l" long:"log" description:"A phone number" required:"false"`
	CallbackURL string `short:"c" long:"callbackurl" description:"Twiml callback URL" required:"false"`
}

func main() {
	_, err := config.LoadDotEnv([]string{"./.env", os.Getenv("ENVPATH")}, 1)
	logutil.FatalErr(err)

	var callOpts CallOptions
	err = envconfig.Process("twilio_demo", &callOpts)
	logutil.FatalErr(err)

	fmtutil.MustPrintJSON(callOpts)

	client, err := authutil.NewClientBasicAuth(callOpts.Sid, callOpts.Token, false)
	logutil.FatalErr(err)

	apiOpts := twilio.TwilioCallsOpts{
		To:          callOpts.To,
		From:        callOpts.From,
		CallbackURL: callOpts.CallbackURL}

	resp, err := twilio.MakeCall(client,
		twilio.BuildTwilioCallURL(callOpts.Sid),
		apiOpts)
	logutil.FatalErr(err)

	fmt.Printf("Called with status [%v]\n", resp.StatusCode)

	fmt.Println("DONE")
}
