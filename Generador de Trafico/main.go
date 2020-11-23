package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"encoding/json"
	"io/ioutil"
	"time"
	"log"
	"net/http"
)

//urlBalanceador: 192.168.99.100:80
//ruta archivo: datos.json

var urlBalanceador string
var cantidadSubrutinas int
var cantidadSolicitudes int
var urlArchivo string

var peticionesEnviadas int =0
var flag bool = false
var timeCurrent = time.Now()

type Casos struct {
    Casos []Caso `json:"casos"`
}

type Caso struct {
    Name string `json:"name"`
    Location string `json:"location"`
	Age int `json:"age"`
	Infectedtype string `json:"infectedtype"`
	State string `json:"state"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	//Toma de datos---------------------------------------------------------------
	fmt.Println("--------------------------------------------------")
	fmt.Println("-----------------Ingreso de Datos-----------------")
	fmt.Println("--------------------------------------------------")
	
	fmt.Print("URL Balanceador: ")
	urlBalanceador, _ := reader.ReadString('\n')
	urlBalanceador = strings.TrimSpace(urlBalanceador)

	fmt.Print("Cantidad de Subrutinas: ")
	temporal, _ := reader.ReadString('\n')
	temporal = strings.TrimSpace(temporal)
	t_cantidadSubrutinas, _ := strconv.ParseFloat(temporal, 64)
	cantidadSubrutinas = int(t_cantidadSubrutinas)

	fmt.Print("Cantidad de Solicitudes: ")
	temporal, _ = reader.ReadString('\n')
	temporal = strings.TrimSpace(temporal)
	t_cantidadSolicitudes, _ := strconv.ParseFloat(temporal, 64)
	cantidadSolicitudes = int(t_cantidadSolicitudes)

	fmt.Print("URL Archivo: ")
	urlArchivo, _ = reader.ReadString('\n')
	urlArchivo = strings.TrimSpace(urlArchivo)

	fmt.Println("")
	fmt.Println(urlBalanceador)
	fmt.Println(cantidadSubrutinas)
	fmt.Println(cantidadSolicitudes)
	fmt.Println(urlArchivo)
	fmt.Println("")

	//Lectura de archivo-------------------------------------------------------------------
	jsonFile, err := os.Open(urlArchivo)

	// error al abrir el archivo
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Archivo abierto correctamente")
	}
	
	defer jsonFile.Close()

	//Se parsean los datos
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var casos Casos
	json.Unmarshal(byteValue, &casos)

	for i := 0; i < len(casos.Casos); i++ {
		fmt.Println("Caso Name: " + casos.Casos[i].Name)
	}

	empezarConcurrencia(cantidadSubrutinas, urlBalanceador, casos.Casos, cantidadSolicitudes)
}

func empezarConcurrencia(concurrence int, url string, data []Caso, solicitudes int){

	peticionesEnviadas = 0
	flag = false


	for i:=0; i< concurrence; i++{

		go peticion(i, url, data, solicitudes, 1000)

	}

	fmt.Println("Los datos fueron enviados.")
	var input string
	fmt.Scanln(&input)

}

func peticion(noPeticion int, url string, data [] Caso, total int, timesiu int)  {

	cont := 0

	for peticionesEnviadas < total {

		sendDataPost(data[cont], url, noPeticion, peticionesEnviadas)

		cont = cont+1

		if cont == len(data){
			cont = 0
		}

		peticionesEnviadas+=1
		time.Sleep((time.Second))
	}
}

func sendDataPost(dataFinal Caso, url string, noConcurrence int, totalPeticiones int) {

	requestBody := strings.NewReader(`{ "name": "`+dataFinal.Name+`", "location":"`+dataFinal.Location+`", "Age": `+ strconv.Itoa(dataFinal.Age) +`, "infectedType":"`+dataFinal.Infectedtype+`", "State":"`+dataFinal.State+`" }`)
	
	res, err := http.Post(
		url,
		"application/json; charset=UTF-8",
		requestBody,
	)

	//-check for response error
	if err != nil {
		fmt.Println("error envio de peticiones")
		log.Fatal( err )
	}

	data, _ := ioutil.ReadAll( res.Body )

	res.Body.Close()

	requestContentType := res.Request.Header.Get( "Content-Type" )
	fmt.Println( "Request content-type:", requestContentType )

	fmt.Println( "%s\n", data )

	fmt.Println("No. concurrencia: ", noConcurrence, " total peticiones: ", peticionesEnviadas, " data: ", requestBody)

}


