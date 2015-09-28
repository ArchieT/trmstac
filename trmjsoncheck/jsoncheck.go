package main

import (
	"fmt"
	"github.com/ArchieT/trmstac/get"
//	"github.com/ArchieT/trmstac/stadata"
	"encoding/json"
)

func main(){
	a := get.Download()
	b,err := json.Marshal(a)
	if err!=nil {
		fmt.Println("damn")
	}
	fmt.Println(string(b))
}
