package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/wanghan/dropbox/storage"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d \
'{"buyer_cert":"","seller_sig":"","seller_cert":"","offer_id":"test_offer","time_expire":0,"payload":"","content":"test"}' \
http://localhost:8081/upload

*/
func UploadRequestHandler(w http.ResponseWriter, r *http.Request) {
	var request UploadRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &request); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	if len(request.Content) > 0 {
		storage.AddPayload(storage.ConnectDB(), request.OfferID, request.Content)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(request); err != nil {
		panic(err)
	}
}

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d \
'{"offer_id":"test_offer"}' \
http://localhost:8081/download

*/
func DownloadRequestHandler(w http.ResponseWriter, r *http.Request) {
	var request DownloadRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &request); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	if len(request.OfferID) > 0 {
		request.Payload = storage.FetchPayload(storage.ConnectDB(), request.OfferID)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(request); err != nil {
		panic(err)
	}
}
