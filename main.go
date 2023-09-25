package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
    "flag"
)


var directoryPtr string = ""

func main(){
    flag.StringVar(&directoryPtr,"d","","path to serve")
    flag.Parse()
    if directoryPtr == "" {
        flag.PrintDefaults()
        return
    }
    fileServer := http.FileServer(http.Dir(directoryPtr))
    http.Handle("/",fileServer)
    fmt.Fprintf(os.Stderr,"Listening on 8000..")
    err := http.ListenAndServe(":8000",nil)
    if err != nil {
        log.Fatal(err)
    }
}
