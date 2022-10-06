package nglclient

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bitly/go-simplejson"
)

func TestShouldAskQuestionOnNGL(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
		}
		if r.URL.EscapedPath() != "/breno" {
			t.Errorf("Expected request to ‘/breno, got ‘%s’", r.URL.EscapedPath())
		}
		reqJson, err := simplejson.NewFromReader(r.Body)
		if err != nil {
			t.Errorf("Error while reading request JSON: %s", err)
		}
		deviceId := reqJson.GetPath("deviceId").MustString()
		if deviceId != FAKE_DEVICE_ID {
			t.Errorf("Expected deviceID to  be %s, got %s", FAKE_DEVICE_ID, deviceId)
		}
		question := reqJson.GetPath("question").MustString()
		if question != "The earth is flat?" {
			t.Errorf("Expected question to  be The earth is flat?, got %s", question)
		}
		fmt.Fprintln(w, "Hello, client")
	}))

	defer ts.Close()
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	c := NGLClient{Writer: writer, URL: ts.URL}

	err := c.AskQuestion("breno", "The earth is flat?")
	if err != nil {
		t.Errorf("Expected request to be successfull")
	}
}
