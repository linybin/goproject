package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"context"
	"github.com/linybin/goproject/protos/order"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
)

const (
	portRPC = ":50051"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	message := "handler the test a"
	w.Write([]byte(message))
}

func main() {

	fmt.Println("start the http server ..")
	http.HandleFunc("/test", getOrderThroughGRPC)
	http.HandleFunc("/order/22/get", getOrder)
	http.HandleFunc("/placeOrder", placeOrder)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)

	}
}
func getOrderThroughGRPC(writer http.ResponseWriter, request *http.Request) {
	id_string := request.URL.Query().Get("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		//return with error
		http.Error(writer, "id is wrong", 422)
		return
	}
	id_32 := int32(id)

	fmt.Println("get order through grpc")
	conn, err := grpc.Dial("localhost:22222", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error occurred %v", err)
	}
	defer conn.Close()
	c := order.NewOrderServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	order, err := c.GetOrder(ctx, &order.GetOrderRequest{Id: id_32})
	if err != nil {

		http.Error(writer, "grpc server down", 500)
		return
	}
	json.NewEncoder(writer).Encode(order)
}
func placeOrder(writer http.ResponseWriter, request *http.Request) {
	orderId := "222"
	fmt.Fprintf(writer, "try to place the order %v", orderId)
	//get the order info, save into db and push to event queue
}

type Order struct {
	UserId string
	Id     int
	From   string
	To     string
	Long   uint64
	Lat    uint64
	status string
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	order := Order{"UX-2382", 22, "tin", "ah", 233, 232, "pending"}
	json.NewEncoder(w).Encode(order)

}
