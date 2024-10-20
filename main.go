package main

import (
	"./client"
	"fmt"
)

func main() {
	// Инициализация клиента
	client := client.NewClient("http", "localhost", "8080")

	// Вызов метода GetVersion
	version, err := client.GetVersion()
	if err != nil {
		fmt.Println("Error fetching version:", err)
		return
	}
	fmt.Println("Server version:", version)

	// Вызов метода DecodeString
	decoded, err := client.DecodeString("SGVsbG8sIFdvcmxkIQ==")
	if err != nil {
		fmt.Println("Error decoding string:", err)
		return
	}
	fmt.Println("Decoded string:", decoded)

	// Вызов метода HardOpWithTimeout
	status, err := client.HardOpWithTimeout()
	if err != nil {
		fmt.Println("Error performing hard operation:", err)
		return
	}
	fmt.Println("Hard operation status code:", status)
}
