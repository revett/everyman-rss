package api

import (
	"log"
)

type combinedLogger struct{}

// Write implements the io.Writer interface. Instead of passing os.Stdout to
// handlers.CombinedLoggingHandler this type will be used, to make sure that log
// messages are formatted correctly.
func (c combinedLogger) Write(p []byte) (n int, err error) {
	log.Printf("INF %s", string(p))

	return len(p), nil
}
