package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main() {
    db, err := sql.Open("mysql", "root:mysql@/asphalt")
    if err != nil {
        panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
    }

    stmtIns, err := db.Prepare("describe post")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtIns.Close()
}
