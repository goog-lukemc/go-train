package main

import (
	"flag"
	"io/ioutil"
	"log"
	"time"

	"sync"

	"github.com/google/uuid"
)

var workers = flag.Int("workers", 1, "Number of Workers")
var testcount = flag.Int("datastrings", 1, "Number of Data Strings")

var wg sync.WaitGroup

func main() {
	start := time.Now()
	flag.Parse()
	s := make(chan string)

	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go writeFile(s)
	}

	for x := 0; x < *testcount; x++ {
		ud, err := uuid.NewUUID()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		s <- ud.String()
	}

	if len(s) == 0 {
		close(s)
	}

	wg.Wait()
	f := time.Since(start)
	log.Printf("Workload Count:%v Goroutine Count:%v Duration:%v", *workers, *testcount, f)
}

func writeFile(itemschan chan string) {
	for data := range itemschan {
		err := ioutil.WriteFile("/Users/ldeakm/temp/"+data+".txt", []byte(data), 0644)
		if err != nil {
			log.Printf("file error%v", err.Error())
		}
	}
	wg.Done()
}
