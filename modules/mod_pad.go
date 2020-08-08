package main

import (
    "fmt"
    "newmod/lib"
)

func main() {
    s1 := "abc123-~ｱｲｳ|"
    s2 := "ａｂｃ１２３－～アイウ｜"
    fmt.Println(s1)
    ns := lib.Convert(s2)
    fmt.Println(ns)
}
