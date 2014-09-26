package server

import "reflect"
import _ "fmt"
import _ "os"
import _ "bufio"
import _ "html/template"

type Model struct {
}

func (*Model) WriteToDB(msg reflect.Value){
}
//s := msg.Elem()
//typeOfT := s.Type()
//for i := 0; i < s.NumField(); i++ {
//        f := s.Field(i)
//        fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
//}

