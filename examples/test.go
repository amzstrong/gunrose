package main

import (
    "gunrose/examples/model"
    "fmt"
    "gunrose/lib"
)

func main() {
    lib.Dbconfig = "/Users/alex/go/src/gunrose/config.json"
    user := new(model.UserModel)
    res, err := user.M().Fields("id").First()
    res2, _ := user.Con.Query("select * from test")

    fmt.Println(res);
    fmt.Println(res2);
    fmt.Println(err);
}
