package client

import (
	"fmt"
	"io/ioutil"
)

func (c *Client) GetVersion() (string, error) {
	url := fmt.Sprintf("%s://%s:%s/version", c.Scheme, c.Host, c.Port)
	resp, err := c.client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
