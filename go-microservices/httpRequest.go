package main

import (
	"fmt"
	"log"
	"net/http"
)
func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Webserver has started")
	fmt.Fprintf(w,"Hello , you have requested %s with token %s\n",r.URL.Path,r.URL.Query().Get("token"))
}

/*func main(){
	http.HandleFunc("/",handler)

	
	fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("Webserver has started")
	http.ListenAndServe(":80",nil)
}*/