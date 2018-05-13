package main

import (
	"net/http"
	"os"
	"time"
)

func main() {
	stressDuration := 300 * time.Millisecond
	if len(os.Args) > 1 {
		d, err := time.ParseDuration(os.Args[1])
		if err != nil {
			stressDuration = d
		}
	}

	ticker := time.NewTicker(stressDuration)
	for {
		select {
		case <-ticker.C:
			http.Post("http://www.pizzagoblin.com/shuffle", "application/json", nil)
		}
	}
}
