package main

import(
   "fmt"
   "testing"
   "flag"
)

func TestHandleRequest(t *testing.T) {
  advanced := flag.Bool("advanced", false, "Test the advanced portion of the excercise")
  fmt.Println(advanced)
  t.Error("for", advanced) 
}
