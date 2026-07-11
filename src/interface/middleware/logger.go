package middleware


import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
	"encoding/json"
)

type responseWriterWrapper struct {
	rw http.ResponseWriter
	multiWriter io.Writer
	statusCode int
}

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request){
		requestLogger(req)
		buffer := &bytes.Buffer{}
		rww := newResponseWriterWrapper(rw, buffer)
		handler.ServeHTTP(rww, req)
		responseLogger(rww, buffer)
	})
}

func requestLogger(req *http.Request) {
	requestBody, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Error reading request body: %v\n", err)
		return
	}

	defer req.Body.Close()

	var formattedBody bytes.Buffer

	json.Indent(&formattedBody, requestBody, "", "  ")
	fmt.Printf(
		"==========================================================================================================\n"+
		"[REQUEST]\n"+
		"Timestamp\t: %s\n"+
		"Method\t:%s\n"+
		"Path\t:%s\n"+
		"Authorization\t:%s\n"+
		"Body\t:%s"+
		"\n++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n",
		time.Now().Format(time.RFC3339),
		req.Method,
		req.URL.Path,
		req.Header.Get("Authorization"),
		formattedBody.String(),
	)
	req.Body = io.NopCloser(bytes.NewBuffer(requestBody))
}

func responseLogger(rww *responseWriterWrapper, buffer *bytes.Buffer){
	var formattedResponseBody bytes.Buffer
	json.Indent(&formattedResponseBody, buffer.Bytes(), "", "  ")

	status := rww.statusCode
	fmt.Printf(
		"[RESPONSE]\n"+
		"Timestamp\t: %s\n"+
		"Status\t:%d\n"+
		"Body\t:%s"+
		"\n==========================================================================================================\n",
		time.Now().Format(time.RFC3339),
		status,
		formattedResponseBody.String(),
	)
} 

func newResponseWriterWrapper(rw http.ResponseWriter, buffer *bytes.Buffer) *responseWriterWrapper {
	return &responseWriterWrapper{
		rw: rw,
		multiWriter: io.MultiWriter(rw, buffer),
	}
}

func (rww *responseWriterWrapper) Header() http.Header {
	return rww.rw.Header()
}

func (rww *responseWriterWrapper) Write(data []byte) (int, error) {
	return rww.multiWriter.Write(data)
}

func (rww *responseWriterWrapper) WriteHeader(statusCode int){
	rww.statusCode = statusCode
	rww.rw.WriteHeader(statusCode)
}
