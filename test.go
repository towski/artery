package test

import (
    "fmt"
    "log"
    "net"
    "errors"
    "net/rpc"
    "net/rpc/jsonrpc"
    "sync"
)

type Args struct {
    A, B int
}

type Quotient struct {
    Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
    *reply = args.A * args.B
    return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
    if args.B == 0 {
        return errors.New("divide by zero")
    }
    quo.Quo = args.A / args.B
    quo.Rem = args.A % args.B
    return nil
}

func startServer() {
    arith := new(Arith)

    server := rpc.NewServer()
    server.Register(arith)

    server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

    l, e := net.Listen("tcp", ":8222")
    if e != nil {
        log.Fatal("listen error:", e)
    }

    for {
        conn, err := l.Accept()
        if err != nil {
            log.Fatal(err)
        }

        go server.ServeCodec(jsonrpc.NewServerCodec(conn))
    }
}

func maino() {
    var wg sync.WaitGroup
    wg.Add(1)

    go startServer()

    conn, err := net.Dial("tcp", "localhost:8222")

    if err != nil {
        panic(err)
    }
    defer conn.Close()

    args := &Args{7, 8}
    var reply int

    c := jsonrpc.NewClient(conn)

    for i := 0; i < 1; i++ {

        err = c.Call("Arith.Multiply", args, &reply)
        if err != nil {
            log.Fatal("arith error:", err)
        }
        fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
    }
    wg.Wait()
}
