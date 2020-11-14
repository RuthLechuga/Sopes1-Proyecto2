package main

import (
	"fmt"
	"net/http"
	"log"  
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type Caso struct {
    Name string `json:"name"`
    Location string `json:"location"`
	Age int `json:"age"`
	Infectedtype string `json:"infectedtype"`
	State string `json:"state"`
}

func main() {
	http.HandleFunc("/", postData)
	http.ListenAndServe(":8080", nil)
}

func postData(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
		case "GET":
			fmt.Fprintf(writer, "Hola Mundo!\n")
			
		case "POST":
			fmt.Fprintf(writer, "Datos recibidos!\n")
			log.Printf("Datos recibidos!")

			if err := request.ParseForm(); err != nil {
				fmt.Fprintf(writer, "ParseForm() err: %v", err)
				return
			}

			reqBody, _ := ioutil.ReadAll(request.Body)
			var caso Caso 
    		json.Unmarshal(reqBody, &caso)
			log.Printf(caso.Name+" => "+strconv.Itoa(caso.Age))
		
		default:
			fmt.Fprintf(writer, "Error en la solicitud!\n")
	}
}
