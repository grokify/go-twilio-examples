package twilio

import (
	"net/http"

	"github.com/grokify/mogo/net/httputilmore"
	"github.com/grokify/twiml"
	"github.com/rs/zerolog/log"
)

func TwimlResponseProcess(w http.ResponseWriter, r *http.Request, res *twiml.Response) {
	// Validate and encode the response.  Validation is done
	// automatically before the response is encoded.
	b, err := res.Encode()
	if err != nil {
		log.Warn().Err(err).Msg("twiml.response_encode_twiml.Response_error]")
		http.Error(w, http.StatusText(http.StatusBadGateway), http.StatusBadGateway)
		return
	}

	log.Info().Str("xml", string(b)).Msg("twiml_response")

	// Write the XML response to the http.ReponseWriter
	if _, err := w.Write(b); err != nil {
		log.Warn().Err(err).Msg("twiml.response_writeXml_error")
		http.Error(w, http.StatusText(http.StatusBadGateway), http.StatusBadGateway)
		return
	}
	w.Header().Set(
		httputilmore.HeaderContentType,
		httputilmore.ContentTypeAppXMLUtf8)
	w.WriteHeader(200)
	log.Info().Msg("twiml.response_success")
}
