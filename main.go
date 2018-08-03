package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http/httputil"
    "os"
)

func DumpRequest(w http.ResponseWriter, req *http.Request) {
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		fmt.Println(w, string(requestDump))
	}
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", DumpRequest)
    log.Fatal(http.ListenAndServe(":"+os.Args[1], router))
}
