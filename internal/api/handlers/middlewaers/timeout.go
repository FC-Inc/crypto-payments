package middlewaers

import (
	"bytes"
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type timeoutWriter struct {
	gin.ResponseWriter
	h           http.Header
	wbuf        bytes.Buffer
	m           sync.Mutex
	timedOut    bool
	wroteHeader bool
	code        int
}

func (tw *timeoutWriter) Header() http.Header { return tw.h }

func (tw *timeoutWriter) Write(b []byte) (int, error) {
	tw.m.Lock()
	defer tw.m.Unlock()
	if tw.timedOut {
		return 0, nil
	}
	return tw.wbuf.Write(b)
}

func (tw *timeoutWriter) WriteHeader(code int) {
	tw.m.Lock()
	defer tw.m.Unlock()
	if tw.timedOut || tw.wroteHeader {
		return
	}
	tw.writeHeader(code)
}

func (tw *timeoutWriter) writeHeader(code int) {
	tw.wroteHeader = true
	tw.code = code
}

func Timeout(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()
		tw := &timeoutWriter{
			ResponseWriter: c.Writer,
			h:              make(http.Header),
		}
		c.Writer = tw
		c.Request = c.Request.WithContext(ctx)
		executed := make(chan struct{})
		go func() {
			c.Next()
			close(executed)
		}()
		select {
		case <-executed:
			tw.m.Lock()
			defer tw.m.Unlock()
			dst := tw.ResponseWriter.Header()
			for k, vv := range tw.Header() {
				dst[k] = vv
			}
			tw.ResponseWriter.WriteHeader(tw.code)
			tw.ResponseWriter.Write(tw.wbuf.Bytes())
		case <-ctx.Done():
			tw.m.Lock()
			defer tw.m.Unlock()
			tw.ResponseWriter.WriteHeader(http.StatusGatewayTimeout)
			c.Abort()
			tw.timedOut = true
		}
	}
}
