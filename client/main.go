package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{}

	// GET /version
	versionResp, err := client.Get("http://localhost:8080/version")
	if err != nil {
		fmt.Println("Error fetching version:", err)
		return
	}
	defer versionResp.Body.Close()
	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//
	//	}
	//}(versionResp.Body)
	versionBody, _ := ioutil.ReadAll(versionResp.Body)
	fmt.Println("Version:", string(versionBody))

	// POST /decode
	inputData := map[string]string{"inputString": "SGVsbG8sIFdvcmxkIQ=="}
	jsonData, _ := json.Marshal(inputData)

	decodeResp, err := client.Post("http://localhost:8080/decode", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error fetching decode:", err)
		return
	}
	defer decodeResp.Body.Close()
	decodeBody, _ := ioutil.ReadAll(decodeResp.Body)
	fmt.Println("Decode response:", string(decodeBody))

	// GET /hard-op with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/hard-op", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	hardOpResp, err := client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("Request timed out")
		} else {
			fmt.Println("Error performing hard-op:", err)
		}
		return
	}
	defer hardOpResp.Body.Close()

	fmt.Println("Hard-op status:", hardOpResp.StatusCode)
	if hardOpResp.StatusCode == http.StatusOK {
		fmt.Println("Hard-op completed successfully")
	} else {
		fmt.Println("Hard-op failed with status:", hardOpResp.StatusCode)
	}
}
