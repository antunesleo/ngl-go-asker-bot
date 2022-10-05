package nglclient

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const FAKE_DEVICE_ID = "4c5c626-a2eb-42e6-b639-689594052958"

type NGLClient struct {
	URL string
}

func (c NGLClient) AskQuestion(user, question string) error {
	url := c.URL + "/" + user
	postBody, _ := json.Marshal(map[string]string{
		"question": question,
		"deviceId": FAKE_DEVICE_ID,
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
		return err
	}

	log.Println("statusCode", resp.StatusCode)
	return nil
}
