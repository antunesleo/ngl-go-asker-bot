package nglclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const FAKE_DEVICE_ID = "4c5c626-a2eb-42e6-b639-689594052958"

type NGLClient struct {
	URL    string
	Writer io.Writer
}

type askQuestionBody struct {
	DeviceId string `json:"deviceId"`
	Question string `json:"question"`
}

func New(url string, writer io.Writer) *NGLClient {
	return &NGLClient{URL: url, Writer: writer}
}

func (c NGLClient) AskQuestion(user, question string) error {
	url := c.URL + "/" + user
	myAskQuestionBody := askQuestionBody{DeviceId: FAKE_DEVICE_ID, Question: question}
	bodyBuffer := new(bytes.Buffer)
	json.NewEncoder(bodyBuffer).Encode(myAskQuestionBody)
	resp, err := http.Post(url, "application/json", bodyBuffer)
	if err != nil {
		fmt.Fprintf(c.Writer, "An Error Occured %v", err)
		return err
	}

	fmt.Fprintln(c.Writer, "statusCode", resp.StatusCode)
	return nil
}
