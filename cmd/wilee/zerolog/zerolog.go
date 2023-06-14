package zerolog

import (
	"os"

	"github.com/rs/zerolog"
)

var zlog = zerolog.New(zerolog.ConsoleWriter{
	Out: os.Stderr,
})
