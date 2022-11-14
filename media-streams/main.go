package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/buaazp/fasthttprouter"
	"github.com/grokify/gohttp/httpsimple"
	"github.com/grokify/mogo/net/urlutil"
	"github.com/grokify/mogo/strconv/strconvutil"
	"github.com/rs/zerolog/log"
)

// HandleHome
func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Example Receive Streaming Media Server\nhttps://github.com/grokify\n")
}

type Service struct {
	BaseUrlWSS    string
	Port          int
	HTTPEngineCfg string
}

func (svc Service) PortInt() int                       { return svc.Port }
func (svc Service) HTTPEngine() string                 { return svc.HTTPEngineCfg }
func (svc Service) RouterFast() *fasthttprouter.Router { return nil }
func (svc Service) Router() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", HandleHome)
	r.HandleFunc("/voice", svc.HandleCall)
	r.HandleFunc("/voice/", svc.HandleCall)
	r.HandleFunc("/media", HandleMediaStream)
	r.HandleFunc("/media/", HandleMediaStream)
	return r
}

func (svc Service) WssUrl() string {
	return urlutil.Join(svc.BaseUrlWSS, "media")
}

func main() {
	svc := Service{
		BaseUrlWSS:    "wss://ringforce.ngrok.io", // ngrok http -subdomain=ringforce 8080
		Port:          strconvutil.AtoiOrDefault(os.Getenv("PORT"), 8080),
		HTTPEngineCfg: ValueOrDefault(os.Getenv("HTTP_ENGINE"), "nethttp"),
	}
	log.Info().
		Int("port", svc.Port).
		Msg("service_starting")
	httpsimple.Serve(svc)
}

func ValueOrDefault(s, defaultString string) string {
	if len(s) > 0 {
		return s
	}
	return defaultString
}
