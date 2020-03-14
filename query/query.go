package main

import (
	"go-urlshortener/common"
	"log"
	"net/http"
	"strings"
)

func main() {
	err := common.ConfigViper("dev", "./properties")
	if err != nil {
		panic(err)
	}
	shortener := createShortUrlQuery()
	http.HandleFunc("/", shortener.getShortUrl)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func (shortener *shortUrlQuery) getShortUrl(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	shortCode := r.URL.Path
	shortCode = strings.ReplaceAll(shortCode, "/", "")

	switch shortCode {
	case "":
		return
	case "favicon.ico":
		return
	}

	shortUrl, err := shortener.shortUrlDao.Get(shortCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, shortUrl.LongURL, http.StatusFound)
}
