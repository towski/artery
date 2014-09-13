package main

import "net/rpc"
import "log"

func main(){
    // Synchronous call
    client, err := rpc.Dial("unix", "/tmp/build.sock")
    if err != nil {
        log.Fatal("dialing:", err)
    }
    err = client.Call("BuildServer.Build", 1, nil)
    if err != nil {
        log.Fatal("error:", err)
    }
}
