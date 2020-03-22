package main

import (
	"fmt"
	"github.com/carvalhorr/go-urlshortener/common"
	"github.com/carvalhorr/go-urlshortener/model"
	"github.com/mattheath/base62"
	"log"
	"net/http"
)

var counter = int64(1)

func main() {
	err := common.ConfigViper("dev", "../properties")
	if err != nil {
		panic(err)
	}
	shortener := createUrlShortener()
	http.HandleFunc("/shorten", shortener.shortenUrl)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func (shortener *urlShortener) shortenUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	longUrl, err := getLongURLFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortCode := getNextCode()

	if err := shortener.persist(longUrl, shortCode); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	returnSuccess(w, shortCode)
}

func (shortener *urlShortener) persist(longUrl, shortCode string) error {
	shortUrl := model.ShortURL{
		ShortURLCode: shortCode,
		LongURL:      longUrl,
	}

	return shortener.shortUrlDao.Add(shortUrl)
}

func getNextCode() string {
	short := base62.EncodeInt64(getNextNumber())
	return short
}

func getLongURLFromRequest(r *http.Request) (string, error) {
	queryParam := r.URL.Query()["url"]
	switch {
	case len(queryParam) == 0:
		return "", fmt.Errorf("expected a URL")
	case len(queryParam) > 1:
		return "", fmt.Errorf("only one URL is expected")
	}
	return queryParam[0], nil
}

func returnSuccess(w http.ResponseWriter, code string) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("{\"shortCode\":\"%s\"}", code)))
}

func getNextNumber() int64 {
	defer func() {
		counter = counter + 1
	}()
	return counter
}
