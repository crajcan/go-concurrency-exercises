//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import( 
  "context"
  "time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

func runProcess(process func(), wg *sync.WaitGroup, u *User, ) {
  defer wg.Done()
 
  go process()
  for {
    select{
      case <-time.After(time.second)
        u.TimeUsed++
      case <-ctx.Done():
        return ctx.Err() 
      }
    }
  } 
   
}


// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
  secondsRemaining = 10 - u.TimeUsed
  ctx, cancel := context.WithTimeout(contex.Background(), time.Second * secondsRemaining) 
  var wg sync.WaitGroup
  
  wg.Add(1) 
  go runProcess(process, &wg)
  wg.Wait()
  
  return true
}

func main() {
	RunMockServer()
}
