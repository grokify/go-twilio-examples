package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/grokify/go-appointment-reminder-demo/demo"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	om "github.com/grokify/oauth2more"
	"github.com/jessevdk/go-flags"
)

const DefaultPort string = "8081"

func call(opts demo.Options) {
	client, err := om.NewClientBasicAuth(opts.Username, opts.Password, false)
	if err != nil {
		log.Fatal(err)
	}

	apiOpts := demo.TwilioCallsOpts{
		To:          opts.To,
		From:        opts.From,
		CallbackURL: opts.CallbackURL}

	resp, err := demo.MakeCall(client, demo.BuildTwilioCallURL(opts.Username), apiOpts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("STATUS [%v]\n", resp.StatusCode)
}

func serve(opts demo.Options) {
	http.HandleFunc("/twiml", demo.CallRequest())

	portStr := ":" + DefaultPort
	if len(os.Getenv("PORT")) > 0 {
		portStr = ":" + os.Getenv("PORT")
	}
	fmt.Printf("Running on [%v]\n", portStr)
	http.ListenAndServe(portStr, nil)
}

func main() {
	err := config.LoadDotEnvSkipEmpty("./.env", os.Getenv("ENVPATH"))
	if err != nil {
		log.Fatal(err)
	}

	opts := demo.Options{}
	_, err = flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	if len(opts.Username) == 0 {
		opts.Username = os.Getenv("TWILIO_SID")
	}
	if len(opts.Password) == 0 {
		opts.Password = os.Getenv("TWILIO_TOKEN")
	}
	if len(opts.To) == 0 {
		opts.To = os.Getenv("TWILIO_DEMO_TO")
	}
	if len(opts.From) == 0 {
		opts.From = os.Getenv("TWILIO_DEMO_FROM")
	}
	if len(opts.Log) == 0 {
		opts.Log = os.Getenv("TWILIO_DEMO_LOG")
	}
	if len(opts.CallbackURL) == 0 {
		opts.CallbackURL = os.Getenv("TWILIO_CALLBACK_URL")
	}
	fmtutil.PrintJSON(opts)

	switch opts.Action {
	case "call":
		call(opts)
	case "serve":
		serve(opts)
	default:
		fmt.Println("E_ACTION_NOT_FOUND (call|serve)")
	}

	fmt.Println("DONE")
}
