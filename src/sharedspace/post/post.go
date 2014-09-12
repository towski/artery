package post

import "reflect"
import "sharedspace/king"
import "fmt"

type Post struct {
    king.King
    Title string
}

func (*Post) hey(){
    fmt.Println("hey")
}

func (*Post) Data(){
    fmt.Println("hey")
}

func StartDBWriter(post_channel chan Post, html_channel chan Post){
    go func() {
        for msg := range post_channel {
            fmt.Println(msg)
            //msg.WriteToDB(reflect.ValueOf(msg))
            _ = reflect.ValueOf(msg)
            html_channel <- msg
        }
    }()
}

func StartHTMLWriter(html_channel chan Post){
    go func() {
        for msg := range html_channel {
            msg.WriteToFile(msg)
        }
    }()
}

var Post_channel = make(chan Post)
var Html_channel = make(chan Post)
func Init()  {
    StartDBWriter(Post_channel, Html_channel)
    StartHTMLWriter(Html_channel)
    // Returns the user with the given id 
}
