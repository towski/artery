package server

import _ "reflect"
import "fmt"
import "log"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "github.com/coopernurse/gorp"
import "net"
import "net/rpc"
import "net/rpc/jsonrpc"
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
            //build_client.Go("BuildServer.Build", 1, 1, nil)
            //msg.WriteToDB(reflect.ValueOf(msg))
        }
    }()
}

/*func (*DataServer) GetPost(id int, post *Post) error {
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
*/

type DataServer struct {
    *rpc.Server
}

func DataClient() *rpc.Client{
    conn, err := net.Dial("unix", "/tmp/data.sock")
    if err != nil {
        log.Fatal("dialing:", err)
    }
    client := jsonrpc.NewClient(conn)
    return client
}

func NewDataServer() (*DataServer){
    os.Remove("/tmp/data.sock")
	//runtime.GOMAXPROCS(4)
    server := rpc.NewServer()
	//server.Register(&DataServer{})
    return &DataServer{server}
}

func (d *DataServer) Start(){
    rpcSocket, err := net.Listen("unix", "/tmp/data.sock")
	if err != nil {
		panic(err)
	}
    //data_client, _ := rpc.Dial("unix", "/tmp/data.sock")
	defer rpcSocket.Close()
    log.Println("Starting data server..")
	for {
        conn, err := rpcSocket.Accept(); 
		if err != nil {
			panic(err)
		}
		go d.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

//var Post_channel = make(chan *Post)
var Dbmap_global *gorp.DbMap

func Init(){
    var db, _ = sql.Open("mysql", "root:mysql@/asphalt")
    var dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
    Dbmap_global = dbmap
}
