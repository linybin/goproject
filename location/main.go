package main

import (
	"net/http"

	"encoding/json"
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
	"fmt"
	"time"
)

func main() {
	log.Printf("server starts")
	http.HandleFunc("/", postLocation)
	if err := http.ListenAndServe(":3333", nil); err != nil {
		panic(err)
	}
}


type location struct {
	// The right side is the name of the JSON variable

	Driverid  string `json:"driverid"`
	TimeStamp string `json:"timestamp"`

	RegionCode string    `json:"region_code"`
	RegionName string    `json:"region_name"`
	City       string    `json:"city"`
	Zipcode    string    `json:"zipcode"`
	Lat        float32   `json:"latitude"`
	Lon        float32   `json:"longitude"`
	MetroCode  int       `json:"metro_code"`
	AreaCode   int       `json:"area_code"`
	CreatedAt  time.Time `json:"created_at"`

}

func postLocation(writer http.ResponseWriter, request *http.Request) {
	//get the json from the body
	switch request.Method {
	case "POST":
		writer.Write([]byte("Post"))
		//get the body
		var l location
		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&l)

		if err != nil {
			http.Error(writer, "error is "+err.Error(), 500)
			return
		}
		//we need to save the post into somewhere
		//for the timebeing we will save into local json
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("ap-southeast-1")},
		)
		svc := dynamodb.New(sess)
		av, err := dynamodbattribute.MarshalMap(l)
		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String("benlin-test-table"),
		}
		_, err = svc.PutItem(input)

		if err != nil {
			fmt.Println("Got error calling PutItem:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		json.NewEncoder(writer).Encode(l)

		return
	case "GET":
		writer.Write([]byte("Get"))
		return

	}


}
