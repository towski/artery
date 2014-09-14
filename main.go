package main

import "net/http"
import _ "html"
import "time"
import "fmt"
import "log"
import _ "reflect"
import _ "unsafe"
import "github.com/towski/artery/post"

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
    http.HandleFunc("/artery/foo.json", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Your changes will appear shortly")
    })
    http.HandleFunc("/artery/bar", func(w http.ResponseWriter, r *http.Request) {
        //fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
        p := &post.Post{Title: "yolo"}
        fmt.Println(p)
        post.Post_channel <- p
        http.Redirect(w, r, "/index.html", http.StatusFound)
        //title := r.URL.Path[len("/edit/"):]
    })
    post.Init()
    log.Fatal(http.ListenAndServe(":8081", nil))
}


