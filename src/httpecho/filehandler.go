package main

// Steven Gantz
// 4/8/2017
// Go-HTTP-Echo
// https://gitlab.com/StevenPG/Go-HTTP-Echo/

import (
	"log"
	"net/http"
	"bytes"
	"io/ioutil"
	"os"
	"time"
)

// JSON as decoded structure
type jsonstring_struct struct {
	JSON string
}

// Takes the method from the request and
// Works appropriately
func buildFolderAndWriteFile(request *http.Request) {
	switch request.Method {
		case "GET":
			log.Println("GET Request received, logging into ./GET")
			os.MkdirAll("GET", 0777)
			appendToFile(buildGETDELETEString(request), "GET/requests.txt")
		case "POST":
			log.Println("POST Request received, logging into ./POST")
			os.MkdirAll("POST", 0777)
			writeJSONToFile(request, "POST")
		case "PUT":
			log.Println("PUT Request received, logging into ./PUT")
			os.MkdirAll("PUT", 0777)
			writeJSONToFile(request, "PUT")
		case "DELETE":
			log.Println("DELETE Request received, logging into ./DELETE")
			os.MkdirAll("DELETE", 0777)
			appendToFile(buildGETDELETEString(request), "DELETE/requests.txt")
		default:
		    log.Println("Invalid Request Type Received... Ignoring...")
	}
}

// Builds the output string for GET and DELETE files
func buildGETDELETEString(request *http.Request) (str string){
	var buffer bytes.Buffer
	buffer.WriteString(request.Method)
	buffer.WriteString(" __HOST: ")
	buffer.WriteString(request.RemoteAddr)
	buffer.WriteString(" __TIME: ")
	buffer.WriteString(time.Now().Format(time.RFC850))
	buffer.WriteString("\n")
	str = buffer.String()
	return
}

func writeJSONToFile(req *http.Request, reqtype string) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
        panic(err)
    }
	err = ioutil.WriteFile(reqtype + "/" + time.Now().Format("2006_01_02_15_04_05") +".txt", []byte(string(body)), 0666)
	if err != nil {
        panic(err)
    }
}

func writeToFile(str string, file string){
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(str); err != nil {
		panic(err)
	}
}

// Appends to input file
func appendToFile(str string, file string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(str); err != nil {
		panic(err)
	}
}

// Clears all folders and contents
func clearAllFolder(){
	log.Println("Clearing all generated server directories...")
	os.RemoveAll("GET")
	log.Println("GET removed")
	os.RemoveAll("POST")
	log.Println("POST removed")
	os.RemoveAll("PUT")
	log.Println("PUT removed")
	os.RemoveAll("DELETE")
	log.Println("DELETE removed")
	log.Println("Complete")
}

func handleFile() {
	log.Println("test")
}
