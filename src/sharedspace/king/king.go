package king 

import "reflect"
import "fmt"
import "os"
import "bufio"
import "html/template"

type King struct {
}

func (*King) WriteToDB(msg reflect.Value){
}
//s := msg.Elem()
//typeOfT := s.Type()
//for i := 0; i < s.NumField(); i++ {
//        f := s.Field(i)
//        fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
//}

func (*King) WriteToFile(data interface{}){
    fmt.Println("Writing to file")
    f, _ := os.Create("/home/towski/code/asphalt/public/index.html")
    w := bufio.NewWriter(f)
    t, err := template.ParseFiles("/home/towski/code/asphalt/templates/index.html")
    _ = err
    //t = template.New("hello template") //create a new template with some name
    //t, _ = t.Parse("hello {{.Name}}!") //parse some content and generate a template, which is an internal representation
    t.Execute(w, data) //merge template ‘t’ with content of ‘p’
    w.Flush()
}

func Init()  {
 //   html_channel := make(chan interface{})
  //  post_channel := make(chan interface{})
//    StartDBWriter(post_channel, html_channel)
//    StartHTMLWriter(html_channel)
    // Returns the user with the given id 
}
