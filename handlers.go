package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func githubHandler(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequestWithContext(r.Context(), "GET", "https://api.github.com/", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Add("Accept", `application/json`)

	client := http.Client{
		Timeout:   time.Duration(1) * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rand.Int()%2 != 0 {
		http.Error(w, "this is a made up error", http.StatusInternalServerError)
		return
	}

	fmt.Printf("client.Do response status : %s \n", resp.Status)

	fmt.Fprint(w, "External call ok\n")
}
