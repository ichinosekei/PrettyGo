package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) DecodeString(input string) (string, error) {
	url := fmt.Sprintf("%s://%s:%s/decode", c.Scheme, c.Host, c.Port)
	inputData := map[string]string{"inputString": input}
	jsonData, err := json.Marshal(inputData)
	if err != nil {
		return "", err
	}

	resp, err := c.client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response struct {
		OutputString string `json:"outputString"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	return response.OutputString, nil
}
