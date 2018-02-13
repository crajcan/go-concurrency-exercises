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
  defer wg.Done()
  L:  
  for {
    tweet, err := stream.Next()
    if err == ErrEOF {
      panic("this is a thing!")
      break L
    } else {
      tweets <- tweet
    }
  }
}

func consumer(tweets <-chan *Tweet, wg *sync.WaitGroup, c chan<- string) {
	for t := range tweets {
		if t.IsTalkingAboutGo() {
      c <- t.Username + "\ttweets about golang\n"
		} else {
			c <- t.Username + "\tdoes not tweet about golang\n"
		}
	}
}


func main() {
	start := time.Now()
	stream := GetMockStream()
  var wg sync.WaitGroup
  var c chan string  
  wg.Add(1)
	
	// Consumer
	tweets := make(chan *Tweet) 
  go consumer(tweets, &wg, c)
  go func(){
    for t := range c {
      fmt.Println(t) 
    } 
  }() 
 
  // Producer
  go producer(stream, tweets, &wg)

  wg.Wait()
	fmt.Printf("Process took %s\n", time.Since(start))
}
