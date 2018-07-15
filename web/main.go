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
	"html/template"
	"os"
	"io/ioutil"
	"html"
	"path/filepath"

)



func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("cnanot find url ")
		os.Exit(1)
	}
	fmt.Println("current directory is " + cwd)
	fmt.Println("start the http server ..")
	http.HandleFunc("/", listOfOptions)
	http.HandleFunc("/order", getOrderThroughGRPC)

	http.HandleFunc("/order/22/get", getOrder)
	http.HandleFunc("/placeOrder", placeOrder)
	http.HandleFunc("/get_price", getPrice)
	http.HandleFunc("/createOrder", createOrder)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)

	}
}
func createOrder(writer http.ResponseWriter, request *http.Request) {
	path := html.EscapeString(request.URL.Path)
	log.Println("request from" + path)
	log.Println("Handler create order started ")
	defer log.Println("Handler stop ")

	conn, err := grpc.Dial("my-order-service:22222", grpc.WithInsecure())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	c := order.NewOrderServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	or, err := c.CreateOrder(ctx, &order.Order{UserId: "222", Id: 22, From: "df", To: "23", Long: 33, Lat: 33, Status: "pending"})
	if err != nil {

		http.Error(writer, err.Error(), 500)
		return
	}
	json.NewEncoder(writer).Encode(or)



}
func getPrice(writer http.ResponseWriter, request *http.Request) {
	path := html.EscapeString(request.URL.Path)
	log.Println("request from" + path)
	log.Println("Handler started ")
	defer log.Println("Handler stop ")
	var netCLient = &http.Client{
		Timeout: time.Second * 5,
	}
	response, err := netCLient.Get("http://my-pricing-service")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Write(body)

}

type mainstruct struct {
}

func listOfOptions(writer http.ResponseWriter, request *http.Request) {
	path := html.EscapeString(request.URL.Path)
	log.Println("request from" + path)
	log.Println("Handler started ")
	defer log.Println("Handler stop ")
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles(filepath.Join(wd, "/main.html"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(writer, mainstruct{})

}

func getOrderThroughGRPC(writer http.ResponseWriter, request *http.Request) {
	path := html.EscapeString(request.URL.Path)
	log.Println("request from" + path)
	log.Println("Handler started ")
	defer log.Println("Handler stop ")
	id_string := request.URL.Query().Get("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		//return with error
		return
	}
	id_32 := int32(id)

	fmt.Println("get order through grpc")
	conn, err := grpc.Dial("my-order-service:22222", grpc.WithInsecure())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	c := order.NewOrderServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	or, err := c.GetOrder(ctx, &order.GetOrderRequest{Id: id_32})
	if err != nil {

		http.Error(writer, err.Error(), 500)
		return
	}
	json.NewEncoder(writer).Encode(or)
}
func placeOrder(writer http.ResponseWriter, _ *http.Request) {
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

func getOrder(w http.ResponseWriter, _ *http.Request) {
	order_item := Order{"UX-2382", 22, "tin", "ah", 233, 232, "pending"}
	json.NewEncoder(w).Encode(order_item)

}
