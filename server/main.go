package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const version = "v1.0.0"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/version", versionHandler)
	mux.HandleFunc("/decode", decodeHandler)
	mux.HandleFunc("/hard-op", hardOpHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	fmt.Println("Server is running on port 8080...")

	// Graceful shutdown
	gracefulShutdown(server)
}

// высылает версию
func versionHandler(w http.ResponseWriter, r *http.Request) {
	// говорим 200 код все хорошо серверу (http.StatusOK)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(version))
	//write, err := w.Write([]byte(version))
	//if err != nil {
	//	return
	//}
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		InputString string `json:"inputString"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(input.InputString)
	if err != nil {
		http.Error(w, "Invalid base64 string", http.StatusBadRequest)
		return
	}

	response := struct {
		OutputString string `json:"outputString"`
	}{
		OutputString: string(decoded),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func hardOpHandler(w http.ResponseWriter, r *http.Request) {
	randomSleep := time.Duration(rand.Intn(11)+10) * time.Second
	time.Sleep(randomSleep)

	if rand.Intn(2) == 0 {
		// кидаем ошибку 500
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Operation completed successfully"))
	}
}

// завершаем работу HTTP-сервера
func gracefulShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	// программа не выполняется пока не получим сигнал
	<-stop
	// даем время, чтобы завершить активные запросы
	// в рамках нашего задание так понимаю не надо давать время на завершение
	// но и не надо использовать server.Close() поэтому таймер
	// с 0 сек программа попытается завершить
	ctx, cancel := context.WithTimeout(context.Background(), 0*time.Second)
	defer cancel()

	// если не получится завершить программу сделаем
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error during server shutdown: %v", err)
	}

	fmt.Println("Server gracefully stopped")
}
