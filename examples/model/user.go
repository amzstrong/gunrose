package model

import (
    "gunrose/lib"
    "github.com/gohouse/gorose"
)

type UserModel struct {
    lib.BaseModel
}

func (u *UserModel)Flag(f string) *UserModel {
    u.InitFlag(f)
    return u
}

func (u  *UserModel) M() *gorose.Database {
    u.InitConnect("local", "user")
    return u.Con.Table(u.TableReal)
}






