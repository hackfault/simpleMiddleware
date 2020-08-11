package main

import (
	"fmt"
	"log"
	"net/http"
)

type logger struct {
	Inner http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//log.Printf("start %s\n", time.Now().String())
	log.Printf("%v: %v - %v", r.RemoteAddr, r.Method, r.RequestURI)
	l.Inner.ServeHTTP(w, r)
	//log.Printf("finish %s\n", time.Now().String())
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello\n")
}

func main() {
	f := http.HandlerFunc(hello)
	l := logger{Inner: f}
	http.ListenAndServe(":8000", &l)
}
