package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

var startedAt = time.Now()

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello, I'm %s. I'm %s.", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: ", err)
	}
	fmt.Fprintf(w, "My family: %s.", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {

	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	fmt.Fprintf(w,"User: %s. Password %s.", user, password)

}

func Healthz(w http.ResponseWriter, r *http.Request) {

	duration := time.Since(startedAt)

	if duration.Seconds() < 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintln("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}

}