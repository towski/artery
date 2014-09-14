package main

import "net/rpc"
import "log"
import "fmt"
import "github.com/towski/artery/post"

func main(){
    // Synchronous call
    client, err := rpc.Dial("unix", "/tmp/data.sock")
    if err != nil {
        log.Fatal("dialing:", err)
    }
    reply := post.Post{}
    err = client.Call("DataServer.GetPost", 1, &reply)
    if err != nil {
        log.Fatal("error:", err)
    }
    fmt.Printf("Post: %s\n", reply.Title)
}
