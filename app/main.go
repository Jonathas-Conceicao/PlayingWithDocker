package main

import (
  "os"
  "fmt"
  // "time"
  "net/http"

  "github.com/gorilla/mux"
  "github.com/garyburd/redigo/redis"
)

var c redis.Conn

func main() {
  c, _ = redis.DialURL(os.Getenv("REDIS_URL"))
  defer c.Close()
  r := mux.NewRouter()
  r.HandleFunc("/", handle).Methods("GET")
  fmt.Fprintln(os.Stdout, "Application started!")
  http.ListenAndServe(":" + os.Getenv("APPLICATION_PORT"), r)
}

func handle (w http.ResponseWriter, r *http.Request) {
  n, _ := c.Do("INCR", "acess")
  fmt.Fprintln(os.Stdout, "Getting a request from:", r.RemoteAddr)
  fmt.Fprintf(w, "Hello World! I have been seen %v time(s).\n", n)
}
