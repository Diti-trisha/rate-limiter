package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	counter int = 1
	lastt       = time.Now()
)

func handler(w http.ResponseWriter, r *http.Request) {
	var possible bool
	tn := time.Now()
	counter++
	if r.Method == http.MethodGet {
		if tn.Sub(lastt) < 1*time.Second {
			if counter > 20 {
				log.Printf("Can not process this for %v", lastt.Add(1*time.Second).Sub(tn))
			} else {
				possible = true
			}
		} else {
			possible = true
			lastt = tn
			counter = 1
		}
	} else {
		if tn.Sub(lastt) < 1*time.Second {
			if counter > 5 {
				http.Error(w, "Can not process", http.StatusTooManyRequests)
				log.Printf("Can not process this for %v", lastt.Add(1*time.Second).Sub(tn))
			} else {
				possible = true
			}
		} else {
			possible = true
			lastt = tn
			counter = 1
		}
	}
	if possible {
		fmt.Fprintln(w, "Processed")
		log.Println("Processed")
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server is online on localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// one request : curl -i -X POST localhost:8080
// multiple request : for /l %i in (1, 1, 10) do curl -i -X POST localhost:8080
