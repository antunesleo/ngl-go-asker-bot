package nglclient

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShouldAskQuestionOnNGL(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
		}
		if r.URL.EscapedPath() != "/breno" {
			t.Errorf("Expected request to ‘/breno, got ‘%s’", r.URL.EscapedPath())
		}

		var myBody askQuestionBody
		json.NewDecoder(r.Body).Decode(&myBody)
		if myBody.DeviceId != FAKE_DEVICE_ID {
			t.Errorf("Expected deviceID to  be %s, got %s", FAKE_DEVICE_ID, myBody.DeviceId)
		}
		if myBody.Question != "The earth is flat?" {
			t.Errorf("Expected question to  be The earth is flat?, got %s", myBody.Question)
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
