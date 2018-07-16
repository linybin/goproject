package main

import (
	"net/http"

	"encoding/json"
)

func main() {
	http.HandleFunc("/", postLocation)
	if error := http.ListenAndServe(":8880", nil); error != nil {
		panic(error)
	}
}

type postResponse struct {
	ok string
}
type location struct {
	Id   int
	Lat  int
	Long int
}

func postLocation(writer http.ResponseWriter, request *http.Request) {
	//get the json from the body

	var l location
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&l)
	if err != nil {
		http.Error(writer, "error is "+err.Error(), 500)
		return
	}
	json.NewEncoder(writer).Encode(l)
}
