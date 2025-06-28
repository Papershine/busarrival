package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		poll()
	}
}

func poll() {
	resp, err := http.Get("https://data.etabus.gov.hk/v1/transport/kmb/route/")
	if err != nil {
		fmt.Println("Error polling:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Polled data:", string(body))
}
