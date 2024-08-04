package main

import (
	"gotik/internal/provider/server"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Zerolog Time Formating > "Mon, 02 Jan 2006 15:04:05 MST"
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:          os.Stdout,
		TimeFormat:   time.RFC1123,
		TimeLocation: time.Local,
	})

	server.StartHTTPServer()
}
