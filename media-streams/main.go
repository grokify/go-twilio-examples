package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/buaazp/fasthttprouter"
	"github.com/grokify/simplego/net/http/httpsimple"
	"github.com/grokify/simplego/net/urlutil"
	"github.com/grokify/simplego/strconv/strconvutil"
	"github.com/rs/zerolog/log"
)

// HandleHome
func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Example Receive Streaming Media Server\nhttps://github.com/grokify\n")
}

type Service struct {
	BaseUrlWss string
	Port       int
	HTTPEngine string
}

func (svc Service) PortInt() int                       { return svc.Port }
func (svc Service) HttpEngine() string                 { return svc.HTTPEngine }
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
	return urlutil.Join(svc.BaseUrlWss, "media")
}

func main() {
	svc := Service{
		BaseUrlWss: "wss://ringforce.ngrok.io", // ngrok http -subdomain=ringforce 8080
		Port:       strconvutil.AtoiOrDefault(os.Getenv("PORT"), 8080),
		HTTPEngine: ValueOrDefault(os.Getenv("HTTP_ENGINE"), "nethttp"),
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
