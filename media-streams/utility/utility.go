package utility

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

func PanicIfErr(err error) {
	if err != nil {
		log.Panic().Err(err)
		panic(err)
	}
}

func SafeClose(c *websocket.Conn) {
	err := c.Close()
	if err != nil {
		log.Warn().Err(err)
	}
}

func DebugLogf(s string, vars ...interface{}) {
	log.Debug().Msg(fmt.Sprintf(s, vars...))
}

func LogWarningf(s string, vars ...interface{}) {
	log.Warn().Msg(fmt.Sprintf(s, vars...))
}
