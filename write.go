package main

import "os"
import "bufio"
import "fmt"
import "html/template"
import "sharedspace/post"

func main(){
    if(len(os.Args) < 2){
        fmt.Println("Need a first argument")
        os.Exit(0)
    }
    file := os.Args[1]
    f, _ := os.Create("./public/" + file)
    w := bufio.NewWriter(f)
    t, err := template.ParseFiles("./templates/" + file)
    _ = err

    post := post.Post{}
    _ = post

    //t = template.New("hello template") //create a new template with some name
    //t, _ = t.Parse("hello {{.Name}}!") //parse some content and generate a template, which is an internal representation
    t.Execute(w, nil) //merge template ‘t’ with content of ‘p’
    w.Flush()
    f.Close()
}
