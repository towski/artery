package post

import _ "reflect"
import "fmt"
import "bufio"
import "os"
import "html/template"
import "net/rpc"
import "net"
import _ "log"

type PostIndex struct {
    Posts []Post
}

func WriteToFile(){
    data := PostIndex{}
    err := data_client.Call("DataServer.GetPostIndex", 1, &data)
    fmt.Println("Writing to file")
    f, _ := os.Create("/home/towski/gopath/src/github.com/towski/artery/public/index.html")
    w := bufio.NewWriter(f)
    t, err := template.ParseFiles("/home/towski/gopath/src/github.com/towski/artery/templates/index.html")
    _ = err
    //t = template.New("hello template") //create a new template with some name
    //t, _ = t.Parse("hello {{.Name}}!") //parse some content and generate a template, which is an internal representation
    t.Execute(w, data) //merge template ‘t’ with content of ‘p’
    w.Flush()
}

type BuildServer struct {
}

func (*BuildServer) Build(id int, ret *int) error {
    WriteToFile()
    return nil
}

func StartBuildServer(){
    var rpcSocket net.Listener
    os.Remove("/tmp/build.sock")
	var err error
	var conn net.Conn
	//runtime.GOMAXPROCS(4)
	rpc.Register(&BuildServer{})
	if rpcSocket, err = net.Listen("unix", "/tmp/build.sock"); err != nil {
		panic(err)
	}
	defer rpcSocket.Close()
	for {
		if conn, err = rpcSocket.Accept(); err != nil {
			panic(err)
		}
		go rpc.ServeConn(conn)
	}
}

