package middlewares

import (
	"alter-io-go/helpers/logger"
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

type bodyDumpResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

type Skipper func(c *gin.Context) bool

func SlogLoggerWithSkipper(skipper Skipper) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Capture request body
		reqBody := captureRequestBody(c)
		maskedReqBody := maskSensitiveFields(reqBody)

		// Capture response body
		resBody := new(bytes.Buffer)
		writer := &bodyDumpResponseWriter{body: resBody, ResponseWriter: c.Writer}
		c.Writer = writer

		// Process the request
		c.Next()

		// Skip logging if the skipper returns true
		if skipper != nil && skipper(c) {
			return
		}

		// Capture and mask response body
		resBodyMap := parseJSON(resBody.Bytes())
		maskedResBody := maskSensitiveFields(resBodyMap)

		// Log the request and response
		logRequestResponse(c, start, maskedReqBody, maskedResBody)
	}
}

func maskSensitiveFields(data map[string]interface{}) map[string]interface{} {
	if data == nil {
		return nil
	}

	// Assuming Document is a struct that has a method throughMap
	doc := &Document{}
	return doc.ProcessMap(data)
}

func captureRequestBody(c *gin.Context) map[string]interface{} {
	var body []byte
	if c.Request.Body != nil {
		body, _ = io.ReadAll(c.Request.Body)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	contentType := c.GetHeader("Content-Type")
	if contentType == "application/json" && len(body) > 0 {
		return parseJSON(body)
	}
	return nil
}

func parseJSON(data []byte) map[string]interface{} {
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil
	}
	return result
}

func logRequestResponse(c *gin.Context, start time.Time, reqBody, resBody map[string]interface{}) {
	status := c.Writer.Status()
	latency := time.Since(start)
	req := c.Request

	logEvent := []slog.Attr{
		slog.Int("status", status),
		slog.String("latency", latency.String()),
		slog.String("method", req.Method),
		slog.String("uri", req.RequestURI),
		slog.String("host", req.Host),
		slog.String("remote_ip", c.ClientIP()),
		slog.Any("headers", req.Header),
		slog.Any("request_body", reqBody),
		slog.Any("response_body", resBody),
	}

	if id := c.GetHeader("X-Request-ID"); id != "" {
		logEvent = append(logEvent, slog.String("request_id", id))
	}

	// Convert []slog.Attr to []any
	logAttrs := make([]any, 0, len(logEvent)*2) // Each slog.Attr becomes a key-value pair
	for _, attr := range logEvent {
		logAttrs = append(logAttrs, attr.Key, attr.Value)
	}

	switch {
	case c.Writer.Status() >= 500:
		logger.Get().Error("Server Error", slog.Group("details", logAttrs...))
	case c.Writer.Status() >= 400:
		logger.Get().Warn("Client Error", slog.Group("details", logAttrs...))
	case c.Writer.Status() >= 300:
		logger.Get().Info("Redirection", slog.Group("details", logAttrs...))
	default:
		logger.Get().Info("Success", slog.Group("details", logAttrs...))
	}
}
