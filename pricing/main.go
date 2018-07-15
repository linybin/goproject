package main

import (

	"net/http"
	"os"
	"log"
	"errors"
)

const version = "v1"



func main(){
	os.Setenv("port", ":8080")

	port, err := getEnv("port")
	if err != nil {
		panic(err)
	}


	http.HandleFunc("/",  price)

	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)

	}
}
func price(writer http.ResponseWriter, _ *http.Request) {
	log.Println("Handler started ")
	defer log.Println("Handler stop ")


	message := "pricing is 22 and version is " + version
	writer.Write([]byte(message))

}

func getEnv(env string) (string, error) {
	to_return := os.Getenv(env)

	if to_return != "" {
		return to_return, nil
	} else {
		return "", errors.New("Env is not set" + env)
	}
}