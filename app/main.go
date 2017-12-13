package main

import (
  "os"
  "fmt"
  "time"
  "net/http"
  "github.com/garyburd/redigo/redis"
)

var c redis.Conn

func main() {
  var err error
  for trys := 5; trys > 0; trys++ {
    c, err = redis.DialURL(os.Getenv("REDIS_URL"))
    if err != nil {
      fmt.Fprintf(os.Stderr, "Try %v failed dial to: %v\n", trys, os.Getenv("REDIS_URL"))
      time.Sleep(5 * time.Second)
    } else {
      fmt.Fprintf(os.Stderr, "Redis initialized!\n")
      break;
    }
  }
  defer c.Close()
  http.HandleFunc("/", handle)
  http.ListenAndServe(":" + os.Getenv("APPLICATION_PORT"), nil)
}

func handle (w http.ResponseWriter, r *http.Request) {
  // n, _ := c.Do("GET", "acess")
  n, _ := c.Do("INCR", "acess")
  fmt.Fprintf(w, "Hello World! I have been seen %v time(s).\n", n)
}
