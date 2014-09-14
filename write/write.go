package write

import "net/rpc"
import "log"

func Client() *rpc.Client{
    client, err := rpc.Dial("unix", "/tmp/build.sock")
    if err != nil {
        log.Fatal("dialing:", err)
    }
}

func main(){
    // Synchronous call
    client := Client()
    err = client.Call("BuildServer.Build", 1, nil)
    if err != nil {
        log.Fatal("error:", err)
    }
}
