package main

import (
	"fmt"
	"github.com/ArchieT/trmstac/get"
//	"github.com/ArchieT/trmstac/stadata"
	"encoding/json"
)

type blad struct {
	err error
}

func main(){
	tabl, cza, err := get.Download()
	data := make(map[string]interface{})
	data["sta"] = *tabl
	data["cza"] = cza
	b,err := json.Marshal(data)
	if err!=nil {
		fmt.Println("error",err)
	}
	fmt.Println(string(b))
}
