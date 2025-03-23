package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	url := flag.String("URL", "", "URL to fetch")
	flag.Parse()

	if *url == "" {
		fmt.Println("Usage: main -url <URL>")
		os.Exit(1)
	}

	fmt.Println("app env: ", os.Getenv("APP_ENV"))

	for {
		resp, err := http.Get(*url)
		if err != nil {
			fmt.Printf("Error fetching URL: %v\n", err)
			fmt.Println("------------------- will called in next 5s -------------------------------")
			time.Sleep(5 * time.Second)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %v\n", err)
		}

		fmt.Printf("HTTP/1.1 %s\n", resp.Status)
		for k, v := range resp.Header {
			fmt.Printf("%s: %s\n", k, v)
		}
		fmt.Println()
		fmt.Println(string(body))
		fmt.Println("------------------- will called in next 5s -------------------------------")
		time.Sleep(5 * time.Second)

	}

}
