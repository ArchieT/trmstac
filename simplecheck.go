package main

import (
	"fmt"
	"github.com/ArchieT/trmstac/get"
//	"github.com/ArchieT/trmstac/stadata"
)

func main(){
	fmt.Println("start")
	a := get.Download()
	fmt.Println("down")
	fmt.Println(a)
	fmt.Println("juz")
}
