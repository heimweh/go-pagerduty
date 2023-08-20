package pagerduty

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestSecureLoggerHandleHeadersLogsContent(t *testing.T) {
	l := newSecureLogger()
	l.SetCanLog(true)
	headers := http.Header{
		"Authorization": []string{"Bearer secretApiKey"},
		"Content-Type":  []string{"application/json"},
	}
	l.handleHeadersLogsContent(headers)

	if !strings.Contains(l.headersContent, "<OBSCURED>iKey") {
		t.Errorf("Authorization header not properly obscured: got %s", l.headersContent)
	}
}

func TestSecureLoggerHandleBodyLogsContent_JSON(t *testing.T) {
	l := newSecureLogger()
	l.SetCanLog(true)
	body := io.NopCloser(bytes.NewReader([]byte(`{"key": "value"}`)))
	_ = l.handleBodyLogsContent(body)

	if !strings.Contains(l.bodyContent, ` "key": "value"`) {
		t.Errorf("JSON body not properly formatted: got %s", l.bodyContent)
	}
}

func TestSecureLoggerHandleBodyLogsContent_NonJSON(t *testing.T) {
	l := newSecureLogger()
	l.SetCanLog(true)
	body := io.NopCloser(bytes.NewReader([]byte(`non-json content`)))
	_ = l.handleBodyLogsContent(body)

	if l.bodyContent != "non-json content\n" {
		t.Errorf("Non-JSON body not handled correctly: got %s", l.bodyContent)
	}
}

func TestSecureLoggerCanLog(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	l := newSecureLogger()
	l.SetCanLog(false)
	req, _ := http.NewRequest("GET", "/abilities", nil)
	l.LogReq(req)

	if buf.String() != "" {
		t.Errorf("Logger should not have log: got %s", buf.String())
	}

	l.SetCanLog(true)
	l.LogReq(req)
	if !strings.Contains(buf.String(), "/abilities") {
		t.Errorf("Request not logged correctly: got %s", buf.String())
	}
}

func TestSecureLoggerLogReq(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	l := newSecureLogger()
	l.SetCanLog(true)
	req, _ := http.NewRequest("GET", "/abilities", nil)
	l.LogReq(req)

	if !strings.Contains(buf.String(), "/abilities") {
		t.Errorf("Request not logged correctly: got %s", buf.String())
	}
}

func TestSecureLoggerLogRes(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	l := newSecureLogger()
	l.SetCanLog(true)
	res := &http.Response{
		Proto:      "HTTP/1.1",
		StatusCode: 200,
		Status:     "OK",
		Header:     make(http.Header),
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"key": "value"}`))),
	}
	l.LogRes(res)

	if !strings.Contains(buf.String(), "HTTP/1.1 200 OK") {
		t.Errorf("Response not logged correctly: got %s", buf.String())
	}
}
