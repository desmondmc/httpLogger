package httpLogger

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/apsdehal/go-logger"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)

		log, err := logger.New("httpLogger", 1, os.Stdout)
		if err != nil {
			panic(err) // Check for error
		}

		msg := fmt.Sprintf("New Request: %s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start))

		log.Notice(msg)
	})
}
