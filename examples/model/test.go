package model

import (
    "gunrose/lib"
    "github.com/gohouse/gorose"
)

type TestMode struct {
    lib.BaseModel
}

func (t *TestMode)Flag(f string) *TestMode {
    t.InitFlag(f)

    return t
}

func (u  *TestMode) M() *gorose.Database {
    u.InitConnect("local", "test")
    return u.Con.Table(u.TableReal)
}







