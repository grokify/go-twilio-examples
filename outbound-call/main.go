package main

import (
	"fmt"
	"log"
	"os"

	"github.com/grokify/go-twilio-examples/twilio"
	om "github.com/grokify/oauth2more"
	"github.com/grokify/simplego/config"
	"github.com/grokify/simplego/fmt/fmtutil"
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
	err := config.LoadDotEnvSkipEmpty("./.env", os.Getenv("ENVPATH"))
	if err != nil {
		log.Fatal(err)
	}

	var callOpts CallOptions
	err = envconfig.Process("twilio_demo", &callOpts)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmtutil.PrintJSON(callOpts)

	client, err := om.NewClientBasicAuth(callOpts.Sid, callOpts.Token, false)
	if err != nil {
		log.Fatal(err)
	}

	apiOpts := twilio.TwilioCallsOpts{
		To:          callOpts.To,
		From:        callOpts.From,
		CallbackURL: callOpts.CallbackURL}

	resp, err := twilio.MakeCall(client,
		twilio.BuildTwilioCallURL(callOpts.Sid),
		apiOpts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Called with status [%v]\n", resp.StatusCode)

	fmt.Println("DONE")
}
