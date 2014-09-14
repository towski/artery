package post

import _ "reflect"
import "fmt"
import "log"
import "github.com/coopernurse/gorp"
import "net"
import "net/rpc"
import "os"

type Post struct {
    Model
    Title string
    Id int64
}

func (*Post) hey(){
    fmt.Println("hey")
}

func (*Post) Data(){
    fmt.Println("hey")
}

func StartDBWriter(post_channel chan *Post, dbmap *gorp.DbMap){
    go func() {
        for msg := range post_channel {
            err := dbmap.Insert(msg)
            if (err != nil){
                fmt.Println(err)
                log.Fatalln("no insert")
            }
            build_client.Go("BuildServer.Build", 1, 1, nil)
            //msg.WriteToDB(reflect.ValueOf(msg))
        }
    }()
}

type DataServer struct {
}

func (*DataServer) GetPost(id int, post *Post) error {
    err := dbmap_global.SelectOne(post, "select * from Post where Id=?", id)
    if (err != nil){
        log.Println("errr %s", err)
    }
    log.Println("p2 row:", post)
    return nil
}

func (*DataServer) GetPostIndex(id int, post_index *PostIndex) error {
    dbmap_global.Select(&post_index.Posts, "select * from Post order by Id")
    log.Println(post_index.Posts)
    return nil
}

func StartDataServer(){
    var rpcSocket net.Listener
    os.Remove("/tmp/data.sock")
	var err error
	var conn net.Conn
	//runtime.GOMAXPROCS(4)
	rpc.Register(&DataServer{})
	if rpcSocket, err = net.Listen("unix", "/tmp/data.sock"); err != nil {
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

var Post_channel = make(chan *Post)
var dbmap_global *gorp.DbMap
var data_client *rpc.Client
var build_client *rpc.Client
func Init(dbmap *gorp.DbMap)  {
    data_client, _ = rpc.Dial("unix", "/tmp/data.sock")
    build_client, _ = rpc.Dial("unix", "/tmp/build.sock")
    dbmap_global = dbmap
    StartDBWriter(Post_channel, dbmap)
    go StartDataServer()
    go StartBuildServer()
    // Returns the user with the given id 
}
