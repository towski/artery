package main

import "net/http"
import "html"
import "time"
import "fmt"
import "log"
import _ "reflect"
import _ "unsafe"
import "sharedspace/post"

type timeHandler struct {
  zone *time.Location
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  tm := time.Now().In(th.zone).Format(time.RFC1123)
  w.Write([]byte("The time is: " + tm))
}

func newTimeHandler(name string) *timeHandler {
  return &timeHandler{zone: time.FixedZone(name, 0)}
}

func fooHandler(){
    log.Fatal("hey")
}

func main (){
    http.Handle("/foo", newTimeHandler("EST"))
    http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
        p := post.Post{Title: "yolo"}
        fmt.Println(p)
        post.Post_channel <- p
        //title := r.URL.Path[len("/edit/"):]
    })
    post.Init()
    log.Fatal(http.ListenAndServe(":8081", nil))
}


