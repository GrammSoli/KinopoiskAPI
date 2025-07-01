package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type Rating struct {
	KP float64 `json:"kp"`
}

type Poster struct {
	URL        string `json:"url"`
	PreviewURL string `json:"previewUrl"`
}
type Movie struct {
	Name   string `json:"name"`
	Year   int    `json:"year"`
	ID     int    `json:"id"`
	Rating Rating `json:"rating"`
	Poster Poster `json:"poster"`
}

type KinopoiskResponse struct {
	Docs []Movie `json:"docs"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	kinopoiskKey := os.Getenv("KINOPOISK_API_KEY")
	fmt.Println(kinopoiskKey)
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		// w.Write([]byte("pong"))
		title := r.URL.Query().Get("title")
		if title == "" {
			http.Error(w, "Параметр 'title' обязателен", http.StatusBadRequest)
			return
		}

		urlKino := "https://api.kinopoisk.dev/v1.4/movie/search?query=" + url.QueryEscape(title)
		req, _ := http.NewRequest("GET", urlKino, nil)
		req.Header.Set("X-API-KEY", kinopoiskKey)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		if resp.StatusCode != http.StatusOK {
			log.Fatalf("Ошибка при запросе к API Kinopoisk: %s", body)
		}
		respStruct := KinopoiskResponse{
			Docs: []Movie{},
		}
		if err := json.Unmarshal(body, &respStruct); err != nil {
			log.Println("Ошибка при разборе JSON:", err)
			http.Error(w, "Ошибка на сервере", http.StatusInternalServerError)
			return
		}
		for _, movie := range respStruct.Docs {
			fmt.Println(movie.Name) // для терминала

			fmt.Fprintf(w,
				"<h2>%s (%d)</h2><p>ID: %d<br>Rating: %.1f</p><img src=\"%s\" width=\"200\"><hr>",
				movie.Name, movie.Year, movie.ID, movie.Rating.KP, movie.Poster.URL)

		}

	})
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}

}
