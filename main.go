package main

import (
  "net/http"
  "log"
  "runtime"
  "io/ioutil"
)

const (
  version = "0.1"
)

func echoContent(r *http.Request)  {
  // r.ParseForm()
  defer r.Body.Close()
  body, _ := ioutil.ReadAll(r.Body)
  log.Println("BODY - " + string(body))
}

func httpHandler(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Server", "HTTP Recevie Server " + version + " (" + runtime.GOOS + ")")
  if r.Method == "POST" {
    echoContent(r)
    w.Header().Set("Content-Type", "charset=utf-8")
    return
  }
  http.Error(w, "post only", http.StatusMethodNotAllowed)
}

func main() {
  http.HandleFunc("/", httpHandler)
	err := http.ListenAndServe(":8000", nil)
  if err != nil {
    log.Fatal("ERROR! ", err)
  }
}
