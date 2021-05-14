package main

import (
	"encoding/json"
	"fmt"
	"github.com/sanan-fataliyev/shorturls/storage"
	"github.com/sanan-fataliyev/shorturls/urlshorten"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	BaseURL = "localhost:8000"
)

var (
	Service = urlshorten.NewService(BaseURL, storage.MapStorage{})
)

// serializers

type CreateReq struct {
	OriginURL string `json:"origin_url"`
}

type CreateRes struct {
	ShortURL string `json:"short_url"`
}

type ErrorRes struct {
	Message string `json:"message"`
}

func createHandler(w http.ResponseWriter, r *http.Request) {

	payload, err := ioutil.ReadAll(r.Body)

	if err != nil {
		// TODO handle
	}
	req := CreateReq{}

	err = json.Unmarshal(payload, &req)
	if err != nil {
		// TODO handle
	}

	shortURL, err := Service.CreateShortURL(req.OriginURL)

	if err != nil {
		// TODO handle
	}

	responsePayload := &CreateRes{ShortURL: shortURL}

	bytes, err := json.Marshal(responsePayload)

	if err != nil {
		// TODO handle
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(bytes)

}


func getHandler(w http.ResponseWriter, r *http.Request) {

	short := strings.TrimLeft(r.RequestURI, "/")
	shortURL := fmt.Sprintf("%s/%s", BaseURL, short)

	originURL, found := Service.GetOriginURL(shortURL)

	if !found {
		// return 404
	}

	http.Redirect(w, r, originURL, http.StatusSeeOther)
}

func main() {

	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/", getHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
