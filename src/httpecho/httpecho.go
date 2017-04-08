package main

// Steven Gantz
// 4/8/2017
// Go-HTTP-Echo
// https://gitlab.com/StevenPG/Go-HTTP-Echo/

import (
	"net/http"
	"fmt"
)

var requestPattern = "/request"
var GETinfoPattern = "/info/get"
var DELETEinfoPattern = "/info/delete"
var resetPattern = "/reset"

var localhostPortConfig = "localhost:8000"

func handleIncomingRequest(w http.ResponseWriter, r *http.Request) {
	// Build folder structure if non-existing
	buildFolderAndWriteFile(r)
}

func displayGETInfo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./GET/requests.txt")
}

func displayDELETEInfo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./DELETE/requests.txt")
}

func displayInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "info")
}

func resetServerInfo(w http.ResponseWriter, r *http.Request) {
	clearAllFolder()
}

func main() {

	// Allocate Web Server Object
	server := http.NewServeMux()
	server.HandleFunc(requestPattern, handleIncomingRequest)

	server.HandleFunc(GETinfoPattern, displayGETInfo)
	server.HandleFunc(DELETEinfoPattern, displayDELETEInfo)

	server.HandleFunc(resetPattern, resetServerInfo)

	// Serve POST files
	go func(){
		panic(http.ListenAndServe("localhost:8001", http.FileServer(http.Dir("./POST"))))
	}()

	// Serve PUT files
	go func(){
		panic(http.ListenAndServe("localhost:8002", http.FileServer(http.Dir("./PUT"))))
	}()

	// Serve HTTP 
	http.ListenAndServe(localhostPortConfig, server)
}
