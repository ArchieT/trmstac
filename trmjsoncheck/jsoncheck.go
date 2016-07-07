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

func main() {
	down, err := get.Download()
	cza := down.Time
	if err != nil {
		fmt.Println("error", err)
	}
	tabl, err := down.ParseSta()
	data := make(map[string]interface{})
	data["sta"] = tabl
	data["cza"] = cza
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(string(b))
}
