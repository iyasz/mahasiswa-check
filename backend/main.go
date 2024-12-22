package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"net/http"
 	"github.com/PuerkitoBio/goquery"
)

func HandleScrape() ([]map[string]string, error) {
	res, err := http.Get("https://pddikti.kemdiktisaintek.go.id/search/elaina")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}
  
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var results []map[string]string
  
	doc.Find("#root div div:nth-child(4) div:nth-child(5) div:nth-child(1) div:nth-child(2) table tbody tr").Each(func(i int, row *goquery.Selection) {
		name := row.Find("td:nth-child(1) span").Text()       // Kolom 1 (Nama)
		nim := row.Find("td:nth-child(2) span").Text()        // Kolom 2 (NIM)
		university := row.Find("td:nth-child(3) span").Text() // Kolom 3 (Universitas)
		major := row.Find("td:nth-child(4) span").Text()      // Kolom 4 (Jurusan)
		detailURL, _ := row.Find("td:nth-child(5) a").Attr("href") // Link detail

		// Simpan data dalam map
		results = append(results, map[string]string{
			"name":       name,
			"nim":        nim,
			"university": university,
			"major":      major,
			"detailURL":  detailURL,
		})

	})

	return results, nil

  }
  

func main() {

	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("keren")
		results, err := HandleScrape()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)

		// allowedOrigin := "http://localhost:5173" 

		// origin := r.Header.Get("Origin")
		// if origin == allowedOrigin {
		// 	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		// 	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// } else {
		// 	http.Error(w, "CORS policy does not allow this origin", http.StatusForbidden)
		// 	return
		// }

		// if r.Method == "OPTIONS" {
		// 	w.WriteHeader(http.StatusOK)
		// 	return
		// }

		// w.Header().Set("Content-Type", "application/json")

		// query := r.URL.Query().Get("q")

		// if query == "" {
		// 	http.Error(w, "Query 'q' is missing", http.StatusBadRequest)
		// 	return
		// }

		// apiUrl := fmt.Sprintf("https://pddikti.kemdiktisaintek.go.id/search/%s", query)

		// client := &http.Client{}
		// req, err := http.NewRequest("GET", apiUrl, nil)
		// if err != nil {
		// 	http.Error(w, "Failed to create request", http.StatusInternalServerError)
		// 	return
		// }
		// req.Header.Set("Accept", "application/json")
		// req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

		// res, err := client.Do(req)
		// if err != nil {
		// 	http.Error(w, "Failed to fetch API", http.StatusInternalServerError)
		// 	return
		// }

		// defer res.Body.Close()

		// body, err := ioutil.ReadAll(res.Body)
		// if err != nil {
		// 	http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		// 	return
		// }

		// fmt.Println("Respon API: ", string(body))

		// w.WriteHeader(res.StatusCode) 
		// w.Write(body)
	})

	fmt.Println("Server started at :3000")
	http.ListenAndServe(":3000", nil)

}
