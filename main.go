package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
    "flag"
)


var directoryPtr string = ""
var port string = ""

func main(){
    flag.StringVar(&directoryPtr,"d","","path to serve")
    flag.StringVar(&port,"p","8000","port to listen on.")
    flag.Parse()
    if directoryPtr == "" {
        flag.PrintDefaults()
        return
    }
    fileServer := http.FileServer(http.Dir(directoryPtr))
    http.Handle("/",fileServer)
    fmt.Fprintf(os.Stderr,"Listening on port %s...",port)
    err := http.ListenAndServe(fmt.Sprintf(":%s",port),nil)
    if err != nil {
        log.Fatal(err)
    }
}
