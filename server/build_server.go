package server

import _ "reflect"
import "bufio"
import "os"
import "html/template"
import "net/rpc"
import "net/rpc/jsonrpc"
import "net"
import "log"
import "path"
import "strings"
import "strconv"
import _ "log"

type PostIndex struct {
    Posts []Post
}

func WriteToFile(template_html string, output_dir string, data interface{}){
    defer func() {
        if r := recover(); r != nil {
            log.Println("Recovered in f", r)
            return
        }
    }()
    stat, err := os.Stat(template_html)
    if err != nil {
        log.Println("Couldn't find file " + template_html)
        return
    }
    if stat.IsDir() {
        log.Println("Is a dir, returning")
        return
    }
    if strings.HasSuffix(output_dir, ".html") == false {
        os.MkdirAll(output_dir, 0777)
    }
    f, err := os.Create(GetOutputFile(template_html, output_dir))
    if err != nil {
        log.Println("Couldn't find output file")
        return
    }
    w := bufio.NewWriter(f)
    t, err := template.ParseFiles(template_html)
    _ = err
    //t = template.New("hello template") //create a new template with some name
    //t, _ = t.Parse("hello {{.Name}}!") //parse some content and generate a template, which is an internal representation
    t.Execute(w, data) //merge template ‘t’ with content of ‘p’
    w.Flush()
}

func GetOutputFile(template_html string, output_dir string) string{
    if(path.Base(template_html) == "id.html"){
        return output_dir
    } else {
        return output_dir + "/" + path.Base(template_html)
    }
}

type GenericBuildServer struct {
}

type GenericHtml struct {
    File string
    Id int
    Class string
    Data interface{}
}

func (*GenericBuildServer) BuildHTML(generic_html GenericHtml, ret *int) error {
    if generic_html.Id != 0 {
        WriteToFile("templates/" + generic_html.Class + "/" + generic_html.File, "public/" + generic_html.Class + "/" + strconv.Itoa(generic_html.Id), generic_html.Data)
    } else {
        WriteToFile("templates/" + generic_html.File, "public", generic_html.Data)
    }
    return nil
}

type BuildServer struct {
    *rpc.Server
}

func NewBuildServer() (*BuildServer){
    os.Remove("/tmp/build.sock")
	//runtime.GOMAXPROCS(4)
    server := rpc.NewServer()
	server.Register(&GenericBuildServer{})
    return &BuildServer{server}
}

func (b *BuildServer) Start(){
    var rpcSocket net.Listener
	var err error
	var conn net.Conn
	if rpcSocket, err = net.Listen("unix", "/tmp/build.sock"); err != nil {
		panic(err)
	}
	//defer server.Close()
    log.Println("Starting build server..")
	for {
		if conn, err = rpcSocket.Accept(); err != nil {
			panic(err)
		}
		go b.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

