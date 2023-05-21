package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	client := &http.Client{Timeout: 10 * time.Second}

	var wg sync.WaitGroup

	responseChan1 := make(chan []byte, 1)
	responseChan2 := make(chan []byte, 1)

	wg.Add(2)

	go func() {
		defer wg.Done()

		resp, err := client.Get("http://localhost:8080/api/v1/async")
		defer resp.Body.Close()

		if err != nil {
			log.Fatal(err)
			return
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		responseChan1 <- b
	}()

	go func() {
		defer wg.Done()

		resp, err := client.Get("http://localhost:8080/api/v1/async")
		defer resp.Body.Close()

		if err != nil {
			log.Fatal(err)
			return
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		responseChan2 <- b
	}()
	wg.Wait()

	resp1 := <-responseChan1
	resp2 := <-responseChan2
	fmt.Println("body 1:", string(resp1))
	fmt.Println("body 2:", string(resp2))
}
