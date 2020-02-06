package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type RequestHandler struct {
}

type ResponseMessage struct {
	CustomerId   string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
}

type RequestMessage struct {
	CustomerId   string `json:"customer_id"`
}

func (handler RequestHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {


	fmt.Println("I invoked a REST Service...")
	respMsg := ResponseMessage{
		CustomerId:   "123",
		CustomerName: "Wow",
	}

	respBytes, err := json.Marshal(respMsg)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	fmt.Fprint(w, string(respBytes))
}

func main() {

	router := mux.NewRouter()

	//handler := RequestHandler{}
	//
	//router.Methods("GET").Handler(handler)

	router.Path("/customer/{id}").Methods("GET").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Customer")


		requestBytes, _ := ioutil.ReadAll(request.Body)

		requestMessage := RequestMessage{}

		err := json.Unmarshal(requestBytes, &requestMessage)
		if err != nil {
			writer.WriteHeader(400)
		}

		fmt.Printf("%+v", requestMessage)

		id := mux.Vars(request)["id"]
		if status == "" {
			fmt.Fprint(writer, "Customer: ", id)
		} else {
			fmt.Fprint(writer, "Customer: ", id, " With Status: ", status)
		}
	})

	router.Path("/product").Methods( "POST").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Product")
		fmt.Fprint(writer, "Product")
	})


	fmt.Println("Starting my server...")
	http.ListenAndServe(":3000", router)

	fmt.Println("Started my server..")
}
