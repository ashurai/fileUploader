package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	hand "github.com/ashurai/fileUploader/handlres"

)

func main(){
	fmt.Println("starting file upload server")
	setRoutes()
}

// setRoutes to deal with all router used in application
func setRoutes(){
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/upload", hand.UploadFile).Methods("POST")
	r.HandleFunc("/files/{page:[0-9]+}", hand.GetFiles).Methods("GET")
	log.Fatal(http.ListenAndServe(":8090", r))
}