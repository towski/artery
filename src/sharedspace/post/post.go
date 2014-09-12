package post

import "reflect"
import "sharedspace/king"
import "fmt"
import "log"
import "github.com/coopernurse/gorp"

type Post struct {
    Title string
    Id int64
    king.King
}

func (*Post) hey(){
    fmt.Println("hey")
}

func (*Post) Data(){
    fmt.Println("hey")
}

func StartDBWriter(post_channel chan *Post, html_channel chan *PostIndex, dbmap *gorp.DbMap){
    go func() {
        for msg := range post_channel {
            err := dbmap.Insert(msg)
            if (err != nil){
                fmt.Println(err)
                log.Fatalln("no insert")
            }
            //msg.WriteToDB(reflect.ValueOf(msg))
            _ = reflect.ValueOf(msg)
            post_index := &PostIndex{}
            dbmap.Select(&post_index.Posts, "select * from Post order by Id")
            html_channel <- post_index
        }
    }()
}

type PostIndex struct {
    Posts []Post
}

func StartHTMLWriter(html_channel chan *PostIndex){
    go func() {
        for msg := range html_channel {
            fmt.Println(msg)
            (&Post{}).WriteToFile(msg)
        }
    }()
}

var Post_channel = make(chan *Post)
var Html_channel = make(chan *PostIndex)
func Init(dbmap *gorp.DbMap)  {
    StartDBWriter(Post_channel, Html_channel, dbmap)
    StartHTMLWriter(Html_channel)
    // Returns the user with the given id 
}
