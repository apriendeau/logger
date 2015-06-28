package logger

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
)

type Wrap struct {
	M http.Handler
}

type loggedResponse struct {
	http.ResponseWriter
	status int
}

func (l *loggedResponse) WriteHeader(status int) {
	l.status = status
	l.ResponseWriter.WriteHeader(status)
}

func (h *Wrap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	lw := &loggedResponse{ResponseWriter: w}
	h.M.ServeHTTP(lw, r)

	t2 := time.Now()
	msg, _ := fmt.Printf("[%d] [%s] %q %v\n", lw.status, r.Method, r.URL.String(), t2.Sub(t1))
	log.WithFields(log.Fields{
		"status":       lw.status,
		"method":       r.Method,
		"url":          r.URL.String(),
		"responseTime": t2.Sub(t1),
	}).Info(msg)
}
