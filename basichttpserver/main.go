package main

import (
	"asciicoolness"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"
)

func main() {

	// handle the request to a route
	http.HandleFunc("/ascii/", asciiHandler)
	// catchall default route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("up"))
		return
	})

	// A little golang worker here to help us know what the Http
	// server is ready.  It serves not purpose other than notification
	go checkhttp()

	// A single line in the http package starts an http server.
	http.ListenAndServe(":8080", nil)
}

// This function handles out http request for the json file
func asciiHandler(w http.ResponseWriter, r *http.Request) {
	// Lets get the string  at the end of the URL
	scount := path.Clean(path.Base(r.URL.Path))

	// Try to convert it to and int
	count, err := strconv.Atoi(scount)

	// if it fails lets notify the user by send it to the Http
	// response body
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	// we use our custom package to get us the data
	data := asciicoolness.ASCIICoolness{
		Length: count + 1,
	}
	result, err := data.MarshalJSON()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	// Set the header so the call knows its json data
	w.Header().Set("Content-Type", "application/json")

	// write the body
	w.Write(result)
	return
}

func checkhttp() {
	for i := 0; i < 10; i++ {
		resp, err := http.Get("http://localhost:8080/")
		if resp != nil {
			if resp.StatusCode == 200 {
				log.Println("Server Is up!")
				return
			}
		}
		if err != nil {
			log.Printf("http check error:%v", err.Error())
		}
		time.Sleep(time.Millisecond * 500)
	}

}
