package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func githubHandler(w http.ResponseWriter, r *http.Request) {
	// Add sleep to make spans more distinct
	time.Sleep(200 * time.Millisecond)

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

	fmt.Printf("Response status : %s \n", resp.Status)

	fmt.Fprint(w, "Github call ok\n")
}
