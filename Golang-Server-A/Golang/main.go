package main

import (
	"fmt"
	"net/http"
	"log"  
	//"io/ioutil"
	"encoding/json"
	//"strconv"
	"google.golang.org/grpc"
	"context"
	"time"
	"bytes"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address     = "192.168.99.100:50051"
)

type Caso struct {
    Name string `json:"name"`
    Location string `json:"location"`
	Age int `json:"age"`
	Infectedtype string `json:"infectedtype"`
	State string `json:"state"`
}

type ResponseModel struct {
	Message string 
}

func main() {
	log.Printf("GO Escuchando...")
	http.HandleFunc("/", postData)
	http.ListenAndServe(":8080", nil)
}

func postData(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
		case "GET":
			fmt.Fprintf(writer, "Hola Mundo!\n")

		case "POST":
			fmt.Println(">>>>Entre al POST<<<<")
			if err := request.ParseForm(); err != nil {
				fmt.Fprintf(writer, "ParseForm() err: %v", err)
				return
			}

			buf := new(bytes.Buffer)
			buf.ReadFrom(request.Body)
			s := buf.String()
			res := ResponseModel{Message: send(s)}	
			
			jsonContent , e2 := json.Marshal(res)
			if e2 != nil {
				panic(e2)
			}

    		fmt.Fprintf(writer, string( jsonContent))

		default:
			fmt.Fprintf(writer, "Error en la solicitud!\n")
	}
}

func send(content string) string {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo establecer la conexión: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: content})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return r.GetMessage()
}