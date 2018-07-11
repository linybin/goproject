package main

import (

	"net/http"
	"os"
)

func main(){
	os.Getenv("port")


	http.HandleFunc("/",  price)
	if err:= http.ListenAndServe(":8080", nil);err!=nil{
		panic(err)

	}
}
func price(writer http.ResponseWriter, request *http.Request) {
	version := os.Getenv("version")
	message := "pricing is 22 and version is " + version
	writer.Write([]byte(message))

}