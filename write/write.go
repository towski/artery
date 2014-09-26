package write

import "net"
import "net/rpc"
import "net/rpc/jsonrpc"
import "log"

func Client() *rpc.Client{
    conn, err := net.Dial("unix", "/tmp/build.sock")
    if err != nil {
        log.Fatal("dialing:", err)
    }
    client := jsonrpc.NewClient(conn)
    return client
}

func main(){
    // Synchronous call
    client := Client()
    err := client.Call("BuildServer.Build", 1, nil)
    if err != nil {
        log.Fatal("error:", err)
    }
}
