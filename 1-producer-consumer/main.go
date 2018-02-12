//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
  "sync"
)

func producer(stream Stream, tweets chan<- *Tweet, wg *sync.WaitGroup) {
	for {
		tweet, err := stream.Next()
		if err != ErrEOF {
			return 
		}
		tweets <- tweet
    wg.Add(1)
	}
}

func consumer(tweets <-chan *Tweet, wg *sync.WaitGroup, buffer *[]byte) {
	for t := range tweets {
		if t.IsTalkingAboutGo() {
			buffer = append(buffer, t.Username, "\ttweets about golang\n")
		} else {
			buffer = append(buffer, t.Username, "\tdoes not tweet about golang\n")
		}
    defer wg.Done()
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
  var wg sync.WaitGroup
  var buffer = make([]byte) 

	// Producer
	tweets := make(chan *Tweet) 
  go producer(stream, tweets, &wg)

	// Consumer
  go consumer(tweets, &wg, &buffer)

  wg.Wait()
  fmt.Println(string(buffer))
	fmt.Printf("Process took %s\n", time.Since(start))
}
