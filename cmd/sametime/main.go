package main

import (
	"flag"
	"io/ioutil"
	"log"
	"time"

	"sync"

	"github.com/google/uuid"
)

// These are our flag variables
var workers = flag.Int("workers", 1, "Number of Workers")
var testcount = flag.Int("datastrings", 1, "Number of Data Strings")

// Wait group from the sync package
var wg sync.WaitGroup

func main() {
	// Setting a time so we can calculate Duration
	start := time.Now()

	// dont forget to parse your flags
	flag.Parse()

	// make a channel to safely communicate with our workders
	s := make(chan string)

	// Create a worker pool
	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go writeFile(s)
	}

	// create a UUID and set it on the channel
	for x := 0; x < *testcount; x++ {
		ud, err := uuid.NewUUID()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		s <- ud.String()
	}

	// make sure the channel is empty before we close it
	// this helps us support buffered channels
	if len(s) == 0 {
		close(s)
	}

	// waits for all our workers to finish
	//wg.Wait()
	f := time.Since(start)
	log.Printf("Workload Count:%v Goroutine Count:%v Duration:%v", *workers, *testcount, f)
}

func writeFile(itemschan chan string) {
	// range the channel watching as the items come in
	for data := range itemschan {
		err := ioutil.WriteFile("/Users/ldeakm/temp/"+data+".txt", []byte(data), 0644)
		if err != nil {
			log.Printf("file error%v", err.Error())
		}
	}
	// gorouting is exiting
	wg.Done()
}
