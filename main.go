package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/github-api", githubHandler)

	http.ListenAndServe(":8080", nil)
}

func githubHandler(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", "https://api.github.com/", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Add("Accept", `application/json`)

	client := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Body : %s \n ", body)
	fmt.Printf("Response status : %s \n", resp.Status)

	fmt.Fprint(w, "Github call ok\n")
}
