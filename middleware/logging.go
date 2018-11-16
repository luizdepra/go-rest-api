package middleware

import (
	"bufio"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/go-playground/pure"
)

type logWriter struct {
	http.ResponseWriter
	status    int
	size      int64
	committed bool
}

// WriteHeader writes HTTP status code.
func (lw *logWriter) WriteHeader(status int) {

	if lw.committed {
		log.Println("response already committed")
		return
	}

	lw.status = status
	lw.ResponseWriter.WriteHeader(status)
	lw.committed = true
}

// Write writes the data to the connection as part of an HTTP reply.
func (lw *logWriter) Write(b []byte) (int, error) {
	lw.size += int64(len(b))
	return lw.ResponseWriter.Write(b)
}

// Status returns the current response's http status code.
func (lw *logWriter) Status() int {
	return lw.status
}

// Size returns the number of bytes written in the response thus far
func (lw *logWriter) Size() int64 {
	return lw.size
}

// Hijack hijacks the current http connection
func (lw *logWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return lw.ResponseWriter.(http.Hijacker).Hijack()
}

// CloseNotify ...
func (lw *logWriter) CloseNotify() <-chan bool {
	return lw.ResponseWriter.(http.CloseNotifier).CloseNotify()
}

var lrpool = sync.Pool{
	New: func() interface{} {
		return new(logWriter)
	},
}

// Logging handle HTTP request logging
func Logging() pure.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			t1 := time.Now()

			lw := lrpool.Get().(*logWriter)
			lw.status = 200
			lw.size = 0
			lw.committed = false
			lw.ResponseWriter = w

			next(lw, r)

			log.Printf("%d [%s] %q %v %d\n", lw.Status(), r.Method, r.URL, time.Since(t1), lw.Size())
		}
	}
}
