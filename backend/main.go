package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

)

func main() {

	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {

		
		allowedOrigin := "http://localhost:5173" 

		origin := r.Header.Get("Origin")
		if origin == allowedOrigin {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		} else {
			http.Error(w, "CORS policy does not allow this origin", http.StatusForbidden)
			return
		}

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		query := r.URL.Query().Get("q")

		if query == "" {
			http.Error(w, "Query 'q' is missing", http.StatusBadRequest)
			return
		}

		apiUrl := fmt.Sprintf("https://api.ryzendesu.vip/api/search/mahasiswa?query=%s", query)

		client := &http.Client{}
		req, err := http.NewRequest("GET", apiUrl, nil)
		if err != nil {
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

		res, err := client.Do(req)
		if err != nil {
			http.Error(w, "Failed to fetch API", http.StatusInternalServerError)
			return
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}


		w.WriteHeader(res.StatusCode) 
		w.Write(body)
	})

	fmt.Println("Server started at :3000")
	http.ListenAndServe(":3000", nil)

}
